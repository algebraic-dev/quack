# Quack

A quake log parser to analyze kill score points

## How to Run

Before running Quack, you need to have golang installed on your system. Once you have it set up you
can follow these steps:

1. Clone the repository
```
> git clone https://github.com/algebraic-sofia/quack
```

2. Navigate to the project and run with a file in the standard input
```
> cd quack
> cat log.txt | go run cmd/quack/main.go
```

## Running tests

Quack comes with a set of tests to ensure it's working. To run these tests you simply need to run
the following command:

```
> go test -v internal/*
```