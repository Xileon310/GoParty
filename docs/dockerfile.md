# Contenedor de Docker

## Motivación

Gracias a docker podemos tener un entorno fácilmente controlable para la ejecución de nuestra aplicación agilizando la ejecución de tests.

## Integración con Task

Para facilitar al usuario el uso de la aplicación, se ha automatizado la ejecución de los tests en el task runner. Esto se ha conseguido editando el fichero Taskfile.yml y añadiendo una nueva directiva que se invoca con la orden ```task docker``` y que es idéntico a utilizar ```docker run -t -v `pwd`:/app/test xileon/goparty```.

## Imagen utilizada

Podría usarse cualquier imagen para la creación, puesto que se puede arrancar cualquier versión, hacer una instalación limpia del lenguaje golang y las dependencias necesarias y bastaría, pero en este caso, se ha optado por usar una de las imágenes oficiales de **golang** para mayor comodidad (ya que la diferencia de eficiencia en nuestra aplicación es imperceptible). Una de las posibilidades es usar scratch o centurylink, y el tamaño de la imagen de docker será igual al tamaño de nuestro programa, aunque al ser imágenes básicas, se debe compilar estáticamente la aplicación con las librerías que necesitemos, lo que se traduce en mayor complicación para el usuario.

Para la creación de la imagen que usará la aplicación había diferentes opciones oficiales que cumplen los requisitos:
- **Imagen basada en Debian**: Esta es la imagen por defecto que debería usar cualquier proyecto go, si no importa el tamaño final de la imagen (esto es importante, pues puede afectar negativamente en el despliegue)
- **Imagen basada en Alpine**: Esta imagen es mucho más recomendada cuando se trata de minimizar el tamaño de la imagen final.

La principal diferencia entre las anteriores imágenes es que la imagen basada en Alpine no hace uso de las librerías ```glibc```y derivados. En su lugar usa ```musl libc```, lo que puede derivar en problemas si requerimos de la primera opción para nuestro programa. Puesto que nuestra aplicación no presenta ningún problema referente a ello, se ha optado por Alpine como imagen, en concreto con Alpine 1.17, aunque sería indiferente hacer uso de la 1.16 o 1.18 también.

## Dockerfile

El archivo Dockerfile contiene los siguientes datos:
- **Imagen**: Se ha usado la imagen oficial de golang, que está basada en alpine.
- **Label**: Se especifica el correo de la persona que mantiene el proyecto.
- **Creación de usuario**: Para evitar el uso de nuestro programa con privilegios de administrador (con los problemas que ello conlleva), se ha creado un usuario, en este caso llamado **goparty** y se cambia a él para hacer uso de la aplicación de forma segura.
- **Directorio de trabajo**: Se ha especificado el directorio de trabajo en el que estará el código (en este caso /app/test). Gracias a la orden WORKDIR, si este directorio no existe, se creará automáticamente.
- **Instalación de dependnecias**: Se instalan las dependencias recogidas en el archivo go.mod y también se instala el gestor de tareas, en este caso Task.
- **Acción**: Se ejecutan los test haciendo uso de Task.

### Dependencias

El programa actualmente depende de las siguientes dependencias:
- **Task**: Se trata del task runner que se ha elegido para el proyecto.
- **Dependencias propias del código**: Se reflejan en el archivo go.mod.

### Buenas prácticas

El archivo Dockerfile se ha desarrollado en todo momento intentando seguir lo mejor posible las buenas prácticas:
- Se ha hecho uso de un nuevo usuario que se crea durante el proceso, para no hacer uso de root.
- Se instalan exclusivamente los programas que son necesarios, con diferentes fines:
    - Reducir lo máximo posible el tamaño de la imagen.
    - Agilizar el proceso de despliegue y ejecución de la aplicación.
- Se hace en todo momento uso del task runner, pues facilita al usuario la ejecución de tareas más tediosas.

## Workflow

Se ha creado un Github Action que crea la imagen del contenedor y la publica en Docker Hub de forma automática.
La idea principal de este workflow es que solo construya la imagen cada vez que se hace push a la rama maestra, pues es en el caso en el que el código ha sido aprobado. Para la ejecución de este objetivo también se ha permitido la construcción de la imagen a partir de los Pull Request (principalmente por comodidad).

Para el desarrollo de este archivo se ha seguido la siguiente [guía](https://docs.docker.com/ci-cd/github-actions/), en la que se realizan diversos pasos: creación de tokens, ejemplos del archivo YAML a usar y explicación de los diferentes campos del mismo.

### Anotación

Puesto que se poseía anteriormente cuenta de Docker Hub por motivos docentes de otras asignaturas, el nick de la cuenta de Github no coincide con el de Docker Hub, en este caso son Xileon310 y xileon respectivamente.

