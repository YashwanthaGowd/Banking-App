banking-app

## Getting Started

### Prerequisites

- Docker
- Docker Compose
- Go

### Installation

1. Clone the repository
2. Navigate to the project directory
3. Run `make build`

### Usage

1. Run `make run`

## Features

This project includes the following features:

- Account Service
- Transaction Service
- Kafka

## Testing

To run the tests, use the following command:

```bash
make test
```

## Deployment

To deploy the application, use the following command:

```bash
make build
```

This will build the Docker images for the services and run them in Docker containers.

## Configuration

The application uses a configuration file to set up the services. The configuration file is located in the `config` directory and is named `config.yml`. You can modify this file to configure the services as needed.

The configuration file supports the following options:

account-service:

- `account-service.server.host`: The host address for the account service.
- `account-service.server.port`: The port number for the account service.
- `account-service.postgres.host`: The host address for the PostgreSQL database.
- `account-service.postgres.port`: The port number for the PostgreSQL database.
- `account-service.postgres.user`: The username for the PostgreSQL database.
- `account-service.postgres.password`: The password for the PostgreSQL database.
- `account-service.postgres.dbname`: The name of the PostgreSQL database.
- `account-service.postgres.sslmode`: The SSL mode for the PostgreSQL database.
- `account-service.postgres.uri`: The URI for the PostgreSQL database.
- `account-service.kafka.brokers`: A list of broker addresses for the Kafka cluster.
- `account-service.kafka.topic`: The topic name for the Kafka messages.
- `account-service.kafka.batch_size`: The batch size for the Kafka messages.
- `account-service.kafka.batch_timeout`: The batch timeout for the Kafka messages.
- `account-service.kafka.required_acks`: The required ACKs for the Kafka messages.
- `account-service.kafka.async`: Whether to use asynchronous processing for the Kafka messages.

transaction-service:

- `transaction-service.server.host`: The host address for the transaction service.
- `transaction-service.server.port`: The port number for the transaction service.
- `transaction-service.mongodb.uri`: The URI for the MongoDB database.
- `transaction-service.kafka.brokers`: A list of broker addresses for the Kafka cluster.
- `transaction-service.kafka.topic`: The topic name for the Kafka messages.
- `transaction-service.kafka.consumer_group`: The consumer group for the Kafka messages.

## Sample API Requests

### Create Account

```bash
curl -X POST "http://localhost:8080/bankingapp/accounts" \
  -H "Content-Type: application/json" \
  -d '{
    "id": "123456",
    "first_name": "John",
    "last_name": "Doe",
    "email": "johndoe@example.com",
    "type": "personal",
    "password": "password",
    "created_at": "2022-01-01T00:00:00Z",
    "updated_at": "2022-01-01T00:00:00Z"
  }'

  HTTP/1.1 200 OK
  Content-Type: application/json
  Date: Mon, 01 Jan 2022 00:00:00 GMT
  Content-Length: 0
```

### Get Account

```bash
curl -X GET "http://localhost:8080/bankingapp/accounts/<accountId>" \
  -H "Content-Type: application/json"

  HTTP/1.1 200 OK
  Content-Type: application/json
  Date: Mon, 01 Jan 2022 00:00:00 GMT
  Content-Length: 123
  {
    "id": "123456",
    "first_name": "John",
    "last_name": "Doe",
    "email": "johndoe@example.com",
    "type": "personal",
    "password": "password",
    "created_at": "2022-01-01T00:00:00Z",
    "updated_at": "2022-01-01T00:00:00Z"
  }
```

### Update Account

```bash
curl -X PUT "http://localhost:8080/bankingapp/accounts/<accountId>" \
  -H "Content-Type: application/json" \
  -d '{
    "first_name": "Jane",
    "last_name": "Doe",
    "email": "janedoe@example.com",
    "type": "personal",
    "password": "password",
    "updated_at": "2022-01-01T00:00:00Z"
  }'

  HTTP/1.1 200 OK
  Content-Type: application/json
  Date: Mon, 01 Jan 2022 00:00:00 GMT
  Content-Length: 0
```

### Disable Account

```bash
curl -X DELETE "http://localhost:8080/bankingapp/accounts/<accountId>" \
  -H "Content-Type: application/json"

  HTTP/1.1 200 OK
  Content-Type: application/json
  Date: Mon, 01 Jan 2022 00:00:00 GMT
  Content-Length: 0
```

### Activate Account

```bash
curl -X PATCH "http://localhost:8080/bankingapp/accounts/<accountId>" \
  -H "Content-Type: application/json" \
  -d '{
    "status": "active"
  }'

  HTTP/1.1 200 OK
  Content-Type: application/json
  Date: Mon, 01 Jan 2022 00:00:00 GMT
  Content-Length: 0
```

### Deposit

```bash
curl -X POST "http://localhost:8080/bankingapp/accounts/deposit" \
  -H "Content-Type: application/json" \
  -d '{
    "id": "<account>",
    "amount": <amount>
  }'

  HTTP/1.1 200 OK
  Content-Type: application/json
  Date: Mon, 01 Jan 2022 00:00:00 GMT
  Content-Length: 0
```

### Withdraw

```bash
curl -X POST "http://localhost:8080/bankingapp/accounts/withdraw" \
  -H "Content-Type: application/json" \
  -d '{
    "id": "<account>",
    "amount": <amount>
  }'

  HTTP/1.1 200 OK
  Content-Type: application/json
  Date: Mon, 01 Jan 2022 00:00:00 GMT
  Content-Length: 0
```

### Get Transactions by Account

```bash
curl -X GET "http://localhost:8080/bankingapp/transactions/history/<accountId>/<count>" \
  -H "Content-Type: application/json"

  HTTP/1.1 200 OK
  Content-Type: application/json
  Date: Mon, 01 Jan 2022 00:00:00 GMT
  Content-Length: 123
  [
    {
      "id": "123456",
      "account": "123456",
      "amount": 100,
      "type": "credit",
      "timestamp": "2022-01-01T00:00:00Z"
    },
    {
      "id": "123456",
      "account": "123456",
      "amount": 100,
      "type": "credit",
      "timestamp": "2022-01-01T00:00:00Z"
    }
  ]
```

### Get Transaction by Id

```bash
curl -X GET "http://localhost:8080/bankingapp/transactions/id/<transactionId>" \
  -H "Content-Type: application/json"

  HTTP/1.1 200 OK
  Content-Type: application/json
  Date: Mon, 01 Jan 2022 00:00:00 GMT
  Content-Length: 123
  {
    "id": "123456",
    "account": "123456",
    "amount": 100,
    "type": "credit",
    "timestamp": "2022-01-01T00:00:00Z"
  }

```

### Get Transactions by Month Range

Use YYYY-MM format for startMonth and endMonth

```bash
curl -X GET "http://localhost:8080/bankingapp/transactions/range/<accountId>/<startMonth>/<endMonth>" \
  -H "Content-Type: application/json"

  HTTP/1.1 200 OK
  Content-Type: application/json
  Date: Mon, 01 Jan 2022 00:00:00 GMT
  Content-Length: 123
  [
    {
      "id": "123456",
      "account": "123456",
      "amount": 100,
      "type": "credit",
      "timestamp": "2022-01-01T00:00:00Z"
    },
    {
      "id": "123456",
      "account": "123456",
      "amount": 100,
      "type": "credit",
      "timestamp": "2022-01-01T00:00:00Z"
    }
  ]
```

## Contributing

Contributions are welcome! If you find any issues or have suggestions for improvements, please open an issue or submit a pull request on the GitHub repository.

## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details. 