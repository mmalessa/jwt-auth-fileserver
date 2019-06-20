# Go Http Fileserver with JWT auth ...and more

## Status
Development version. DO NOT USE!

## License
GPL  
// TODO

## JWT Auth Conception
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

## Generating self signed test SSL key / crt
```sh
openssl req -x509 -nodes -days 3650 -newkey rsa:2048 -keyout ./ssl/test.key -out ./ssl/test.crt
```