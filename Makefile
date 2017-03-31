.PHONY: build

build:
	docker build -t jabley/hello-go-web-builder -f Dockerfile.build .
	docker run --rm jabley/hello-go-web-builder | docker build -t jabley/hello-go-web -f Dockerfile.run -
