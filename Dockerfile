FROM golang:1.20

RUN go version
ENV $GOPATH=/

WORKDIR /app

COPY ./ /app
 
RUN go mod download
 
RUN go build -o webapp ./cmd/web
 
CMD ["./webapp"]