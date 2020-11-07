
BINARY=api-creator-backend
test:
	go test -v -cover -covermode=atomic ./...

build:
	go build -o ${BINARY} app/api-creator-backend/api-creator-backend.go

build-documents:
	go build -o api-creator-documents app/api-creator-documents/api-creator-documents.go

unittest:
	go test -short  ./...

clean:
	if [ -f ${BINARY} ] ; then rm ${BINARY} ; fi

docker:
	docker build -t ${BINARY} .

run:
	docker-compose -f ./docker/docker-compose.yml up --build -d

stop:
	docker-compose -f ./docker/docker-compose.yml down --volumes

.PHONY: test docker run stop build make