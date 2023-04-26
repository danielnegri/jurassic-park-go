# Use the offical golang image to create a binary.
# This is based on Debian and sets the GOPATH to /go.
# https://hub.docker.com/_/golang
FROM golang:1.19-buster as builder

RUN set -x \
    && apt-get update \
    && apt-get install -y build-essential ca-certificates git-core zip \
    && rm -rf /var/lib/apt/lists/*

RUN set -x \
   && go install github.com/AlekSi/gocov-xml@v1.1.0 \
   && go install github.com/axw/gocov/gocov@v1.1.0 \
   && go install github.com/t-yuki/gocover-cobertura@latest \
   && go install github.com/tebeka/go2xunit@v1.4.10

# Create and change to the app directory.
WORKDIR /go/src/github.com/danielnegri/jurassic-park-go

# Retrieve application dependencies.
# This allows the container build to reuse cached dependencies.
# Expecting to copy go.mod and if present go.sum.
COPY go.* ./
RUN go mod download

# Copy local code to the container image.
COPY . ./

# Build the binary.
RUN set -x \
    && make testall \
    && make release-binary \
    && mkdir -p /usr/share/jurassic-park-go \
    && cp -r ./release/bin /usr/share/jurassic-park-go/. \
    && cp -r ./results /usr/share/jurassic-park-go/. \
    && echo "Build complete."


# Use the official Debian slim image for a lean production container.
# https://hub.docker.com/_/debian
# https://docs.docker.com/develop/develop-images/multistage-build/#use-multi-stage-builds
FROM debian:buster-slim
RUN set -x \
    && apt-get update \
    && DEBIAN_FRONTEND=noninteractive apt-get install -y build-essential ca-certificates \
    && rm -rf /var/lib/apt/lists/*

# Copy the binary to the production image from the builder stage.
COPY --from=builder /usr/share/jurassic-park-go /usr/share/jurassic-park-go
RUN ln -s /usr/share/jurassic-park-go/bin/admin /usr/bin/admin

# Run the web service on container startup.
CMD ["admin", "version"]
