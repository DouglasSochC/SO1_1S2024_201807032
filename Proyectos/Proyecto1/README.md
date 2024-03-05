# Monitoreo y Señales de Procesos

_Este es el primer proyecto del curso de Sistemas Operativos 1, el cual se tiene como objetivo principal implementar un sistema de monitoreo de recursos del sistema y gestión de procesos. El sistema resultante permitirá obtener información clave sobre el rendimiento del computador, procesos en ejecución y su administración a través de una interfaz amigable._

<!-- [📑 Enunciado](enunciado.pdf) -->

## 🚀 Comenzando

### 📋 Requerimientos para desarrollo

* [Golang 1.22.0](https://go.dev/dl/)
```console
go version
```

* [Node 20.11.1](https://nodejs.org/en/download/)
```console
node --version
```

* GCC 13.1.0
```console
gcc --version
```

### 📋 Requerimientos para la ejecución del proyecto

* Docker 25.0.3
```console
docker -v
```

### ⚙️ Ejecución 

Dado que los módulos se implementaron en C, se utiliza un Makefile para facilitar la compilación y gestión del proyecto. Para compilar y gestionar los módulos, se recomienda ejecutar el siguiente comando en el directorio 'modulos'

```console
make all
```

> Otros comandos make [aqui](referencia_comandos.md)

Dado que el proyecto se ejecuta a través de Docker, unicamente se debe de realizar el siguiente comando en la raiz del proyecto:

```console
docker compose up
```

## 📖 Documentación

### 🔠 Titulo 1

abc

### 📑 Titulo 2
abc -->