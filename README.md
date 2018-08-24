# OAUTH2 Server using JWT Authentication

[![Build Status](https://travis-ci.org/hemanik/go-oauth2-jwt.svg?branch=master)](https://travis-ci.org/hemanik/go-oauth2-jwt)


This service implements [OAuth 2.0](https://tools.ietf.org/html/rfc6749#section-4.1) specification with GO, JWT-go and gorilla.

## Protocol Flow

JSON Web Tokens (JWT) are an approach to authentication that provide a way for clients to authenticate every request without having to maintain a session or repeatedly pass login credentials to the server.

``` text
 +--------+                                                      +---------------+
 |        |                                                      |               |
 |        |--(A)------ Authorization Grant---------------------->| Authorization |
 | Client |           POST /api/authenticate                     |     Server    |
 |        |           username=...&password=...                  |               | 
 |        |<-(B)---------- Access Token -------------------------|               |
 |        |               & Refresh Token                        |               |
 |        |                 HTTP 200 OK                          |               |         
 |        |               {Token: `..JWT..`,                     |               |
 |        |                Refresh: `..JWT`}                     |               |                     |        |                                                      |               |          
 |        |                                  +--------------+    |               |
 |        |--(C)-- Access Token ------------>|   Resource   |    |               |
 |        |        GET /api/resource         |    Server    |    |               |
 |        |  Authorization: Bearer `..JWT..` |              |    |               |          
 |        |                                  |              |    |               |
 |        |<-(D)-- Protected Resource -------|              |    |               |
 |        |          HTTP 200 OK             +--------------+    |               |
 |        |     {Message: `Authenticated`}                       |               |
 |        |                                                      |               |     
 |        |--(E)-------- Refresh Token ------------------------->|               |
 |        |<-(F)--------- Access Token --------------------------|               | 
 +--------+                                                      +---------------+
```

The flow illustrated in Figure includes the following steps:

  (A)  The client requests an access token by authenticating with the authorization server and presenting an authorization grant.
  (B)  The authorization server authenticates the client and validates the authorization grant, and if valid, issues an access token and a refresh token.
  (C)  The client makes a protected resource request to the resource server by presenting the access token.
  (D)  The resource server validates the access token, and if valid, serves the request.
  (E)  The client requests a new access token by authenticating with  the authorization server and presenting the refresh token.  The client authentication requirements are based on the client type and on the authorization server policies.
  (F)  The authorization server authenticates the client and validates the refresh token, and if valid, issues a new access token.

## Get Started

### Prerequisites

You need to have following packages in place:

* `git`
* `make`
* `go` (`>= v1.10.2`)
* link:https://github.com/Masterminds/glide[`glide`] for dependency management

Assuming that you have all the link:https://golang.org/doc/install[Golang prerequisites] in place (such as `$GOPATH`), clone the repository first:

[source,bash]
----
$ git clone https://github.com/fabric8-services/go-oauth2-jwt $GOPATH/src/github.com/fabric8-services/go-oauth2-jwt
----

### Install and execute

Execute make target `install` for installing the dependencies.
  ```language:shell
  $ make install
  ```

after this step, you can do this:
- go run server.go

try it with curl:

```language:shell
curl localhost:8000/api/authenticate -X POST -d "{"username": "developer", "password": "developer"}'
```

### Todos
- Add Refresh Token Flow
- Add Authorization Code Flow wherein a code is sent first (if required)