# Go_boolean_service
> The role of this service is to create, delete and update boolean values with authorisation with token.
## Installation
### Docker(locally)
To create docker image 
``` bash
  docker build . -t go-boolean 
```
To run docker container
```bash
  docker run -p 8080:8080 go-boolean -e DOCK_USR=_ -e DOCK_PASS=_ -e DOCK_DB=_ -e DOCK_PORT= _
```
### Docker(from DockerHub)
```bash
  docker pull raghavtayal/go-boolean
```
```bash
  docker run -p 8080:8080 raghavtayal/go-boolean -e DOCK_USR=_ -e DOCK_PASS=_ -e DOCK_DB=_ -e DOCK_PORT= _
```
### Source
  To build the app locally. In root directory use :
  ```bash
    go mod download
  ```
``` bash
  go build .
```
  To run the app
  ```bash
    ./Go_boolean_service
  ```

## API
### HTTP
#### `POST /`
Create a boolean and return authentication token.
```bash
  curl -X POST http://localhost:8080  
  --header "Content-Type: application/json" --data '{"value": true, "key": "bool_key"}'
```
```bash
  {
  "id":"0e7aba578a6d-bc3d-9066-eaf5ec13e126",
  "key":"bool_key",
  "token":"eyJhbGciOiJIUzI1NiIsImtpZCI6InNpZ25pbl8xIiwidHlwIjoiSldUIn0.eyJleHAiOjE2MDA4ODM4NTUsImlkIjoiMGU3YWJhNTc4YTZkLWJjM2QtOTA2Ni1lYWY1ZWMxM2UxMjYifQ.x1bjQdauu0FzBNBrubmsnJQRDQKEuHHH-cTLxovYxeE",
  "value":true
  }
```
#### `GET /:id`
After authentication with token return the boolean with corresponding **ID**
```bash
  curl http://localhost:8080/[id]
  curl http://localhost:8080/"0e7aba578a6d-bc3d-9066-eaf5ec13e126"
```
```bash
  {
  "id":"0e7aba578a6d-bc3d-9066-eaf5ec13e126",
  "key":"bool_key",
  "value":true
  }
```
#### `PATCH /:id`
After authentication with token, updates the boolean corresponding to the given **ID**
```bash
curl -X PATCH http://localhost:8080/[id]
  --header "Content-Type: application/json" 
  --data '{"value": true, "key": "new_bool"}' 
  --header "Authorization: Bearer <token> "
  
  curl -X PATCH http://localhost:8080/"0e7aba578a6d-bc3d-9066-eaf5ec13e126" 
  --header "Content-Type: application/json" 
  --data '{"value": true, "key": "bool_newKey"}' 
  --header "Authorization: Bearer eyJhbGciOiJIUzI1NiIsImtpZCI6InNpZ25pbl8xIiwidHlwIjoiSldUIn0.eyJleHAiOjE2MDA4ODM4NTUsImlkIjoiMGU3YWJhNTc4YTZkLWJjM2QtOTA2Ni1lYWY1ZWMxM2UxMjYifQ.x1bjQdauu0FzBNBrubmsnJQRDQKEuHHH-cTLxovYxeE"
```
```bash
  {
  "id":"0e7aba578a6d-bc3d-9066-eaf5ec13e126",
  "key":"new_bool",
  "value":true
  }
```
#### `DELETE /:id`
After authentication with token, deletes the boolean corresponding to the given **ID**
```bash
  curl -X DELETE http://localhost:8080/[id] --header "Authorization: Bearer <token>"

  curl -X DELETE http://localhost:8080/"0e7aba578a6d-bc3d-9066-eaf5ec13e126" 
  --header "Authorization: Bearer eyJhbGciOiJIUzI1NiIsImtpZCI6InNpZ25pbl8xIiwidHlwIjoiSldUIn0.eyJleHAiOjE2MDA4ODM4NTUsImlkIjoiMGU3YWJhNTc4YTZkLWJjM2QtOTA2Ni1lYWY1ZWMxM2UxMjYifQ.x1bjQdauu0FzBNBrubmsnJQRDQKEuHHH-cTLxovYxeE"
```
```bash
  HTTP 204 No Content
```
