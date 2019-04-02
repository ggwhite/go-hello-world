build:
	GOOS=linux go build -o hello
	docker build -t ggwhite/go-hello-world .