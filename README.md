# BocaFast
Aplicacion destinada a la reserva de ciertos alimentos para centros publicos como institutos o universidades

Estructura de directorios:
Esta estructura creada esta basada en las buenas practicas de Golang. 
- El archivo "main.go" esta almacenado en cmd/server y sera donde estará la lógica de arranque de la aplicación.
- La carpeta config es la encargada de almacenar variables de entorno y configuraciones de la aplicación.
- En la carpeta internal se almacerá la lógica de negocio (entidades, repositories, services...).
- En la carpeta pkg se almacenará "código reutilizable", como conexión a base de datos o utilidades varias.
- En la carpeta test se indicarán los test unitarios

Paquetes instalados:
- Gin: Gestion API
- UUID: Ofrece la posibilidad de obtener un ID unico