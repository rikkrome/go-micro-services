# Go Micro Services

My goal for this repository is to build multiple microservices in Golang using a single codebase. Using [Standard Golang Project Layout](https://github.com/golang-standards/project-layout) 

## Get Setup
```
go work init
go work use ./services/auth-service
```

# podman
Start the podman server
```
podman machine start
```

Load compose
```
podman-compose -f compose.yaml up -d
```

# Docs
- [user_service](./services/user-service/docs/README.md)
- [auth-service](./services/auth-service/docs/README.md)

