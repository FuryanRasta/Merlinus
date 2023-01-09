mage-env:
	docker build --build-arg UID=$(shell id -u) --build-arg GID=$(shell id -g) --tag magelabs/mage-env mage-env

mage-python:
	docker build --tag magelabs/mage-python mage-python

mage-cosmovisor:
	docker build \
		--build-arg UID=$(shell id -u) --build-arg GID=$(shell id -g) \
		--build-arg MAGE_VERSION=$(MAGE_VERSION) \
		--tag magelabs/mage-cosmovisor:$(MAGE_VERSION) \
		mage-cosmovisor

mage-builder:
	docker build --tag magelabs/builder -f mage-builder/Dockerfile $(CONTEXT)

all: mage-env mage-python mage-cosmovisor mage-builder

.PHONY: all mage-env mage-python mage-cosmovisor mage-builder
