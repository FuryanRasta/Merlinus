# To build the Mage image, just run:
# > docker build -t mage .
#
# Simple usage with a mounted data directory:
# > docker run -it -p 26657:26657 -p 26656:26656 -v ~/.mage:/root/.mage mage mage init
# > docker run -it -p 26657:26657 -p 26656:26656 -v ~/.mage:/root/.mage mage mage start
#
# If you want to run this container as a daemon, you can do so by executing
# > docker run -td -p 26657:26657 -p 26656:26656 -v ~/.mage:/root/.mage --name mage mage
#
# Once you have done so, you can enter the container shell by executing
# > docker exec -it mage bash
#
# To exit the bash, just execute
# > exit
FROM alpine:edge

# Install ca-certificates
RUN apk add --update ca-certificates

# Install bash
RUN apk add --no-cache bash

# Copy over binaries from the build-env
COPY --from=magelabs/builder:latest /code/build/mage /usr/bin/mage

EXPOSE 26656 26657 1317 9090

# Run mage by default, omit entrypoint to ease using container with mage
CMD ["mage"]
