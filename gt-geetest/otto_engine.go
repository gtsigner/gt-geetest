package geetest

import (
    "github.com/robertkrimen/otto"
    "sync"
)

type GeeJsEngine struct {
    locker sync.Mutex
    vm     *otto.Otto
    //caller
    _submitSlideSimpleString func(string) string
}

func NewOttoEngine() *GeeJsEngine {
    engine := &GeeJsEngine{}
    engine.vm = otto.New()
    engine._load()
    return engine
}

func (engine *GeeJsEngine) _load() error {
    codes, _ := readFile("ex_code.js")
    _, err := engine.vm.Run(codes)
    if err != nil {
        return err
    }
    engine._submitSlideSimpleString = func(s string) string {
        res, err := engine.vm.Call("submitSlideSimpleString", nil, s)
        if err != nil {
            return ""
        }
        return res.String()
    }
    return nil
}
func (engine *GeeJsEngine) SubmitSlideSimpleString(str string) string {
    engine.locker.Lock()
    defer engine.locker.Unlock()
    return engine._submitSlideSimpleString(str)
}
