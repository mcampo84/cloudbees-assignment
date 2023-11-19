#!/bin/bash

# Define variables
HOST="localhost:50051"  # Change this to your gRPC server's host and port
PROTO_FILE="./internal/proto/tickets.proto"  # Replace with the path to your .proto file

# Define test data
USER_FIRST_NAME="John"
USER_LAST_NAME="Doe"
USER_EMAIL="john@example.com"
FROM="London"
TO="France"
PRICE=20.0

# Create JSON payload for the PurchaseTicket request
PAYLOAD=$(cat <<EOF
{
  "user": {
    "first_name": "$USER_FIRST_NAME",
    "last_name": "$USER_LAST_NAME",
    "email": "$USER_EMAIL"
  },
  "from": "$FROM",
  "to": "$TO",
  "price": $PRICE
}
EOF
)

# Invoke PurchaseTicket endpoint using grpcurl
grpcurl -plaintext -proto $PROTO_FILE -d "$PAYLOAD" $HOST trainTicketing.TrainTicketService/PurchaseTicket