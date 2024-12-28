FROM golang:1.23

WORKDIR /app

COPY . .

RUN go mod tidy

RUN go build -o app main.go

ENV PORT=58080

EXPOSE ${PORT}

ENTRYPOINT [ "./app" ]
