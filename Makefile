# Docker Commands Makefile

.PHONY: build up down restart logs clean help

# Default target
help:
	@echo "Available commands:"
	@echo "  build    - Build the Docker images"
	@echo "  up       - Start the application and database"
	@echo "  down     - Stop and remove containers"
	@echo "  restart  - Restart the application"
	@echo "  logs     - Show application logs"
	@echo "  db-logs  - Show database logs"
	@echo "  shell    - Access application container shell"
	@echo "  clean    - Remove all containers, images, and volumes"
	@echo "  rebuild  - Clean build and start"

# Build Docker images
build:
	docker compose build

# Start services
up:
	docker compose up -d

# Stop services
down:
	docker compose down

# Restart services
restart:
	docker compose restart

# Show application logs
logs:
	docker compose logs -f weather-bot

# Show database logs
db-logs:
	docker compose logs -f postgres

# Access application container shell
shell:
	docker compose exec weather-bot sh

# Clean everything
clean:
	docker compose down -v --rmi all --remove-orphans
	docker system prune -f

# Rebuild everything
rebuild: clean build up
