# Tests unitarios

Para la realización de los tests unitarios que se piden en el objetivo se ha creado un fichero denominado actividad_test.go en el que se declaran los tests correspondientes a la entidad Actividad y al objeto valor Zona. Se ha seguido el estándard de go que puede ser consultado [aquí](https://go.dev/doc/tutorial/add-a-test) en el que se hace uso de la biblioteca testing y se define la sintaxis que deben seguir las funciones: `TestXxxXXX (t *testing.T)`, entre otras cosas.

El estándar de Go define por defecto el uso de la biblioteca testing, pero existen más alternativas:
- **Testify**: Permite la definición de características para uso local y generación automática de datos para los tests.
- **gocheck**: Define funciones de aserción, para, por ejemplo, la correcta definición de un objeto.
- **Ginkgo and Gomega**: Se trata de un framework bastante pesado (pero con muchas utilidades) de testeo de BDD y funciones de aserción.
- **GoConvey**: Otro framework de testeo de BDD con una interfaz web de fácil uso.

La elección de testing viene motivada por dos motivos principales:
- **Documentación oficial**: La documentación oficial de go hace uso de ella por su simplez y porque viene ya integrada en el propio entorno de go. Esto promueve que sea fácil de utilizar (ya que es fácilmente localizable en la documentación de go) y que no sea necearia instalaciones adicionales.
- **Sencillez del proyecto**: En este caso la sencillez de la aplicación permite el uso de cualquiera de las bibliotecas habilitadas para testeo, pero tiene menor sentido hacer uso de frameworks pesados llenos de utilizadades que no usaremos.

Realmente la diferencia radicaba entre usar Testify y testing para este proyecto, y nos decantamos por testing por seguir los estándares marcados para go (que no quiere decir que sea lo óptimo) y por la sencillez de los tests, además de no tener la necesidad de hacer uso de funciones especiales.

Por otro lado se ha seguido el convenio de valor-error propio de go en las clases Actividad y Zona. Las funciones de creación de un objeto realizarán diversas comprobaciones, en el caso de haber un error, se devolverá el objeto y el error que ha ocurrido. En caso de que la ejecución no tenga ningún problema, se devolverá `nil`, que será el valor que se tomará de referencia para las comprobaciones en los tests.

## Uso de Task

Gracias al task runner elegido en el objetivo 3, podemos automatizar la ejecución de tests con la siguiente orden:
	`task test`

Nos aparecerá por pantalla la ejecución de cada test y su duración.

Por último añadir que nuestro archivo actividad_test.go se encuentra en `/pkg/actividad/actividad_test.go` y en el archivo Taskfile.yml se encuentra la ruta `/pkg/...`. Es por ello que el anterior comando funciona correctamente.

## Principio F.I.R.S.T.

Los tests desarrollados siguen el principio F.I.R.S.T. (Fast, Independent, Repeteable, Self-validating, Timely) por los siguientes motivos:
- **Fast**: Tras la ejecución en repetidas ocasiones de los tests, el tiempo estimado es cercano a 0.001s en una cantidad de 6 tests, por lo que nos arroja una aproximación de una ejecución de 600 tests en 1s.
- **Independent**: Cada test es independiente de los demás.
- **Repeteable**: El resultado de las pruebas es independiente del servidor en el que se ejecuta. En este caso se ha realizado tanto localmente como en el contenedor de Docker del Objetivo 5.
- **Self-validating**: Los tests son completamente compatibles con la ejecución automática que se desarrolla en el Objetivo 5.
- **Timely**: Los tests han sido desarollados antes de la perfilación del código.