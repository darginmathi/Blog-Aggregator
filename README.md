# Blog-Aggregator

Gator is a CLI tool for following and fetching RSS feeds

## Requirements

1. **PostgreSQL**:
   - You'll need a PostgreSQL database to store feed data.
      - `sudo apt update`
      - `sudo apt install postgresql postgresql-contrib`
      - `sudo passwd postgres`

2. **Go**:
   - The program is written in Go, so you'll need Go installed.
      - `sudo apt update`
      - `sudo apt install golang-go`

3. **Goose**:
   - `go install github.com/pressly/goose/v3/cmd/goose@latest`

## Installation
  - `git clone https://github.com/darginmathi/Blog-Aggregator`
  - `cd Blog-Aggregator`
  - `go install .`

## Setup
  - Start the Postgres server in the background
    - `sudo service postgresql start`
    - `sudo -u postgres psql`
    - `CREATE DATABASE gator;`
    - `\c gator`
    - `ALTER USER postgres PASSWORD 'postgres';`

  -  Migrate the tables go to src/schema use this command:
    - `goose postgres postgres://postgres:postgres@localhost:5432/gator up`

  - Create and set config file
    - `nano ~/.gatorconfig.json`
    - `{"db_url":"postgres://postgres:postgres@localhost:5432/gator?sslmode=disable"}`

# Commands

### `gator register`

**Description:**
Registers the User

**Usage:**
```bash
gator <command> [options] [arguments]


