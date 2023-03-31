package inject

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_Decorate(t *testing.T) {
	type t1 struct {
		count int
	}
	c := NewContainer()

	inst := t1{
		count: 0,
	}
	c.Register(Decorate(func() t1 {
		inst.count = 1
		return inst
	}, func(a t1) t1 {
		a.count = 100
		return a
	}))

	err := c.Invoke(func(t1 t1) {
		require.Equal(t, 100, t1.count)
	})
	require.NoError(t, err)

}

func Test_NonSingleton(t *testing.T) {
	type t1 struct {
		count int
	}
	c := NewContainer()

	invocationCount := 0
	inst := t1{
		count: 0,
	}
	c.Register(func() t1 {
		invocationCount++
		inst.count++
		return inst
	})

	err := c.Invoke(func(t1 t1) {
		require.Equal(t, 1, t1.count)
	})
	require.NoError(t, err)
	require.Equal(t, 1, invocationCount)

	err = c.Invoke(func(t1 t1) {
		require.Equal(t, 2, t1.count)
	})
	require.NoError(t, err)
	require.Equal(t, 2, invocationCount)
}

func Test_Singleton(t *testing.T) {
	type t1 struct {
		count int
	}
	c := NewContainer()

	invocationCount := 0
	inst := t1{
		count: 0,
	}
	c.Register(Singleton(func() t1 {
		invocationCount++
		inst.count++
		return inst
	}))

	err := c.Invoke(func(t1 t1) {
		require.Equal(t, 1, t1.count)
	})
	require.NoError(t, err)
	require.Equal(t, 1, invocationCount)

	err = c.Invoke(func(t1 t1) {
		require.Equal(t, 1, t1.count)
	})
	require.NoError(t, err)
	require.Equal(t, 1, invocationCount)
}
