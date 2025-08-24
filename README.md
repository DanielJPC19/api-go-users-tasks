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

### Obtención de Driver

Dado que haremos uso de la librería database/sql, debemos de conseguir un Driver para conectarnos con MySQL, en nuestro caso utilizaremos uno que se puede hallar en la documentación de Go. Para obtenerlo, ya que no es puro de Go, debemos de ejecutar el siguiente comando:

```bash
go get -u github.com/go-sql-driver/mysql
```

### Creación del .env

Creamos el ```.env``` definiendo la dirección hacia MYSQL, usando el siguiente formato:

```.env
<usuario>:<contraseña>@tcp(<conexion/contenedor/ip>:3306)/<base de datos>
```

Donde, el **usuario**, la **contraseña** y la **base de datos** son las ya definidas en nuestro ```docker-compose.yml```. Adicionalmente, dado que estamos ejecutando la aplicación de Go de manera local y no haciendole un contenedor, pondremos en la **conexión/condenedor/ip** *localhost*.