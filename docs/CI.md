# Integración Continua

## Requisitos

Para realizar la tarea de ejecución de tests se deben elegir dos sistemas de Integración Continua que satisfagan los siguientes requisitos:

- **Docker Hub**: Los sistemas deben ser capaces de hacer uso de imágenes de docker que estén subidas a Docker Hub de forma sencilla. Esto facilitará la puesta en marcha del contenedor actualizado cada vez que se modifique el código y que será usado para las pruebas del proyecto. Puesto que en el anterior objetivo, se usó dicho registro para almacenar la imagen que se irá actualizando a lo largo del proyecto, se buscarán CIs que sean compatibles.
- **Github Checks**: Los sistemas deben tener soporte para gestionar el workflow desde Github y obtener mensajes de información del mismo.
- **Facilidad de uso**: El sistema debe ser fácil de configurar y que no requiera instalación en un servidor. Esto no es un requisito obligatorio, pero facilitará en gran parte el despliegue de la misma. Respecto al servidor, el motivo principal es que debido a los recursos que se poseen, es preferible usar servidores externos que estén funcionando siempre.
- **Gratuito**: Ya que este proyecto es completamente docente se buscará un sistema que no sea de pago o en su defecto, que contenga una versión freemium o para estudiantes.


## Sistemas de Integración Continua

Se ha hecho una comparativa entre los sistemas de Integración Continua más famosos, todos ellos serían compatibles con nuestro proyecto, puesto que satisfacen los anteriores requisitos. Esta sección se desglosará en diferentes apartados que se han considerado oportunos para la elección.

Los CIs considerados son:
- **Travis CI**: Uno de los proyectos referentes en la industria, se trata de un servicio de CI con una fácil y rápida configuración que puede ser sincronizada fácilmente con Github. Tiene una API muy útil y una interfaz de línea de comandos. No ha sido elegido para este proyecto puesto que la versión gratuita necesita de la introducción obligatoria de una tarjeta de crédito y los créditos gratuitos proporcionados son solo válidos durante 30 días. Por lo que no aparecerá en las siguientes secciones ya que no se ha podido tener acceso a él, pero en una comparativa se considera fundamental su mención, puesto que es uno de los más famosos.
- **Circle CI**: Otro de los proyectos más famosos en la industria, con una versión gratuita sin requerimientos y fácil configuración.
- **Github Actions**: Una alternativa simple para la CI de nuestro proyecto, gratuito y sin depender de registro en plataformas externas a Github.
- **Git Lab**: Además de CI sirve como gestor de repositorios, así como alojamiento de wikis y sistema de seguimiento de errores, entre otras cosas. Por último, destacar que es de código abierto.
- **Semaphore CI**: Destaca por su gran cantidad de plantillas creadas para la realización del despliegue de tu proyecto.
- **Buddy**: Quizás sea el menos famoso de la lista, pero se trata de un sistema de Integración Continua con mucho potencial.

A continuación, se compararán y se ordenarán de mayor a menor importancia cada CI en diferentes categorías oportunas.

### Precio

Como se ha comentado anteriormente, todos son gratuitos o en su defecto, tienen versiones freemium o para estudiantes de las que nos podemos beneficiar. Principalmente todos tienen la misma importancia en este proyecto, a excepción de Semaphore CI y Git Lab, que te limitan los minutos que se tienen para CI. Esto se traduce en un gran inconveniente de cara a la asignatura, pues nos interesaría tener el mismo CI hasta acabar la misma.

El orden quedaría de la siguiente forma:
1. Circle CI = Github Actions = Buddy
2. Git Lab
3. Semaphore CI

La principal diferencia entre el segundo y tercer puesto es que Git Lab nos proporciona 400 minutos en caso de un proyecto privado (solo accesible para nosotros) y 50.000 minutos o 34 días para proyectos públicos, mientras que Semaphore CI 1300 minutos en cualquier caso.

Por otro lado, cualquiera de las alternativas nos permiten los suficientes jobs (acciones) para satisfacer los objetivos de la asignatura, ya que la mayoría presenta restricciones en esta parte, por ejemplo, Buddy sólo permite 5 proyectos, mientras que Semaphore CI sólo permite un job al mismo tiempo.

### Interfaz

Sin duda, uno de los principales motivos para la elección de un CI ha sido que tuviese una interfaz intuitiva, es decir, que siendo un usuario novato del propio CI, pudiese encontrar fácilmente todo lo que me propusiese.

En este caso, el orden ha quedado de la siguiente forma:

1. Buddy
2. Circle CI
3. Git Lab
4. Semaphore CI
5. Github Actions

La interfaz más intuitiva (y agradable a la vista) es la de Buddy, seguida de la de Circle CI, aunque esta no es más agradable que la de los demás, pero sí destaca por ser muy sencilla. Por otro lado, la interfaz de Semaphore CI es bastante confusa a la hora de moverse y encontrar diferentes opciones, sobre todo al usar su workflow (aunque es muy útil, algo que se tratará en el siguiente punto). Por último, Github Actions ni siquiera tiene una interfaz gráfica intuitiva que pueda competir con los demás.

### Facilidad de Uso

Puede parecer que este punto y el anterior son iguales, pero no. Se valora la facilidad con la que se ha podido crear el workflow en cada caso, puesto que se ha probado para todos.

Respecto a la facilidad de uso, el orden ha quedado:

1. Buddy
2. Semaphore CI
3. Circle CI
4. Git Lab
5. Github Actions

Además de tener una interfaz que te guía en todo el proceso, se puede hacer absolutamente todo desde la interfaz del proyecto en Buddy, sin necesidad de crear un archivo YAML (te lo crea automáticamente, aunque se puede usar su Web UI, lo que te permite usar Buddy sin añadir el archivo de configuración al repositorio). En este caso, se podría repetir la información anterior para Semaphore CI, pero está en el segundo puesto porque me ha resultado difícil de comprender la gestión del workflow mediante bloques (aunque podría haberse creado un .yml normal). Los demás tienen una configuración por fichero YAML, en la que hay que destacar que la integración con Git Lab ha sido la más tediosa de todas, mientras que Circle CI y Github Actions se han adaptado correctamente y de forma sencilla al proyecto.

Todos tienen plantillas para facilitar muchas acciones comunes a la mayoría de usuarios. Destacaríamos en ese caso Git Lab porque tiene un mayor número de plantillas, aunque eso no signifique que sean mejores. Dejando en últimos lugares a Semaphore CI y Github Actions.

### Performance

Aunque parece algo insignificante, puesto que el despliegue de una aplicación realmente depende de su peso y procesos que se deban ejecutar, en la puesta a punto de cada CI se ha detectado una velocidad a la hora de ejecución diferente (en la mayoría de casos insignificante).

1. Circle CI
2. Semaphore CI
3. Buddy
4. Github Actions 
5. Git Lab

Realmente los 4 primeros CIs tienen velocidades muy parecidas, pero la sorpresa ha sido Git Lab. En su página web avisan de que los cambios realizados en el repositorio podrían tardar hasta 180s en reflejarse a la hora de saltar el workflow definido en el mismo, aunque en este proyecto no ha sido el caso, sí hemos visto tiempos de hasta 45 segundos entre un commit hecho y la realización de los tests.

### Documentación

Es difícil responder a las preguntas: ¿Es buena esta documentación? ¿Este CI tiene mejor documentación que otro? Por lo que esta clasificación será la más subjetiva de todas, y se basará en la facilidad para encontrar información que me haya permitido desarrollar los workflows correspondientes en cada CI.

1. Circle CI
2. Semaphore CI
3. Github Actions
4. Buddy 
5. Git Lab

Pese a que Buddy destaca en otros aspectos, su documentación me ha decepcionado,ya que es bastante pobre. Por el lado contrario, Circle CI tiene una gran documentación, sobre todo respecto a la integración con Github y la Checks Api.

### Tabla comparativa con resumen

A continuación, se mostrará una tabla que resume las anteriores secciones:

| **CI** 			| **Precio** | **Interfaz** | **Facilidad de Uso** | **Performance** | **Documentación** |
| ------------------| ---------- | ------------ | -------------------- | --------------- | ----------------- |
| Circle CI 		| 1 		 | 2 			| 2 				   | 1 				 | 1 				 |
| Semaphore CI 		| 3 		 | 4 			| 3 				   | 2 				 | 2 				 |
| Buddy 			| 1 		 | 1 			| 1 				   | 3 				 | 4 				 |
| Github Actions 	| 1 		 | 5 			| 4 				   | 4 				 | 3 				 |
| Git Lab 			| 2 		 | 3 			| 5 				   | 5 				 | 5 				 |

Si hiciésemos una media de los valores anteriores (siendo 1 el mejor y 5 el peor) quedaría:

Circle CI: 1.4
Semaphore CI: 2.8
Buddy: 2
Github Actions: 3.4
Git Lab: 4

Como podemos apreciar, los CIs que tienen mejores puntuaciones son Circle CI y Buddy. Principalmente en el caso de proyectos sencillos o que nos preocupase el hecho de registrar cuentas en plataformas externas a nuestro gestor de repositorios la mejor integración podría ser Github Actions (considerando que no queremos un sistema que dependa de instalación en servidor). Semaphore CI tiene un muy buen workflow al igual que Buddy, pero su interfaz no es tan buena e intuitiva (pero sí efectiva si se conoce la herramienta).

## Requisitos del objetivo

Puesto que se elaboró el objetivo antes de que estableciese el requisito de no usar Circle CI por su popularidad entre los alumnos (en ordinaria), he dejado el archivo correspondiente. Por lo que se han usado 3 CI (aunque 2 de ellos hacen lo mismo).

La primera opción fue usar Git Lab (como se puede ver en el progreso del repositorio), pero no fui capaz de hacerlo funcionar (y su documentación no ayudó), por lo que pensé en Travis CI, pero no pude usarlo porque al realizar el registro pedía un parámetro (VAT ID) que no tenía. Por último, utilicé Github Actions ya que Semaphore CI da muy pocos minutos (poco más de 20h).

**NOTA:** Por alguna extraña razón, simplemente haciendo pull del contenedor en Git Lab y Buddy, me da error de task, mientras que en Circle CI no pasa. Por lo que se ha optado a que Buddy simplemente ejecute la versió 1.16.13.

## Conclusión

Una vez analizados algunos de los principales sistemas de IC, se ha elegido **Circle CI**, **Github Actions** y **Buddy** para la elaboración de los workflows de la asignatura.

Para la integración y elaboración en cada CI se han seguido las siguientes guías: [Circle CI](https://circleci.com/docs/2.0/enable-checks/) y [Buddy](https://buddy.works/guides/golang).

## Versiones utilizadas

Para este objetivo nos interesa principalmente testear 2 versiones de go puesto que su ciclo de vida de actualizaciones es de 2 versiones, es decir, actualmente sólo tienen soporte 1.16 y 1.17 (y derivadas), por lo que cuando salga la versión 1.18, la versión más antigua (1.16) dejará de tener soporte. Estas versiones se pueden comprobar [aquí](https://go.dev/doc/devel/release).

Las versiones usadas serán entonces:
- 1.16.13
- 1.17.6 (usada en la imagen de Docker)

Puesto que son las últimas versiones derivadas de las principales (1.16 y 1.17 respectivamente).

## Tareas

- **Buddy**: Se ponen en marcha los tests comprobando el funcionamiento en la versión 1.16.13 (última versión de golang de 1.16). En este caso es una tarea muy simple, pues sólo hay que asignar el proyecto de Github en Buddy, crear un workflow y especificar la versión que se desea (por defecto testea el proyecto). De querer hacerlo con más versiones, podríamos ejecutarlas todas concurrentemente.
- **Circle CI**: Se ponen en marcha los tests haciendo uso del task runner en el contenedor de pruebas Docker que se encuentra en Docker Hub. En este caso la imagen se basa en la imagen oficial de golang con versión 1.17.6, usando la distribución de Linux alpine.
- **Github Actions**: Se ponen en marcha los tests comprobando el funcionamiento en la versión 1.16.13 (última versión de golang de 1.16). En este caso es una tarea muy simple, pues sólo hay que asignar el proyecto de Github en Buddy, crear un workflow y especificar la versión que se desea (por defecto testea el proyecto). De querer hacerlo con más versiones, podríamos ejecutarlas todas concurrentemente.

## Problemas en la resolución del objetivo

El principal problema que se ha encontrado en la resolución ha sido que Circle CI ha funcionado de forma diferente a los demás CIs (Buddy o GitLab), cuando se hace uso de la imagen cargada en DockerHub, en los dos últimos CIs el task runner no encuentra el fichero Taskfile. Sinceramente he investigado todo lo que he podido pero no he encontrado solución a esto. Por otro lado la alternativa que había planteado es correr los tests de forma normal con el comando ```docker run -t -v `pwd`:/app/test xileon/goparty```, que en esencia es lo que realiza el task runner en el Taskfile.

Esta alternativa ha tenido diferentes problemas en los CIs:
- **Buddy**: la herramienta **docker** no estaba instalada en la imagen, por lo que se ha intentado instalar en tiempo de ejecución (algo como última alternativa, pues retrasaba la ejecución de los tests de forma notoria). Una vez instalado, se intentaba levantar con los gestores de procesos (service / systemctl), pero en ningún caso se conseguía, por lo que no se ha podido ejecutar.
- **Git Lab**: para la ejecución de comandos de Docker en Git Lab hay varias alternativas, pero la única que podría valer es la de instalar GitLab Runner en nuestro proyecto, que es un programa de CI/CD que envía la información a Git Lab. Ante la imposibilidad de realizar la integración sin necesidad de instalar ningún programa externo, se ha descartado.