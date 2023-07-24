Proyecto para la asignatura Arquitectura de Software 2 de la Universidad Catolica de Cordoba.

Desarrollaremos un sistema de publicación de clasificados, mediante el cual las empresas inmobiliarias puedan cargar sus bases de datos mediante el posteo de un
archivo json con la información relacionada a las propiedades, los usuarios/navegantes puedan buscar esos clasificados desde la home del sitio en base a una oración 
y traiga los resultados priorizados que permitan ver el detalle de la publicación y enviar un mensaje con la información de contacto al vendedor.
Para esto, se piden desarrollar 5 microservicios:

1. búsqueda
Get:
- search=:searchQuery ejemplo http://localhost:8081/search=casa

2. usuarios
Post:
  - /user
    ejemplo

      http://localhost:9000/user

      {
          "username": "sofia",
          "email": "sofia@gmail.com",
          "password": "1234"
      }

  - /login
    ejemplo

      http://localhost:9000/login
      {
          "username": "leandro",
          "password": "1234"
      }

GET:
  /user/:id
    ejemplo
      http://localhost:9000/user/1

3. mensajes
POST:
- /messages
      http://localhost:9001/messages

      {
          "userid": 1,
          "itemid": "ABC123",
          "content": "Hola, este es un mensaje de ejemplo.",
          "createdat": "2023-06-02"
      }
  
  GET:
  - /messages/:id
      http://localhost:9001/messages/1
  
4. publicaciones

http://localhost:8090/items

[
  {
   
    "title": "Departamento sin amoblar",
    "userid": 2,
    "image": "sdasd",
    "currency": "USD",
    "price": 48000,
    "state": 1,
    "condition": "Nuevo",
    "address": "obispo trejo"
  },
  {
  
    "title": "Departamento amoblado",
    "userid": 3,
    "image": "asjdkb",
    "currency": "USD",
    "price": 40000,
    "state": 1,
    "condition": "Antiguo",
    "address": "obispo Salguero"
  }
]

5. frontend

El Frontend debia contener la vista de inicio con el input de búsqueda, el listado de Items, el detalle de la publicación.
En la implementacion, el frontend simplemente se comunica con el servicio BÚSQUEDA a traves del request http - GET Query - que se especifico anteriormente,
obtiene la informacion de los items y la muestra, cargando tambien las imagenes correspondientes





Uso general del proyecto:

Debemos correr el siguiente comando en la carpeta frontend
- npm install
- npm start

Luego debemos correr los servicios de memcached, solr, mongo y rabbit en docker

Finalmente los 4 servicios (users, messages, items, search) deben ser corridos en consola:
- go run main.go
