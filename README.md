# Train Ticketing System (Golang gRPC)

This repository contains the implementation of a train ticketing system using gRPC in Golang.

## Setup
1. Clone the repository.
2. Install Golang and necessary dependencies.
3. Run `go run cmd/server/main.go` to start the gRPC server.

## gRPC Endpoints
- `PurchaseTicket`
- `UserDetails`
- `UsersBySection`
- `RemoveUser`
- `ModifyUserSeat`

## Usage
- Purchase a ticket: [Implement gRPC client code to call PurchaseTicket method]
- Retrieve user details: [Implement gRPC client code to call UserDetails method]
- View users by section: [Implement gRPC client code to call UsersBySection method]
- Remove a user: [Implement gRPC client code to call RemoveUser method]
- Modify a user's seat: [Implement gRPC client code to call ModifyUserSeat method]

**Note:** This implementation utilizes gRPC APIs, not RESTful ones. In-memory storage is used; data will be lost on server restart.
