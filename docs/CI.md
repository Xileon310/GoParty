# Integración Continua

## Requisitos

Para realizar la tarea de ejecución de tests se deben elegir dos sistemas de Integración Continua que satisfagan los siguientes requisitos:

- **Docker Hub**: Los sistemas deben ser capaces de hacer uso de imágenes de docker que estén subidas a Docker Hub de forma sencilla. Esto facilitará la puesta en marcha del contenedor actualizado cada vez que se modifique el código y que será usado para las pruebas del proyecto.
- **Github Checks**: Los sistemas deben tener soporte para gestionar el workflow desde Github y obtener mensajes de información del mismo.
- **Facilidad de uso**: El sistema debe ser fácil de configurar y que no requiera instalación en un servidor.
- **Gratuito**: Puesto que este proyecto es completamente docente se buscará un sistema que no sea de pago o en su defecto, que contenga una versión freemium o para estudiantes.
- **Matrix jobs**: Se usará el sistema de matrix para ejecutar los tests en las dos versiones de Go que actualmente están bajo soporte (1.16 y 1.17) y que pueden ser consultadas [aquí](https://endoflife.date/go). Cada versión "major" o estable de Golang tiene soporte hasta que dos nuevas versiones "major" son lanzadas.

## Sistemas de Integración Continua

Se ha hecho una comparativa entre los sistemas de Integración Continua más famosos y a continuación se muestran las diferencias más relevantes.

- **Jenkins**: Proyecto open-source escrito en Java, tiene una interfaz muy intuitiva disponible en Linux, Windows y Mac. Tiene más de 1000 plugins para soportar la integración con muchas herramientas. Su mayor defecto y por lo que no es candidado para ser elegido es porque necesita instalación.
- **TeamCity**: Desarrollado por el equipo de Jetbrains es uno de los sistemas referentes en la industria de la Integración Continua. Tiene una versión gratuita para pequeños proyectos con no más de 100 configuraciones de construcción. Principalmente no ha sido elegido para este proyecto porque para usarse con golang hay que realizar diversas configuraciones que pueden llegar a ser tediosas (añadir el soporte manualmente desde su página web, editar el Makefile con los comandos necesarios, añadir el flag -json cada vez que se use ```go test```, etc.).
- **Travis CI**: Otro de los proyectos referentes en la industria, se trata de un servicio de IC con una fácil y rápida configuración que puede ser sincronizada fácilmente con Github. Tiene una API muy útil y una interfaz de línea de comandos. No ha sido elegido para este proyecto puesto que la versión gratuita necesita de la introducción obligatoria de una tarjeta de crédito y los créditos gratuitos proporcionados son solo válidos durante 30 días.
- **Circle CI**: Este proyecto tiene una configuración muy simple y permite la sincronización directa con Github. Su versión gratuita no tiene ningún requerimiento y los créditos proporcionados se renuevan cada mes. Es un posible candidato para la elección del sistema en este proyecto.
- **Github Actions**: Evidentemente tiene una sincronización con Github total y permite almacenar secretos para el inicio de sesión de Docker. Su utilización es sencilla y es un posible candidato para la elección del sistema en este proyecto.
- **Buildbot**: Algo menos conocido que los anteriores pero también muy útil, es un framework de Integración Continua que automatiza la compilación y realización de tests en los proyectos. Es un proyecto open-source y tiene soporte para varias plataformas de testeo con diferentes configuraciones. No ha sido elegido para este proyecto principalmente porque requiere de instalación (aunque es sencilla, puesto que solo se debe instalar el maestro y el trabajador con el instalador de python ```pip install buildbot``` y ```pip install buildbot-worker```).

## Conclusión

Una vez analizados algunos de los principales sistemas de IC, se ha elegido Circle CI y Github Actions puesto que son los únicos que cumplen todos los requisitos anteriormente mencionados.

## Tareas

- **Github Actions**: Se ponen en marcha los tests comprobando el funcionamiento en las dos versiones de Go elegidas haciendo uso de una job matrix.
- **Circle CI**: Se ponen en marcha los tests haciendo uso del task runner en el contenedor de pruebas Docker que se encuentra en Docker Hub.