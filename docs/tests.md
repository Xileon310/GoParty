# Tests unitarios

Para la realización de los tests unitarios que se piden en el objetivo se ha creado un fichero denominado actividad_test.go en el que se declaran los tests correspondientes a la entidad Actividad y al objeto valor Zona. Se ha seguido el estándard de go que puede ser consultado [aquí](https://go.dev/doc/tutorial/add-a-test) en el que se hace uso de la biblioteca testing y se define la sintaxis que deben seguir las funciones: `TestXxxXXX (t *testing.T)`, entre otras cosas.

Por otro lado se ha seguido el convenio de valor-error propio de go en las clases Actividad y Zona. Las funciones de creación de un objeto realizarán diversas comprobaciones, en el caso de haber un error, se devolverá el objeto y el error que ha ocurrido. En caso de que la ejecución no tenga ningún problema, se devolverá `nil`, que será el valor que se tomará de referencia para las comprobaciones en los tests.

## Uso de Task

Gracias al task runner elegido en el objetivo 3, podemos automatizar la ejecución de tests con la siguiente orden:
	`task test`

Nos aparecerá por pantalla la ejecución de cada test y su duración.

Por último añadir que nuestro archivo actividad_test.go se encuentra en `/pkg/actividad/actividad_test.go` y en el archivo Taskfile.yml se encuentra la ruta `/pkg/...`. Es por ello que el anterior comando funciona correctamente.