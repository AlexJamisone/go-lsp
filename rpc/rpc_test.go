package rpc_test

import (
	"educationlsp/rpc"
	"testing"
)
type EncodingExaple struct {
    Testing bool
}
func TestEncode(t *testing.T) {
    expected := "Content-Length: 16\r\n\r\n{\"Testing\":true}"
    actual := rpc.EncodeMessage(EncodingExaple{Testing:true})
    if expected != actual {
        t.Fatalf("Expected: %s, Actual: %s", expected, actual)
    }
}
