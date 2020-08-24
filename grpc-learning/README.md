# build

docker build -t my-go-app .

visit localhost:8081

# view list

docker image ls

```
REPOSITORY   TAG        IMAGE ID            CREATED             SIZE
my-go-app    latest     408d6846bafd        5 seconds ago       355MB
```

# run

docker run -it -p 8080:8081 my-go-app

visit localhost:8080
