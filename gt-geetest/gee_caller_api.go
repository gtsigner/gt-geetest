package geetest

import (
    "encoding/json"
    "errors"
    "time"
    "zhaojunlike/common/chttp"
    "zhaojunlike/gt-axios/base"
)

type CallerApi struct {
    base.IBApi
    BaseURL string
}

func NewCallerApi(base base.IBApi) *CallerApi {
    api := &CallerApi{}
    api.IBApi = base
    api.BaseURL = "http://127.0.0.1/engine/gt"
    return api
}

//注册信息然后拿到AES之类的信息
func (api *CallerApi) FirstRegKey(data *GeeData) (*GeeFirstRegKeyRes, error) {
    params := map[string]interface{}{}
    params["lang"] = data.Lang
    params["gt"] = data.Gt
    params["challenge"] = data.Challenge
    params["time"] = data.Time
    params["ua"] = data.UA

    u := api.BaseURL + "/firstRegKey"
    conf := chttp.NewPostConfig(u, params, false)
    conf.Timeout = time.Minute
    conf.ParseFun = func(r *chttp.Res, respStr string) interface{} {
        res := apiCallerFirstRegKeyRes{}
        err := json.Unmarshal(r.Buffer, &res)
        if err != nil {
            return nil
        }
        r.Success = res.Success
        return res
    }
    res := api.Request(conf)
    if res.Success {
        data := res.Data.(apiCallerFirstRegKeyRes)
        regKeyRes := &GeeFirstRegKeyRes{
            W:       data.Data.W,
            AesKey:  data.Data.AesKey,
            EAesKey: data.Data.EAesKey,
            I:       data.Data.I,
            Gt:      data.Data.Gt,
            Ch:      data.Data.Ch,
            T:       data.Data.T,
        }
        return regKeyRes, nil
    }
    return nil, errors.New(res.Message)
}

//请求Ajax的Ver数据拿到一些便宜
func (api *CallerApi) RequestAjaxVer(data *GeeData) (*GeeReqVerRes, error) {
    params := make(map[string]interface{})
    params["lang"] = data.Lang
    params["gt"] = data.Gt
    params["challenge"] = data.Challenge
    params["time"] = data.Time
    params["aesKey"] = data.AesKey
    params["eAesKey"] = data.EAesKey
    params["ua"] = data.UA
    params["i"] = data.RegI
    params["c"] = data.RegC
    params["s"] = data.RegS

    u := api.BaseURL + "/ReqVer"
    conf := chttp.NewPostConfig(u, params, false)
    conf.Timeout = time.Minute

    conf.ParseFun = func(r *chttp.Res, respStr string) interface{} {
        res := apiCallerReqVerRes{}
        err := json.Unmarshal(r.Buffer, &res)
        if err != nil {
            return nil
        }
        r.Success = res.Success
        return res
    }
    res := api.Request(conf)
    if res.Success {
        data := res.Data.(apiCallerReqVerRes)
        regKeyRes := &GeeReqVerRes{
            W:    data.Data.W,
            Move: data.Data.Move,
            Lp:   data.Data.Lp,
        }
        return regKeyRes, nil
    }
    return nil, errors.New(res.Message)
}

//提交的的时候
func (api *CallerApi) SubmitSlide(data *GeeData) (string, error) {
    if data.GeeSliderSubmitData == nil {
        return "", errors.New("slider params empty")
    }
    params := make(map[string]interface{})
    params["lang"] = data.Lang
    params["gt"] = data.Gt
    params["challenge"] = data.Challenge
    params["time"] = data.Time
    params["aesKey"] = data.AesKey
    params["eAesKey"] = data.EAesKey
    params["c"] = data.RegC
    params["s"] = data.RegS
    params["offset"] = data.Offset

    u := api.BaseURL + "/submitSlide"
    conf := chttp.NewPostConfig(u, params, false)
    conf.Timeout = time.Minute
    conf.ParseFun = func(r *chttp.Res, respStr string) interface{} {
        res := apiCallerSubmitSlideRes{}
        err := json.Unmarshal(r.Buffer, &res)
        if err != nil {
            return nil
        }
        r.Success = res.Success
        return res
    }
    res := api.Request(conf)
    if res.Success {
        data := res.Data.(apiCallerSubmitSlideRes)
        return data.Data, nil
    }
    return "", errors.New(res.Message)
}

//获取参数
func (api *CallerApi) SubmitSlideSimple(data *GeeData) (*GeeSubmitSliderRes, error) {
    if data.GeeSliderSubmitData == nil {
        return nil, errors.New("slider params empty")
    }
    params := make(map[string]interface{})
    params["lang"] = data.Lang
    params["gt"] = data.Gt
    params["challenge"] = data.Challenge
    params["time"] = data.Time
    params["c"] = data.RegC
    params["s"] = data.RegS
    params["offset"] = data.Offset

    u := api.BaseURL + "/submitSlideSimple"
    conf := chttp.NewPostConfig(u, params, false)
    conf.Timeout = time.Minute
    conf.ParseFun = func(r *chttp.Res, respStr string) interface{} {
        res := apiCallerSubmitSlideSimpleRes{}
        err := json.Unmarshal(r.Buffer, &res)
        if err != nil {
            return nil
        }
        r.Success = res.Success
        r.Message = res.Message
        return res
    }
    res := api.Request(conf)
    if res.Success {
        data := res.Data.(apiCallerSubmitSlideSimpleRes)
        resp := &GeeSubmitSliderRes{
            Aa:           data.Data.Aa,
            UserResponse: data.Data.UserResponse,
            PassTime:     data.Data.PassTime,
        }
        return resp, nil
    }
    return nil, errors.New(res.Message)
}
