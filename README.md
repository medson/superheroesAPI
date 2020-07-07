# Super Heroes API

![Go](https://img.shields.io/badge/go-1.13-blue.svg)

Structure for app :open_file_folder: :octocat:

### Guide

1. [Installing necessary dependencies](#installing-necessary-dependencies)
2. [Running the app](#running-the-app)
3. [Documentation](#documentation)
4. [Tests](#tests)
5. [Decisions](#decisions)

## Installing Necessary Dependencies

First of all, you need to install Docker and docker-compose in case you don't have them already installed. For this, you can access this [link](https://docs.docker.com/install/) and install Docker and [here](https://docs.docker.com/compose/install/) for docker-compose.

## Running the App

Now we can run the app. Go to project root folder, where is located `Makefile` and run `make run-app` to bring services up. Docker may need to build images for the first time, and this operation may take some time to complete. Example:

### Running Makefile

```bash
$ make run-app
Running the application
Creating superheroesapi_postgres_1 ... done
Creating superheroesapi_api_1      ... done
...
...
...
postgres_1  | 2020-07-07 14:22:18.927 UTC [1] LOG:  database system is ready to accept connections
api_1       | 2020/07/07 14:22:20 Database Migrated
api_1       |         _______ __
api_1       |   ____ / ____(_) /_  ___  _____   HOST   0.0.0.0  OS    LINUX
api_1       | _____ / /_  / / __ \/ _ \/ ___/   PORT   3000     CORES 4
api_1       |   __ / __/ / / /_/ /  __/ /       TLS    FALSE    MEM   7.6G
api_1       |     /_/   /_/_.___/\___/_/1.12.4  ROUTES 6        PPID  0
```

### The following addresses can be used to access running services:

-   [http://localhost:3000/api/v1/supers](http://localhost:300/api/v1/supers), for the API
-   [http://localhost:3000/swagger](http://localhost:3000/swagger), for swagger

Response data types and request optional parameters, can be seen at [api/README.md](api/README.md) File.

## Documentation

-   For documentation, all code documentation stay above structs and functions definitions.
-   For API usage, was used the Swagger.io, after run the app you can access the Swagger [here](http://localhost:3000/swagger) to see how to use the application routes.

## Tests

-   For run the tests, just go to root folder of this project and run `make test`.

## Decisions

### Architecture

So, for this api, I decided to use `DDD` architecture pattern, thinking like big appplications,if we will have features such as user authentication, transactions, purchases, and consequently we will start to involve important business rules here, the proposal that `DDD` brings, to reduce complexity, facilitate maintenance, extension and understanding of the code , it goes very well. How framework, was used Go-fiber, very simple to use and despite being new, it has several interesting features. It resembles ExpressJS from the NodeJS community.

### Database

For the database was used GORM with PostgreSQL. Some reasons for choosing Postgres, it could operate and protect data integrity at table level transactions, which make it less vulnerable to data corruption. Also, Postgres has more friendly types, like native uuid field, boolean for bool values(in MySQL is tinyint), jsonb, array, enums, among others. It implements Multiversion Concurrency Control (MVCC) without read locks Postgres supports parallel query plans that can use multiple CPUs/cores. For end, is open source.

### Testing

For the tests, I preferred to use the standard Go library, without using asserts directly, as is the most common in others languages. To run the tests it is only necessary to run the command `make test` in the same directory of the `Makefile`.

### Documentation

For the documentation, I decided to use Swagger and also keep the code documentation using the follow pattern: Title, parameters, a brief description, return types and their description, and also exceptions, if any.

### Linting

About linting, was used a editor config with some go patterns.
