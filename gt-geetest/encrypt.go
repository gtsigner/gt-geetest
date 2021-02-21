package geetest

import (
    "zhaojunlike/adyen"
    "zhaojunlike/common"
)

const geeRsaKey = `00C1E3934D1614465B33053E7F48EE4EC87B14B95EF88947713D25EECBFF7E74C7977D02DC1D9451F79DD5D1C10C29ACB6A9B4D6FB7D0A0279B6719E1772565F09AF627715919221AEF91899CAE08C0D686D748B20A3603BE2318CA6BC2B59706592A9219D0BF05C9F65023A21D2330807252AE0066D59CEEFA5F2748EA80BAB81`

var (
    adRsa = adyen.NewRsa()
)

func init() {
    _ = adRsa.SetPublicKey(geeRsaKey, 65537)
}

func encryptGeeData(aesKey string, data interface{}, eAesKey string) (string, string, error) {
    bk := []byte(aesKey)
    str, err := common.JSONStringify(data)
    if err != nil {
        return "", "", err
    }
    res := common.AesCBCPKCS5PaddingEncrypt([]byte(str), bk, []byte("0000000000000000"))
    if eAesKey == "" {
        eAesKey, _ = adRsa.Encrypt(aesKey, "")
    }
    return base64Encode(res), eAesKey, nil
}
