# IV-Project
Repositorio para la asignatura de IV.

## Idea principal

La idea principal es desarrollar una aplicación en la que las personas puedan organizar eventos y conocerse entre sí.

Por otro lado, se pretende establecer una clasificación en diferentes categorías (deportes, actividades culturales, etc.) para facilitar la búsqueda y creación a los usuarios.

Esta aplicación dispondría de un sistema en el que cualquier usuario podría crear una actividad con ciertas características y otros usuarios pudiesen apuntarse a ella, consiguiendo por tanto, que diferentes usuarios se conociesen entre sí y pudiesen realizar la actividad que de otra forma no hubiese sido posible, que a su vez, representa la lógica de negocio de la aplicación.

## Motivación

La motivación principal es el hecho de facilitar a las personas realizar actividades en grupo. En muchas ocasiones una persona decide hacer actividades que no se pueden realizar en solitario (jugar a algún deporte cooperativo) o que no están disponibles en su zona (esquiar, rafting, escalada, etc.) pero no conoce a nadie para desarrollarlas.

Es cierto que esta idea puede ser desarrollada en cualquier red social como Facebook o Instagram, pero de esta forma daríamos una mayor facilidad al usuario, centralizando todas las actividades en la misma aplicación y estableciendo comunicaciones organizador-participante de forma más rápida.

## Documentación

Toda la documentación referente al proyecto se puede encontrar [aquí](docs)

## Issues

La información detallada de las Issues del proyecto se encuentran [aquí](docs/issues.md)

## Elección lenguaje de programación 

Se ha elegido Go como lenguaje de programación para el desarrollo del proyecto porque es un lenguaje que utiliza tipado estático, es sencillo y es eficiente tanto en optimización de la memoria como en velocidad de ejecución.

## Task runner

La información más detallada acerca de la información del task runner elegido se encuentra [aquí](docs/task-runner.md).

Puesto que el lenguaje de programación elegido es Go, se ha buscado un task runner para su ejecución. En este caso se ha elegido [Task](https://taskfile.dev/#/), aunque una alternativa podría haber sido go-task(https://leandroveronezi.github.io/go-task/). Principalmente se ha elegido la primera opción por la extensa documentación y mayor utilidad, ya que go-task es bastante sencillo (aunque podría haber servido para este objetivo perfectamente).

1. Instalación de **Go** --> Se puede hacer desde [aquí](https://go.dev/doc/install).

2. Instalación de Task --> La forma más sencilla es usando un script como el siguiente, que instalará el programa en /usr/local/bin (puede requerir permisos de sudo).

	```	shell

	sh -c "$(curl --location https://taskfile.dev/install.sh)" -- -d -b /usr/local/bin

	```

3. Clonar este repositorio y acceder a él.

4. Usando Task
Una vez que hemos clonado el repositorio y estamos dentro de él, podemos utilizar las siguientes acciones:
* **task install** --> Genera un ejecutable del proyecto en $GOPATH/bin
* **task run** --> Ejecuta el proyecto.
* **task check** --> Comprueba la sintaxis de las entidades programadas.
* **task installdeps** --> Incluye las dependencias que se encuentran en go.mod a la hora de instalar o ejecutar el programa.
* **task test** --> Ejecuta todos los tests del proyecto.	

## Test unitarios

Toda la información referente a los tests unitarios desarrollados para este proyecto se puede encontrar [aquí](docs/tests.md).