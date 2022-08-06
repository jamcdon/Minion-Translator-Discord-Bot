FROM golang

ENV TOKEN ""

RUN apk --no-cache add gcc

WORKDIR /app
COPY ./app/go.mod ./
COPY ./app/go.sum ./

RUN go mod download

COPY ./app/main.go ./
COPY ./translate.db /

RUN go build -o /mtbot

CMD [ "sh", "-c", "/mtbot", "-t ", "$TOKEN" ]
