FROM golang:1.19-alpine

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download
RUN go mod tidy
COPY *.go ./
COPY internal/ ./internal

RUN go build -o /todo-list

EXPOSE 8000

CMD [ "/todo-list" ]
