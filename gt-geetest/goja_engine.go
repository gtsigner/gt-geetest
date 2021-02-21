package geetest

import (
    "github.com/dop251/goja"
    "sync"
)

//引擎需要并发
type GoJaEngine struct {
    locker                   sync.Mutex
    vm                       *goja.Runtime
    _submitSlideSimpleString func(string) string
}

func NewGoJaEngine() *GoJaEngine {
    engine := &GoJaEngine{}
    engine.vm = goja.New()
    engine._load()
    return engine
}
func (engine *GoJaEngine) _load() error {
    codes, _ := readFile("ex_code.js")
    _, err := engine.vm.RunString(string(codes))
    if err != nil {
        return err
    }
    //加载函数列表
    var fn func(string) string
    err = engine.vm.ExportTo(engine.vm.Get("submitSlideSimpleString"), &fn)
    if err != nil {
        return err
    }
    engine._submitSlideSimpleString = fn
    return nil
}
func (engine *GoJaEngine) SubmitSlideSimpleString(str string) string {
    engine.locker.Lock()
    res := engine._submitSlideSimpleString(str)
    engine.locker.Unlock()
    return res
}
