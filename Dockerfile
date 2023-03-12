ARG  DISTROLESS_IMAGE=gcr.io/distroless/static:nonroot
ARG AppListen=80

# Use the official Go image as the base image
FROM golang:alpine3.17 AS build

# Set the working directory to /app
WORKDIR /app

# Install deps
RUN apk update --no-cache && apk add pkgconf openssl-dev gcc libc-dev

# Copy the source code to the container
COPY . .

ENV GO111MODULE=on
RUN go mod download

# Build the binary
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -installsuffix 'static' -o /go/bin/file-manager .

# Use a lightweight image for the final stage
FROM ${DISTROLESS_IMAGE}

USER 65532:65532

# Copy the binary from the previous stage
COPY --from=build /go/bin/file-manager /go/bin/file-manager
COPY --from=build /lib/libssl.so.3     /lib/libssl.so.3
COPY --from=build /lib/libcrypto.so.3     /lib/libcrypto.so.3
COPY --from=build /lib/ld-musl-x86_64.so.1  /lib/ld-musl-x86_64.so.1

# Expose the port that the application will listen on
EXPOSE ${AppListen}

# Run the binary
ENTRYPOINT ["/go/bin/file-manager"]
