package geetest

type geeFirstRegParams struct {
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
