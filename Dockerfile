FROM golang:alpine

ARG TOKEN

RUN apk --no-cache gcc

WORKDIR /app
COPY ./app/go.mod ./
COPY ./app/go.sum ./

RUN go mod download

COPY ./app/main.go ./
COPY ./app/translate.db ./

RUN go build -o /mtbot

CMD [ "/mtbot -t " ]
