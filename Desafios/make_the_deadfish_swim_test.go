package desafios

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestMakeTheDeadfishSwim(t *testing.T) {

	require.Equal(t, Parse("ooo"), []int{0, 0, 0})
	require.Equal(t, Parse("ioioio"), []int{1, 2, 3})
	require.Equal(t, Parse("idoiido"), []int{0, 1})
	require.Equal(t, Parse("ssdizbsnsdkkincuybidds"), []int{})
	require.Equal(t, Parse("codewars"), []int{0})
}
