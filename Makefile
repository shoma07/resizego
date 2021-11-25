build:
	go build -buildmode=c-shared -o lib/resizeman/resizeman.bundle ext/resizeman/main.go

# fake out clean and install
clean:
install:

.PHONY: build
