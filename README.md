# Go Gin MySQL Backend Framework

This project is a backend framework built using Go, the Gin web framework, and MySQL. It handles common backend tasks such as authentication, CRUD operations, and JWT-based user authentication.

## Table of Contents

- [Go Gin MySQL Backend Framework](#go-gin-mysql-backend-framework)
  - [Table of Contents](#table-of-contents)
  - [Introduction](#introduction)
  - [Features](#features)
  - [Project Structure](#project-structure)
  - [Installation](#installation)
  - [Usage](#usage)
  - [API Endpoints](#api-endpoints)
    - [Authentication](#authentication)
    - [Products](#products)
    - [Users](#users)
  - [Environment Variables](#environment-variables)
  - [Contributing](#contributing)
  - [License](#license)

## Introduction

The objective of this project is to provide a robust backend framework that leverages Go's performance, the simplicity of the Gin web framework, and the reliability of MySQL for data storage. The framework supports user authentication, CRUD operations for various entities, pagination, and secure JWT-based session management.

## Features

- User registration and login with JWT authentication.
- CRUD operations for products.
- Pagination support for API endpoints.
- Middleware for request logging, authentication, and error handling.
- Environment-based configuration management.

## Project Structure

project/
│
├── .env(.env.test)
├── main.go
├── go.mod
├── .gitignore
├── controllers/
│ ├── auth_controller.go
│ ├── product_controller.go
│ └── user_controller.go
├── models/
│ ├── product.go
│ └── user.go
├── db/
│ └── db.go
├── middleware/
│ └── auth_middleware.go
└── routes/
└── routes.go


## Installation

1. Clone the repository:

    ```sh
    git clone hhttps://github.com/noxilaman/backendbygo.git
    cd project
    ```

2. Install dependencies:

    ```sh
    go mod download
    ```

3. Create and configure the `.env` file:

    ```sh
    touch .env
    ```

    Add the following variables:

    ```
    DB_DRIVER=mysql
    DB_CONNECTION=[database user]:[database password]@tcp([database ip]:[database port])/[database name]?charset=utf8mb4&parseTime=True&loc=Local
    PORT=backend port
    JWT_SECRET=your_secret_key
    ```

4. Initialize the database:

    ```sh
    go run main.go
    ```

## Usage

1. Run the server:

    ```sh
    go run main.go
    ```

2. The server will start on the port specified in the `.env` file (default is `8080`). Access the API at `http://localhost:8080`.

## API Endpoints

### Authentication

- **POST** `/login`: User login.
- **POST** `/register`: User registration.
- **POST** `/refresh`: Refresh JWT token.

### Products

- **GET** `/products`: Get all products with pagination support.
- **GET** `/products/:id`: Get a specific product by ID.
- **POST** `/products`: Create a new product.
- **PUT** `/products/:id`: Update a product by ID.
- **DELETE** `/products/:id`: Delete a product by ID.

### Users

- **GET** `/users`: Get all users.
- **GET** `/users/:id`: Get a specific user by ID.

## Environment Variables

The following environment variables need to be set in the `.env` file:

- `DB_DRIVER`: Database Driver.
- `DB_CONNECTION`: Database connection.
- `PORT`: Backend Port.
- `JWT_SECRET`: Secret key for JWT.

## Contributing

Contributions are welcome! Please open an issue or submit a pull request for any changes.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
