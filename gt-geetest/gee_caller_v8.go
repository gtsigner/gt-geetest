package geetest

import (
    "encoding/base64"
    "fmt"
    "github.com/thoas/go-funk"
    "strconv"
    "zhaojunlike/common"
)

var (
    reqVerD    []interface{}
    reqVerDStr = `[
        ['move', [411, 1001], 0],
        ['blur', 349],
        ['move', [83, -66], 44],
        ['move', [4, -1], 33],
        ['move', [10, -8], 17],
        ['move', [20, -13], 16],
        ['move', [43, -38], 17],
        ['move', [66, -58], 16],
        ['move', [652, -423], 19369],
        ['move', [-40, 25], 10],
        ['move', [-67, 37], 18],
        ['move', [-49, 22], 17],
        ['move', [-30, 12], 15],
        ['move', [-12, 4], 17],
        ['move', [-4, 0], 17],
        ['move', [-1, 0], 17],
        ['move', [-1, 0], 24], ['move', [-1, -1], 9], ['move', [-5, -3], 17], ['move', [-7, -4], 16], ['move', [-5, -3], 17],
        ['move', [-8, -10], 18], ['move', [-6, -9], 16], ['move', [-7, -7], 16], ['move', [-4, -9], 17], ['move', [-4, -4], 17],
        ['move', [-3, -3], 16], ['move', [-4, -4], 17], ['move', [-2, -1], 16], ['move', [0, -1], 17], ['move', [-1, -1], 17],
        ['move', [0, -1], 20], ['move', [-1, 0], 13], ['focus', 42], ['down', [0, 0], 4], ['up', [0, 0], 30], ['move', [0, 1], 2313],
        ['move', [0, 5], 15], ['move', [5, 14], 18], ['move', [8, 15], 16], ['move', [9, 22], 17], ['move', [14, 26], 16],
        ['move', [12, 22], 18], ['move', [10, 21], 16], ['move', [7, 19], 15], ['move', [6, 17], 18], ['move', [2, 11], 16],
        ['move', [1, 7], 17], ['move', [0, 3], 16], ['move', [-1, 3], 17], ['move', [-2, 1], 17], ['move', [-1, 3], 16],
        ['move', [0, 3], 17], ['move', [-1, 4], 17], ['move', [-1, 1], 16], ['move', [0, 1], 17], ['move', [-1, 1], 17],
        ['down', [0, 0], 216], ['up', [0, 0], 55],
    ]`
)

func init() {
    _ = common.JSONParse(reqVerDStr, &reqVerD)
}

//使用V8来执行
type GeeCallerV8 struct {
}

func NewCallerV8() *GeeCallerV8 {
    api := &GeeCallerV8{}
    return api
}

//注冊第一個參數
func (api *GeeCallerV8) FirstRegKey(base *GeeData) (*GeeFirstRegKeyRes, error) {
    var data = &geeFirstRegParams{
        Gt:         base.Gt,
        Challenge:  base.Challenge,
        Offline:    false,
        NewCaptcha: 1,
        Product:    "bind",
        Width:      "300px",
        Pure:       1,
        Protocol:   "https://",
        Voice:      "/static/js/voice.1.1.3.js",
        Slide:      "/static/js/slide.7.3.9.js",
        Geetest:    "/static/js/geetest.6.0.9.js",
        Fullpage:   "/static/js/fullpage.8.5.7.js",
        Click:      "/static/js/click.2.5.6.js",
        Type:       "fullpage",
        StaticServers: []string{
            "static.geetest.com/",
            "dn-staticdown.qbox.me/",
        },
        AspectRadio: struct {
            Click int64 `json:"click"`
            Slide int64 `json:"slide"`
            Voice int64 `json:"voice"`
        }{
            Click: 128,
            Voice: 128,
            Slide: 103,
        },
        Pencil: "/static/js/pencil.1.0.1.js",
        Cc:     12,
        Ww:     true,
        I:      "28987!!84086!!CSS1Compat!!246!!-1!!-1!!-1!!-1!!4!!-1!!-1!!-1!!86!!19!!11!!2!!-1!!-1!!-1!!-1!!-1!!-1!!-1!!-1!!37!!2!!-1!!-1!!-1!!0!!0!!0!!0!!1920!!943!!1920!!1050!!zh-CN!!zh-CN,zh!!-1!!1!!24!!" + base.UA + "!!1!!1!!1920!!1080!!1920!!1050!!1!!1!!1!!-1!!Win32!!0!!-8!!" + common.Md5Encrypt(funk.RandomString(5)) + "!!" + common.Md5Encrypt(funk.RandomString(5)) + "!!mhjfbmdgcfjbbpaeojofohoefgiehjai,pepflashplayer.dll,internal-pdf-viewer!!0!!-1!!0!!12!!Arial,ArialBlack,ArialNarrow,Calibri,Cambria,CambriaMath,ComicSansMS,Consolas,Courier,CourierNew,Georgia,Helvetica,Impact,LucidaConsole,LucidaSansUnicode,MicrosoftSansSerif,MSPGothic,MSSansSerif,MSSerif,PalatinoLinotype,SegoePrint,SegoeScript,SegoeUI,SegoeUILight,SegoeUISemibold,SegoeUISymbol,Tahoma,Times,TimesNewRoman,TrebuchetMS,Verdana,Wingdings!!" + strconv.FormatInt(common.CreateTimestamp(), 10) + "!!-1,-1,1,0,0,0,0,16,74,3,4,11,26,443,443,478,659,659,660,-1!!-1!!-1!!249!!109!!28!!206!!21!!false!!false",
    }
    if base.AesKey == "" {
        base.AesKey = base64.StdEncoding.EncodeToString(RandomBytes(16))
    }
    result, eAesKey, err := encryptGeeData(base.AesKey, data, "")
    if err != nil {
        return nil, err
    }
    res := &GeeFirstRegKeyRes{
        Gt:      base.Gt,
        Ch:      base.Challenge,
        W:       result + eAesKey,
        AesKey:  base.AesKey,
        EAesKey: eAesKey,
        I:       data.I,
        T:       base.Time,
    }
    return res, nil
}

func (api *GeeCallerV8) RequestAjaxVer(base *GeeData) (*GeeReqVerRes, error) {
    const gStr = "44937magic data100270magic dataCSS1Compatmagic data187magic data-1magic data-1magic data-1magic data-1magic data4magic data-1magic data-1magic data1magic data44magic data19magic data11magic data5magic data-1magic data-1magic data-1magic data-1magic data-1magic data-1magic data-1magic data-1magic data158magic data1magic data-1magic data-1magic data-1magic data41magic data96magic data0magic data0magic data1977magic data1004magic data1879magic data954magic datazh-CNmagic datazh-CN,zhmagic data-1magic data0.949999988079071magic data24magic dataMozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/63.0.3239.26 Safari/537.36 Core/1.63.6756.400 QQBrowser/10.3.2473.400magic data1magic data1magic data1920magic data1080magic data1920magic data1050magic data1magic data1magic data1magic data-1magic dataWin32magic data0magic data-8magic data5682b327da75e9e40955e76447258302magic data10bef12cfc3912d19616d7f5fba076femagic datainternal-pdf-viewer,mhjfbmdgcfjbbpaeojofohoefgiehjai,internal-nacl-plugin,pepflashplayer.dll,widevinecdmadapter.dllmagic data0magic data-1magic data0magic data12magic dataArial,ArialBlack,ArialNarrow,Calibri,Cambria,CambriaMath,ComicSansMS,Consolas,Courier,CourierNew,Georgia,Helvetica,Impact,LucidaConsole,LucidaSansUnicode,MicrosoftSansSerif,MSGothic,MSPGothic,MSSansSerif,MSSerif,PalatinoLinotype,SegoePrint,SegoeScript,SegoeUI,SegoeUILight,SegoeUISemibold,SegoeUISymbol,Tahoma,Times,TimesNewRoman,TrebuchetMS,Verdana,Wingdingsmagic data' + new Date().getTime() + 'magic data-1,-1,7,0,0,0,0,4,624,215,216,218,300,668,669,859,-1,-1,-1,-1magic data-1magic data-1magic data278magic data33magic data32magic data147magic data24magic datafalsemagic datafalse"
    var data = map[string]interface{}{
        "lang":          base.Lang,
        "type":          "fullpage",
        "tt":            "M:_8Pjp/.(38NAPjp8ON9Up8Pjp8Pjp..*M8AA.(5b((n((b59(59A((55.G.-,(-be(99((b((:-BB/(((,5D55,9.,,((b((((,(B-:(((1(((((-A9A9AEOSo7?IEKE11K)O2K2VVG)*1-2OES/S))lKE-0Q?L-KfMUNd.c?eC3NjK0MbOUM9KD3:)3DC0cFb3SYhNj-1RkI:K)O2K)*2M9NkM9K:OCVlhLM9OLTmG8-1/S-25*IO2OR11Z/FKScDRjRj-1S)O0-Dp)O2M9OEM9L1?)*)*2n?.VK)*2K)*2ZGFj/1S2G5FjKDRjS)*3kc-9/01DH/-3G0OERV1qK3GDS2mGTVaRcO)(@j:l)(X.5M57(),b8(,((,e((8(5,(e5(b555(e,e((nb((n(b(((,n,(q((R(beb(b)BBBBA(,(b5Gn?(J.TC))U--BQM9G))jKGS)*UFg6,3)(j1-C)),9c(9ME/-(*6.b9K/L))Q65-*8)(?bU-)1E-)1?M9-NMb1cM99fB1-,3)(?**(E1(/,M)(?-51)MU(I-*6DOoqI-7b9VT5E-)5P2:6)5?/)(U-)M91b3)(?-1-*M9/,MQ6.2AAU-,9U/,/*(U-)1*-5-,:UY9.cQL1LEW5*f_3O9P/:@mNA*)(?(j/)1)M9(((((,qqM(0qqb((((,qb((1*CI(,b(8e5,5q8,bbb(5(8(q,n((b(((e(,,e58be55((b((((88Y11(-N1)Mb,)(?-)55-*Mj*)(PMM1E1*(?M9-j-*?)(U/)ME(I/*Kb95PbM4)(94*M9/*()-)(E-(-)(E-(M91?/)(U1)MM(?b90)MY-)19-N,)(M-N(?-*/)()O1(919b9-)NU*)M9M9-11)111)MU**(9//()M9(E-(-1-)1fBTW),)OUVeAd(c(E/(bE-5/)MM),b8b)qqqM(8qb",
        "light":         "INPUT_0",
        "s":             common.Md5Encrypt(gStr + "!"),
        "h":             common.Md5Encrypt(gStr + "!!"),
        "hh":            common.Md5Encrypt(gStr + "!!!"),
        "hi":            common.Md5Encrypt(base.RegI),
        "vip_order":     "",
        "ct":            "",
        "ep":            ep(base),
        "captcha_token": "newage",
        "passtime":      common.CreateTimestamp() - base.Time,
    }
    data["rp"] = common.Md5Encrypt(base.Gt + base.Challenge + fmt.Sprintf("%v", data["passtime"]))
    result, _, err := encryptGeeData(base.AesKey, data, base.EAesKey)
    if err != nil {
        return nil, err
    }
    move := []interface{}{"move", 411, 1001, common.CreateTimestamp(), "pointermove"}
    lp := []interface{}{"up", 770, 544, common.CreateTimestamp() + 23247, "pointerup"}
    res := &GeeReqVerRes{W: result, Move: move, Lp: lp}
    return res, nil
}

func (api *GeeCallerV8) SubmitSlide(base *GeeData) (string, error) {
    res, err := api.SubmitSlideSimple(base)
    if err != nil {
        return "", err
    }
    tm := common.CreateTimestamp() - int64(funk.RandomInt(3000, 6000))
    var data = map[string]interface{}{
        "lang":         base.Lang,
        "userresponse": res.UserResponse,
        "passtime":     res.PassTime,
        "imgload":      funk.RandomInt(80, 120),
        "aa":           res.Aa,
        "ep": map[string]interface{}{
            "v":  "8.5.7",
            "f":  common.Md5Encrypt(base.Gt + base.Challenge),
            "te": false,
            "me": true,
            "tm": performance(tm),
        },
    }
    data["rp"] = common.Md5Encrypt(base.Gt + base.Challenge[:32] + fmt.Sprintf("%v", data["passtime"]))
    result, eAesKey, err := encryptGeeData(base.AesKey, data, base.EAesKey)
    if err != nil {
        return "", err
    }
    return result + eAesKey, nil
}

func (api *GeeCallerV8) SubmitSlideSimple(data *GeeData) (*GeeSubmitSliderRes, error) {
    return submitSlideSimple(engine, data)
}
