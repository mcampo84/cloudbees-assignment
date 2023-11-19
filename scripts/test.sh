#!/bin/bash

# Define variables
HOST="localhost:50051"  # Change this to your gRPC server's host and port
PROTO_FILE="./internal/proto/tickets.proto"  # Replace with the path to your .proto file

# PurchaseTicket request data
USER_ID=123
FROM="London"
TO="France"
PRICE=20.0

# Create JSON payload for the PurchaseTicket request
PAYLOAD=$(cat <<EOF
{
  "user_id": $USER_ID,
  "from": "$FROM",
  "to": "$TO",
  "price": $PRICE
}
EOF
)

# Invoke PurchaseTicket endpoint using grpcurl
grpcurl -plaintext -proto $PROTO_FILE -d "$PAYLOAD" $HOST trainTicketing.TrainTicketService/PurchaseTicket
