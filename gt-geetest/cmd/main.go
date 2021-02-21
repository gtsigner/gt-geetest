package main

import (
    "fmt"
    "zhaojunlike/common"
    "zhaojunlike/common/chttp"
    "zhaojunlike/common/logger"
    "zhaojunlike/gt-axios/axios"
    "zhaojunlike/gt-axios/base"
    geetest "zhaojunlike/gt-geetest"
)

func main() {
    options := axios.NewOptions()
    options.Proxy = chttp.LocalProxy
    bapi := base.NewBApi(options)
    bapi.UserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.66 Safari/537.36"
    api := geetest.NewApiV2(bapi)

    gee := &geetest.GeeData{
        GeeBaseData: geetest.GeeBaseData{
            ApiServer:  "api-na.geetest.com",
            Gt:         "1e505deed3832c02c96ca5abe70df9ab",
            UA:         bapi.UserAgent,
            Challenge:  "cf00b731f1834041ec45a1f0c79f2480",
            Lang:       "en-us",
            ClientType: "web",
            Time:       common.CreateTimestamp(),
        },
        Static: "",
        Type:   "voice",
    }
    res := api.GetVoice(gee)
    logger.Info(res.RespStr)
    logger.Info("Input Validate:")
    {
        _, _ = fmt.Scanln(&gee.Value)
    }
    err := api.CheckVoice(gee)
    if err != nil {
        logger.Error(err)
        return
    }
    logger.Info(gee.Gt, gee.Validate, gee.Challenge)
}
