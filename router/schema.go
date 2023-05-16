package router

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/dagger/dagger/core/pipeline"
	"github.com/dagger/graphql"
	"github.com/opencontainers/go-digest"
	"github.com/vito/progrock"
)

type LoadedSchema interface {
	Name() string
	Schema() string
}

type ExecutableSchema interface {
	LoadedSchema
	Resolvers() Resolvers
	Dependencies() []ExecutableSchema
}

type Resolvers map[string]Resolver

type Resolver interface {
	_resolver()
}

type ObjectResolver map[string]graphql.FieldResolveFn

func (ObjectResolver) _resolver() {}

type ScalarResolver struct {
	Serialize    graphql.SerializeFn
	ParseValue   graphql.ParseValueFn
	ParseLiteral graphql.ParseLiteralFn
}

func (ScalarResolver) _resolver() {}

type StaticSchemaParams struct {
	Name         string
	Schema       string
	Resolvers    Resolvers
	Dependencies []ExecutableSchema
}

func StaticSchema(p StaticSchemaParams) ExecutableSchema {
	return &staticSchema{p}
}

var _ ExecutableSchema = &staticSchema{}

type staticSchema struct {
	StaticSchemaParams
}

func (s *staticSchema) Name() string {
	return s.StaticSchemaParams.Name
}

func (s *staticSchema) Schema() string {
	return s.StaticSchemaParams.Schema
}

func (s *staticSchema) Resolvers() Resolvers {
	return s.StaticSchemaParams.Resolvers
}

func (s *staticSchema) Dependencies() []ExecutableSchema {
	return s.StaticSchemaParams.Dependencies
}

type Context struct {
	context.Context
	ResolveParams graphql.ResolveParams

	// Vertex is a recorder for sending logs to the request's vertex in the
	// progress stream.
	Vertex *progrock.VertexRecorder
}

// Pipelineable is any object which can return a pipeline.Path.
//
// It is used to construct a Progrock recorder group that is passed via ctx to
// the resolver.
type Pipelineable interface {
	PipelinePath() pipeline.Path
}

// Digestible is any object which can return a digest of its content.
//
// It is used to record the request's result as an output of the request's
// vertex in the progress stream.
type Digestible interface {
	Digest() (digest.Digest, error)
}

// ToResolver transforms any function f with a *Context, a parent P and some args A that returns a Response R and an error
// into a graphql resolver graphql.FieldResolveFn.
func ToResolver[P any, A any, R any](f func(*Context, P, A) (R, error)) graphql.FieldResolveFn {
	return func(p graphql.ResolveParams) (any, error) {
		recorder := progrock.RecorderFromContext(p.Context)

		var args A
		argBytes, err := json.Marshal(p.Args)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal args: %w", err)
		}
		if err := json.Unmarshal(argBytes, &args); err != nil {
			return nil, fmt.Errorf("failed to unmarshal args: %w", err)
		}

		parent, ok := p.Source.(P)
		if !ok {
			parentBytes, err := json.Marshal(p.Source)
			if err != nil {
				return nil, fmt.Errorf("failed to marshal parent: %w", err)
			}
			if err := json.Unmarshal(parentBytes, &parent); err != nil {
				return nil, fmt.Errorf("failed to unmarshal parent: %w", err)
			}
		}

		if pipelineable, ok := p.Source.(Pipelineable); ok {
			recorder = pipelineable.PipelinePath().RecorderGroup(recorder)
			p.Context = progrock.RecorderToContext(p.Context, recorder)
		}

		vtx, err := queryVertex(recorder, p)
		if err != nil {
			return nil, err
		}

		ctx := Context{
			Context:       p.Context,
			ResolveParams: p,
			Vertex:        vtx,
		}

		res, err := f(&ctx, parent, args)
		if err != nil {
			vtx.Done(err)
			return nil, err
		}

		if edible, ok := any(res).(Digestible); ok {
			dg, err := edible.Digest()
			if err != nil {
				return nil, fmt.Errorf("failed to compute digest: %w", err)
			}

			vtx.Output(dg)
		}

		vtx.Done(nil)

		return res, nil
	}
}

func PassthroughResolver(p graphql.ResolveParams) (any, error) {
	return ToResolver(func(ctx *Context, parent any, args any) (any, error) {
		if parent == nil {
			parent = struct{}{}
		}
		return parent, nil
	})(p)
}

func ErrResolver(err error) graphql.FieldResolveFn {
	return ToResolver(func(ctx *Context, parent any, args any) (any, error) {
		return nil, err
	})
}

func queryVertex(recorder *progrock.Recorder, params graphql.ResolveParams) (*progrock.VertexRecorder, error) {
	dig, err := queryDigest(params)
	if err != nil {
		return nil, fmt.Errorf("failed to compute query digest: %w", err)
	}

	var inputs []digest.Digest

	name := params.Info.FieldName
	if len(params.Args) > 0 {
		name += "("
		args := []string{}
		for name, val := range params.Args {
			if dg, ok := val.(Digestible); ok {
				d, err := dg.Digest()
				if err != nil {
					return nil, fmt.Errorf("failed to compute digest for param %q: %w", name, err)
				}

				inputs = append(inputs, d)

				// display digest instead
				val = d
			}

			jv, err := json.Marshal(val)
			if err != nil {
				return nil, fmt.Errorf("failed to marshal arg %s: %w", name, err)
			}

			args = append(args, fmt.Sprintf("%s: %s", name, jv))
		}
		name += strings.Join(args, ", ")
		name += ")"
	}

	if edible, ok := params.Source.(Digestible); ok {
		id, err := edible.Digest()
		if err != nil {
			return nil, fmt.Errorf("failed to compute digest: %w", err)
		}

		inputs = append(inputs, id)
	}

	return recorder.Vertex(
		dig,
		name,
		progrock.WithInputs(inputs...),
		progrock.Internal(),
	), nil
}

func queryDigest(params graphql.ResolveParams) (digest.Digest, error) {
	type subset struct {
		Source any
		Field  string
		Args   map[string]any
	}

	payload, err := json.Marshal(subset{
		Source: params.Source,
		Field:  params.Info.FieldName,
		Args:   params.Args,
	})
	if err != nil {
		return "", err
	}

	return digest.SHA256.FromBytes(payload), nil
}