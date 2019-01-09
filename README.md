# prisma-echo-api

Sample API using Echo (Go) and Prisma (with Postgres).

## Getting Started

Here's what you need to do to run this application. This project assumes you have already installed Go, Docker and Docker Compose.

### Install the Prisma CLI

Homebrew:

```sh
brew tap prisma/prisma
brew install prisma
```

NPM:

```sh
npm install -g prisma
```

### Run Docker Compose

This will start the `prisma` and `postgres` images:

```sh
docker-compose up -d
```

### Deploy Prisma

```bash
prisma deploy
```

Your endpoint should now be live at [http://localhost:4466](http://localhost:4466)

## Run

Once everything is installed and deployed, run the application using:

```sh
go run main.go
```

There are four endpoints:

```
POST        /users          Create a new user
GET         /users          Get all users
GET         /users/:id      Get user by their ID
DELETE      /users/:id      Delete user by their ID
```

## Built With

* [Go](https://golang.org/)
* [Echo](https://echo.labstack.com/) - Web framework used
* [Prisma](https://www.prisma.io/) - Database framework used
* [Docker](https://www.docker.com/)