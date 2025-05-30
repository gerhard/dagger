<?php

/**
 * This class has been generated by dagger-php-sdk. DO NOT EDIT.
 */

declare(strict_types=1);

namespace Dagger;

/**
 * A Dagger module.
 */
class Module extends Client\AbstractObject implements Client\IdAble
{
    /**
     * The dependencies of the module.
     */
    public function dependencies(): array
    {
        $leafQueryBuilder = new \Dagger\Client\QueryBuilder('dependencies');
        return (array)$this->queryLeaf($leafQueryBuilder, 'dependencies');
    }

    /**
     * The doc string of the module, if any
     */
    public function description(): string
    {
        $leafQueryBuilder = new \Dagger\Client\QueryBuilder('description');
        return (string)$this->queryLeaf($leafQueryBuilder, 'description');
    }

    /**
     * Enumerations served by this module.
     */
    public function enums(): array
    {
        $leafQueryBuilder = new \Dagger\Client\QueryBuilder('enums');
        return (array)$this->queryLeaf($leafQueryBuilder, 'enums');
    }

    /**
     * The generated files and directories made on top of the module source's context directory.
     */
    public function generatedContextDirectory(): Directory
    {
        $innerQueryBuilder = new \Dagger\Client\QueryBuilder('generatedContextDirectory');
        return new \Dagger\Directory($this->client, $this->queryBuilderChain->chain($innerQueryBuilder));
    }

    /**
     * A unique identifier for this Module.
     */
    public function id(): ModuleId
    {
        $leafQueryBuilder = new \Dagger\Client\QueryBuilder('id');
        return new \Dagger\ModuleId((string)$this->queryLeaf($leafQueryBuilder, 'id'));
    }

    /**
     * Interfaces served by this module.
     */
    public function interfaces(): array
    {
        $leafQueryBuilder = new \Dagger\Client\QueryBuilder('interfaces');
        return (array)$this->queryLeaf($leafQueryBuilder, 'interfaces');
    }

    /**
     * The name of the module
     */
    public function name(): string
    {
        $leafQueryBuilder = new \Dagger\Client\QueryBuilder('name');
        return (string)$this->queryLeaf($leafQueryBuilder, 'name');
    }

    /**
     * Objects served by this module.
     */
    public function objects(): array
    {
        $leafQueryBuilder = new \Dagger\Client\QueryBuilder('objects');
        return (array)$this->queryLeaf($leafQueryBuilder, 'objects');
    }

    /**
     * The container that runs the module's entrypoint. It will fail to execute if the module doesn't compile.
     */
    public function runtime(): Container
    {
        $innerQueryBuilder = new \Dagger\Client\QueryBuilder('runtime');
        return new \Dagger\Container($this->client, $this->queryBuilderChain->chain($innerQueryBuilder));
    }

    /**
     * The SDK config used by this module.
     */
    public function sdk(): SDKConfig
    {
        $innerQueryBuilder = new \Dagger\Client\QueryBuilder('sdk');
        return new \Dagger\SDKConfig($this->client, $this->queryBuilderChain->chain($innerQueryBuilder));
    }

    /**
     * Serve a module's API in the current session.
     *
     * Note: this can only be called once per session. In the future, it could return a stream or service to remove the side effect.
     */
    public function serve(?bool $includeDependencies = null): void
    {
        $leafQueryBuilder = new \Dagger\Client\QueryBuilder('serve');
        if (null !== $includeDependencies) {
        $leafQueryBuilder->setArgument('includeDependencies', $includeDependencies);
        }
        $this->queryLeaf($leafQueryBuilder, 'serve');
    }

    /**
     * The source for the module.
     */
    public function source(): ModuleSource
    {
        $innerQueryBuilder = new \Dagger\Client\QueryBuilder('source');
        return new \Dagger\ModuleSource($this->client, $this->queryBuilderChain->chain($innerQueryBuilder));
    }

    /**
     * Forces evaluation of the module, including any loading into the engine and associated validation.
     */
    public function sync(): ModuleId
    {
        $leafQueryBuilder = new \Dagger\Client\QueryBuilder('sync');
        return new \Dagger\ModuleId((string)$this->queryLeaf($leafQueryBuilder, 'sync'));
    }

    /**
     * Retrieves the module with the given description
     */
    public function withDescription(string $description): Module
    {
        $innerQueryBuilder = new \Dagger\Client\QueryBuilder('withDescription');
        $innerQueryBuilder->setArgument('description', $description);
        return new \Dagger\Module($this->client, $this->queryBuilderChain->chain($innerQueryBuilder));
    }

    /**
     * This module plus the given Enum type and associated values
     */
    public function withEnum(TypeDefId|TypeDef $enum): Module
    {
        $innerQueryBuilder = new \Dagger\Client\QueryBuilder('withEnum');
        $innerQueryBuilder->setArgument('enum', $enum);
        return new \Dagger\Module($this->client, $this->queryBuilderChain->chain($innerQueryBuilder));
    }

    /**
     * This module plus the given Interface type and associated functions
     */
    public function withInterface(TypeDefId|TypeDef $iface): Module
    {
        $innerQueryBuilder = new \Dagger\Client\QueryBuilder('withInterface');
        $innerQueryBuilder->setArgument('iface', $iface);
        return new \Dagger\Module($this->client, $this->queryBuilderChain->chain($innerQueryBuilder));
    }

    /**
     * This module plus the given Object type and associated functions.
     */
    public function withObject(TypeDefId|TypeDef $object): Module
    {
        $innerQueryBuilder = new \Dagger\Client\QueryBuilder('withObject');
        $innerQueryBuilder->setArgument('object', $object);
        return new \Dagger\Module($this->client, $this->queryBuilderChain->chain($innerQueryBuilder));
    }
}
