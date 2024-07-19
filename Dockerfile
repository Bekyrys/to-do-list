FROM golang:1.19-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./
COPY internal/ ./internal

RUN go build -o /todo-list

EXPOSE 8000

CMD [ "/todo-list" ]
