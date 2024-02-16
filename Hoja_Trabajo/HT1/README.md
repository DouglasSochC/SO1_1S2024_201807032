#### Link del video: 

https://drive.google.com/file/d/1srOm64Ztqekhr7HI-mSSQPIdGdHf7wVT/view?usp=sharing

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

Eliminar el modulo (.ko)
```
sudo rmmod [nombre_modulo.ko]
```