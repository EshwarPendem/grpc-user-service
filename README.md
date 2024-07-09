# gRPC User Service

This is a Golang gRPC service that provides specific functionalities for managing user details and includes a search capability. The primary functionalities are as follows:
- A map of users with id as key is used as a mock databse. The sample data is in database/user_database.go, the same is provided 
   at the end of this file.
- 3 gRPC endpoints are implemented
   1. GetUserById - to get user based on user_id provided in request body.
   2. GetUsersByIds - to get multiple users with multiple ids in request body.
   3. SearchUser - returns list of users that match the criteria mentioned in request body. 
               you can provide one or more user attributes for searching and 
               all users who match all provided attributes will be returned.

## sample user data
   {"id": 1, "fname": "Steve", "city": "LA", "phone": 1234567890, "height": 5.8, "Married": true}

## Prerequisites

- Go (https://go.dev/dl/) (at least Go 1.22.4) needs to be installed.
- Postman needs to be installed
- (optional) install 'grpcurl'(https://github.com/fullstorydev/grpcurl/releases) homebrew tool to use endpoints from terminal. It can be installed on windows as well.
   for macOS grpcurl can be installed as a brew tool.
   for windows after extracting the zip. The bin folder path needs to be added in system environment variables.

## Getting Started

1. Clone the repository
   ```bash
   git clone https://github.com/EshwarPendem/grpc-user-service.git
   cd grpc-user-service
   ```

2. Build and run the application locally:

   ```bash
   go mod tidy
   go build -o grpc-service (for macOS)
   go build -o grpc-service.exe (for windows)
   ./grpc-service
   ```

   The gRPC service will be accessible on port `8081`.

3. Access the gRPC service endpoints.

   a. Via Postman
      -> open new request in postman (by clicking 'new' button beside 'import') and select the type as gRPC
      -> provide 'localhost:8081' in url 
      -> click select a method
      -> click on import .proto file and provide the path for .proto file in downloaded folder. (optional)
      -> select the endpoint and provide message(json request body)
      -> click on invoke.

   sample request body :- 
      1. GetUserById
      ```bash
         {"user_id": {"value": 2}}
      ```
      2. GetUserByIds
      ```bash
         {"user_ids": [1,2,4,9]}
      ```
      3. SearchUser
      ```bash 
         {"city":{"value":"New York"}, "phone":{"value":1234567890}}
      ```
   similary you can provide other attributes ("user_id","height","married") for search
   Note: at least one attribute should be provided for search phone number should be 10 digitsheight, user_id, and phone should be positive values.

   b. Via grpcurl
      -> install grpcurl and verify installation.

   sample curl commands :- 
      Note:- For mac os '\' characters may not be required.
      1.GetUserById
      ```bash
      grpcurl -v -d '{\"user_id\":5}' -plaintext localhost:8081 user.UserService/GetUserById 
      ```
      2.GetUsersByIds
      ```bash
      grpcurl -v -d '{\"user_ids\":[5,7,10,9]}' -plaintext localhost:8081 user.UserService/GetUsersByIds 
      ```
      3.SearchUser
      ```bash
      grpcurl -v -d '{\"city\":\"New York\", \"phone\":1234567890}' -plaintext localhost:8081 user.UserService/SearchUser
      ```
## Running Tests

```bash
cd service
go test -v
```
## Users
These example values cover a range of scenarios:

Users 1-15 are added in database.
Users 1, 2, and 3 share the same city and phone number.
Users 4 and 5 share the same height and marital status.

User 1:

Id: 1
Fname: "Alice"
City: "New York"
Phone: 1234567890
Height: 5.6
Married: true

User 2:

Id: 2
Fname: "Bob"
City: "New York"
Phone: 1234567890
Height: 5.9
Married: false

User 3:

Id: 3
Fname: "Carol"
City: "New York"
Phone: 1234567890
Height: 5.4
Married: true

User 4:

Id: 4
Fname: "David"
City: "Los Angeles"
Phone: 9876543210
Height: 6.0
Married: false

User 5:

Id: 5
Fname: "Emily"
City: "Los Angeles"
Phone: 9876543210
Height: 6.1
Married: true

User 6:

Id: 6
Fname: "Frank"
City: "Chicago"
Phone: 5551112222
Height: 5.8
Married: false

User 7:

Id: 7
Fname: "Grace"
City: "Houston"
Phone: 9998887777
Height: 5.5
Married: true

User 8:

Id: 8
Fname: "Henry"
City: "Miami"
Phone: 3334445555
Height: 5.9
Married: false

User 9:

Id: 9
Fname: "Isabella"
City: "Seattle"
Phone: 1112223333
Height: 5.6
Married: true

User 10:

Id: 10
Fname: "Jack"
City: "San Francisco"
Phone: 7778889999
Height: 6.2
Married: false

User 11:

Id: 11
Fname: "Kate"
City: "Boston"
Phone: 4445556666
Height: 5.7
Married: true

User 12:

Id: 12
Fname: "Liam"
City: "Denver"
Phone: 6667778888
Height: 5.8
Married: false

User 13:

Id: 13
Fname: "Mia"
City: "Portland"
Phone: 2223334444
Height: 5.4
Married: true

User 14:

Id: 14
Fname: "Noah"
City: "Austin"
Phone: 8889990000
Height: 6.0
Married: false
User 15:


Id: 15
Fname: "Olivia"
City: "Phoenix"
Phone: 5556667777
Height: 5.5
Married: true
