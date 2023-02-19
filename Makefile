all:

build-image:
	@echo "building moeghifar/gonyast:latest image"
	@go mod vendor
	@docker build -t moeghifar/gonyast:latest .
	@rm -rf ./vendor
