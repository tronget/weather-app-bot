# Weather App Bot 🌤️

A Telegram bot that provides weather information with internationalization support and PostgreSQL database integration.

## Features

- 🌍 **Weather Information**: Get current weather data for any city using OpenWeather API
- 🗣️ **Internationalization (i18n)**: Multi-language support with automatic language detection
- 👤 **User Management**: Automatic user registration and language preference persistence
- 🗄️ **Database Integration**: PostgreSQL database for storing user preferences and data
- 🐳 **Docker Support**: Complete containerized deployment with Docker Compose
- 🔧 **Configuration Management**: Environment-based configuration for different deployment scenarios


## Usage Examples

### Getting Weather Information
Simply send a city name to the bot:
```
London
```
```
New York
```
```
Tokyo
```

The bot will respond with current weather information including temperature, weather conditions, and other relevant data in your preferred language.

## Quick Start

### Docker Deployment (Recommended)
For Docker deployment instructions, see [README-Docker.md](README-Docker.md).

### Local Development

1. **Clone the repository:**
   ```bash
   git clone https://github.com/tronget/weather-app-bot.git
   cd weather-app-bot
   ```

2.  **Configure your environment:**

    - Copy the example environment file:

      ```bash
      cp .env.example .env
      ```
    -   Edit the `.env` file by adding your API keys for Telegram and OpenWeather.

4. **Install dependencies:**
   ```bash
   go mod tidy
   ```

5. **Run the application:**
   ```bash
   go run main.go
   ```

## Architecture

### Components

- **Telegram Bot Handler**: Manages incoming messages and user interactions
- **Weather Service**: Interfaces with OpenWeather API for weather data
- **Database Layer**: PostgreSQL integration for user management
- **Internationalization**: JSON-based translation system
- **Configuration Management**: Environment-based configuration loading

### Database Schema

The application uses PostgreSQL with the following main table:

```sql
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(255) UNIQUE NOT NULL,
    lang_code VARCHAR(10) NOT NULL DEFAULT 'en',
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);
```

## Internationalization (i18n)

The bot features a robust internationalization system:

- **Automatic Language Detection**: Detects user's Telegram language settings
- **Persistent Preferences**: Stores language preferences in the database
- **JSON-based Translations**: Translation files located in `internal/locales/json/`
- **Fallback Support**: Falls back to English if preferred language is unavailable
- **Runtime Loading**: Supports multiple deployment environments (local, Docker)


## Configuration

### Required Environment Variables

- `TELEGRAM_APIKEY` - Your Telegram Bot API token from @BotFather
- `OPENWEATHER_APIKEY` - Your OpenWeather API key

### Database Configuration

- `DB_USER` - Database username (default: default_user)
- `DB_HOST` - Database hostname (default: postgres)
- `DB_PORT` - Database port (default: 5432)
- `DB_NAME` - Database name (default: weather_app_bot)
- `DB_PASSWORD` - Database password (default: default_password)

## API Keys Setup

### Telegram Bot API
1. Message @BotFather on Telegram
2. Create a new bot using `/newbot`
3. Copy the provided API token

### OpenWeather API
1. Sign up at [OpenWeatherMap](https://openweathermap.org/api)
2. Generate an API key
3. Use the free tier for development

## Development

### Project Structure
```
weather-app-bot/
├── internal/
│   ├── botutil/          # Telegram bot utilities
│   ├── config/           # Configuration management
│   ├── locales/          # Internationalization
│   ├── network/          # API and database layers
│   └── weather/          # Weather service logic
├── Dockerfile            # Container configuration
├── docker-compose.yml    # Multi-service deployment
├── Makefile             # Development commands
└── init.sql             # Database initialization
```

### Available Make Commands
```bash
make help       # Show available commands
make build      # Build Docker images
make up         # Start services
make down       # Stop services
make logs       # Show application logs
make clean      # Clean everything
```

## Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Add tests if applicable
5. Submit a pull request

## License

This project is open source and available under the [MIT License](LICENSE).

## Support

For issues, questions, or contributions, please visit the [GitHub repository](https://github.com/tronget/weather-app-bot).