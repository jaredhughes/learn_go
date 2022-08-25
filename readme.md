[![Test Authentication Service](https://github.com/jaredhughes/learn_go/actions/workflows/test_auth.yml/badge.svg)](https://github.com/jaredhughes/learn_go/actions/workflows/test_auth.yml)

# Learn Go

> A collection of Go services developed as part of the [Working with Microservices in Go](https://www.udemy.com/course/working-with-microservices-in-go/) course.

## Overview

### Project Root

The `/project` directory contains the Docker Compose and Makefiles. This is nested due to project initialization and dependency resolution nuances of VSCode. Prefer opening project as a workspace using the file committed in the repository root.

### Authentication

A naive implementation that accepts a username/password combination via HTTP POST, queries a connected Postgres instance, and verifies password match. Also demonstrates basic model services (`User` struct) and corresponding tests.

### Broker

Service handles all requests from frontend. Includes implementation of logging via gRPC and AMQP

### Frontend

Basic HTML rendering with simple JavaScript to test all exposed server endpoints.

### Listener

An AMQP implementation using RabbitMQ.

### Logger

Simple logging service that persists to Mongo, accepting events via REST, AMQP, RPC, and gRPC.

### Mailer

A transactional mail service implementing Mailhog via a singular HTTP POST endpoint (`/send`).

## Getting Started

#### Prerequistes

* [Docker Desktop](https://www.docker.com/products/docker-desktop/)
* [Make](https://www.gnu.org/software/make/)
* [Visual Studio Code](https://code.visualstudio.com/download) (recommended)
* [Go 1.19](https://go.dev/doc/install) (optional if not using Docker)

### Setup

1. Open [learn_go.code-workspace](learn_go.code-workspace) in VSCode
2. `cd project`
3. `make up_build`
