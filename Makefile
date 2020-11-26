
ADMIN=api-creator-admin
API_SERVER=api-creator-apiserver
test:
	go test -v -cover -covermode=atomic ./...

build-admin:
	go build -o ${ADMIN} app/${ADMIN}/${ADMIN}.go

build-apiserver:
	go build -o ${API_SERVER} app/${API_SERVER}/${API_SERVER}.go

unittest:
	go test -short  ./...

clean:
	if [ -f ${ADMIN} ] ; then rm ${ADMIN} ; fi

docker:
	docker build -t ${ADMIN} .

build-run:
	docker-compose -f ./docker/docker-compose.yml up --build -d

run:
	docker-compose -f ./docker/docker-compose.yml up -d

stop:
	docker-compose -f ./docker/docker-compose.yml down --volumes

.PHONY: test docker run stop build make