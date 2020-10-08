<p align="center">
  <img  alt="docker for dummies" height="128px" width="128px" src="https://miro.medium.com/max/1200/1*i2skbfmDsHayHhqPfwt6pA.png">
</p>

<h3 align="center"> Microservices in GO </h3>

## Objective
- Make a microservices architecture in GO which is ready for production.
- Create a dummy coffee shop for this purpose.

## Environment
- Server: GO
- Client: React
- Database: Not decided yet

## Commands
`go run main.go` - To run the server
`curl -v localhost:9090 -d "World"` - To send API calls from terminal to go runtime

### GET
`curl -v localhost:9090`
### POST
`curl -v localhost:9090 -XPOST -d {"id": 1, "name": "Tea", "description": "Cuppa Tea", "price": 10}`
### PUT
`curl -v localhost:9090/2 -XPUT -d {"name": "Frappuccino", "description": "Cuppa frappuccino", "price": 100}`

## Hosted links
- Will update when hosted

## Production setup
- Will update when setup is complete
