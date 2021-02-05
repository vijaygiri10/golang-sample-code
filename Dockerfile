FROM golang:1.15-alpine as builder

ENV APP_HOME /go/src/myapp
RUN mkdir -p $APP_HOME

WORKDIR $APP_HOME
COPY . $APP_HOME

# ENV GOOS linux
# ENV GOARCH amd64

RUN go mod download
RUN go mod verify
RUN go build -o server -ldflags="-w -s" *.go

# https://hub.docker.com/_/alpine
# https://docs.docker.com/develop/develop-images/multistage-build/#use-multi-stage-builds
FROM alpine:latest  
RUN apk --no-cache add ca-certificates

ENV APP_HOME /go/src/myapp

RUN mkdir -p $APP_HOME
WORKDIR $APP_HOME

# Copy the binary to the from the builder stage.
COPY --chown=0:0 --from=builder $APP_HOME/server $APP_HOME

# Copy the config file from the builder stage
COPY --chown=0:0 --from=builder $APP_HOME/config/staging.json $APP_HOME

# Run the web service on container startup.
EXPOSE 8080
ENTRYPOINT ["./server", "staging.json"]
