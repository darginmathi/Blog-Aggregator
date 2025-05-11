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
  - `git clone https://github.com/darginmathi/gator`
  - `cd gator`
  - `go install .`

## Setup
  - Start the Postgres server in the background
    - `sudo service postgresql start`
    - `sudo -u postgres psql`
    - `CREATE DATABASE gator;`
    - `\c gator`
    - `ALTER USER postgres PASSWORD 'postgres';`

  -  Migrate the tables go to src/schema use this command
    - `goose postgres postgres://postgres:postgres@localhost:5432/gator up`

  - Create and set config file
    - `nano ~/.gatorconfig.json`
    - `{"db_url":"postgres://postgres:postgres@localhost:5432/gator?sslmode=disable"}`

# Commands

### `gator register`

**Description:**
Registers the User.

**Usage:**
```bash
gator register <username>
```
### `gator login`

**Description:**
Loging into the user account.

**Usage:**
```bash
gator login <username>
```

### `gator addfeed`

**Description:**
Adding a feed to the database and folloing the feed(current user).

**Usage:**
```bash
gator addfeed <feedname> <url>
```

### `gator follow`

**Description:**
Follow the feed.

**Usage:**
```bash
gator follow <url>
```

### `gator agg`

**Description:**
Scrapeing the feeds the user is following at the interval input. format(10s, 2m, 1h)

**Usage:**
```bash
gator agg <time>
```

### `gator browse`

**Description:**
Displays a number of posts scraped by the agg fn by order of newly published and default quantity being 2 if not specified.

**Usage:**
```bash
gator browse <quantity>
```

### `gator unfollow`

**Description:**
Unfollow a feed.

**Usage:**
```bash
gator unfollow <url>
```

### `gator feeds`

**Description:**
List all the feeds in the database.

**Usage:**
```bash
gator feeds
```

### `gator following`

**Description:**
Lists all the feeds the user is following

**Usage:**
```bash
gator following
```

