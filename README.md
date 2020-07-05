# plan_cuentas_mongo_crud

El API plan_cuentas_mongo_crud proporciona interfaces para la manipulación (CRUD) de los datos almacenados en una base de datos no relacional MongoDB (rubros, apropiaciones, fuentes de financiamiento, cpd, crp, vigencias)

Esta API representa la capa de datos del sistema de gestión financiero KRONOS.

## Especificaciones Técnicas

### Tecnologías Implementadas y Versiones

* [Golang](https://github.com/udistrital/introduccion_oas/blob/master/instalacion_de_herramientas/golang.md)
* [BeeGo](https://github.com/udistrital/introduccion_oas/blob/master/instalacion_de_herramientas/beego.md)
* [MongoDB]()
* [Docker](https://docs.docker.com/engine/install/ubuntu/)
* [Docker Compose](https://docs.docker.com/compose/)

### Variables de Entorno
```shell
FINANCIERA_MONGO_CRUD_PORT = [descripción]
FINANCIERA_MONGO_CRUD_DB_URL = [descripción]
FINANCIERA_MONGO_CRUD_DB_NAME = [descripción]
FINANCIERA_MONGO_CRUD_DB_USER = [descripción]
FINANCIERA_MONGO_CRUD_DB_PASS = [descripción]
FINANCIERA_MONGO_CRUD_DB_AUTH = [descripción]
```
**NOTA:** Las variables se pueden ver en el fichero conf/app.conf y están identificadas con FINANCIERA_MONGO_CRUD_...

### Ejecución del Proyecto
```shell
#1. Obtener el repositorio con Go
go get github.com/udistrital/plan_cuentas_mongo_crud

#2. Moverse a la carpeta del repositorio
cd $GOPATH/src/github.com/udistrital/plan_cuentas_mongo_crud

# 3. Moverse a la rama **dev**
git pull origin dev && git checkout dev

# 4. alimentar todas las variables de entorno que utiliza el proyecto.
FINANCIERA_MONGO_CRUD_PORT=8080 FINANCIERA_MONGO_CRUD_DB_URL=127.0.0.1:27017 FINANCIERA_MONGO_CRUD_SOME_VARIABLE=some_value bee run
```
![Vista previa](images/terminal_api_view.png)

### Ejecución Dockerfile
```shell
# docker build --tag=covid-api . --no-cache
# docker run -p 80:80 covid-api
```

### Ejecución docker-compose
```shell
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
### Ejecución Pruebas

Pruebas unitarias
```shell
# Not Data
```
## Estado CI

| Develop | Relese 0.0.1 | Master |
| -- | -- | -- |
| 1 | 2 | 3 |


## Licencia

This file is part of plan_cuentas_mongo_crud.

plan_cuentas_mongo_crud is free software: you can redistribute it and/or modify it under the terms of the GNU General Public License as published by the Free Software Foundation, either version 3 of the License, or (at your option) any later version.

plan_cuentas_mongo_crud is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU General Public License for more details.

You should have received a copy of the GNU General Public License along with plan_cuentas_mongo_crud. If not, see https://www.gnu.org/licenses/.
