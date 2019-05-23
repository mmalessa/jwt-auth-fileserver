# Go Http Fileserver with JWT auth

## Status
Development version. Do not use!

## License
GPL  
// TODO

## Conception
```
[App Frontend] (with JWT authentication / authorization)
    |
    |   request for file (JWT token in header)                   file
    ---------------------------------------------->[File Sever]---------> [browser]
                                                    |      ^
                                                    |      |
                 ask for authorization (JWT token)  |      |
[Auth Server] <--------------------------------------      |
    |                                                      |
    |           agreement or not                           |
    -------------------------------------------------------|

```

## Installation
```sh
git clone ...
cd go-http-fileserver-jwt
go build -o go-http-fileserver-jwt main.go
```