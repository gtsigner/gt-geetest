package geetest

import "errors"

var (
    ErrGetGeeImg         = errors.New("get image error")
    ErrParseGeeImg       = errors.New("parse image error")
    ErrGetOffsetFail     = errors.New("get offset error")
    ErrExecuteError      = errors.New("execute result error")
    ErrSubmitParamsError = errors.New("submit params error")
    ErrNotSupportGeeType = errors.New("not support gee type")
)
