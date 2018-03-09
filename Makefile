DOCKER_IMAGE         ?= tchaudhry/hash-svc
DOCKER_IMAGE_TAG     ?= master
DOCKER_IMAGE_TAG_ARM ?= arm

all: docker docker-arm

docker:
	@echo ">> Building Docker Image"
	@docker build -t $(DOCKER_IMAGE):$(DOCKER_IMAGE_TAG) .

docker-push:
	@docker login -u $(DOCKER_USERNAME) -p $(DOCKER_PASSWORD)
	@docker push $(DOCKER_IMAGE):$(DOCKER_IMAGE_TAG)

docker-arm:
	@echo ">> Building Docker Image-ARM"
	@docker build -t $(DOCKER_IMAGE):$(DOCKER_IMAGE_TAG_ARM) -f Dockerfile-ARM .

docker-push-arm:
	@docker login -u $(DOCKER_USERNAME) -p $(DOCKER_PASSWORD)
	@docker push $(DOCKER_IMAGE):$(DOCKER_IMAGE_TAG_ARM)
