#!/bin/bash

# Define variables
HOST="localhost:50051"  # Change this to your gRPC server's host and port
PROTO_FILE="./internal/proto/tickets.proto"  # Replace with the path to your .proto file
PRICE=29.00

CONTINUE="y"
while [ "$CONTINUE" == "y" ]; do
    echo "Purchase a Train Ticket"

    read -p "Enter your first name: " USER_FIRST_NAME
    read -p "Enter your last name: " USER_LAST_NAME
    read -p "Enter your email: " USER_EMAIL
    read -p "Enter the travel origin: " FROM
    read -p "Enter the travel destination: " TO

    # Create JSON payload for the PurchaseTicket request
    PAYLOAD=$(cat <<EOF
    {
        "first_name": "$USER_FIRST_NAME",
        "last_name": "$USER_LAST_NAME",
        "email": "$USER_EMAIL",
        "from": "$FROM",
        "to": "$TO",
        "price": $PRICE
    }
EOF
)

    # Invoke PurchaseTicket endpoint using grpcurl
    grpcurl -plaintext -proto $PROTO_FILE -d "$PAYLOAD" $HOST trainTicketing.TrainTicketService/PurchaseTicket

    read -p "Do you want to purchase another ticket? (y/n): " CONTINUE
    if [ "$CONTINUE" != "y" ]; then
        break
    fi
done
