package geetest

import (
    "testing"
    "zhaojunlike/common/logger"
)

func TestJsonpParse(t *testing.T) {
    bytes, err := readFile("ex_code.js")
    logger.Info(string(bytes), err)
}
