DOCKER_TAG ?= wemusic-golang-app:development
BUILD_TYPE ?= development
CONTAINER_NAME ?= wemusic-golang

build:
	docker build -t ${DOCKER_TAG} .

run:
	docker run -p 8080:8080 --name ${CONTAINER_NAME} ${DOCKER_TAG}

stop:
	docker stop ${CONTAINER_NAME}

clean_container:
	docker rm -f ${CONTAINER_NAME}

clean_image:
	docker rmi ${DOCKER_TAG}

dev:
	make build
	make clean_container
	make run
	make show
