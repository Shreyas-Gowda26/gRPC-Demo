# gRPC Demo

A simple gRPC application built with Go to learn the fundamentals of:

* Protocol Buffers (protobuf)
* gRPC Services
* Client-Server Communication
* Code Generation using `protoc`
* Unary RPC Calls

## Project Structure

```text
grpc-learning/
│
├── client/
│   └── main.go
│
├── server/
│   └── main.go
│
├── proto/
│   ├── calculator.proto
│   ├── calculator.pb.go
│   └── calculator_grpc.pb.go
│
├── go.mod
└── README.md
```

## Features

* Calculator Service
* Add two numbers using gRPC
* Protocol Buffer based communication
* Auto-generated Go code from `.proto` definitions

## Technologies Used

* Go
* gRPC
* Protocol Buffers

## Getting Started

### Install Dependencies

```bash
go mod tidy
```

### Run Server

```bash
go run server/main.go
```

### Run Client

```bash
go run client/main.go
```

## Learning Objectives

This project was created to understand:

* How gRPC works
* How Protocol Buffers generate code
* Client-server communication using RPC
* Building services in Go using gRPC

## Author

Shreyas G Gowda
