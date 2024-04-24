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
* [Desplegar proyecto](#desplegar-proyecto)
* [Documentacion](#documentacion)

## ⭐ Comenzando <div id='comenzando'></div>

### 📋 Requerimientos <div id='requerimientos'></div>

* [Python 3.12.0](https://www.python.org/downloads/)

    Python es un lenguaje de programación de alto nivel, interpretado y de propósito general. Destacado por su legibilidad y simplicidad en la sintaxis, permite a los programadores expresar conceptos en menos líneas de código comparado con otros lenguajes. La versión 3.12.0 es la última versión estable, que incluye mejoras en las funcionalidades del lenguaje y correcciones de errores.

    ```console
    python --version
    ```


* [Golang 1.21.6](https://go.dev/doc/install)

    Go, también conocido como Golang, es un lenguaje de programación creado por Google. Es un lenguaje compilado, tipado estáticamente que facilita la construcción de software de manera eficiente y concurrente. La versión 1.21.6 incluye actualizaciones de rendimiento y seguridad, así como nuevas características para mejorar el desarrollo.

    ```console
    go version
    ```

* [GCloudCLI 471.0.0](https://cloud.google.com/sdk?hl=es-419)

    GCloud CLI (Google Cloud Command Line Interface) es una herramienta que permite a los desarrolladores gestionar los recursos de Google Cloud Platform (GCP) desde la línea de comando. Ofrece comandos para desplegar y manejar aplicaciones, manejar almacenamiento en la nube, configurar redes, entre otros. La versión 471.0.0 trae las últimas actualizaciones y características compatibles con GCP.
    ```console
    gcloud version
    ```

#### Paquetes adicionales

* [Protoc](https://www.geeksforgeeks.org/how-to-install-protocol-buffers-on-windows/)

    Protoc es el compilador de Protocol Buffers, un sistema de serialización de datos estructurado desarrollado por Google, usado ampliamente en servicios de comunicación y almacenamiento de datos. Protoc se utiliza para generar código fuente a partir de archivos de definición .proto en varios lenguajes de programación, en este caso es utilizado para una comunicación grpc en golang.

* [gRPC para Golang](https://grpc.io/docs/languages/go/quickstart/)

    gRPC es un marco de trabajo moderno y de alto rendimiento para la comunicación entre servicios, que usa HTTP/2 como protocolo de transporte y Protocol Buffers como mecanismo de serialización. La versión para Golang permite a los desarrolladores de Go construir sistemas distribuidos y escalables de manera eficiente.

* Kubectl

    Kubectl es una herramienta de línea de comando para interactuar con clusters de Kubernetes. Permite a los usuarios desplegar aplicaciones, inspeccionar y manejar recursos del cluster, y ver logs. Es esencial para la gestión de clusters Kubernetes y es mantenido por Google como parte de su conjunto de herramientas de Google Cloud.

    ```console
    gcloud components install kubectl
    ```

## 📖 Documentacion <div id='documentacion'></div>

### 🎡 Arquitectura



### Realización de graficos en Grafana

### 📑 Preguntas
abc

## 🚀 Desplegar proyecto <div id='desplegar-proyecto'></div>

Dado que las imágenes de cada módulo se encuentran en Docker Hub, solo necesitas ejecutar los manifiestos en el siguiente orden. Asegúrate de que la consola esté ubicada en la ruta raíz de este proyecto antes de proceder.

1. Creación del namespace

    ```console
    kubectl create -f namespace.yaml
    ```

2. Creación del pod de MongoDB

    ```console
    kubectl create -f Database/mongodb.yaml
    ```

3. Creación del pod de Redis

    ```console
    kubectl create -f Database/redis.yaml
    ```

4. Creación de Kafka con Strimzi

    * Creación del operador

        ```console
        kubectl create -f 'https://strimzi.io/install/latest?namespace=so1-p2-201807032' -n so1-p2-201807032
        ```

    * Creación del volumen

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

6. Creación del servicio y pod del producer WASM (No realizado)

    <!-- ```console
    ``` -->

7. Creación del servicio y pods del deployment consumer

    > Nota 1: Recuerda crear la base de datos y la colección en MongoDB

    > Nota 2: Recuerda ajustar la variable de entorno **MONGO_HOST** con la dirección IP del servicio de MongoDB, así como la variable **REDIS_ADDR** con la dirección IP del servicio de Redis.

    ```console
    kubectl create -f Deployment/deployment.yaml
    ```

8. Creación del Horizontal Pod Autoscaler (HPA)

    ```console
    kubectl create -f Deployment/hpa-deployment.yaml
    ```

9. Implementacion de Grafana

    ```console
    kubectl create -f Grafana/grafana.yaml
    ```

10. Creación de Ingress

    ```console
    kubectl create -f Ingress/ingress.yaml
    ```
