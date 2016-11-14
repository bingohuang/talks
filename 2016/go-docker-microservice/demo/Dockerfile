# Start from a Debian image with the latest version of Go installed
# and a workspace (GOPATH) configured at /go.
FROM golang:1.7.3

# Copy the local package files to the container's workspace.
ADD . /go/src/app

# install inside the container.
RUN go install -v app

# Document that the service listens on port 8080.
EXPOSE 8080

# Run the app command by default when the container starts.
#ENTRYPOINT /go/bin/app
