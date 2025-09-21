# Weather Bot — Docker Setup

This guide explains how to run the Weather Bot application using Docker with environment variables managed via a local `.env` file.

## Prerequisites

- Docker and Docker Compose installed
- Telegram Bot token
- OpenWeather API key

## Quick start

1) Clone and enter the project
```bash
git clone https://github.com/tronget/weather-app-bot.git
cd weather-app-bot
```

2) Create your local environment file
```bash
cp .env.example .env
```

3) Edit `.env` and set your secrets and configuration
- TELEGRAM_APIKEY
- OPENWEATHER_APIKEY
- DB_* values (optional if you’re okay with the defaults)
- DB_NETWORK should stay as `weather-bot-network` unless you change the compose network name

4) Start the stack
```bash
make up
# or
docker compose up -d
# (older CLI)
docker-compose up -d
```

5) Follow logs
```bash
make logs
# or
docker compose logs -f weather-bot
# (older CLI)
docker-compose logs -f weather-bot
```

## How it works

The stack includes:
- PostgreSQL (service: `postgres`)
   - Image: `postgres:15-alpine`
   - Initializes from `init.sql` if present
   - Persists data in the `postgres_data` volume
- Weather Bot (service: `weather-bot`)
   - Built from the project `Dockerfile` (multi-stage Go build)
   - Waits for the database to become healthy before starting

Docker Compose automatically loads variables from the project’s `.env` file and injects them into services via `environment:`.

## Environment variables

Copy `.env.example` to `.env` and adjust as needed. Key variables:

- API
   - TELEGRAM_APIKEY — Telegram Bot token
   - OPENWEATHER_APIKEY — OpenWeather API key
- Database
   - DB_USER — database username (default: `default_user`)
   - DB_PASSWORD — database password (default: `default_password`)
   - DB_HOST — database hostname (default: `postgres`)
   - DB_PORT — database port (default: `5432`)
   - DB_NAME — database name (default: `weather_app_bot`)
   - DB_NETWORK — compose network name (default: `weather-bot-network`, must match the defined network in docker-compose.yml)


## Make targets (recommended)

- make build — Build images
- make up — Start the stack in the background
- make down — Stop and remove containers
- make restart — Restart the app
- make logs — Tail app logs
- make db-logs — Tail database logs
- make shell — Shell into the app container
- make clean — Remove containers, images, and volumes
- make rebuild — Clean build and start

## Common Docker Compose commands

- docker compose build — Build images
- docker compose up -d — Start services
- docker compose down — Stop services
- docker compose ps — Show service status
- docker compose logs -f weather-bot — Follow app logs
- docker compose logs -f postgres — Follow DB logs

(If your Docker uses the legacy plugin, replace `docker compose` with `docker-compose`.)

## Database initialization

If `init.sql` is present in the project root, it will be mounted and executed at first database startup to create/seed required structures. Update `init.sql` as your schema evolves—no schema is hardcoded in the application or docs.

Access the database:
```bash
docker compose exec postgres psql -U "${DB_USER}" -d "${DB_NAME}"
# (older CLI)
docker-compose exec postgres psql -U "${DB_USER}" -d "${DB_NAME}"
```

## Troubleshooting

- Check service status
  ```bash
  docker compose ps
  # or
  docker-compose ps
  ```

- View all logs
  ```bash
  docker compose logs
  # or
  docker-compose logs
  ```

- Restart services
  ```bash
  make restart
  ```

- Reset the database (removes data)
  ```bash
  docker compose down -v
  docker compose up -d
  # or
  docker-compose down -v
  docker-compose up -d
  ```

## Development (without Docker)

You can run the app locally by exporting the same variables you see in `.env.example` (or by creating a `.env` that your local tooling loads). Ensure you point DB_* to your local database instance.
