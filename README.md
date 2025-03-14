# Project Setup

## 1. Start the Docker Container

Run the following command to start the Docker container:

```bash
docker compose up -d
```

After this, create the table and user using the SQL scripts from the `main.sql` file.

## 2. Install Dependencies

To install the project dependencies, simply run:

```bash
go mod tidy
```

This command will fetch all the required dependencies listed in the `go.mod` file and ensure that your project is ready to run.

## 3. Run the App

To run the application, execute:

```bash
go run main.go
```

The server will start and be ready for use.
