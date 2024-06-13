package handler

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestHandler_jsonError(t *testing.T) {
	msg := "Hello, error!"
	result := jsonError(msg)
	require.Equal(t, []byte(`{"message":"Hello, error!"}`), result)
}
