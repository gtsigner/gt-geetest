package geetest

import (
    "sync"
    "testing"
    "zhaojunlike/common"
    "zhaojunlike/common/logger"
)

func TestGeeCallerV8_RequestAjaxVer(t *testing.T) {
    //submitSlideSimple()
    data := &GeeData{}
    _ = common.JSONParse(`{
	"c": [12, 58, 98, 36, 43, 95, 62, 15, 12],
	"challenge": "398cdca8b3984c28299dd82547f00f7138",
	"gt": "1e505deed3832c02c96ca5abe70df9ab",
	"lang": "en",
	"offset": 118,
	"s": "694b3870",
	"time": 1606900305371
}`, data)
    res, err := submitSlideSimple(NewOttoEngine(), data)
    logger.Info(res, err)
}
func BenchmarkGeeJsEngine_Load(b *testing.B) {
    data := &GeeData{}
    _ = common.JSONParse(`{
	"c": [12, 58, 98, 36, 43, 95, 62, 15, 12],
	"challenge": "398cdca8b3984c28299dd82547f00f7138",
	"gt": "1e505deed3832c02c96ca5abe70df9ab",
	"lang": "en",
	"offset": 118,
	"s": "694b3870",
	"time": 1606900305371
}`, data)

    //BenchmarkGeeJsEngine_Load-16    	      32	  36594522 ns/op
    //BenchmarkGeeJsEngine_Load-16    	      32	  36188491 ns/op
    //BenchmarkGeeJsEngine_Load-16    	       9	 115222167 ns/op
    //BenchmarkGeeJsEngine_Load-16    	       9	 115555244 ns/op
    for i := 0; i < 1000; i++ {
        res, err := submitSlideSimple(engine, data)
        logger.Info(res, err)
    }
}

func BenchmarkGeeJsEngine_GoEngine1(b *testing.B) {
    wg := sync.WaitGroup{}

    for i := 0; i < b.N; i++ {
        wg.Add(1)
        logger.CreateSafeGo(func() {
            data := &GeeData{}
            _ = common.JSONParse(`{
	"c": [12, 58, 98, 36, 43, 95, 62, 15, 12],
	"challenge": "398cdca8b3984c28299dd82547f00f7138",
	"gt": "1e505deed3832c02c96ca5abe70df9ab",
	"lang": "en",
	"offset": 118,
	"s": "694b3870",
	"time": 1606900305371
}`, data)
            res, err := submitSlideSimple(engine, data)
            logger.Info(res, err)
            wg.Done()
        }, nil)
    }

    wg.Wait()

    //BenchmarkGeeJsEngine_Load-16    	      32	  36594522 ns/op
    //BenchmarkGeeJsEngine_Load-16    	      32	  36188491 ns/op
    //BenchmarkGeeJsEngine_Load-16    	       9	 115222167 ns/op
    //BenchmarkGeeJsEngine_Load-16    	       9	 115555244 ns/op

}

//并发有问题
func BenchmarkGeeJsEngine_LoadV2(b *testing.B) {
    wg := sync.WaitGroup{}

    for i := 0; i < b.N; i++ {
        wg.Add(1)
        logger.SafeGoWithParams(func(i int) {
            data := &GeeData{}
            _ = common.JSONParse(`{
	"c": [12, 58, 98, 36, 43, 95, 62, 15, 12],
	"challenge": "398cdca8b3984c28299dd82547f00f7138",
	"gt": "1e505deed3832c02c96ca5abe70df9ab",
	"lang": "en",
	"offset": 118,
	"s": "694b3870",
	"time": 1606900305371
}`, data)
            res, err := submitSlideSimple(engine, data)
            logger.Info(i, res, err)
            wg.Done()
        }, nil, i)
    }

    wg.Wait()

    //BenchmarkGeeJsEngine_Load-16    	      32	  36594522 ns/op
    //BenchmarkGeeJsEngine_Load-16    	      32	  36188491 ns/op
    //BenchmarkGeeJsEngine_Load-16    	       9	 115222167 ns/op
    //BenchmarkGeeJsEngine_Load-16    	       9	 115555244 ns/op

}
