package geetest

import (
    "testing"
    "zhaojunlike/common/logger"
)

func TestCallerApi_FirstRegKey(t *testing.T) {
    data := map[string]string{}
    enc, eaes, err := encryptGeeData("1e352aba132db671", data, "")
    logger.Info(enc, eaes, err)
}
