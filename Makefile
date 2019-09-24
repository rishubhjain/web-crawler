APPNAME = web-crawler
VERSION = 0.0.1-dev

test:
	go test ./... -v -cover

setup:
	GO111MODULE=on go mod vendor

build:
	go build -o ${APPNAME} .

clean:
	rm -f ${APPNAME}

install: build
	sudo install -d /usr/local/bin
	sudo install -c ${APPNAME} /usr/local/bin/${APPNAME}

uninstall:
	sudo rm /usr/local/bin/${APPNAME}
