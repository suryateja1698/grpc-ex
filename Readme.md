
```
go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
```


To run protoc command

```
export PATH="$PATH:$(go env GOPATH)/bin"
protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative protos/players.proto

```


## Procedure

1. In `protos/players.proto`, we define request, response, models & add method signatures with arguments and return value.
2. Run the protoc command, it will generate required files.
3. In `cmd/app/server`
   *  register service and add methods `AddPlayer` and `GetPlayer`. 
   * `AddPlayer` will just add player details from request in the array & `GetPlayer` will return the player details based on ID
4. In `cmd/app/client`
   * Connect to the server
   * Call `AddPlayer` and `GetPlayer` to add players and get player details.
5. Run Server using `make server` and client by `make client`
