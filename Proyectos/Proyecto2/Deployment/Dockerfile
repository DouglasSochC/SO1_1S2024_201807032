# Imagen oficial de golang
FROM golang:latest

# Establece el directorio de trabajo dentro del contenedor
WORKDIR /go/src/consumer

# Copia los archivos go mod y sum
COPY go.mod go.sum ./

# Descarga e instala las dependencias
RUN go mod download

# Copia el resto de los archivos de la aplicación
COPY . .

# Construye la aplicación para producción
RUN go build -o main .

# Puerto en el que la aplicación se ejecutara
EXPOSE 3000

# Comando para iniciar la aplicación cuando se ejecute el contenedor
CMD ["./main"]
