package v1

import (
    "testing"
)

func Test_V1_0(t *testing.T) {
    if (true != true) {
        t.Error("false positive for true test")
    }
}
