# Last traded price

Welcome to the repository of last traded price project!

## Description

## Features

1.**Docker Support**: The repository includes a Dockerfile and Docker Compose configuration to facilitate running the
system in Docker containers.

## Setup and Usage

### Prerequisites

- Docker
- Docker Compose

### Running the Application

1. **Clone the Repository**

    ```bash
    git clone https://github.com/sergey4qb/ltp.git
    cd rate
    ```
2. **Set Environment Variables**

   Before running the application, go to `ltp/configs` and configure the environment file. Add the following environment
   variables:

- `GIN_MODE`: The mode in which the Gin web framework runs. Default is `release`.
- `HTTP_PORT`: The port on which the HTTP server runs. Default is `8081`.

3. **Build and Run with Docker Compose**

    ```bash
    docker-compose up --build
    ```


# Endpoints

The following is a list of endpoints provided by the service along with their descriptions:

## Price

GET /api/v1/ltp

    Description: Retrieves last traded price for BTC/USD, BTC/CHF, BTC/EUR.
Response:
```
{
  "ltp": [
    {
      "pair": "string",
      "amount": "string"
    }
  ]
}

```