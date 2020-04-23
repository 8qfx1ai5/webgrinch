

# build container image from container manifest file
.PHONY: build
build:
	docker build -t viewcrypt-alpha -f ./build/container-image/Dockerfile .


# run container based on an image
.PHONY: run
run:
	docker run --rm -p "80:8080" viewcrypt-alpha
