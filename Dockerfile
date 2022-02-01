FROM golang:1.17 as docker-learn
ENV GO111MODULE=on
WORKDIR /server
COPY go.mod go.sum /server/
RUN go mod download

COPY . .
ARG HOST
ENV DB_HOST ${HOST}
ARG PASS
ENV DB_PASS ${PASS}
ARG NAME
ENV DB_NAME ${NAME}
ARG PORT
ENV DB_PORT ${PORT}
ARG USER
ENV DB_USER ${USER}
RUN CGO_ENABLED=0 go build -o bin/docker-learn main.go

FROM golang:1.17-alpine
WORKDIR /
COPY . /server
COPY --from=docker-learn /server/bin .
EXPOSE 8080
ENTRYPOINT ["./docker-learn"]
