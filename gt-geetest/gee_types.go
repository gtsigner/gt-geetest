package geetest

type apiGeeReqSliderPhpRegParam struct {
    AspectRadio struct {
        Click int64 `json:"click"`
        Slide int64 `json:"slide"`
        Voice int64 `json:"voice"`
    } `json:"aspect_radio"`
    Cc            int64    `json:"cc"`
    Challenge     string   `json:"challenge"`
    Click         string   `json:"click"`
    Fullpage      string   `json:"fullpage"`
    Geetest       string   `json:"geetest"`
    Gt            string   `json:"gt"`
    I             string   `json:"i"`
    NewCaptcha    int64    `json:"new_captcha"`
    Offline       bool     `json:"offline"`
    Pencil        string   `json:"pencil"`
    Product       string   `json:"product"`
    Protocol      string   `json:"protocol"`
    Pure          int64    `json:"pure"`
    Slide         string   `json:"slide"`
    StaticServers []string `json:"static_servers"`
    Type          string   `json:"type"`
    Voice         string   `json:"voice"`
    Width         string   `json:"width"`
    Ww            bool     `json:"ww"`
}

//基础数据
type GeeBaseData struct {
    ApiServer  string  `json:"api_server"`
    Gt         string  `json:"gt"`
    UA         string  `json:"ua"`
    Challenge  string  `json:"challenge"`
    Lang       string  `json:"lang"`
    AesKey     string  `json:"aes_key"`
    EAesKey    string  `json:"eAesKey"`
    Time       int64   `json:"time"`
    RegI       string  `json:"reg_i"`
    RegC       []int64 `json:"reg_c"`
    RegS       string  `json:"reg_s"`
    Ch         string  `json:"ch"`
    ClientType string  `json:"client_type"`
}
type GeeSliderSubmitData struct {
    Move   interface{} `json:"move"`
    Offset int         `json:"offset"`
    Lp     []interface{}
}

type GeeConfig struct {
    ApiURL string `json:"api_url"`
}
