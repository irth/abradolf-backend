.PHONY: build run

build:
	go build github.com/irth/abradolf-backend/cmd/abradolf

run: build
	./abradolf

clean: 
	rm abradolf
