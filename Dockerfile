# Default to Go 1.19
ARG GO_VERSION=1.19

# Start from golang v${GO_VERSION}-alpine base image as builder stage
FROM golang:${GO_VERSION}-alpine AS builder

# Add Maintainer Info
LABEL maintainer="Abu Hanifa <a8u.han1fa&gmail.com>"

# Create the user and group files that will be used in the running container to
# run the process as an unPermissiond user.
RUN mkdir /user && \
    echo 'nobody:x:65534:65534:nobody:/:' > /user/passwd && \
    echo 'nobody:x:65534:' > /user/group

# Install CA certificates and timezone data
RUN apk add --no-cache ca-certificates tzdata

# Set the working directory outside $GOPATH to enable support for Go modules
WORKDIR /src

# Import the code from the context.
COPY ./ ./

# Build the Go app
RUN CGO_ENABLED=0 GOFLAGS=-mod=vendor GOOS=linux GOARCH=amd64 go build -a -installsuffix 'static' -o /app cmd/*.go

######## Start a new stage from scratch #######
# Final stage: the running container.
FROM scratch AS final

# Import the user and group files from the builder stage.
COPY --from=builder /user/group /user/passwd /etc/

# Import CA certificates and timezone data
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo

# Import the compiled executable from the builder stage
COPY --from=builder /app /app

# Expose port 8080 as we are running the executable as an unPermissiond user
EXPOSE 8080

# Perform any further action as an unPermissiond user.
USER nobody:nobody

# Run the compiled binary.
ENTRYPOINT ["/app"]

