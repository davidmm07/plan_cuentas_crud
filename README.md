# plan_cuentas_crud

Api CRUD para negocio de plan de cuentas del sistema de gestion financiera kronos.

## Especificaciones Técnicas

### Tecnologías Implementadas y Versiones
* [Golang](https://github.com/udistrital/introduccion_oas/blob/master/instalacion_de_herramientas/golang.md)
* [BeeGo](https://github.com/udistrital/introduccion_oas/blob/master/instalacion_de_herramientas/beego.md)
* [Docker](https://docs.docker.com/engine/install/ubuntu/)
* [Docker Compose](https://docs.docker.com/compose/)

### Variables de Entorno
```shell
# Ejemplo que se debe actualizar acorde al proyecto
PLAN_CUENTAS_CRUD_DB_USER = [descripción]
PLAN_CUENTAS_CRUD_HTTP_PORT = [descripción]
```
**NOTA:** Las variables se pueden ver en el fichero conf/app.conf y están identificadas con PLAN_CUENTAS_CRUD_...

### Ejecución del Proyecto
```shell
#1. Obtener el repositorio con Go
go get github.com/udistrital/plan_cuentas_crud

#2. Moverse a la carpeta del repositorio
cd $GOPATH/src/github.com/udistrital/plan_cuentas_crud

# 3. Moverse a la rama **develop**
git pull origin develop && git checkout develop

# 4. alimentar todas las variables de entorno que utiliza el proyecto.
PLAN_CUENTAS_CRUD_HTTP_PORT=8080 PLAN_CUENTAS_CRUD_DB_HOST=127.0.0.1:27017 PLAN_CUENTAS_CRUD_SOME_VARIABLE=some_value bee run
```

### Ejecución Dockerfile
```shell
# docker build --tag=plan_cuentas_crud . --no-cache
# docker run -p 80:80 plan_cuentas_crud
```

### Ejecución docker-compose
```shell
#1. Clonar el repositorio
git clone -b develop https://github.com/udistrital/plan_cuentas_crud

#2. Moverse a la carpeta del repositorio
cd plan_cuentas_crud

#3. Crear un fichero con el nombre **custom.env**
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

## Modelo de datos
[Modelo de Datos Relacional](modelobd.png)  
[Script SQL](https://drive.google.com/file/d/10c-NiUE-869AT5GYyGMmFyMn3eOcGLU5/view?usp=sharing)


## Estado CI
| Develop | Relese 0.0.1 | Master |
| -- | -- | -- |
|1 |2 |3 |

## Licencia
This file is part of [plan_cuentas_crud](LICENSE).

plan_cuentas_crud is free software: you can redistribute it and/or modify it under the terms of the GNU General Public License as published by the Free Software Foundation, either version 3 of the License, or (at your option) any later version.

plan_cuentas_crud is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU General Public License for more details.

You should have received a copy of the GNU General Public License along with necesidades_crud. If not, see https://www.gnu.org/licenses/.
