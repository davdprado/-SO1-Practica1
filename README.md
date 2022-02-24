# Practica1

_Practica 1 del Sistemas Operativos 1, 1er Semestre 2022_

- [Construir el entorno](#inicio)
- [Creando nuestras imagenes](#images)
- [Docker-Compose](#compose)

<a name="inicio"></a>

### Pre-requisitos 📋

_Se necesitan tener instalado y poseer conocimientos basicos de las siguientes herramientas y/o lenguajes de programacion:_

- Tener instalado [Docker](https://docs.docker.com/engine/install/ubuntu/)
- [MongoDB](https://docs.mongodb.com/)
- [React](https://es.reactjs.org/docs/create-a-new-react-app.html)
- [Go](https://go.dev/learn/)

## Empezando ⚙️

_Para empezar vamos a tener que hacer una API-REST, una aplicacion en react y una base de datas hecha en mongodb._

- [Mongo](mongo)
- [React](react)
- [Golang](go)

<a name="mongo"></a>

#### MogoDB

_Para crear nuestra base de datos utilizaresmos [MongoDB](https://docs.mongodb.com/) iniciaremos por descargar la imagen de docker para crear nuestro contenedor_

```
docker pull mongo
#para ver que si se descargo la imagen podemos hacer el siguiente comando
docker images
```

_Ya tenemos nuestra imagen es hora de crear un contenedor, para eso corremos el siguiente comando_

```
docker run -d -p 27017:27017 -v /rutadevolumen:/data/db --network pra1 --name nombrecontenedor -e MONGO_INITDB_ROOT_USERNAME=usuario -e MONGO_INITDB_ROOT_PASSWORD=contrasenia mongo
```

_los parametros son los siguientes:_

- `-v` crea un volumen para que nuestra data quede almacenada aun si se elimina el contenedor.
- `-p` superpone el puerto donde lo queremos.
- `-d` corre el contenedor en 2do plano.
- `-e` manda variables de entorno.
- `--name` nombre que queremos que tenga nuestro contenedor.
- `--network` escoge una red de contenedores

_Listo! Creamos nuestro contenedor que esta alojado en el puerto 27017_

_ya solo creamos nuestra database y una coleccion en mongo en nuestro ide favorito_

<a name="react"></a>

#### React

_Para la aplicacion web hice una calculadora simple que guarde los operandos, simbolo, resultado y fecha._

_El codigo lo puede consultar en el siguiente enlace:_

- [Calculadora simple React](https://github.com/davdprado/-SO1-Practica1/tree/main/front/calculadora)

<a name="go"></a>

#### Go

_Para la APIREST utilize go, el cual sera encargado de recibir, enviar y guardarlos en la base de datos; puede consultar el codigo en el siguiente [enlace](https://github.com/davdprado/-SO1-Practica1/tree/main/go)._

---

<a name="images"></a>

## Creando nuestras Imagenes

_Para poder correr contenedores con nuestras aplicaciones hechas necesitamos hacerlas imagenes para que posteriormente cualquiera las pueda utilizar como contenedores, para eso es necesario subir nuestras imagenes a [Docker Hub](https://hub.docker.com/) asi mismo crear un usuario para subir nuestras imagenes._

### Imagen de nuestra API

_Como utilizamos go para nuestra APIREST tendremos que crear nuestra imagen y subirla a [DockerHub](https://hub.docker.com/)._

_Ya que hicimos nuestra api procedemos a irnos a nuestra carpeta raiz y crear un archivo llamado [Dockerfile](https://github.com/davdprado/-SO1-Practica1/blob/main/go/Dockerfile) el cual es el encargado de crear nuestras imagenes; El codigo de nuestro dockerfile para go seria el siguiente:_

```
# aqui se pone la imagen que queremos utilizar
FROM golang:1.17.1-alpine3.14

#le decimos que dentro de esa imagen cree una carpeta llamada /backen/go o como ud prefiera
RUN mkdir -p /backend/go

#le decimos que se dirija a la carpeta que acabamos de crear
WORKDIR /backend/go

#copiamos todo lo que esta en este momento en nuestro directorio raiz
COPY . .

#y que corra el siguiente comando en la terminal
RUN go mod download

#el puerto donde queremos que se este ejecutando
EXPOSE 8080

#y que inicie la aplicacion
CMD ["go", "run", "main.go"]
```

_Con esto ya tenemos nuestro dockerfile, ahora falta crear nuestra imagen que es muy simple, seguimos en la carpeta raiz donde esta nuestro dockerfile y corremos el siguiente codigo:_

```
docker build -t nombredelaimagen .
```

_con esto crea nuestra imagen localmente ahora solo falta subirla al hub_

```
docker login
```

_este comando nos iniciara la sesion de docker hub y nos pedira credenciales, a continuacion corremos lo siguiente:_

```
docker tag nombredelaimagen usuarioDockerHub/nombredelaimagen

docker push usuarioDockerHub/nombredelaimagen
```

_y ya con eso subimos nuestra imagen a dockerhub_

_Para el caso de nuestra imagen en react se procede a hacer los mismos pasos solo que en este caso se utilizara la imagen [Nginx](https://hub.docker.com/_/nginx).\_

<a name="compose"></a>

## Docker-Compose

_El docker compose es un documento que nos simplifica correr varios comandos en consola ya que este se simplifica bastante, para esto es necesario instalar docker-compose._

```
sudo curl -L "https://github.com/docker/compose/releases/download/1.27.4/docker-compose-$(uname -s)-$(uname -m)" -o /usr/local/bin/docker-compose
```

_Ahora se le agregan los permisos necesarios para ejecutarlo:_

```
sudo chmod +x /usr/local/bin/docker-compose
```

_se crea un documento llamado docker-compose.yml, donde iran nuestras instrucciones para correr nuestros comandos necesarios puedes consultar el [docker-compose](https://github.com/davdprado/-SO1-Practica1/blob/main/docker-compose.yml) de este repositorio para ver un ejemplo._

_Para correr el docker-compose es necesario ir a la carpeta donde se encuentra el archivo y correr el siguiente comando_

```
docker-compose up
```

_El docker compose up hace todas las instrucciones del docker-compose.yml y las ejecuta en un contenedor, para apagar ese contenedor solo se corre el siguiente comando:_

```
docker-compose down
```

## Notas

_para mayor informacion sobre el uso correcto del dockerfile y el docker-compose se recomienda ver la documentacion oficial_

## Construido con 🛠️

- [Docker](https://docs.docker.com/engine/install/ubuntu/) - Herramienta para los Contenedores.
- [VSCODE](https://code.visualstudio.com/) - Herramienta de trabajo.
- [React](https://es.reactjs.org/) -Herramienta para nuestra aplicacion Web
- [Golang](https://www.google.com/url?sa=t&rct=j&q=&esrc=s&source=web&cd=&cad=rja&uact=8&ved=2ahUKEwjXmpfb_Jj2AhWMVTABHb6CDDwQFnoECAMQAQ&url=https%3A%2F%2Fgo.dev%2F&usg=AOvVaw2GopNGe_pMsOKwFeS7coZ3) - Herramienta para nuestra APIREST
- [Mongo](mongo) -Gestor de base de datos

_Todo esto fue hecho en Linux/Ubuntu version 20.04 lts_

## Referencias

- [Taller de Docker](https://github.com/sergioarmgpl/taller-docker)

- [Golan y Mongo](https://programmerclick.com/article/53431054920/)

## Autores ✒️

- **David Roberto Diaz Prado** - [GithubProfile](https://github.com/davdprado)
  _201807420_
