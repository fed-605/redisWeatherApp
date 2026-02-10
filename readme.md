# redisWeather

A small Go HTTP service for fetching current weather with Redis caching. Data source: Visual Crossing Weather API.

## Features
- Get current weather by location
- Redis response caching (TTL 1 minute)
- Simple health-check endpoint

## Requirements
- Go `1.25.1`
- Redis

## launching the application
1. Copy `.env.example` to `.env` and fill in the values.
2. Start Redis locally.
3. Run the service:

```bash
make run
```

## Environment variables
```env
PORT=8080
API_KEY=your_api_key
PROVIDER_BASE_URL=https://weather.visualcrossing.com/VisualCrossingWebServices/rest/services/timeline
REDIS_ADDR=localhost:6379
```

## API
- `GET /weather/{location}` — current weather by location
- `GET /healthz` — service health check

Example request:

```bash
curl http://localhost:8080/weather/Moscow
```

Example response (trimmed):

```json
{
  "resolvedAddress": "Moscow, Russia",
  "timezone": "Europe/Moscow",
  "currentConditions": {
    "datetime": "12:00:00",
    "temp": 10.5,
    "precip": 0
  }
}
```
