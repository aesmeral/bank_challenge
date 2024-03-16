## Backend Challenge Project in GO

This project is to learn the Go language, Gin framework, and GORM (Go ORM). Also to learn more about docker and docker-compose, mostly creating the dockerfile and also how everything is connected.


### How to run locally

1. Install postgres https://www.postgresql.org/
2. Set up your environment variables. For me, Currently I have them in my .zshrc file since I plan on using docker compose to set them up.

    ```sh
    export BANK_HOST=localhost
    export BANK_USER=postgres
    export BANK_PASSWORD=postgres
    export BANK_DB=bank_challenge
    export BANK_PORT=5432
    ```

3. Run the necessary migrations (currently just one)

    ```sh
    > go run migrate/2024_03_16_migrate.go
    ```
4. Build and run the app

    ```sh
    > go build .
    > ./bank_challenge

    # alternatively you can also just run the app
    > go run .
    ```

### How to run using docker-compose

- TODO