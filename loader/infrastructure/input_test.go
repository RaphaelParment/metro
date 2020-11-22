package infrastructure

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestFindLineExtensions(t *testing.T) {
	number := "7"
	expected := []string{"7_1", "7_2"}

	files, err := findLineExtensions(number)
	require.NoError(t, err)
	require.Equal(t, expected, files)
}
