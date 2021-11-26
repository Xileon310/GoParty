# Task Runner

Se usará Task como task runner del proyecto.

## Sobre Go

El lenguaje de programación elegido usa un task runner implícito (go). Gracias a este task runner podemos construir, instalar programas e incluso ejecutar tests, pero tiene unos límites, por lo que usaremos un nuevo task runner para superarlos.

## Alternativas

Estas limitaciones pueden ser cubiertas por diferentes aplicaciones, en concreto podemos destacar GNU make, Task, go-task, Realize y Mask.

Todas podrían resolver el problema para este proyecto, pero en proyectos más amplios alguna podría presentar más ventajas que las demás.

### GNU make

Se trata de la utilidad más popular para automatizar las tareas de un proyecto. Es muy simple de utilizar (simplemente invocando a ```make```), pero es bastante complicado de implementar de forma ordenada para proyectos considerablemente grandes.

Este task runner usa un archivo conocido como Makefile, y gracias a su [documentación](https://www.gnu.org/software/make/manual/make.html), sabemos que esta herramienta tiene muchas herramientas para automatizar procesos como edición de cadenas, condiciones (if/else), bucles, funciones, etc.

Normalmente el problema radica en la pobre estructura del código de las tareas, que suelen mezclar bash con el lenguaje que estemos usando en ese momento.

<center><img src="https://tsh.io/wp-content/uploads/2021/04/make-syntax.png" alt="Sintaxis Makefile"></center>

### Task

Task es un task runner que tiene como objetivo ser simple y fácil de utilizar, al contrario de GNU make. 

Esta herramienta está construida en go y su fichero principal Taskfile.yml se basa en un formato YAML que facilita la estructuración de la información. Una vez creado dicho fichero, podremos ejecutar la utilidad con la orden ```task``` y argumentos que queramos usar. 

Sus características principales son:
* Instalación sencilla, se agrega un binario a nuestro $PATH. También se puede instalar usando Homebrew, Snapcraft y Scoop.
* Multiplataforma, se puede usar en Windows, MacOs y Linux.
* Excelente generación de código, puedes evitar que una tarea no se ejecute si un grupo de archivos no han cambiado desde la última vez.

<center><img src="https://tsh.io/wp-content/uploads/2021/04/taskfile-syntax.png" alt="Sintaxis Taskfile.yml"></center>

### go-task

Realmente hay poca información acerca de este task runner, podemos destacar que las tareas pueden agruparse y ejecutar el programa simplemente para el grupo elegido, o principalmente ejecutar el programa en la función que queramos.
Por último añadir que se ejecuta en la propia línea de órdenes con ```go-test```, no es necesario que haya un fichero en nuestro directorio como los anteriores.

No ha sido elegido por su pobre documentación y escasa utilidad dentro de sus opciones.

### Realize

Realize, al igual que Task, está construido en Go y para su utilización necesitamos disponer de un fichero en formato YAML, en este caso .realize.yml.

Podemos destacar las siguientes características:
* Capacidad de gestión de varios proyectos al mismo tiempo. Se pueden mirar por extensiones y rutas personalizadas.
* Variables de entorno personalizadas para el proyecto.
* Capacidad de ejecutar comandos personalizados antes y después de que un archivo cambie o de forma global.
* Capacidad de exportar registros y errores a un archivo externo.
* Panel rediseñado que muestra errores de construcción, salidas de consola y advertencias.

En este caso no se ha elegido este task runner por la complejidad que tiene el archivo .realize.yml. Este archivo tiene una larga extensión porque almacena configuraciones que en este proyecto ni nos molestaremos en usar, la utilidad ```realize start``` lo crea y lo inicializa y con utilidades como ```realize add <project>``` podemos añadir proyectos sin editar el .realize.yml.
Además, dispone de un panel rediseñado que tampoco usaremos en nuestra aplicación web.

### Mask

Mask es un CLI task runner que se define en un archivo de tipo markdown. La ventaja que tiene este tipo de formato es que se producirá un destacado de sintaxis por bloques en muchos editores y renderizadores, lo que ayudará al usuario a identificar cada parte del código. El fichero se llama maskfile.md.

Podemos destacar las siguientes características:
* Argumentos posicionales, se definen entre paréntesis y más tarde se inyectan en los scripts como variables de entorno.
* Flags, que se trata de variables que puedes definir incluso con comandos del sistema u otros scripts y que podrás usar en diferentes fases. También la misma variable tendrá una descripción de para qué se usa. (Realmente esto no es más que una variable que coge un valor y un comentario en la misma variable, solo que en este caso, **mask se encarga de validarlo**).
* Subcomandos, que se rigen por la jerarquía de **#** (## para un comando, ### para subcomandos dentro de ese comando).
* Soporte para otros lenguajes de scripting, lo que da mayor versatilidad a tu código.
* Multiplataforma, en el mismo comando se pueden añadir órdenes para bash y powershell, lo que hace mucho más portable tu proyecto.
* Posibilidad de ejecutar mask dentro de un comando que ya esté en maskfile.md.
* Variables de entorno útiles que maskfile añade automáticamente para facilitar tareas.
* Facilidad para documentar el código

En este caso Mask hubiese sido una buena opción a usar por su facilidad para documentar código principalmente, pero prefiero la estructura que te da un fichero YAML a la hora de leer la configuración a usar.

## Motivación

Una de las principales limitaciones es la ejecución de varios tests simultáneos, con la utilidad "go test 'ruta-test'" podemos ejecutar un test, pero en el caso de varios tests sería molesto especificar cada ruta relativa desde la raíz de nuestro proyecto. Este problema también puede ser llevado igual a las limitaciones para construir el programa "go build 'ruta-src'".

Task necesita de un fichero para realizar estas acciones, este fichero irá creciendo a medida que crezca nuestro proyecto.

Principalmente se ha elegido Task por la extensa documentación y la facilidad con la que es posible utilizar la aplicación. Este proceso se realiza gracias al Taskfile.yml.

## Uso

Su uso principal se basa en un archivo con formato YAML llamado Taskfile.yml que se encuentra en la raíz de nuestro proyecto. En este fichero se incluirán las líneas de texto necesarias para referenciar a los directorios que contengan nuestros archivos y los comandos necesarios para realizar la acción pertinente.

## Gestor de dependecias

Principalmente las dependencias se encuentran en el archivo go.mod, y el propio task runner de go es el que lo edita cuando nosotros hacemos uso de los comandos de go para manejo de dependencias.
