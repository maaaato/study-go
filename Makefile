
build:
	GOOS=linux GOARCH=amd64 go build -o monitor
	docker build . -t monitor

drun:
	docker run -it --rm  --entrypoint /bin/ash monitor