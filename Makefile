#
#
#

all: image

rpm:
	./scripts/build-rpm.sh

image:
	./scripts/build-image.sh

push_image:
	./scripts/push-image.sh

clean:
	rm -rf build
