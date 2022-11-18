FROM golang:1.19-buster

RUN go version
ENV GOPATH=/

COPY ./ ./

# build go app
RUN go mod download
RUN go build -o avito ./cmd/main.go

CMD ["./avito"]