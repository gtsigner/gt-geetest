package geetest

import (
    "context"
    "fmt"
    "regexp"
    "testing"
    "time"
    "zhaojunlike/common"
    "zhaojunlike/common/chttp"
    "zhaojunlike/gt-axios/axios"
    "zhaojunlike/gt-axios/base"
    "zhaojunlike/common/logger"
)

func TestGeeSlider(t *testing.T) {
    options := axios.NewOptions()
    options.Proxy = chttp.LocalProxy
    bapi := base.NewBApi(options)
    bapi.UserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.66 Safari/537.36"
    api := NewApiV2(bapi)
    api.vm = NewCallerV8()

    gee := &GeeData{
        GeeBaseData: GeeBaseData{
            ApiServer:  "api-na.geetest.com",
            Gt:         "1e505deed3832c02c96ca5abe70df9ab",
            UA:         bapi.UserAgent,
            Challenge:  "e2ed9bccdc1183a678beb056ffe66ffa",
            Lang:       "en",
            ClientType: "web",
            Time:       common.CreateTimestamp(),
        },
        Static: "",
        Type:   "",
    }
    ctx, _ := context.WithTimeout(context.Background(), time.Minute)
    err := api.ExecuteBypass(ctx, gee)
    if err != nil {
        logger.Error(err)
        return
    }
    logger.Info(gee.Gt, gee.Validate, gee.Challenge)
}
func TestGeeSliderByPass(t *testing.T) {
    options := axios.NewOptions()
    options.Proxy = chttp.LocalProxy
    bapi := base.NewBApi(options)
    bapi.UserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.66 Safari/537.36"
    api := NewApiV2(bapi)
    api.vm = NewCallerV8()

    gee := &GeeData{
        GeeBaseData: GeeBaseData{
            ApiServer:  "api-na.geetest.com",
            Gt:         "1e505deed3832c02c96ca5abe70df9ab",
            UA:         bapi.UserAgent,
            Challenge:  "e2ed9bccdc1183a678beb056ffe66ffa",
            Lang:       "en",
            ClientType: "web",
            Time:       common.CreateTimestamp(),
        },
        Static: "",
        Type:   "",
    }
    ctx, _ := context.WithTimeout(context.Background(), time.Minute)
    err := api.ExecuteBypass(ctx, gee)
    if err != nil {
        logger.Error(err)
        return
    }
    logger.Info(gee.Gt, gee.Validate, gee.Challenge)
}
func TestGeeClick(t *testing.T) {
    options := axios.NewOptions()
    options.Proxy = chttp.LocalProxy
    bapi := base.NewBApi(options)
    bapi.UserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.66 Safari/537.36"
    api := NewApiV2(bapi)
    api.vm = NewCallerV8()

    gee := &GeeData{
        GeeBaseData: GeeBaseData{
            ApiServer:  "api.geetest.com",
            Gt:         "babddc6959f35b45ad8c3947c9a00c5c",
            UA:         bapi.UserAgent,
            Challenge:  "39654bc8bc473f49f51f6fe62df8a2a7",
            Lang:       "zh-hk",
            ClientType: "web",
            Time:       common.CreateTimestamp(),
        },
        Static: "",
        Type:   "",
    }
    ctx, _ := context.WithTimeout(context.Background(), time.Minute)
    err := api.ExecuteBypass(ctx, gee)
    if err != nil {
        logger.Error(err)
        return
    }
    logger.Info(gee.Gt, gee.Validate, gee.Challenge)
}

//测试拿到Vioce
func TestGeeVoice(t *testing.T) {
    options := axios.NewOptions()
    options.Proxy = chttp.LocalProxy
    bapi := base.NewBApi(options)
    bapi.UserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.66 Safari/537.36"
    api := NewApiV2(bapi)
    api.vm = NewCallerV8()

    gee := &GeeData{
        GeeBaseData: GeeBaseData{
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

func TestParseCallbackJSONP(t *testing.T) {
    str := `geetest_1587016719120({"status": "success", "data": {"fullpage": "/static/js/fullpage.8.9.3.js", "geetest": "/static/js/geetest.6.0.9.js", "static_servers": ["static.geetest.com/", "dn-staticdown.qbox.me/"], "maze": "/static/js/maze.1.0.1.js", "slide": "/static/js/slide.7.7.0.js", "type": "fullpage", "aspect_radio": {"beeline": 50, "click": 128, "pencil": 128, "slide": 103, "voice": 128}, "beeline": "/static/js/beeline.1.0.1.js", "click": "/static/js/click.2.8.9.js", "voice": "/static/js/voice.1.2.0.js", "pencil": "/static/js/pencil.1.0.3.js"}})`
    str = JsonpParse("geetest_1587016719120", str)
    data := struct {
    }{}
    err := common.JSONParse(str, &data)
    fmt.Println(err)
}

func TestRe(t *testing.T) {
    html := `
        initGeetest({
        api_server: 'api-na.geetest.com',
        gt: '1e505deed3832c02c96ca5abe70df9ab',
        challenge: 'f8b95137ec18ad43d2dda0c87b997671',
        product: 'float',
        offline:  0 ,
        new_captcha: 1,
        lang: 'zh-cn',
        http: 'https' + '://'
    }, handlerEmbed);
    };
`
    re := regexp.MustCompile(`(?m)gt:\s*'(.*?)'`)

    for i, match := range re.FindAllStringSubmatch(html, -1) {
        fmt.Println(match, "found at index", i)
    }
}
