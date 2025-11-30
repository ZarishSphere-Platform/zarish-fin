# Zarish-Fin (Financial Management Module)

Financial module for ZarishSphere Platform, providing comprehensive ERP accounting capabilities.

## Features

- **Journal Entry Management**: Double-entry bookkeeping system.
- **Chart of Accounts**: Configurable account hierarchy.
- **Multi-currency Support**: Handle transactions in multiple currencies.
- **Reporting**: Trial Balance, Balance Sheet, and P&L (planned).
- **Integration**: Seamlessly integrates with HIS for billing and HRA for payroll.

## Architecture

The system follows the standard ZarishSphere architecture:

- **API Layer**: RESTful endpoints using Gin.
- **Service Layer**: Business logic for accounting rules.
- **Repository Layer**: Data persistence with GORM.
- **Models**: Database schemas for financial data.

## API Endpoints

### Journal Entries

- `POST /api/v1/journal-entries` - Create journal entry
- `GET /api/v1/journal-entries/:id` - Get journal entry by ID
- `GET /api/v1/journal-entries` - List journal entries (paginated)

### Accounts

- `GET /api/v1/accounts` - List chart of accounts
- `POST /api/v1/accounts` - Create new account

## Running Locally

### Prerequisites

- Go 1.23+
- PostgreSQL 15+

### Steps

1. **Configure Database**: Ensure PostgreSQL is running.
2. **Run Server**:
   ```bash
   go run cmd/server/main.go
   ```
   Server runs on port **8081**.

## Docker Support

Build and run using Docker Compose:

```bash
docker-compose up --build zarish-fin
```

## Technology Stack

- **Language**: Go 1.23
- **Framework**: Gin
- **ORM**: GORM
- **Database**: PostgreSQL
