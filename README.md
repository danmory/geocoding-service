# Geocoding Service

The service to retrieve coordinates of a place via its name or vice verse.

The service uses Dadata API.

The Geocoding Service has microservice architecture:

1. Auth service
2. Retrieval service

## Auth Service

Service is used for authentication.

### Endpoints

1. /register - register new account
2. /login - login to existing account

### Usage

1. Clone the repository
2. Enter service directory
3. Install Docker, Docker-Compose
4. Initialize .env file with the following fields:

    ```dotenv
        APP_ADDRESS=:8080
        DATABASE_URL=postgres://user:pass@database:5432/dbname
        APP_SECRET=xxx

        POSTGRES_PASSWORD=xxx
        POSTGRES_USER=xxx
        POSTGRES_DB=xxx
    ```

5. Run the command

    `` $ docker-compose up -d --build ``

6. The service is accessable via localhost:8080
