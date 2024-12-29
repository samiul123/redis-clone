# Redis-Clone

Redis-Clone is a simplified implementation of a Redis server designed to explore and understand the internal workings of Redis. This project serves as both an educational tool and a practical exercise to learn and apply the Go programming language.

## Features

Redis-Clone supports the following commands using the RESP (REdis Serialization Protocol):

- **SET**: Stores a key-value pair.
- **GET**: Retrieves the value of a key.
- **HSET**: Stores a key-value pair within a specified hash.
- **HGET**: Retrieves the value of a key within a hash.
- **HGETALL**: Retrieves all key-value pairs within a hash.

## How to Run

### Prerequisites

- Go programming language installed on your system.
- A Redis client or a custom implementation for testing.

### Steps to Run the Server

1. Clone the repository:

    ```bash
    git clone <repository-url>
    cd redis-clone
    ```

2. Run the server:

    ```bash
    go run app/*.go
    ```

3. The server will start listening at `localhost:6379`.

### Steps to Connect and Use

1. Install a Redis client or use a custom implementation.
2. Run the Redis client and connect to `localhost:6379`.
3. Execute any of the supported commands (e.g., `SET`, `GET`, `HSET`, etc.).

### Example Commands

- **SET Command**:
    ```bash
    SET key value
    ```

- **GET Command**:
    ```bash
    GET key
    ```

- **HSET Command**:
    ```bash
    HSET hash key value
    ```

- **HGET Command**:
    ```bash
    HGET hash key
    ```

- **HGETALL Command**:
    ```bash
    HGETALL hash
    ```

## Purpose

The Redis-Clone project was developed as a learning tool for:
- Understanding the design and implementation of a key-value store.
- Practicing concurrency and network programming in Go.
- Exploring the RESP protocol.

## Acknowledgments

This project is inspired by [Redis](https://redis.io/), a powerful and efficient in-memory data structure store.
A special thanks to the tutorial [Build Redis From Scratch](https://www.build-redis-from-scratch.dev/en/introduction), which provided valuable guidance and insights into building this project.
Working on expanding the project further on top of this.
