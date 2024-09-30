APP?=simpleservice
PORT?=8080

clean:
	rm -f ${APP}
build: clean
	go build -o ${APP}
run: build
	PORT=${PORT} ./${APP}
test:
	go test -v -race ./..
