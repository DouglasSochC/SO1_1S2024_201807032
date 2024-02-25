## Comandos utilizados

Limpiar archivos generados por el make.

```
make clean
```

Generar archivos para implementar el modulo.

```
make all
```

Borrar el búfer del registro de mensajes del kernel.

```
sudo dmesg -C
```

Mostrar los mensajes del kernel del sistema.

```
sudo dmesg
```

#### Para utilizar dentro del proyecto donde se generaron los archivos por medio del make.

Implementar el modulo (.ko)
```
sudo insmod [nombre_modulo.ko]
```

> **Nota**: El modulo implementado se encuentra en la siguiente ruta: /proc

Eliminar el modulo (.ko)
```
sudo rmmod [nombre_modulo.ko]
```