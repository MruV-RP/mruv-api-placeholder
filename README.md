# mruv-api-placeholder
Implementation of MruV API server with placeholder data.

## How to run?
- **With docker:** 
``` docker-compose up ```
- **With go** 
``` go run main.go ```

## How to build?
Just simple type:
```
go buid
```
In project root directory.

## Configuration
You can configure this app by environment variables:
- HOST - host on which you want to serve API
- GRPC_PORT - port on which you want to serve gRPC API
- HTTP_PORT - port on which you want to serve gRPC-gateway REST API translation
