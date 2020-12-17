# Bequest Backend Assignment

## Environment

This project is intialized using:

- [Gin](https://github.com/gin-gonic/gin)
- [GORM](https://github.com/jinzhu/gorm)
- [GoDotEnv](https://github.com/joho/godotenv)
- [Docker](https://www.docker.com/)
- [swag](https://github.com/swaggo/swag)

## Quick Start

```bash
# Spin up the gin server and database
$ make run
```

## APIs Endpoint
After spinning up the App, you should be able to see the “pong” message return by navigating to this endpoint: http://localhost/api/ping

## Environment variables

> Do not commit .env file to git. I have done so here is to easily demonstrate my works,
> but in a production system you must keep it privately.

Configure database connection info and secret etc. which already populate in '.env'

## For developers

### Steps
```bash
# Spin up database in dev
$ make start-dev-db

# Go run
$ go run main.go  
```