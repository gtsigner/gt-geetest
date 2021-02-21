module zhaojunlike/gt-geetest

go 1.13

require (
	github.com/disintegration/imaging v1.6.2
	github.com/dlclark/regexp2 v1.4.0 // indirect
	github.com/dop251/goja v0.0.0-20201107160812-7545ac6de48a
	github.com/go-sourcemap/sourcemap v2.1.3+incompatible // indirect
	github.com/rakyll/statik v0.1.7
	github.com/robertkrimen/otto v0.0.0-20200922221731-ef014fd054ac
	github.com/thoas/go-funk v0.7.0
	gopkg.in/sourcemap.v1 v1.0.5 // indirect
	zhaojunlike/adyen v1.0.0
	zhaojunlike/common v1.0.0
	zhaojunlike/gt-axios v1.0.0
	zhaojunlike/gt-cmap v1.0.0
)

replace zhaojunlike/gt-axios => ../../zhaojunlike/gt-axios

replace zhaojunlike/common => ../../zhaojunlike/common

replace zhaojunlike/adyen => ../../zhaojunlike/adyen

replace zhaojunlike/gt-cmap => ../../zhaojunlike/gt-cmap
