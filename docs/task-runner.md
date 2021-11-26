# Task Runner

Se usará Task como task runner del proyecto.

## Sobre Go

El lenguaje de programación elegido usa un task runner implícito (go). Gracias a este task runner podemos construir, instalar programas e incluso ejecutar tests, pero tiene unos límites, por lo que usaremos un nuevo task runner para superarlos.

## Alternativas

Estas limitaciones pueden ser cubiertas por diferentes aplicaciones, en concreto podemos destacar make, Task o go-task.

Todas podrían resolver el problema para este proyecto, pero en proyectos más amplios alguna podría presentar más ventajas que las demás.

## Motivación

Una de las principales limitaciones es la ejecución de varios tests simultáneos, con la utilidad "go test 'ruta-test'" podemos ejecutar un test, pero en el caso de varios tests sería molesto especificar cada ruta relativa desde la raíz de nuestro proyecto. Este problema también puede ser llevado igual a las limitaciones para construir el programa "go build 'ruta-src'".

Task necesita de un fichero para realizar estas acciones, este fichero irá creciendo a medida que crezca nuestro proyecto.

Principalmente se ha elegido Task por la extensa documentación y la facilidad con la que es posible utilizar la aplicación. Podemos indicar y enviar variables a nuestro proyecto, de esta forma no tendríamos que estar creando actividades en el terminal de forma tediosa, simplemente añadiríamos los datos a un fichero y usaríamos dichas variables para ello. Este proceso se realiza gracias al Taskfile.yml.

## Características de Task

* Instalación sencilla, se agrega un binario a nuestro $PATH. También se puede instalar usando Homebrew, Snapcraft y Scoop.
* Multiplataforma, se puede usar en Windows, MacOs y Linux.
* Excelente generación de código, puedes evitar que una tarea no se ejecute si un grupo de archivos no han cambiado desde la última vez.

## Uso

Su uso principal se basa en un archivo con formato YAML llamado Taskfile.yml que se encuentra en la raíz de nuestro proyecto. En este fichero se incluirán las líneas de texto necesarias para referenciar a los directorios que contengan nuestros archivos y los comandos necesarios para realizar la acción pertinente.

## Gestor de dependecias

Principalmente las dependencias se encuentran en el archivo go.mod, y el propio task runner de go es el que lo edita cuando nosotros hacemos uso de los comandos de go para manejo de dependencias.
