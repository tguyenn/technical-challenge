# Introduction

Hello! This user management application consists of
1. Front facing GoLang CLI that allows the user to interact with the database contents
1. PostgreSQL server

The CLI supports actions like creating, reading, updating and deleting database entries.

The PostgreSQL comes prepopulated with some sample user data, which consists of employee's ID, name, email, and password. Note that database contents are stored in a volume and can be viewed with `docker volume inspect postgres_data`

# Setup
If this is your first time using this project, then please build the docker image in the `techincal_challenge` directory:

```bash
docker-compose build
```

# Launching the application
Simply launch the containers, and you will have a CLI after waiting a moment.

```bash
docker-compose run -it app
```

