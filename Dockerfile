# seteamos la imagen base de nuestro container
# al ser esta una aplicación de Golang, usamos eso

FROM golang:1.20 

WORKDIR /app

COPY go.mod .
COPY main.go .

#instalamos dependencias

RUN go get

# usamos el comando para construir el container y lo dejamos en bin

RUN go build -o bin .

ENTRYPOINT [ "/app/bin" ]