PROJECT_PATH=$(shell pwd)
MODULE_NAME=file-nginx

BUILD_NUM_FILE=build.num
BUILD_NUM=$$(cat ./build.num)
APP_VERSION=$$(cat ./version.txt)
IMAGE_REPOSITORY="kjuiop/uploader-nginx"
TARGET_VERSION=$(APP_VERSION).$(BUILD_NUM)

all: docker-build


docker-build:
	@echo "TARGET_VERSION : $(TARGET_VERSION)"
	docker build -f Dockerfile --tag $(IMAGE_REPOSITORY):$(TARGET_VERSION) .

docker-run:
	@echo "TARGET_VERSION : $(TARGET_VERSION)"
	docker run -p 8080:80 -p 443:443 --name $(MODULE_NAME) $(IMAGE_REPOSITORY):$(TARGET_VERSION)

docker-rm:
	docker stop $(MODULE_NAME)
	docker rm $(MODULE_NAME)

docker-push:
	@echo "TARGET_VERSION : $(TARGET_VERSION)"
	docker push $(IMAGE_REPOSITORY):$(TARGET_VERSION)

build_num:
	@echo $$(($$(cat $(BUILD_NUM_FILE)) + 1 )) > $(BUILD_NUM_FILE)
	@echo "BUILD_NUM      : $(BUILD_NUM)"

target-version:
	@echo "========================================"
	@echo "APP_VERSION    : $(APP_VERSION)"
	@echo "BUILD_NUM      : $(BUILD_NUM)"
	@echo "TARGET_VERSION : $(TARGET_VERSION)"
	@echo "========================================"
