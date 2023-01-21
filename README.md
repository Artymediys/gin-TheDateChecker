# TheDateChecker

## Project goal
### Part 1 - gin framework application

Develop a web application in the gin framework that contains a single route `/when/:year`. In the handler of that route, you need to output a string that shows how many days are left or have passed until January 1st of the year specified in the route parameter.

For example:  
`/when/2000` Must print the number of days from 01.01.2000 to today, in the format: `Days gone: 1234`

For example:  
`/when/2030` Should output how many days left to 01.01.2030, in the format: `Days left: 5432`

The HTTP-response status should be 200.

#### Error handling

Errors in the application must be logged in the console log.

#### Running the application

To run the application, create a Makefile with two commands:

* `run` - starts the application
* `build` - compiles the application

### Part 2 - middleware in the gin framework

Create your own middleware in this application which checks for HTTP headers. If it contains an `X-PING` header with a `ping` value, add an `X-PONG` response header with a `pong` value to your service response.

### Optional Part 3

Try to organize your project's structure and code following best practices for organizing your application code and application structure in Go.

## Tech stack
[![Go](https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white)](https://go.dev)
[![Gin](https://svgur.com/i/pgK.svg)](https://gin-gonic.com)


## Structure
* `bin/` - build files
* `cmd/` - main application
* `internal/` - internal application files
* `pkg/` - public files

## How to build
**Use the Makefile for build and run the application**  
* `make build` - builds the application and saves the output binary file to bin/app.out
* `make run` - builds the application and runs the bin-file
* `make clean` - cleans object files from package source directories (for example: app.out)