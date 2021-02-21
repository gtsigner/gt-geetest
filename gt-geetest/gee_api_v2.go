package geetest

import (
    "bytes"
    "context"
    "errors"
    "fmt"
    "github.com/disintegration/imaging"
    "github.com/thoas/go-funk"
    "image"
    "net/url"
    "time"
    "zhaojunlike/common"
    "zhaojunlike/common/chttp"
    "zhaojunlike/gt-axios/base"
)

var (
    supportsGeeTypes = []string{
        "slide",
        //"click",
    }
)

type GeeApiV2 struct {
    base.IBApi
    Debug  bool
    vm     IGeeCall //适配器
    config GeeConfig
}

func NewApiV2(base base.IBApi) *GeeApiV2 {
    api := &GeeApiV2{
        Debug:  false,
        config: GeeConfig{ApiURL: "https://api.geetest.com"},
    }
    api.Init(base)
    return api
}

func (api *GeeApiV2) SetVM(base base.IBApi) {
    api.vm = NewCallerV8()
}
func (api *GeeApiV2) Init(base base.IBApi) {
    api.IBApi = base
}
func (api *GeeApiV2) headers() map[string]string {
    return map[string]string{
        "x-requested-with": "XMLHttpRequest",
        "sec-fetch-dest":   "empty",
        "sec-fetch-mode":   "cors",
        "sec-fetch-site":   "same-origin",
        "user-agent":       api.GetUserAgent(),
    }
}

func (api *GeeApiV2) ExecuteBypass(context context.Context, gee *GeeData) error {
    res := api.GetPhpRegData(gee)
    if !res.Success {
        return errors.New(res.Message)
    }
    php := res.Data.(apiGeeResFirstRegPhp)
    gee.RegC = php.Data.C
    gee.RegS = php.Data.S

    select {
    case <-context.Done():
        return context.Err()
    default:
        break
    }

    {
        res := api.GTReqVerType(gee)
        if !res.Success {
            return errors.New(res.Message)
        }
        data := res.Data.(apiGeeResAjaxVer)
        gee.Type = data.Data.Result
    }

    //判断是否支持了
    if !funk.ContainsString(supportsGeeTypes, gee.Type) {
        return errors.New("not support gee:" + gee.Type)
    }

    select {
    case <-context.Done():
        return context.Err()
    default:
        break
    }

    //开始获取图片
    {
        res := api.GetRegParams(gee)
        if !res.Success {
            return errors.New(res.Message)
        }
        data := res.Data.(apiGeeResGetNext)
        gee.SetRegParamsResp(data)
        off, err := api.getSliderOffset(data.Bg, data.Fullbg)
        if err != nil {
            return err
        }
        gee.GeeSliderSubmitData.Offset = off - 3

        {
            res := api.VerifySlider(gee)
            if !res.Success {
                return errors.New(res.Message)
            }
            data := res.Data.(apiGeeResVerifySlider)
            gee.Validate = data.Validate
            gee.Score = data.Score
            gee.SecCode = data.Validate + "|jordan"
        }
    }
    return nil
}

func (api *GeeApiV2) ExecuteBypassV1(context context.Context, gee *GeeData) error {
    res := api.GetPhpRegData(gee)
    if !res.Success {
        return errors.New(res.Message)
    }
    php := res.Data.(apiGeeResFirstRegPhp)
    gee.RegC = php.Data.C
    gee.RegS = php.Data.S

    select {
    case <-context.Done():
        return context.Err()
    default:
        break
    }

    {
        res := api.GTReqVerType(gee)
        if !res.Success {
            return errors.New(res.Message)
        }
        data := res.Data.(apiGeeResAjaxVer)
        gee.Type = data.Data.Result
    }

    //判断是否支持了
    if !funk.ContainsString(supportsGeeTypes, gee.Type) {
        return errors.New("not support gee:" + gee.Type)
    }

    select {
    case <-context.Done():
        return context.Err()
    default:
        break
    }

    //开始获取图片
    {
        res := api.GetRegParams(gee)
        if !res.Success {
            return errors.New(res.Message)
        }
        data := res.Data.(apiGeeResGetNext)
        gee.SetRegParamsResp(data)
        off, err := api.getSliderOffset(data.Bg, data.Fullbg)
        if err != nil {
            return err
        }
        gee.GeeSliderSubmitData.Offset = off - 3

        {
            res := api.VerifySlider(gee)
            if !res.Success {
                return errors.New(res.Message)
            }
            data := res.Data.(apiGeeResVerifySlider)
            gee.Validate = data.Validate
            gee.Score = data.Score
            gee.SecCode = data.Validate + "|jordan"
        }
    }
    return nil
}

func (api *GeeApiV2) GetVoice(data *GeeData) *chttp.Res {
    params := &url.Values{}
    jsonp := api._callbackStr()
    params.Add("gt", data.Gt)
    params.Add("challenge", data.Challenge)
    params.Add("lang", data.Lang)
    params.Add("type", data.Type)

    //44136dfbe144dd3bef1924211345816f
    //1ba8273459fffcf914355a8741317d4cff44b8ca

    params.Add("callback", jsonp)
    u := "https://" + data.ApiServer + "/get.php?" + params.Encode()
    conf := chttp.NewConfig(u)
    conf.Headers = api.headers()
    conf.ParseFun = func(r *chttp.Res, respStr string) interface{} {
        r.RespStr = JsonpParse(jsonp, respStr)
        res := apiGeeResFirstRegPhp{}
        err := common.JSONParse(r.RespStr, &res)
        if err != nil {
            r.Message = "parse json error"
            return nil
        }
        r.Success = res.Status == "success"
        r.Message = res.Error
        return res
    }
    return api.Request(conf)
}
func (api *GeeApiV2) CheckVoice(data *GeeData) *chttp.Res {
    jsonp := api._callbackStr()
    params := &url.Values{}
    params.Add("gt", data.Gt)
    params.Add("challenge", data.Challenge)
    params.Add("lang", data.Lang)
    params.Add("a", data.Value)
    params.Add("type", data.Type)
    params.Add("callback", jsonp)

    u := "https://" + data.ApiServer + "/ajax.php?" + params.Encode()
    conf := chttp.NewConfig(u)
    conf.Headers = api.headers()
    conf.ParseFun = func(r *chttp.Res, respStr string) interface{} {
        r.RespStr = JsonpParse(jsonp, respStr)
        res := apiGeeResVerifySlider{}
        err := common.JSONParse(r.RespStr, &res)
        if err != nil {
            return nil
        }
        r.Success = res.Message == "success"
        return res
    }
    return api.Request(conf)
}

//通过这个进行注册
func (api *GeeApiV2) GetPhpRegData(data *GeeData) *chttp.Res {
    params := &url.Values{}
    reg, err := api.vm.FirstRegKey(data)
    if err != nil {
        return &chttp.Res{Code: "GEE-ERROR", Message: "Engine RegKey Error"}
    }
    data.SetFirstRegRes(reg)

    jsonp := api._callbackStr()

    params.Add("gt", data.Gt)
    params.Add("challenge", data.Challenge)
    params.Add("lang", data.Lang)
    params.Add("w", reg.W)
    params.Add("pt", "0")
    params.Add("client_type", data.ClientType)

    params.Add("callback", jsonp)
    u := "https://" + data.ApiServer + "/get.php?" + params.Encode()
    conf := chttp.NewConfig(u)
    conf.Headers = api.headers()
    conf.ParseFun = func(r *chttp.Res, respStr string) interface{} {
        r.RespStr = JsonpParse(jsonp, respStr)
        res := apiGeeResFirstRegPhp{}
        err := common.JSONParse(r.RespStr, &res)
        if err != nil {
            r.Message = "parse json error"
            return nil
        }
        r.Success = res.Status == "success"
        r.Message = res.Error
        return res
    }
    return api.Request(conf)
}

//获取注册类型
func (api *GeeApiV2) GTReqVerType(data *GeeData) *chttp.Res {
    //请求nodeapi 获取ReqVer执行结果
    reqVer, err := api.vm.RequestAjaxVer(data)
    if err != nil {
        return &chttp.Res{Code: "GEE-ERROR", Message: "ReqVer Error"}
    }
    data.SetReqVerRes(reqVer)

    //获取类型
    params := url.Values{}
    jsonp := api._callbackStr()
    params.Add("gt", data.Gt)
    params.Add("challenge", data.Challenge)
    params.Add("lang", data.Lang)
    params.Add("pt", "0")
    params.Add("client_type", data.ClientType)
    params.Add("w", reqVer.W)
    params.Add("callback", jsonp)
    u := "https://" + data.ApiServer + "/ajax.php?" + params.Encode()

    conf := chttp.NewConfig(u)
    conf.Headers = api.headers()
    conf.ParseFun = func(r *chttp.Res, respStr string) interface{} {
        res := apiGeeResAjaxVer{}
        r.RespStr = JsonpParse(jsonp, respStr)
        err := common.JSONParse(r.RespStr, &res)
        if err != nil {
            r.Message = "parse json error"
            return nil
        }
        r.Success = res.Status == "success"
        r.Message = res.Error
        return res
    }
    return api.Request(conf)
}

//# 获取注册参数
func (api *GeeApiV2) GetRegParams(data *GeeData) *chttp.Res {
    jsonp := api._callbackStr()
    params := url.Values{}
    params.Add("is_next", "true")
    params.Add("type", data.Type)
    params.Add("gt", data.Gt)
    params.Add("challenge", data.Challenge)
    params.Add("lang", data.Lang)
    params.Add("https", "true")
    params.Add("isPC", "true")
    params.Add("bg_color", "transparent")
    params.Add("width", "260px")
    params.Add("area", ".treasure-captcha-gt")
    params.Add("product", "custom")
    params.Add("api_server", data.ApiServer)
    params.Add("protocol", "https://")
    params.Add("callback", jsonp)

    u := "https://" + data.ApiServer + "/get.php?" + params.Encode()
    conf := chttp.NewConfig(u)
    conf.Headers = api.headers()
    conf.ParseFun = func(r *chttp.Res, respStr string) interface{} {
        r.RespStr = JsonpParse(jsonp, respStr)
        res := apiGeeResGetNext{}
        err := common.JSONParse(r.RespStr, &res)
        if err != nil {
            r.Message = "parse json error"
            return nil
        }
        r.Success = res.Status == ""
        r.Message = res.Error
        return res
    }
    return api.Request(conf)
}

//获取图片的偏移量
func (api *GeeApiV2) getSliderOffset(bgPath string, bgfPath string) (int, error) {
    baseStatic := "https://static.geetest.com/"
    var rbg image.Image
    var rbgf image.Image
    {
        //保存FullBg
        conf := chttp.NewConfig(baseStatic + bgPath)
        conf.Timeout = 10 * time.Second
        r := api.Request(conf)
        if !r.Ok {
            return 0, ErrGetGeeImg
        }
        //解析背景图失败
        reader := bytes.NewBuffer(r.Buffer)
        bg, err := imaging.Decode(reader)
        if err != nil {
            return 0, ErrParseGeeImg
        }
        rbg = bg
    }

    //获取背景图
    {
        conf := chttp.NewConfig(baseStatic + bgfPath)
        conf.Timeout = 10 * time.Second
        r := api.Request(conf)
        if !r.Ok {
            return 0, ErrGetGeeImg
        }
        //设置一个背景图
        reader := bytes.NewBuffer(r.Buffer)
        bgf, err := imaging.Decode(reader)
        if err != nil {
            return 0, ErrParseGeeImg
        }
        rbgf = bgf
    }
    return geeOffsetV2(rbg, rbgf)
}

//提交识别
func (api *GeeApiV2) VerifySlider(data *GeeData) *chttp.Res {
    jsonp := api._callbackStr()
    w, err := api.vm.SubmitSlide(data)
    if err != nil {
        return &chttp.Res{Code: "ERROR", Message: "Submit Slider Error"}
    }

    params := &url.Values{}
    params.Add("gt", data.Gt)
    params.Add("challenge", data.Challenge)
    params.Add("lang", data.Lang)
    params.Set("client_type", data.ClientType)
    params.Add("w", w)
    params.Add("pt", "0")
    params.Add("callback", jsonp)

    u := "https://" + data.ApiServer + "/ajax.php?" + params.Encode()
    conf := chttp.NewConfig(u)
    conf.Headers = api.headers()
    conf.ParseFun = func(r *chttp.Res, respStr string) interface{} {
        r.RespStr = JsonpParse(jsonp, respStr)
        res := apiGeeResVerifySlider{}
        err := common.JSONParse(r.RespStr, &res)
        if err != nil {
            return nil
        }
        r.Success = res.Message == "success"
        return res
    }
    return api.Request(conf)
}

//获取配置
func (api *GeeApiV2) GetType(data *GeeData) *chttp.Res {
    params := url.Values{}
    jsonp := api._callbackStr()

    params.Add("gt", data.Gt)
    params.Add("callback", jsonp)
    u := api.config.ApiURL + "/gettype.php?" + params.Encode()
    conf := chttp.NewConfig(u)
    conf.ParseFun = func(r *chttp.Res, respStr string) interface{} {
        res := ApiGtTypeRes{}
        respStr = JsonpParse(jsonp, respStr)
        err := common.JSONParse(respStr, res)
        if err != nil {
            return nil
        }
        r.Success = true
        return res
    }
    res := api.Request(conf)
    if res.Success {

    }
    return res
}

func (api *GeeApiV2) _callbackStr() string {
    return fmt.Sprintf("geetest_%v", common.CreateTimestamp())
}
