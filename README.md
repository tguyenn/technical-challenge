notes:
- what is ORM?
    - ORM (Object-Relational Mapping) is a programming technique that allows you to interact with a relational database (i.e. postgresql)
- 

todo: 
- need to make database persistent between container sessions
    - smth smth mounting volume?

- populate build script
    - what else do i need to do besides "docker-compose build"??
- populate startup script
    - startup script should start up both containers (prob smth docker-compose thing) and then expose a CLI interface for user to interact with the database

Need 2 containers:
1. front facing Go client that takes CLI user input (GORM)
1. postgresql

# Setup
If this is your first time using this project, then please build the docker image in the `techincal_challenge` directory:

```bash
bash build.sh

```

# Running the Docker Container
Simply run the startup script and be on your merry way :)

```bash
bash startup.sh
```

