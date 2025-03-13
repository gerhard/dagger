package drivers

import (
	"context"
	"crypto/tls"
	"errors"
	"fmt"
	"net"
	"net/url"
	"time"

	"github.com/dagger/dagger/internal/cloud"
	"github.com/dagger/dagger/internal/cloud/auth"
)

func init() {
	register("dagger-cloud", &daggerCloudDriver{})
}

// daggerCloudDriver creates and manages a remote Dagger Engine, then connects to it
type daggerCloudDriver struct{}

type daggerCloudConnector struct {
	EngineSpec cloud.EngineSpec
}

func (dc *daggerCloudConnector) Connect(ctx context.Context) (net.Conn, error) {
	serverAddr := dc.EngineSpec.URL

	// Extract hostname for SNI
	host, _, err := net.SplitHostPort(serverAddr)
	if err != nil {
		return nil, fmt.Errorf("failed to split host and port from '%s': %w", serverAddr, err)
	}

	clientCert, err := dc.EngineSpec.TLSCertificate()
	if err != nil {
		return nil, fmt.Errorf("failed to unserialize cert: %w", err)
	}

	tlsConfig := &tls.Config{
		Certificates: []tls.Certificate{*clientCert},
		ServerName:   host,             // Required for SNI and server certificate validation
		MinVersion:   tls.VersionTLS12, // Recommended
	}

	dialer := net.Dialer{Timeout: 10 * time.Second}
	rawConn, err := dialer.DialContext(ctx, "tcp", serverAddr)
	if err != nil {
		return nil, fmt.Errorf("failed to dial TCP to '%s': %w", serverAddr, err)
	}
	// No defer rawConn.Close() here yet, as tls.Client will take ownership

	// Create TLS Client Connection
	// tls.Client wraps the rawConn. If Handshake fails, rawConn should be closed.
	// If Handshake succeeds, closing tlsConn will close rawConn.
	tlsConn := tls.Client(rawConn, tlsConfig)

	// Use HandshakeContext for better cancellation and timeout control
	handshakeCtx, cancelHandshake := context.WithTimeout(ctx, 15*time.Second) // Handshake timeout
	defer cancelHandshake()

	if err := tlsConn.HandshakeContext(handshakeCtx); err != nil {
		rawConn.Close() // Ensure raw connection is closed if handshake fails
		return nil, fmt.Errorf("TLS handshake failed with '%s': %w", serverAddr, err)
	}

	return tlsConn, nil
}

func (d *daggerCloudDriver) Provision(ctx context.Context, _ *url.URL, opts *DriverOpts) (Connector, error) {
	client, err := cloud.NewClient(ctx)
	if err != nil {
		return nil, errors.New("Please run `dagger login <org>` first.")
	}

	_, err = client.User(ctx)
	if err != nil {
		return nil, err
	}

	org, err := auth.CurrentOrg()
	if org == nil {
		return nil, errors.New("Please select which org you want to create an Engines in by running `dagger login <org>.")
	}

	return d.create(ctx, client)
}

func (d *daggerCloudDriver) create(ctx context.Context, client *cloud.Client) (*daggerCloudConnector, error) {
	engineSpec, err := client.Engine(ctx)
	if err != nil {
		return nil, err
	}
	return &daggerCloudConnector{EngineSpec: *engineSpec}, nil
}
