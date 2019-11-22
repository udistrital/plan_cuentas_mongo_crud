# plan_cuentas_mongo_crud

Api integrada con MongoDB para el negocio de plan de cuentas (rubros, apropiaciones, fuentes de financiamiento, productos, CDP, CRP, modificaciones presupuestales).

* Integration with CI

## Como ejecutar

### Requisitos previos
* Docker
* Docker Compose 

*Se se utiliza Windows como sistema operativo, se recomienda utilizar **Git bash** para ejecutar todos los comandos posteriores*
### Comenzar a desarrollar
#### Ejecutar con Docker

1. Clonar el repositorio
```sh
git clone -b dev https://github.com/udistrital/plan_cuentas_mongo_crud
```

2. Moverse a la carpeta del repositorio
```sh
cd plan_cuentas_mongo_crud
```

3. Crear un fichero con el nombre **custom.env**
```sh
touch custom.env
```

*En windows ejecutar:* ` ni custom.env`

4. Crear la network **back_end** para los contenedores
```sh
docker network create back_end
```

5. Ejecutar el compose del contenedor
```sh
docker-compose up --build
```

6. Comprobar que los contenedores estén en ejecución
```sh
docker ps 
```

#### Ejecutar directamente con Go
Para ejecutar el proyecto directamente con el lenguaje **Go** es necesario tener preinstalado en su equipo:

Prerequisitos: 

* Go
* beego
* MongoDB

1. Obtener el repositorio con Go
```sh
go get github.com/udistrital/plan_cuentas_mongo_crud
```

2. Moverse a la carpeta del repositorio
```sh
cd $GOPATH/src/github.com/udistrital/plan_cuentas_mongo_crud
```

3. Moverse a la rama **dev**
```sh
git pull origin dev && git checkout dev
```

4. Enlistar todas las variables de entorno que utiliza el proyecto. Las variables se pueden ver en el fichero **conf/app.conf** y están identificadas con **${FINANCIERA_MONGO_CRUD_...}**
```sh
FINANCIERA_MONGO_CRUD_PORT=8080 
FINANCIERA_MONGO_CRUD_DB_URL=127.0.0.1:27017
FINANCIERA_MONGO_CRUD_SOME_VARIABLE=some_value
.
.
.
bee run
```
*Nota: El comando anterior es una sola linea, no se ejecuta uno por uno (sin saltos de linea).*

![Vista previa](images/terminal_api_view.png)

## Servicios

Para la documentación de esta API, se utiliza swagger. Si quieres ver una documentación exaustiva de todos los servicios que provee esta API, una vez ejecutado el contenedor o el proyecto de beego, ve a la dirección http://127.0.0.1:8082/swagger/swagger-1/

## Derechos de Autor

This program is free software: you can redistribute it 
and/or modify it under the terms of the GNU General Public 
License as published by the Free Software Foundation, either
version 3 of the License, or (at your option) any later
version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

### UNIVERSIDAD DISTRITAL FRANCISCO JOSÉ DE CALDAS

### OFICINA ASESORA DE SISTEMAS

### 2019
### 
