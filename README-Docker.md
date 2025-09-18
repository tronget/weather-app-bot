# Weather Bot Docker Setup

This guide explains how to run the Weather Bot application using Docker.

## Prerequisites

- Docker and Docker Compose installed on your system
- API keys for Telegram Bot and OpenWeather API

## Quick Start

1. **Clone and navigate to the project:**
   ```bash
   cd /path/to/weather-app-bot
   ```

2. **Configure environment variables:**
   Update the `.env` file with your API keys:
   ```env
   TELEGRAM_APIKEY=your_telegram_bot_token_here
   OPENWEATHER_APIKEY=your_openweather_api_key_here
   ```

3. **Start the application:**
   ```bash
   make up
   # or
   docker-compose up -d
   ```

4. **Check logs:**
   ```bash
   make logs
   # or
   docker-compose logs -f weather-bot
   ```

## Available Commands

### Using Makefile (Recommended)
- `make build` - Build the Docker images
- `make up` - Start the application and database
- `make down` - Stop and remove containers
- `make restart` - Restart the application
- `make logs` - Show application logs
- `make db-logs` - Show database logs
- `make shell` - Access application container shell
- `make clean` - Remove all containers, images, and volumes
- `make rebuild` - Clean build and start

### Using Docker Compose Directly
- `docker-compose build` - Build images
- `docker-compose up -d` - Start services in background
- `docker-compose down` - Stop services
- `docker-compose logs -f weather-bot` - Follow application logs
- `docker-compose logs -f postgres` - Follow database logs

## Architecture

The application consists of two services:

### 1. PostgreSQL Database (`postgres`)
- **Image:** postgres:15-alpine
- **Port:** 5432
- **Database:** weather_app_bot
- **Initialization:** Automatically creates tables from `init.sql`
- **Data persistence:** Uses Docker volume `postgres_data`

### 2. Weather Bot Application (`weather-bot`)
- **Built from:** Dockerfile (multi-stage Go build)
- **Depends on:** PostgreSQL service
- **Environment:** Production-ready Alpine Linux container

## Database Schema

The application automatically creates the following table:

```sql
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(255) UNIQUE NOT NULL,
    lang_code VARCHAR(10) NOT NULL DEFAULT 'en',
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);
```

## Environment Variables

The application uses these environment variables:

### API Configuration
- `TELEGRAM_APIKEY` - Your Telegram Bot API token
- `OPENWEATHER_APIKEY` - Your OpenWeather API key

### Database Configuration (Docker)
- `DB_USER` - Database username (default: tronget)
- `DB_HOST` - Database hostname (default: postgres)
- `DB_PORT` - Database port (default: 5432)
- `DB_NAME` - Database name (default: weather_app_bot)
- `DB_PASSWORD` - Database password (default: postgres)

## Troubleshooting

### Check Service Status
```bash
docker-compose ps
```

### View All Logs
```bash
docker-compose logs
```

### Restart Services
```bash
make restart
```

### Access Database
```bash
docker-compose exec postgres psql -U tronget -d weather_app_bot
```

### Clean Installation
```bash
make clean
make up
```

## Development

For local development without Docker, use the legacy environment variables in `.env`:
- `DATABASE_HOST=localhost`
- `DATABASE_USER`, `DATABASE_PASSWORD`, etc.

## Health Checks

The PostgreSQL service includes health checks to ensure the database is ready before starting the application. The weather-bot service will wait for the database to be healthy before starting.

## Data Persistence

Database data is persisted in a Docker volume named `postgres_data`. To completely reset the database:

```bash
docker-compose down -v  # This removes volumes
docker-compose up -d
```
