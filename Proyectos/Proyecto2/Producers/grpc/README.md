# GRPC

gRPC es un marco de comunicación de procedimiento remoto de código abierto inicialmente desarrollado por Google. Es utilizado para habilitar la comunicación entre servicios de servidor y cliente a través de diferentes lenguajes de programación. Funciona de manera eficiente en entornos distribuidos y soporta tanto conexiones locales como globales.

## Generacion de compilados proto

Para generar los compilados tanto del cliente como del servidor, es necesario abrir una consola en la raiz del proyecto y ejecutar los siguientes comandos. Esto permitirá generar los compilados correctamente.

### Para el cliente

```console
protoc --go_out=Producers/grpc/cliente/proto/. --go-grpc_out=Producers/grpc/cliente/proto/. Producers/grpc/cliente/proto/client.proto
```

### Para el servidor

```console
protoc --go_out=Producers/grpc/servidor/proto/. --go-grpc_out=Producers/grpc/servidor/proto/. Producers/grpc/servidor/proto/server.proto
```