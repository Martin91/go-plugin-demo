all: build_en build_ch build_main

build_en:
	go build -buildmode=plugin -o en.so plugin/en/greeter.go

build_ch:
	go build -buildmode=plugin -o ch.so plugin/ch/greeter.go

build_main:
	go build -o greeter greeter.go
