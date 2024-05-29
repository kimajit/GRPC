# GRPC
This project implements a gRPC service for managing user details. It includes functionalities like fetching user data based on user ID, retrieving a list of user details based on a list of IDs, and searching for user details based on specific criteria.

## Prerequisites

Before running the application, ensure you have the following installed:


- Go 
- Docker 
- ProtoC 


## Dependencies

####  Protobuf support for Go.
```
github.com/golang/protobuf
```
#### gRPC framework for Go.
```
google.golang.org/grpc
```

## Project Structure 
```

├── main.go
├── client
| ├──client.go
├── server
│ ├── server.go
├── go.mod
├── go.sum
├── proto
| ├── userdata.proto

```
- **`main.go`**: This file contains the main entry point for running the gRPC server.

- **`client.go`**: Contains the client code to interact with the gRPC server.

- **`server/server.go`**: Implements the gRPC server methods for user data management.

- **`userdata.proto`**: This Protobuf file defines the service and message types used in the gRPC communication.

### Generate .pb.go files from the proto file

```
protoc --go_out=. --go-grpc_out=. ./proto/data.proto  
```


### Runnig locally  

```
go mod tidy  

go run main.go 

```
* This command will start the server 
* For Testing the server you can navigate to client folder and Run 

``` 
cd client 
go run client.go

```

### Runnig Via Docker 

```
docker build -t userdata .  
docker run -p 50051:50051 userdata
```
* These command will build the docker file and start the server 
* For Testing the server you can navigate to client folder and Run 

``` 
cd client 
go run client.go

```
---------------
#### Unit testing can be found in server/dataServer_test.go