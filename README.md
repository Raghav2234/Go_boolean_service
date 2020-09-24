# Go_boolean_service
> The role of this service is to create, delete and update boolean values.
## Installation
## 

## API
### HTTP
#### `POST /`
Create a boolean and return authentication token.
```bash
  curl -X POST http://localhost:8080  
  --header "Content-Type: application/json" --data '{"value": true, "key": "bool_1"}'
```
```bash
  {
  "id":"0e7aba578a6d-bc3d-9066-eaf5ec13e126",
  "key":"bool_1",
  "token":"eyJhbGciOiJIUzI1NiIsImtpZCI6InNpZ25pbl8xIiwidHlwIjoiSldUIn0.eyJleHAiOjE2MDA4ODM4NTUsImlkIjoiMGU3YWJhNTc4YTZkLWJjM2QtOTA2Ni1lYWY1ZWMxM2UxMjYifQ.x1bjQdauu0FzBNBrubmsnJQRDQKEuHHH-cTLxovYxeE","value":true
  }
```
#### `GET /:id`
After authentication with token return the boolean with corresponding **ID**
```bash
```
#### `PATCH /:id`
After authentication with token, updates the boolean corresponding to the given **ID**
```bash
```
#### `DELETE /:id`
After authentication with token, deletes the boolean corresponding to the given **ID**
```bash
```
