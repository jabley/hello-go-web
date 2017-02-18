.PHONY: build

build:
	docker build -t hello-go-web-builder -f Dockerfile.build .
	docker run --rm hello-go-web-builder | docker build -t hello-go-web -f Dockerfile.run -
