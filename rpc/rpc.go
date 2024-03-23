package rpc

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
)

func EncodeMessage(msg any) string {
    content, err := json.Marshal(msg)
    if err !=nil {
        panic(err)
    }
    return fmt.Sprintf("Content-Length: %d\r\n\r\n%s", len(content), content)
}
func DecodeMessage(msg []byte) (int, error) {
    header, content, found := bytes.Cut(msg, []byte{'\r', '\n', '\r', '\n'})
    if !found {
        return 0, errors.New("Did not find separator")
    }
    contentLenghByte := header[len("Content-Length: "):]
    contentLengh, err := strconv.Atoi(string(contentLenghByte))
    if err != nil {
        return 0, err
    }
    _ = content
    return contentLengh, nil
}
