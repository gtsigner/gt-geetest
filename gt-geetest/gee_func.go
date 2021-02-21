package geetest

import (
    "zhaojunlike/common"
)

type urlConfig struct {
    getType string
    getPHP  string
    ajaxPHP string
}
type geeSlider struct {
    FullBg string
    Bg     string
    Offset int64
}
type GeeData struct {
    GeeBaseData
    *GeeSliderSubmitData
    Static   string `json:"static"`
    Type     string `json:"type"`
    Validate string `json:"validate"`
    Score    string `json:"score"`
    SecCode  string `json:"seccode"`
    Value    string `json:"value"`
}

//设置注册的返回结果
func (gee *GeeData) SetFirstRegRes(res *GeeFirstRegKeyRes) {
    gee.AesKey = res.AesKey
    gee.EAesKey = res.EAesKey
    gee.RegI = res.I
    gee.Ch = res.Ch
    gee.Time = res.T
}

//设置第二次返回的参数
func (gee *GeeData) SetReqVerRes(ver *GeeReqVerRes) {
    if gee.GeeSliderSubmitData == nil {
        gee.GeeSliderSubmitData = &GeeSliderSubmitData{}
    }
    gee.GeeSliderSubmitData.Move = ver.Move
    gee.GeeSliderSubmitData.Lp = ver.Lp
}

//设置获取参数的时候的响应
func (gee *GeeData) SetRegParamsResp(data apiGeeResGetNext) {
    gee.RegC = data.C
    gee.RegS = data.S
    gee.Challenge = data.Challenge
}

func ep(base *GeeData) map[string]interface{} {
    var ac = common.Md5Encrypt("eed9d1a9b09fca987c13197c128e4a3e")
    return map[string]interface{}{
        "ts":  common.CreateTimestamp(),
        "v":   "8.5.7",
        "ip":  "192.168.0.109",
        "f":   common.Md5Encrypt(base.Gt + base.Challenge),
        "ci":  nil,
        "de":  false,
        "te":  false,
        "me":  true,
        "ven": "Google Inc.",
        "ren": "ANGLE (NVIDIA GeForce GTX 1080 Direct3D11 vs_5_0 ps_5_0)",
        "ac":  ac,
        "pu":  false,
        "ph":  false,
        "ni":  false,
        "se":  false,
        "fp":  "",
        "lp":  "",
        "em":  em(),
        "tm":  performance(base.Time),
        "by":  -1,
    }
}
func em() map[string]interface{} {
    return map[string]interface{}{
        "ph": 0, "cp": 0, "ek": "11",
        "wd": 0, "nt": 0, "si": 0, "sc": 0,
    }
}
func performance(tm int64) map[string]int64 {
    return map[string]int64{
        "a": tm,
        "b": 0,
        "c": 0,
        "d": 0,
        "e": 0,
        "f": tm,
        "g": tm + 10,
        "h": tm + 30,
        "i": tm + 30,
        "j": tm + 200,
        "k": tm + 60,
        "l": tm + 250,
        "m": tm + 300,
        "n": tm + 304,
        "o": tm + 307,
        "p": tm + 640,
        "q": tm + 640,
        "r": tm + 700,
        "s": tm + 1140,
        "t": tm + 1140,
        "u": tm + 1141,
    }
}
func VerifyEp(base *GeeData) map[string]interface{} {
    return map[string]interface{}{
        "ts": common.CreateTimestamp(),
        "v":  "8.5.7",
        "f":  common.Md5Encrypt(base.Gt + base.Challenge),
        "te": false,
        "me": true,
        "tm": performance(base.Time),
    }
}
