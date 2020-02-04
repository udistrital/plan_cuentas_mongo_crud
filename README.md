# plan_cuentas_mongo_crud


El API plan_cuentas_mongo_crud, generada con beego, proporciona interfaces para la manipulación(CRUD) de los datos almacenados en una base de datos no relacional MongoDB (rubros, apropiaciones, fuentes de financiamiento, cpd, crp, vigencias), esta API hace representa la capa de datos del sistema de gestión financiero KRONOS.

## Especificaciones Técnicas

### Requisitos previos
* Docker
* Docker Compose

*Si se utiliza Windows como sistema operativo, se recomienda utilizar **Git bash** para ejecutar todos los comandos posteriores*

### Variables de Entorno

```sh
FINANCIERA_MONGO_CRUD_PORT = [descripción]
FINANCIERA_MONGO_CRUD_DB_URL = [descripción]
FINANCIERA_MONGO_CRUD_DB_NAME = [descripción]
FINANCIERA_MONGO_CRUD_DB_USER = [descripción]
FINANCIERA_MONGO_CRUD_DB_PASS = [descripción]
FINANCIERA_MONGO_CRUD_DB_AUTH = [descripción]
```

### Ejecución del proyecto con Docker

```sh
#1. Clonar el repositorio
git clone -b dev https://github.com/udistrital/plan_cuentas_mongo_crud

#2. Moverse a la carpeta del repositorio
cd plan_cuentas_mongo_crud

#3. Crear un fichero con el nombre **custom.env**
# En windows ejecutar:* ` ni custom.env`
touch custom.env

#4. Crear la network **back_end** para los contenedores
docker network create back_end

#5. Ejecutar el compose del contenedor
docker-compose up --build

#6. Comprobar que los contenedores estén en ejecución
docker ps
```

### Ejecución del proyecto directamente con Go
Para ejecutar el proyecto con el lenguaje **Go** es necesario tener instalado en su equipo:

Prerequisitos:

* Go
* beego
* MongoDB

```sh
#1. Obtener el repositorio con Go
go get github.com/udistrital/plan_cuentas_mongo_crud

#2. Moverse a la carpeta del repositorio
cd $GOPATH/src/github.com/udistrital/plan_cuentas_mongo_crud

# 3. Moverse a la rama **dev**
git pull origin dev && git checkout dev

# 4. listar todas las variables de entorno que utiliza el proyecto. Las variables se pueden ver en el fichero **conf/app.conf** y están identificadas con **${FINANCIERA_MONGO_CRUD_...}**

FINANCIERA_MONGO_CRUD_PORT=8080 FINANCIERA_MONGO_CRUD_DB_URL=127.0.0.1:27017 FINANCIERA_MONGO_CRUD_SOME_VARIABLE=some_value bee run
```

![Vista previa](images/terminal_api_view.png)

### Servicios

Para la documentación de esta API, se utiliza swagger. Si quieres ver una documentación exaustiva de todos los servicios que provee esta API, una vez ejecutado el contenedor o el proyecto de beego, ve a la dirección http://127.0.0.1:8082/swagger/swagger-1/


## Licencia

This file is part of plan_cuentas_mongo_crud.

plan_cuentas_mongo_crud is free software: you can redistribute it and/or modify it under the terms of the GNU General Public License as published by the Free Software Foundation, either version 3 of the License, or (at your option) any later version.

Foobar is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU General Public License for more details.

You should have received a copy of the GNU General Public License along with Foobar. If not, see https://www.gnu.org/licenses/.
