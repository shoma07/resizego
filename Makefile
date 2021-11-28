build:
	go build -buildmode=c-shared -o lib/resizego/resizego.so ext/resizego/main.go

# fake out clean and install
clean:
install:

.PHONY: build
