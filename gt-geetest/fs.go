package geetest

import (
    "errors"
    "github.com/rakyll/statik/fs"
    "io/ioutil"
    "net/http"
    "path"
    _ "zhaojunlike/gt-geetest/assets/statik" //静态资源
)

var (
    errNeedLoad = errors.New("you need load filesystem first")
    FileSystem  http.FileSystem
)

func init() {
    err := load()
    if err != nil {
        panic(err)
    }
}

func load() (err error) {
    FileSystem, err = fs.NewWithNamespace("gtbase")
    return err
}

func readFile(ps string) (data []byte, err error) {
    f, err := FileSystem.Open(path.Join("/", ps))
    if err != nil {
        return data, err
    }
    defer f.Close()
    buf, err := ioutil.ReadAll(f)
    if err != nil {
        return data, err
    }
    return buf, err
}
