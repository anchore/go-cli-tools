package inject

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_Container(t *testing.T) {
	type data struct {
		name string
	}

	type data1 struct {
		name string
	}

	type data2 struct {
		name string
	}

	d := &data{
		name: "1 name",
	}

	d1 := data1{
		name: "d1",
	}

	c := NewContainer()

	c.Bind(d, d1)

	c.Register(func(d data, d2 *data1) data2 {
		return data2{
			name: strings.ReplaceAll(d.name, "1", "2"),
		}
	})

	err := c.Invoke(func(d data2) {
		require.NotNil(t, d)
		require.Equal(t, "2 name", d.name)
	})

	require.NoError(t, err)

	type data3 struct {
		name string
	}

	type data4 struct {
		name string
	}

	c2 := NewContainer(c)

	c2.Register(func(d *data, d2 data2) *data3 {
		return &data3{
			name: d.name + ", " + d2.name,
		}
	})

	var d4 data4
	err = c2.Invoke(func(d data3) {
		d4 = data4{
			name: d.name,
		}
	})
	require.NoError(t, err)

	require.Equal(t, "1 name, 2 name", d4.name)

	c3 := NewContainer(c2)

	type err1 struct{}

	c3.Register(func() (err1, error) {
		return err1{}, fmt.Errorf("an error")
	})

	err = c3.Invoke(func(err1 err1) {})
	require.Error(t, err)

	var d1v2 data1
	err = c3.Invoke(func(d1 *data1) {
		d1v2 = *d1
	})
	require.NoError(t, err)
	require.Equal(t, "d1", d1v2.name)

	err = c3.Invoke(func(d1 *data1) error {
		return fmt.Errorf("direct error")
	})
	require.Error(t, err)

	c4 := NewContainer(c3)
	err = c4.Invoke(func(err1 err1) {})
	require.Error(t, err)

	dv3, err := c4.Resolve(data{})
	require.NoError(t, err)
	require.Equal(t, *d, dv3)

	dv4, err := c4.Resolve(&data{})
	require.NoError(t, err)
	require.Equal(t, d, dv4)
}

func Test_InjectingContainer(t *testing.T) {
	c := NewContainer()

	err := c.Invoke(func(c Container) {
		require.NotNil(t, c)
	})
	require.NoError(t, err)
}

func Test_UnresolvedDependency(t *testing.T) {
	c := NewContainer()

	type unbound struct{}

	err := c.Invoke(func(u unbound) {})

	require.Error(t, err)
	require.Contains(t, err.Error(), "unbound")
	require.Contains(t, err.Error(), "Test_UnresolvedDependency")
	require.Contains(t, err.Error(), "container_test.go")
}
