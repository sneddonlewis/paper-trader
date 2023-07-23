# Paper Trader

## Dev Requirements

[Go](https://golang.org/),
[Node.js](https://nodejs.org/),
[Docker](https://www.docker.com/), and
[Docker Compose](https://docs.docker.com/compose/)

## Backend (Development)

```sh
docker-compose -f docker-compose-dev.yml up
cd server
go run server.go
```
Web API served @ http://localhost:8080.
Runs a PostgreSQL database (so make sure port 5432 is free)

## Frontend (Development)
```sh
cd webapp
npm install
npm run watch
```
Angular SPA served @ http://localhost:3000.
 
## Backend (Production)

```sh
docker-compose up
```

## Frontend (Production)
```shell
cd webapp
npm run build
```