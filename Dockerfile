FROM golang:1.17 as docker-learn
ENV GO111MODULE=on
ENV GOPUBLIC=github.com/athioushranjans
WORKDIR /server
COPY go.mod go.sum /server/
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 go build -o bin/docker-learn main.go

FROM golang:1.17-alpine
WORKDIR /
COPY . /server
COPY --from=docker-learn /server/bin .
EXPOSE 8080
CMD ["./docker-learn"]
