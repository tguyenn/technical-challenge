# Introduction

Hello! This user management application consists of 2 containers:
1. Front facing GoLang application that allows the user to interact with the database contents
1. PostgreSQL server

The GoLang CLI application supports actions like creating, reading, updating and deleting database entries.

The PostgreSQL comes prepopulated with some sample user data, which consists of employee's ID, name, email, and password. Note that database contents are stored in a volume and persist between container sessions.

# Setup
If this is your first time using this project, then please build the docker image in the `techincal_challenge` directory using your terminal of choice:

```bash
docker-compose build
```

# Launching the services
Simply launch the containers, and you will have a CLI after waiting a moment:

```bash
docker-compose run -it app
```

# Exiting the application
When prompted to pick an action, choose the `E` option to exit. This will put you back into your system's terminal.

Make sure to run the following command to kill any remaining services:

```bash
docker-compose down
```

# Debugging
If you run into an issue like the following, the problem may be caused by Git when cloning the repository.


```bash
postgres-1  | /usr/bin/env: ‘bash\r’: No such file or directory
postgres-1  | /usr/bin/env: use -[v]S to pass options in shebang lines
postgres-1 exited with code 127
```

To resolve, please change the way Git handles line endings by running this config command:
```bash
git config --global core.autocrlf false

```