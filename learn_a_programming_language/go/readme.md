# Go learnings

## Install Go

See [here](https://go.dev/wiki/Ubuntu)

## VS Code setup

Install alongside the Go [extension](https://marketplace.visualstudio.com/items?itemName=golang.Go)

## Running Go code

`go run main.go`

### Build Go code

`go build main.go`

###

go install - installs packages

go get - downloads someone elses packages

go test - run go tests

## Go Files

### Package

Can be executables or reusable libraries

#### executable

`package main` suggests make a executable package... e.g. main.exe

Also must include `func main(){}`

#### reusable

`package something_else` suggests create a library...

### Import

Import allows you to include someone elses package / library (reusable code as above) in your application...

[Reference](https://golang.org/pkg)

## Projects

[Card Game](./projects/card_game/readme.md)

## Type conversion

new_type(type to convert)
