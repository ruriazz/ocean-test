# Take Home Test - Ocean Innovation

This project was developed as a take-home test for the Backend Engineer position at Ocean Innovation. The primary objective was to create an API for generating and validating OTPs using the Go programming language.

## Table of Contents

- [Overview](#overview)
- [API Documentation](#api-documentation)
- [Getting Started](#getting-started)
- [Environment Variables](#environment-variables)
- [Running the Project](#running-the-project)
- [Live Demo](#live-demo)

## Overview

This project provides two main API endpoints:
1. Generate OTP
2. Validate OTP

These APIs are designed and implemented following clean architecture principles to ensure scalability and maintainability.

## API Documentation

Detailed API documentation can be found in the Postman collection linked below:

[API Documentation on Postman](https://documenter.getpostman.com/view/12324981/2sA3kaAy9k)

## Getting Started

To get a local copy up and running, follow these steps.

### Prerequisites

Ensure you have the following installed:
- Go version 1.22.5
- Redis (for OTP storage)

### Installation

1. Clone the repository
    ```sh
    git clone https://github.com/ruriazz/ocean-test.git
    cd ocean-test
    ```

2. Install dependencies
    ```sh
    go mod tidy
    ```

## Environment Variables

Create a `local.env` file in the root directory of your project and set the following environment variables:

```env
DEBUG=

REDIS_HOST=
REDIS_PORT=
REDIS_USER=
REDIS_PASS=

META_BUSINESS_ID=
META_ACCESS_TOKEN=
```


### Variables Description

- `DEBUG`: A boolean flag indicating whether the application should run in debug mode. Set to `true` for more verbose logging and additional debug information.
- `REDIS_HOST`: The hostname of the Redis server.
- `REDIS_PORT`: The port of the Redis server.
- `REDIS_USER`: The username for the Redis server (if any).
- `REDIS_PASS`: The password for the Redis server (if any).
- `META_BUSINESS_ID`: The Meta (Facebook) Business ID used for WhatsApp API integration.
- `META_ACCESS_TOKEN`: The access token for authenticating with the Meta (Facebook) API.

## Running the Project

To run the project, use the following command:

```sh
go run main.go
```
The server will start and be accessible at http://localhost:8000.



## Usage
### Generate OTP

Endpoint: POST `/otp`
Request Body:
```json
{
    "whatsapp_number": "string"
}
```

### Validate OTP
Endpoint: POST `/otp/:whatsapp_number`
Request Body:

```json
{
    "token": "string",
    "otp": "string"
}
```

## Live Demo
You can check out the live demo of this project at the following URL:
[Live Demo](https://ocean-test.ruriazz.com)