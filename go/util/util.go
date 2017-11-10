package util

import (
    "bytes"
)

func IsRawbodyArray(ba []byte) bool {
    return bytes.HasPrefix(ba, []byte("[")) && bytes.HasSuffix(ba, []byte("]"))
}

func WrapJsonAsRoot(js []byte) []byte {
    var b bytes.Buffer
    b.Write([]byte(`{"root" : `))
    b.Write(js)
    b.WriteString("}")
    return b.Bytes()
}
