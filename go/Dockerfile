FROM golang:latest

WORKDIR /app
COPY Directo-al-Modelaje/go.mod .
COPY Directo-al-Modelaje/go.sum .
COPY . .

RUN go build -o main .

EXPOSE 8080

CMD ["./main"]

