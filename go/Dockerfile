FROM golang:latest

WORKDIR /app
<<<<<<< HEAD

COPY go.mod go.sum ./

RUN go mod download

=======
COPY Directo-al-Modelaje/go.mod .
COPY Directo-al-Modelaje/go.sum .
>>>>>>> parent of 4621b8b (.)
COPY . .

RUN go build -o main .

<<<<<<< HEAD
CMD ["./main"]
=======
EXPOSE 8080

CMD ["./main"]

>>>>>>> parent of 4621b8b (.)
