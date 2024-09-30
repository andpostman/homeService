APP?=simpleservice
PORT?=8080
PROJECT?=https://github.com/andpostman/homeService/tree/dev

RELEASE?=0.0.1
COMMIT?=$(shell git rev-parse --short HEAD)
BUILD_TIME?=$(shell date -u '+%Y-%m-%d_%H:%M:%S')


clean:
	rm -f ${APP}
build: clean
	go build \
		-ldflags "-s -w -X ${PROJECT}/version.BuildTime=${BUILD_TIME} \
		-X ${PROJECT}/version.Commit=${COMMIT} -X ${PROJECT}/version.Release=${RELEASE}" \
		-o ${APP}
run: build
	PORT=${PORT} ./${APP}
test:
	go test -v -race ./..
