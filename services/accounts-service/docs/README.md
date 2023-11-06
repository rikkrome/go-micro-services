# Accounts `/docs`

## APIs

### `/accounts/signup`
####  Request
```bash
curl --location --request POST 'http://localhost:8080/accounts/signup' \
--header 'Content-Type: application/json' \
--data-raw '{
    "firstname": "ricky",
    "lastname": "romero",
    "username": "rickyromero",
    "email": "contact@romeroricky.com",
    "password": "123456789"
}' | jq '.'
```
#### Response 
```json
{
  "id": "3ca22775-88d3-421e-8ede-c37731b3e6d8",
  "accessToken": "...",
  "accessTokenExpiry": 1699504468,
  "refreshToken": "...",
  "refreshTokenExpiry": 1701837268
}
```

### `/accounts/login`
####  Request
```bash
curl --location --request POST 'http://localhost:8080/accounts/login' \
--header 'Content-Type: application/json' \
--data-raw '{
    "email": "contact@romeroricky.com",
    "password": "123456789"
}' | jq '.'
```
#### Response 
```json
{
  "id": "3ca22775-88d3-421e-8ede-c37731b3e6d8",
  "accessToken": "...",
  "accessTokenExpiry": 1699504270,
  "refreshToken": "...",
  "refreshTokenExpiry": 1701837070
}
```

### `/accounts/mine`
####  Request
```bash
curl --location --request GET 'http://localhost:8080/accounts/mine' \
--header 'Authorization: Bearer AUTH_TOKEN' \
| jq '.'
```
#### Response 
```json
{
  "id": "3ca22775-88d3-421e-8ede-c37731b3e6d8",
  "firstname": "ricky",
  "lastname": "romero",
  "username": "rickyromero",
  "email": "contact@romeroricky.com"
}
```