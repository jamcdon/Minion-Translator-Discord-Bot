FROM golang

ENV TOKEN ""

WORKDIR /app
COPY ./app/go.mod ./
COPY ./app/go.sum ./

RUN go mod download

COPY ./app/main.go ./
COPY ./translate.db ./

RUN go build -o /app/mtbot

CMD [ "sh", "-c", "/app/mtbot", "-t ", "$TOKEN" ]
