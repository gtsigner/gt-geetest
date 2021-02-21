package geetest

type GeeFirstRegKeyRes struct {
    AesKey  string `json:"aesKey"`
    EAesKey string `json:"eAesKey"` //Rsa加密aesKey
    Ch      string `json:"ch"`
    Gt      string `json:"gt"`
    I       string `json:"i"`
    T       int64  `json:"t"`
    W       string `json:"w"`
}
type GeeReqVerRes struct {
    Lp   []interface{} `json:"lp"`
    Move []interface{} `json:"move"`
    W    string        `json:"w"`
}
type GeeSubmitSliderRes struct {
    Aa           string `json:"aa"`
    ImgLoad      int64  `json:"imgload"`
    PassTime     int64  `json:"passtime"`
    UserResponse string `json:"userresponse"`
}

type IGeeCall interface {
    FirstRegKey(data *GeeData) (*GeeFirstRegKeyRes, error)
    RequestAjaxVer(data *GeeData) (*GeeReqVerRes, error)
    SubmitSlide(data *GeeData) (string, error)
    SubmitSlideSimple(data *GeeData) (*GeeSubmitSliderRes, error)
}
