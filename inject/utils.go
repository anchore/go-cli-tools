package inject

import (
	"fmt"
	"reflect"
)

func Invoke[T any](c Container, fn any) (T, error) {
	c2, ok := c.(*container)
	var t T
	if !ok {
		return t, fmt.Errorf("unable to get return value from invoke, argument is not *container: %+v", c)
	}
	out, err := c2.invoke(reflect.ValueOf(fn))
	if !out.CanInterface() {
		return t, fmt.Errorf("unable get value from return: %+v", out)
	}
	v := out.Interface()
	if tv, ok := v.(T); ok {
		t = tv
	} else {
		typ := reflect.TypeOf(t)
		return t, fmt.Errorf("unable convert return value to expected type: %s", typeName(typ))
	}
	return t, err
}

func MustInvoke[T any](c Container, fn any) T {
	out, err := Invoke[T](c, fn)
	if err != nil {
		panic(err)
	}
	return out
}

func Decorate[T any](provider any, decorator func(T) T) any {
	baseFunc := reflect.ValueOf(provider)
	t := baseFunc.Type()
	validateProvider(t, provider)
	v := reflect.MakeFunc(t, func(args []reflect.Value) (results []reflect.Value) {
		results = baseFunc.Call(args)
		// TODO check for error returns
		out, ok := results[0].Interface().(T)
		if ok {
			out = decorator(out)
			results[0] = reflect.ValueOf(out)
		}
		return
	})
	return v.Interface()
}

func Singleton(provider any) any {
	value := nilValue
	baseFunc := reflect.ValueOf(provider)
	t := baseFunc.Type()
	validateProvider(t, provider)
	v := reflect.MakeFunc(baseFunc.Type(), func(args []reflect.Value) (results []reflect.Value) {
		if value != nilValue {
			results = make([]reflect.Value, t.NumOut(), t.NumOut())
			results[0] = value
			return
		}
		results = baseFunc.Call(args)
		value = results[0]
		return
	})
	return v.Interface()
}
