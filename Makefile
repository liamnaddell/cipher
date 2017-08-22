DEP=$(GOPATH)/bin/dep
GO=go
FINAL=/usr/local/bin
VERSION=0.0.1-beta
LDFLAGS=-ldflags "-X main.Version=$(VERSION)"

all:
	# Dep is golang's official dependancy management tool, the /vendor direcotry(created by dep ensure) can act as another GOPATH source directory
	$(DEP) ensure
	cd cmd/cipher; $(GO) build $(LDFLAGS)
	mv cmd/cipher/cipher .
	

install:
	cp cipher $(FINAL)

clean:
	rm -f cipher
	rm -rf vendor/

uninstall:
	rm $(FINAL)/cipher
