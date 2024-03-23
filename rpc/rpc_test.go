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
func TestDecode(t *testing.T) {
    incomingMessage := "Content-Length: 15\r\n\r\n{\"Method\":\"hi\"}"
    method, content, err := rpc.DecodeMessage([]byte(incomingMessage))
    contenLegth := len(content)
    if err != nil {
        t.Fatal(err)
    }
    if contenLegth != 15 {
        t.Fatalf("Expected: 15, Got: %d", contenLegth)
    }
    if method != "hi" {
        t.Fatalf("Expected: 'hi', Got: %s", method)
    }
}
