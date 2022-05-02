# Go RESTful API v2

Esta aplicación fue diseñada con la intención de aplicar y promover las mejores practicas que siguen los [principios SOLID](https://en.wikipedia.org/wiki/SOLID) y [Clean Architecture](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html).

La aplicación contiene las siguientes features:

* RESTful endpoints
* Standard CRUD
* JWT-based authentication
* Log estructurado con información contextual
* Manejo de errores
* Migración de Base de datos
* Data validation
* Test optimizados

Utilizo los siguientes Go packages:

* Routing: [ozzo-routing](https://github.com/go-ozzo/ozzo-routing)
* Database access: [ozzo-dbx](https://github.com/go-ozzo/ozzo-dbx)
* Database migration: [golang-migrate](https://github.com/golang-migrate/migrate)
* Data validation: [ozzo-validation](https://github.com/go-ozzo/ozzo-validation)
* Logging: [zap](https://github.com/uber-go/zap)
* JWT: [jwt-go](https://github.com/dgrijalva/jwt-go)

## Empecemos

Algo obvio pero a veces no tanto, debemos instalar Go siguiendo las [instrucciones de instalación](https://golang.org/doc/install). La App utiliza **Go 1.14 o mayor**.

[Docker](https://www.docker.com/get-started) es necesario si se quiere configurar una propia base de datos en un contenedor.
Requerimos de la versión **Docker 17.05 o mayor** para el multi-stage build support.

Una vez tengamos Go y Docker, utilizaremos los siguientes comandos:

```shell
# descargar la app
git clone https://github.com/Matias-Guevara/Go/API-Rest-v2.git

cd API-Rest-v2

# Iniciar un servidor de PostgreSQL en un Docker container
make db-start

# Agregar data a la bd para testear
make testdata

# Correr el servidor
make run

# Tambien se puede correr el server con live reloading, es muy util durante el desarrollo
# requerimos de fswatch (https://github.com/emcrisostomo/fswatch)
make run-live
```

En este momento, tendras una API RESTful corriendo en `http://127.0.0.1:8080`. La cual tiene los siguientes endpoints:

* `GET /healthcheck`: un servicio de chequeo de estado del servidor (es necesario si es que implementamos un server cluster)
* `POST /v1/login`: autentifica al usuario y genera un JWT (JSON Web Token)
* `GET /v1/albums`: retorna una lista de albums paginados
* `GET /v1/albums/:id`: retorna información detallada del album
* `POST /v1/albums`: crea un album
* `PUT /v1/albums/:id`: actualiza un album existente
* `DELETE /v1/albums/:id`: elimina un album

Pruebe la URL `http://localhost:8080/healthcheck`, deberia salir algo como `"OK v1.0.0"` en pantalla.

Si tenes `cURL` o alguna herramienta para APIs (ej. [Postman](https://www.getpostman.com/)), podes probar escenarios más complejos:

```shell
# autentifica al usuario via: POST /v1/login
curl -X POST -H "Content-Type: application/json" -d '{"username": "demo", "password": "pass"}' http://localhost:8080/v1/login
# deberia devolver a JWT como: {"token":"...JWT token here..."}

# con el token que acabamos de obtener, tenemos acceso a los albums: GET /v1/albums
curl -X GET -H "Authorization: Bearer ...JWT token here..." http://localhost:8080/v1/albums
# deberia devolver una lista de albums en formato JSON
```

## Diseño del Proyecto
 
```
.
├── cmd                  carpeta principal del proyecto
│   └── server           la API server application
├── config               configuración de archivos y para distintos environments
├── internal             codigo privado
│   ├── album            album-related features
│   ├── auth             authentication feature
│   ├── config           configuration library
│   ├── entity           definición de entidades y lógica del dominio
│   ├── errors           manejo de errores
│   ├── healthcheck      healthcheck feature (chequeo del estado del servidor)
│   └── test             helpers para testeo
├── migrations           database migrations
├── pkg                  codigo publico
│   ├── accesslog        access log middleware
│   ├── dbcontext        soporte para transacciones de la BD
│   ├── log              structured and context-aware logger
│   └── pagination       lista paginada
└── testdata             test data scripts
```

Las carpetas principales del proyecto `cmd`, `internal`, `pkg` son comunes en la mayoria de proyectos hechos en Go, como se explica en 
[Standard Go Project Layout](https://github.com/golang-standards/project-layout).

Dentro de `internal` y `pkg`, los paquetes estan estructurados por funcionalidad, acá aplico un poco de 
[screaming architecture](https://blog.cleancoder.com/uncle-bob/2011/09/30/Screaming-Architecture.html).

Además, dentro de cada "paquete funcional", el código esta organizado por capas (API, servicio, repositorio), siguiendo la estructura descripta en 
[clean architecture](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html).

### Configuraciones

La configuración de la aplicación se encuentra en `internal/config/config.go`. Cuando se inicia la aplicación,
carga la configuración desde el archivo, así como las variables de entorno. La ruta del archivo de configuración 
se especifica a través del argumento de línea de comandos `-config`, que por defecto es `./config/local.yml`. Las configuraciones
especificadas en las variables de entorno deben nombrarse con el prefijo `APP_` y en mayúsculas. Cuando una configuración
se especifica tanto en un fichero de configuración como en una variable de entorno, esta última tiene prioridad. 

El directorio `config` contiene los ficheros de configuración con nombres de entornos diferentes. Por ejemplo
`config/local.yml` corresponde al entorno de desarrollo local y se utiliza cuando se ejecuta la aplicación 
mediante `make run`.

## Deployment

La aplicación se puede ejecutar como un contenedor de Docker. Puedes usar `make build-docker` para construir la aplicación 
en una imagen docker. El contenedor docker se inicia con el script `cmd/server/entryscript.sh` que lee 
la variable de entorno `APP_ENV` para determinar qué archivo de configuración utilizar. Por ejemplo
si `APP_ENV` es `qa`, la aplicación se iniciará con el archivo de configuración `config/qa.yml`.

También puedes ejecutar `make build` para construir un binario ejecutable llamado `server`. A continuación, inicie el servidor de la API con el siguiente comando
comando:

```shell
./server -config=./config/prod.yml
```

```
