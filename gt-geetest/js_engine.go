package geetest

import (
    "zhaojunlike/common"
    "zhaojunlike/common/logger"
)

type IJsEngine interface {
    SubmitSlideSimpleString(str string) string
}

var engine IJsEngine

func init() {
    engine = NewGoJaEngine()
}

func submitSlideSimple(engine IJsEngine, data *GeeData) (res *GeeSubmitSliderRes, err error) {
    if data == nil || data.GeeSliderSubmitData == nil {
        return nil, ErrSubmitParamsError
    }
    //异常了
    defer logger.SafeGoRecoverV2(true, func(interface{}) {
        if res == nil {
            err = ErrExecuteError
        }
    })
    params := make(map[string]interface{})
    params["lang"] = data.Lang
    params["gt"] = data.Gt
    params["challenge"] = data.Challenge
    params["time"] = data.Time
    params["c"] = data.RegC
    params["s"] = data.RegS
    params["offset"] = data.Offset

    paramsStr, err := common.JSONStringify(params)
    if err != nil {
        return nil, err
    }
    resStr := engine.SubmitSlideSimpleString(paramsStr)
    res = &GeeSubmitSliderRes{}
    err = common.JSONParse(resStr, res)
    return res, err
}
