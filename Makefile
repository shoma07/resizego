build:
	go build -buildmode=c-shared -o lib/resizego/resizego.bundle ext/resizego/main.go

# fake out clean and install
clean:
install:

.PHONY: build
