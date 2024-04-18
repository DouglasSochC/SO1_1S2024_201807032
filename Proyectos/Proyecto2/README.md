# Sistema distribuido de votaciones

En este proyecto universitario del curso Sistemas Operativos 1

## Introduccion

El principal objetivo de este proyecto es establecer un sistema de votación para un certamen de bandas de música guatemalteca. Se planea dirigir tráfico a través de archivos de votación hacia varios servicios (grpc y wasm) que se encargarán de encolar los datos recibidos. Además, se implementarán consumidores que monitorearán el sistema de colas para transferir los datos a una base de datos en Redis. Estos datos serán visualizados en tiempo real en paneles de control. Asimismo, se utilizará una base de datos MongoDB para almacenar registros, los cuales podrán ser consultados mediante una aplicación web.

## Objetivos

* Implementar un sistema distribuido con microservicios en kubernetes.
* Encolar distintos servicios con sistemas de mensajerías.
* Utilizar Grafana como interfaz gráfica de dashboards.

## Indice

* [Comenzando](#comenzando)
    * [Requerimientos](#requerimientos)
* [Entorno de desarrollo](#entorno-desarrollo)
    * [Para Locust](#para-locust)
    * [Para el Producer GRPC](#para-producer-grpc)
* [Desplegar proyecto](#desplegar-proyecto)
* [Documentacion](#documentacion)

## ⭐ Comenzando <div id='comenzando'></div>

### 📋 Requerimientos <div id='requerimientos'></div>

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

#### Paquetes adicionales

* [Protoc](https://www.geeksforgeeks.org/how-to-install-protocol-buffers-on-windows/)

* [gRPC para Golang](https://grpc.io/docs/languages/go/quickstart/)

* Kubectl

    ```console
    gcloud components install kubectl
    ```

## ⚙️ Entorno de desarrollo <div id='entorno-desarrollo'></div>

Después de haber instalado todos los requisitos del proyecto, aquí tienes una guía y un conjunto de comandos útiles que te servirán si decides trabajar en el proyecto.

### Para Locust <div id='para-locust'></div>

#### Creación de entorno virtual

Se utilizara un entorno virtual para levantar Locust con el fin de aislar las dependencias, evitar conflictos entre versiones, y garantizar que Locust tenga su propio entorno reproducible.

Instalar el modulo **virtualenv**

```console
pip install virtualenv
```

Ahora dentro de la carpeta **Locust** se debe de realizar lo siguiente:

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

4. Ejecutar

    ```console
    locust -f traffic.py
    ```

### Para el Producer GRPC <div id='para-producer-grpc'></div>

#### Generacion de compilados proto

Para generar los compilados tanto del cliente como del servidor, es necesario abrir una consola en la raiz del proyecto y ejecutar los siguientes comandos. Esto permitirá generar los compilados correctamente.

1. Para el cliente

    ```console
    protoc --go_out=Producers/grpc/cliente/proto/. --go-grpc_out=Producers/grpc/cliente/proto/. Producers/grpc/cliente/proto/client.proto
    ```

2. Para el servidor

    ```console
    protoc --go_out=Producers/grpc/servidor/proto/. --go-grpc_out=Producers/grpc/servidor/proto/. Producers/grpc/servidor/proto/server.proto
    ```

#### Subir imagen a Docker Hub

1. Se inicia sesión

    ```console
    docker login
    ```

2. Se crea el tag de la imagen

    ```console
    docker tag mi-aplicacion tu_nombre_de_usuario/mi-aplicacion:version
    ```

3. Se sube la imagen

    ```console
    docker push tu_nombre_de_usuario/mi-aplicacion:version
    ```

## 🚀 Desplegar proyecto <div id='desplegar-proyecto'></div>

Dado que las imágenes de cada módulo se encuentran en Docker Hub, solo necesitas ejecutar los manifiestos en el siguiente orden. Asegúrate de que la consola esté ubicada en la ruta raíz de este proyecto antes de proceder.

1. Creación del namespace

    ```console
    kubectl create -f namespace.yaml
    ```

2. Creación del pod de MongoDB

    <!-- ```console
    kubectl create -f Database/mongodb.yaml
    ``` -->

3. Creación del pod de Redis

    <!-- ```console
    kubectl create -f Database/redis.yaml
    ``` -->

4. Creación de Kafka con Strimzi

    * Creación del operador

        ```console
        kubectl create -f 'https://strimzi.io/install/latest?namespace=so1-p2-201807032' -n so1-p2-201807032
        ```

    * Creación del Zookeeper

        ```console
        kubectl apply -f https://strimzi.io/examples/latest/kafka/kafka-persistent-single.yaml -n so1-p2-201807032
        ```

    * Creación del topic

        ```console
        kubectl create -f Kafka/topic.yaml
        ```

5. Creación del servicio y pod del producer GRPC

    ```console
    kubectl create -f Producers/grpc/grpc.yaml
    ```

6. Creación del servicio y pod del producer WASM

    ```console
    ```

7. Creación de Ingress

    ```console
    kubectl create -f Ingress/ingress.yaml
    ```

## 📖 Documentacion <div id='documentacion'></div>

### 🎡 Arquitectura

abc

### 📑 b
abc