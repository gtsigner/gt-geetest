package geetest

type ApiGtRes struct {
    Challenge string `json:"challenge"`
    Gt        string `json:"gt"`
    Success   int64  `json:"success"`
}

type ApiGtTypeRes struct {
    Data struct {
        AspectRadio struct {
            Beeline int64 `json:"beeline"`
            Click   int64 `json:"click"`
            Pencil  int64 `json:"pencil"`
            Slide   int64 `json:"slide"`
            Voice   int64 `json:"voice"`
        } `json:"aspect_radio"`
        Beeline       string   `json:"beeline"`
        Click         string   `json:"click"`
        Fullpage      string   `json:"fullpage"`
        Geetest       string   `json:"geetest"`
        Maze          string   `json:"maze"`
        Pencil        string   `json:"pencil"`
        Slide         string   `json:"slide"`
        StaticServers []string `json:"static_servers"`
        Type          string   `json:"type"`
        Voice         string   `json:"voice"`
    } `json:"data"`
    Status string `json:"status"`
}

type ApiParseRes struct {
    Offset  int64   `json:"offset"`
    Message string  `json:"message"`
    T       float64 `json:"t"`
}

type apiGeeResBase struct {
    Error     string `json:"error"`
    ErrorCode string `json:"error_code"`
    UserError string `json:"user_error"`
    Status    string `json:"status"`
}
type apiGeeResFirstRegPhp struct {
    apiGeeResBase
    Data struct {
        APIServer     string   `json:"api_server"`
        C             []int64  `json:"c"`
        Feedback      string   `json:"feedback"`
        S             string   `json:"s"`
        StaticServers []string `json:"static_servers"`
        Theme         string   `json:"theme"`
        ThemeVersion  string   `json:"theme_version"`
    } `json:"data"`
}
type apiGeeResAjaxVer struct {
    apiGeeResBase
    Data struct {
        Result string `json:"result"`
    } `json:"data"`
}

//获取data
type apiGeeResGetNext struct {
    apiGeeResBase
    APIServer  string  `json:"api_server"`
    Benchmark  bool    `json:"benchmark"`
    Bg         string  `json:"bg"`
    C          []int64 `json:"c"`
    Challenge  string  `json:"challenge"`
    Clean      bool    `json:"clean"`
    Feedback   string  `json:"feedback"`
    Fullbg     string  `json:"fullbg"`
    Fullpage   bool    `json:"fullpage"`
    Gt         string  `json:"gt"`
    Height     int64   `json:"height"`
    HideDelay  int64   `json:"hide_delay"`
    HTTPS      bool    `json:"https"`
    I18nLabels struct {
        Cancel       string `json:"cancel"`
        Close        string `json:"close"`
        Error        string `json:"error"`
        Fail         string `json:"fail"`
        Feedback     string `json:"feedback"`
        Forbidden    string `json:"forbidden"`
        Loading      string `json:"loading"`
        Logo         string `json:"logo"`
        ReadReversed bool   `json:"read_reversed"`
        Refresh      string `json:"refresh"`
        Slide        string `json:"slide"`
        Success      string `json:"success"`
        Tip          string `json:"tip"`
        Voice        string `json:"voice"`
    } `json:"i18n_labels"`
    ID            string   `json:"id"`
    Link          string   `json:"link"`
    Logo          bool     `json:"logo"`
    Mobile        bool     `json:"mobile"`
    Product       string   `json:"product"`
    S             string   `json:"s"`
    ShowDelay     int64    `json:"show_delay"`
    Slice         string   `json:"slice"`
    So            int64    `json:"so"`
    StaticServers []string `json:"static_servers"`
    Template      string   `json:"template"`
    Theme         string   `json:"theme"`
    ThemeVersion  string   `json:"theme_version"`
    Type          string   `json:"type"`
    Version       string   `json:"version"`
    Width         string   `json:"width"`
    Xpos          int64    `json:"xpos"`
    Ypos          int64    `json:"ypos"`
}

type apiGeeResVerifySlider struct {
    apiGeeResBase
    Message  string `json:"message"`
    Score    string `json:"score"`
    Success  int64  `json:"success"`
    Validate string `json:"validate"`
}
