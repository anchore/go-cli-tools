package inject

import (
	"fmt"
	"reflect"
)

// Container is the interface used by the simple injection container
type Container interface {
	// Register registers a provider function
	Register(providers ...any)

	// Bind binds a value directly into the container
	Bind(values ...any)

	// Resolve returns the resolved value for a given type, e.g. value, err := c.Resolve(Type{})
	Resolve(typ any) (any, error)

	// Invoke invokes a function with injected parameters that may an error as the last value
	Invoke(fn any) error
}

// ProviderNotFoundError is a standard error that can be returned from a container.Get call
type ProviderNotFoundError struct {
	typ reflect.Type
}

func (n ProviderNotFoundError) Error() string {
	return fmt.Sprintf("provider not found for: %s", typeName(n.typ))
}

var _ error = (*ProviderNotFoundError)(nil)

// ProviderNotFound can be used to check if a Get call returns a "not found" error
var ProviderNotFound = ProviderNotFoundError{}
