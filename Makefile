BINARY=simplepush
OS=linux openbsd darwin windows

VERSION=`git describe --tags --always --dirty`

LDFLAGS=-ldflags "-X main.Version=${VERSION}"

all: test build

get-deps:

build: get-deps
	cd cmd/simplepush && go build ${LDFLAGS} -o ../../${BINARY}

test: get-deps
	go test

install:
	cd cmd/simplepush && go install ${LDFLAGS}

release:
	cd cmd/simplepush && \
	for os in ${OS}; do \
		GOOS=$$os GOARCH=amd64 go build ${LDFLAGS} -o ../../${BINARY}-$$os-amd64; \
	done
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 cd cmd/simplepush && go build ${LDFLAGS} -o ../../${BINARY}-alpine-amd64

doc: build

clean:
	for os in ${OS}; do \
		f=${BINARY}-$$os-amd64; \
		if [ -f $$f ] ; then rm $$f ; fi; \
	done

	rm -f ${BINARY} ${BINARY}-alpine-amd64

distclean: clean
	rm -f ${BINARY}.1

.PHONY: clean install
