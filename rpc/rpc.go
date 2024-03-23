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
type BaseMessage struct {
    Method string `json:"method"`
}
func DecodeMessage(msg []byte) (string, []byte, error) {
    header, content, found := bytes.Cut(msg, []byte{'\r', '\n', '\r', '\n'})
    if !found {
        return "", nil, errors.New("Did not find separator")
    }
    contentLenghByte := header[len("Content-Length: "):]
    contentLengh, err := strconv.Atoi(string(contentLenghByte))
    if err != nil {
        return "", nil, err
    }
    _ = content
    var baseMessage BaseMessage
    if err := json.Unmarshal(content[:contentLengh], &baseMessage); err != nil {
        return "", nil, err
    }
    return baseMessage.Method, content[:contentLengh], nil
}
func Split(data []byte, _ bool) (advance int, token []byte, err error) {
    header, content, found := bytes.Cut(data, []byte{'\r', '\n', '\r', '\n'})
    if !found {
        return 0, nil, nil
    }
    contentLenghByte := header[len("Content-Length: "):]
    contentLengh, err := strconv.Atoi(string(contentLenghByte))
    if err != nil {
        return 0, nil, err
    }
    if len(content) < contentLengh {
        return 0, nil, nil
    }
    totalLength := len(header) + 4 + contentLengh
    return  totalLength, data[:totalLength], nil
}
