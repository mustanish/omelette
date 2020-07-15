# omelette

A simple golang skeleton built for writing REST APIs following best practices.

Contains basic CRUD operations, middlewares, test cases and schema validation.

## Requirements

- Docker and Docker Compose

## Getting Started

- Clone this repo `https://github.com/mustanish/omelette` then
- Run `docker-compose -f docker-compose.yml up` inside project root directory then
- Open `http://localhost:3000`

## Code Layout

The directory structure of the application:

    app/              App sources
        handlers/     App handlers go here
        connectors/   App connectors go here
        middlewares/  App middlewares go here
        responses/    App responses go here
        routes/       App routes go here
        schemas/      App schemas go here(Validation and Database)

    config/           Configuration directory
        config.go     Main app configuration file

    helpers/          Helper functions can be written

    tests/            Test cases can be written here

## Available routes

    https://www.getpostman.com/collections/42dba8a3c1243d76facb
