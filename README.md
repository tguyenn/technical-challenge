# Introduction

Hello! This user management application consists of 2 containers:
1. GoLang application that allows interaction with database contents via HTTP requests
1. PostgreSQL server

* The GoLang application uses a CLI to support actions like creating, reading, updating and deleting database entries.
* The PostgreSQL comes prepopulated with some sample user data, which consists of employee's ID, name, email, and password. 
* Note that database contents are stored in a volume and persist between container sessions.

# Setup
Clone this repository and ensure the Docker daemon is running:
```bash
git clone git@github.com:tguyenn/technical-challenge.git
cd technical-challenge
```

### Option 1 - Build images from source
```bash
docker-compose build
```

### Option 2 - Pull images from DockerHub
```bash
docker pull tguyen/technical-challenge-app:latest
docker pull tguyen/technical-challenge-postgres:latest
```

# Launching the services
Simply launch the services, and the CLI will be ready after a moment:
```bash
docker-compose run --service-ports -it app
```

# Exiting the application
When prompted to pick an action, choose the `E` option to exit. This will put you back into your system's terminal.

Run the following command in your terminal to ensure all services (i.e. PostgreSQL) end:
```bash
docker-compose down
```
