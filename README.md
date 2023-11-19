# Train Ticketing System

This project implements a Train Ticketing System using gRPC.

### Approach

I adopted a domain-driven approach, leveraging interfaces and adapters to enhance code flexibility and facilitate future modifications. Uber fx modules were used for dependency injection and management, ensuring a streamlined implementation.

The project scope appeared larger than the initial Cloudbees estimation of 2 hours. Therefore, I focused on implementing a high-quality solution for a single endpoint rather than delivering a subpar, feature-complete product. Prioritizing quality over functionality minimizes technical debt in the long run.

### Purpose of Endpoints

#### PurchaseTicket
- **Status:** Implemented
- **Description:** Submits a purchase for a train ticket.
- **Request Input:** Ticket information including first and last name, email, travel origin, destination, and price.
- **Response Output:** Receipt with user details, travel information, and price.

#### ViewReceipt
- **Status:** Unimplemented
- **Description:** Retrieves the receipt for a user.
- **Request Input:** User information.
- **Response Output:** Receipt with user details, travel information, and price.

#### ViewPassengerManifest
- **Status:** Unimplemented
- **Description:** Retrieves the passenger manifest.
- **Request Input:** None.
- **Response Output:** Manifest containing sections and users.

#### CancelTicket
- **Status:** Unimplemented
- **Description:** Cancels a ticket for a user.
- **Request Input:** User information.
- **Response Output:** Empty response.

#### ChangeSeat
- **Status:** Unimplemented
- **Description:** Modifies a user's seat.
- **Request Input:** User information, requested section, and seat number.
- **Response Output:** Receipt with updated seat information.

### Running the Application

To start the application, execute the following command from the project root:
```shell
make start
```

**Note:** You will have to manually terminate the process which listens on port 50051 by finding it via the `ps` command, and killing it.
