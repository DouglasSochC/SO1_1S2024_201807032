# Docker

1. Se inicia sesión

    ```console
    docker login
    ```

2. Construir imagen.

    El comando debe de ejecutarse donde esta el Dockerfile

    ```console
    docker build -t tu_nombre_de_usuario/mi-aplicacion:version_nueva .
    ```

3. Se crea el tag de la imagen

    ```console
    docker tag tu_nombre_de_usuario/mi-aplicacion tu_nombre_de_usuario/mi-aplicacion:version
    ```

4. Se sube la imagen

    ```console
    docker push tu_nombre_de_usuario/mi-aplicacion:version
    ```