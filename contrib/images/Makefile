desmos-env:
	docker build --build-arg UID=$(shell id -u) --build-arg GID=$(shell id -g) --tag desmoslabs/desmos-env desmos-env

mage-python:
	docker build --tag desmoslabs/desmos-python mage-python

mage-cosmovisor:
	docker build \
		--build-arg UID=$(shell id -u) --build-arg GID=$(shell id -g) \
		--build-arg MAGE_VERSION=$(MAGE_VERSION) \
		--tag desmoslabs/desmos-cosmovisor:$(MAGE_VERSION) \
		mage-cosmovisor

mage-builder:
	docker build --tag desmoslabs/builder -f mage-builder/Dockerfile $(CONTEXT)

all: desmos-env mage-python mage-cosmovisor mage-builder

.PHONY: all desmos-env mage-python mage-cosmovisor mage-builder
