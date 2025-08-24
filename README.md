# API Go Users-Tasks

## Paso a paso realizado en este proyecto

### Inicicialización de un proyecto en Golang

Primero, para iniciar el proyecto y definir el nombre del módulo, debemos de ejecutar el siguiente comando:

```bash
go mod init <nombre_del_modulo>
```

En nuestro caso, hacemos uso como nombre del módulo user-task. Por consiguiente, se crea un archivo ```go.mod```, donde define nuestra versión de Go y el nombre del módulo.

### Uso de Base de Datos

Para nuestro caso, haremos uso de una base de datos en MySQL, creando un contenedor Docker para todo lo que vamos a realizar. Cremos un docker compose para manejar nuestros módulos. Se habilita un cliente web para visualizar los cambios en la base de datos.

> Recuerde autenticarse en la base de datos con las credenciales puestas en el docker compose, siendo el usuario **user-task** y la contraseña **2021**.