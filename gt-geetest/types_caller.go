package geetest

type apiCallerBaseRes struct {
    Message string `json:"message"`
    Success bool   `json:"success"`
}
type apiCallerFirstRegKeyRes struct {
    apiCallerBaseRes
    Data struct {
        W       string `json:"w"`
        AesKey  string `json:"aesKey"`
        EAesKey string `json:"eAesKey"` //Rsa加密aesKey
        I       string `json:"i"`
        Gt      string `json:"gt"`
        Ch      string `json:"ch"`
        T       int64  `json:"t"`
    }
}
type apiCallerReqVerRes struct {
    apiCallerBaseRes
    Data struct {
        W    string        `json:"w"`
        Move []interface{} `json:"move"`
        Lp   []interface{} `json:"lp"`
    }
}
type apiCallerSubmitSlideRes struct {
    apiCallerBaseRes
    Data string
}
type apiCallerSubmitSlideSimpleRes struct {
    apiCallerBaseRes
    Data struct {
        Aa           string `json:"aa"`
        ImgLoad      int64  `json:"imgload"`
        PassTime     int64  `json:"passtime"`
        UserResponse string `json:"userresponse"`
    }
}
