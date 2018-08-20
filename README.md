OAUTH2 example with GO, JWT-go and gorilla

[![Build Status](https://travis-ci.org/hemanik/go-oauth2-jwt.svg?branch=master)](https://travis-ci.org/hemanik/go-oauth2-jwt)

# Get Started
- Install Glide https://github.com/Masterminds/glide
- Install dependencies
  ```language:shell
  $ glide install
  ```

after this step, you can do this:
- go run server.go

try it with curl:

```language:shell
curl localhost:8000/api/authenticate -X POST -d "{"username": "developer", "password": "developer"}'
```