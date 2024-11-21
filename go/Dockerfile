# Usa una imagen base de Go
FROM golang:1.20 AS builder

# Setea el directorio de trabajo
WORKDIR /app

# Copia el go.mod y go.sum y descarga las dependencias
COPY go.mod go.sum ./
RUN go mod download

# Copia el resto del código fuente
COPY . .

# Compila la aplicación
RUN go build -o main ./go/main.go

# Usa una imagen base ligera para el contenedor final
FROM alpine:latest

# Setea el directorio de trabajo
WORKDIR /app

# Copia el binario de la etapa de construcción
COPY --from=builder /app/main .

# Copia los archivos HTML
COPY front/html /app/front/html

# Expone el puerto
EXPOSE 8080

# Comando para ejecutar la aplicación
CMD ["./main"]