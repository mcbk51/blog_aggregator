# Blog Aggregator

A command-line RSS feed aggregator that allows you to follow and read blog posts from multiple sources.

## Prerequisites

Before running the blog aggregator, you'll need to have the following installed:

### 1. PostgreSQL
- Install PostgreSQL on your system
- Create a database for the blog aggregator
- Ensure PostgreSQL is running and accessible

### 2. Go
- Install Go (version 1.19 or later recommended)
- Ensure Go is properly configured in your PATH

## Installation

Install the blog aggregator CLI using Go:

```bash
go install github.com/mcbk51/blog_aggregator@latest
```

This will install the `blog_aggregator` binary to your `$GOPATH/bin` directory (or `$HOME/go/bin` if `GOPATH` is not set).

## Configuration

Before using the blog aggregator, you need to set up a configuration file:

1. Create a configuration file (the program will look for it in the expected location)
2. Add your PostgreSQL database connection string to the config file
3. The config file should include your database URL in the format:
   ```
   DBURL=postgres://username:password@localhost:5432/database_name?sslmode=disable
   ```

## Usage

Run the blog aggregator with various commands:

```bash
blog_aggregator <command> [args...]
```

### Available Commands

#### User Management
- **`register <username>`** - Create a new user account
- **`login <username>`** - Log in as an existing user
- **`users`** - List all registered users
- **`reset`** - Reset user data

#### Feed Management
- **`addfeed <feed_url>`** - Add a new RSS feed (requires login)
- **`feeds`** - List all available feeds
- **`follow <feed_url>`** - Follow a specific feed (requires login)
- **`following`** - Show feeds you're currently following (requires login)
- **`unfollow <feed_url>`** - Unfollow a specific feed (requires login)

#### Content
- **`agg`** - Aggregate and fetch new posts from feeds
- **`browse`** - Browse and read posts from your followed feeds (requires login)

### Example Workflow

1. Register a new user:
   ```bash
   blog_aggregator register marco
   ```

2. Log in:
   ```bash
   blog_aggregator login marco
   ```

3. Add and follow some feeds:
   ```bash
   blog_aggregator addfeed https://example.com/rss
   blog_aggregator follow https://example.com/rss
   ```

4. Aggregate new posts:
   ```bash
   blog_aggregator agg
   ```

5. Browse your feeds:
   ```bash
   blog_aggregator browse
   ```

## Notes

- Most commands require you to be logged in first
- The program maintains state between runs using the configuration file
- Make sure your PostgreSQL database is running before using the application
