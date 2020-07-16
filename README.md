# Super Heroes API

![Go](https://img.shields.io/badge/go-1.13-blue.svg)

Structure for app :open_file_folder: :octocat:

### Guide

1. [Installing necessary dependencies](#installing-necessary-dependencies)
2. [Running the app](#running-the-app)
3. [Documentation](#documentation)
4. [Tests](#tests)

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

-   [http://localhost:3000/api/v1/supers](http://localhost:300/api/v1/supers), for the API.
-   [http://localhost:3000/swagger](http://localhost:3000/swagger), for swagger.


## Documentation

-   For documentation, all code documentation stay above structs and functions definitions.
-   For API usage, was used the Swagger.io, after run the app you can access the Swagger [here](http://localhost:3000/swagger) to see how to use the application routes.

## Tests

-   For run the tests, just go to root folder of this project and run `make test`.
