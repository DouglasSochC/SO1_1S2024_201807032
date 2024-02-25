# XSQL

_Este es el primer proyecto del curso de Sistemas Operativos 1, el cual se tiene como objetivo principal implementar un sistema de monitoreo de recursos del sistema y gestión de procesos. El sistema resultante permitirá obtener información clave sobre el rendimiento del computador, procesos en ejecución y su administración a través de una interfaz amigable._

<!-- [📑 Enunciado](enunciado.pdf) -->

## 🚀 Comenzando

### 📋 Requerimientos

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

### ⚙️ Ejecucion

#### Backend

Dado que el backend está implementado en Golang, es esencial compilar la aplicación antes de ejecutarla. Por lo tanto, se recomienda ejecutar el siguiente comando en el directorio correspondiente (/backend).

```console
go build main.go
```

Una vez hecho lo anterior, simplemente se ejecuta el compilado

```console
./main
```

#### Frontend

Dado que el frontend está implementado en ReactJS, simplemente se ejecuta el siguiente comando en el directorio correspondiente (/frontend).

```console
npm start
```

#### Modulos

Dado que los módulos se implementaron en C, se utiliza un Makefile para facilitar la compilación y gestión del proyecto. Para compilar y gestionar los módulos, se recomienda ejecutar el siguiente comando en el directorio correspondiente

```console
make all
```

## 📖 Documentacion

### 🔠 Gramatica

abc

### 📑 Reportes
abc -->