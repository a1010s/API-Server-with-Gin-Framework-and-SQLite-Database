BINARY_NAME=go-apiserver
DOCKER_IMAGE_NAME=go-apiserver-image

build:
	go build -o bin/$(BINARY_NAME)

#run: build
#	./bin/$(BINARY_NAME)

run:
	go run main.go

test:
	go test -v ./...

docker-build: build
	docker build -t $(DOCKER_IMAGE_NAME) .

docker-destroy:
	docker rmi $(DOCKER_IMAGE_NAME) --force
