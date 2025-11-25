# Zarish-Fin (Financial Management Module)

Financial module for ZarishSphere Platform, providing ERP accounting capabilities.

## Features

- Journal Entry Management (Account Moves)
- Chart of Accounts
- Multi-currency support
- GORM-based data persistence

## API Endpoints

### Journal Entries
- `POST /api/v1/journal-entries` - Create journal entry
- `GET /api/v1/journal-entries/:id` - Get journal entry by ID
- `GET /api/v1/journal-entries` - List journal entries (paginated)

## Running

```bash
go run cmd/server/main.go
```

Server runs on port **8081**.

## Models

- **AccountMove**: Journal entries (invoices, bills, manual entries)
- **AccountMoveLine**: Individual debit/credit lines
- **Account**: Chart of accounts

## Technology Stack

- Go 1.21+
- Gin Web Framework
- GORM ORM
- PostgreSQL
