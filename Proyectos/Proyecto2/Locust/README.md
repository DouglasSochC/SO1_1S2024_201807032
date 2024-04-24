# Locust

Locust es una herramienta de prueba de carga de código abierto escrita en Python. Se utiliza principalmente para probar el rendimiento de sistemas web y determinar cómo manejan la carga bajo condiciones de uso simuladas. La característica distintiva de Locust es que permite escribir escenarios de prueba en Python puro, lo que proporciona una gran flexibilidad y hace que las pruebas sean fácilmente legibles y mantenibles.

## Creación de entorno virtual

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