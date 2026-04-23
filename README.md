# Gator - Blog Aggregator

Gator is a CLI-based RSS feed aggregator built in Go. It allows users to register, follow their favorite blogs, and scrape the latest posts directly into a PostgreSQL database for easy browsing.

## Prerequisites

Before you can run Gator, ensure you have the following installed on your system:

### 1. PostgreSQL
You need a running PostgreSQL instance. You can install it via your package manager:
- **Ubuntu/Debian:** `sudo apt install postgresql`
- **macOS (Homebrew):** `brew install postgresql`

Create a database for the project:
```bash
createdb gator
```

### 2. Goose
Goose is used for database migrations. Install it using Go:
```bash
go install github.com/pressly/goose/v3/cmd/goose@latest
```

### 3. SQLC
This project uses SQLC to generate type-safe Go code from SQL:
```bash
go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
```

## Installation

1. Clone the repository.
2. Build and install the binary:
   ```bash
   go install .
   ```
   *Note: Ensure your `$GOPATH/bin` is in your system's `PATH`.*

## Setup

### 1. Configuration
Gator stores its configuration in a `.gatorconfig.json` file in your home directory. You don't need to create it manually, but you should ensure your database connection string is correct.

### 2. Migrations
Run the migrations to set up your database schema:
```bash
cd sql/schema
goose postgres "postgres://username:password@localhost:5432/gator" up
```

## Usage

Gator is controlled via CLI commands.

### User Management
- **Register a new user:**
  ```bash
  gator register <username>
  ```
- **Login:**
  ```bash
  gator login <username>
  ```
- **List all users:**
  ```bash
  gator users
  ```

### Feed Management
- **Add a new feed:**
  ```bash
  gator addfeed <name> <url>
  ```
- **Follow a feed:**
  ```bash
  gator follow <url>
  ```
- **List all feeds:**
  ```bash
  gator feeds
  ```

### Browsing Posts
- **Aggregate (Scrape) feeds:**
  ```bash
  gator agg <time_between_reqs>
  ```
- **Browse latest posts:**
  ```bash
  gator browse <optional_limit>
  ```
