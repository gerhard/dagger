# This file generated by `mix dagger.gen`. Please DO NOT EDIT.
defmodule Dagger.LocalModuleRef do
  @moduledoc "A reference to a module loaded from a path locally relative to a directory."
  use Dagger.Core.QueryBuilder
  @type t() :: %__MODULE__{}
  defstruct [:selection, :client]

  (
    @doc "A unique identifier for this LocalModuleRef."
    @spec id(t()) :: {:ok, Dagger.LocalModuleRefID.t()} | {:error, term()}
    def id(%__MODULE__{} = local_module_ref) do
      selection = select(local_module_ref.selection, "id")
      execute(selection, local_module_ref.client)
    end
  )

  (
    @doc ""
    @spec module_source_path(t()) :: {:ok, Dagger.String.t()} | {:error, term()}
    def module_source_path(%__MODULE__{} = local_module_ref) do
      selection = select(local_module_ref.selection, "moduleSourcePath")
      execute(selection, local_module_ref.client)
    end
  )
end
