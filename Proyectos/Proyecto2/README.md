# Sistema distribuido de votaciones

_En este proyecto universitario del curso Sistemas Operativos 1, se tiene como objetivo principal implementar un sistema de votaciones para un concurso de bandas de música guatemalteca; el propósito de este es enviar tráfico por medio de archivos con votaciones creadas hacia distintos servicios (grpc y wasm) que van a encolar cada uno de los datos enviados, así mismo se tendrán ciertos consumidores a la escucha del sistema de colas para enviar datos a una base de datos en Redis; estos datos se verán en dashboards en tiempo real. También se tiene una base de datos de Mongodb para guardar los logs, los cuales serán consultados por medio de una aplicación web._

## Introduccion

## Objetivos

## 🚀 Comenzando

### 📋 Requerimientos

* [Python 3.12.0](https://www.python.org/downloads/)
```console
python --version
```

* [Golang 1.21.6](https://go.dev/doc/install)
```console
go version
```

* [GCloudCLI 471.0.0](https://cloud.google.com/sdk?hl=es-419)
```console
gcloud version
```

* [gRPC para Golang](https://grpc.io/docs/languages/go/quickstart/)

* [Protoc](https://www.geeksforgeeks.org/how-to-install-protocol-buffers-on-windows/)

<!-- ### ⚙️ Ejecucion

Se utilizara un entorno virtual para levantar el proyecto con el fin de aislar las dependencias, evitar conflictos entre versiones, y garantizar que el proyecto tenga su propio entorno reproducible.

Instalar el modulo **virtualenv**

```console
pip install virtualenv
```

Ahora dentro de la carpeta del proyecto se debe de realizar lo siguiente:

1. Creación del entorno virtual, en este caso llama **venv**

    ```console
    virtualenv venv
    ```

2. Activar entorno virtual

    ```console
    source venv/Scripts/activate
    ```

3. Instalar las dependencias del proyecto

    ```console
    pip install -r requirements.txt
    ```

4. Ejecutar el proyecto

    ```console
    py main.py
    ```

## 📖 Documentacion

### 🔠 a

abc

### 📑 b
abc -->