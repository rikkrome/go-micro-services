# `/docs` user-service

Design and user documents (in addition to your godoc generated documentation).


# Get Setup
```
go work use ./services/user-service
```

## podman

### Access database
1. psql cli
```
psql -h 127.0.0.1 -p 5432 -U postgres user_service
```

2. podman cli
```
podman exec -it d94ec31696e9 psql -h 127.0.0.1 -p 5432 -U postgres user_service
```
