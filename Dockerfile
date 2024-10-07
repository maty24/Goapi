# Etapa 1: Compilación de la aplicación
FROM golang:1.21-alpine AS builder

# Instalación de dependencias necesarias
RUN apk add --no-cache git

WORKDIR /app

# Copia y descarga de dependencias
COPY go.mod go.sum ./
RUN go mod download

# Copia del código fuente
COPY . .

# Compilación del binario
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o main ./cmd

# Etapa 2: Crear la imagen final de producción
FROM scratch

WORKDIR /app

# Copia del binario y del archivo .env desde el builder
COPY --from=builder /app/main .
COPY .env .

# Añadir etiquetas para mejor trazabilidad
LABEL maintainer="TuNombre <tuemail@dominio.com>"
LABEL version="1.0"
LABEL description="Una API escrita en Go"

# Variables de entorno
ENV GIN_MODE=release
ENV PORT=8080

# Exponer el puerto
EXPOSE ${PORT}

# Comando de inicio de la aplicación
CMD ["./main"]
