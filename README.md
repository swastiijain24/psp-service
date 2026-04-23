# PSP Service (NPCI Gateway)

## 📌 Overview

The **PSP Service** acts as the **gateway between external Payment Service Providers (PSPs)** and the NPCI network in a UPI-like distributed payment system.

It is responsible for:
* Accepting requests from PSP applications
* Authenticating PSPs via API keys
* Resolving VPAs to destination banks
* Publishing transaction requests to the NPCI event bus (Kafka)
* Storing asynchronous responses for client polling
* Providing APIs for transaction status retrieval

This service models the **entry point into the NPCI switch**, similar to real-world UPI infrastructure.

##  🔑  Responsibilities

### 1. API Gateway for PSPs

* Exposes REST APIs for payments and account operations
* Authenticates requests using API keys

### 2. VPA Resolution

* Resolves Virtual Payment Address (VPA) to:

  * destination bank
  * account identifier
* Enables correct routing of transactions across banks

### 3. Event Producer (Kafka)

* Publishes **payment request events** to Kafka
* Ensures decoupled and scalable processing

### 4. Asynchronous Response Handling

* Does **not return final status immediately**
* Stores responses received from Kafka in Redis
* Clients **poll for transaction status**

### 5. Response Mapping (Redis)

* Maps:

  ```id="3slw6n"
  transaction_id → latest transaction status
  ```
* Enables fast, low-latency status lookup
---

## 📨 Kafka Topics

| Topic Name         | Purpose                    |
| ------------------ | -------------------------- |
| `payment_request`  | Payment initiation events  |
| `payment_response` | Final transaction outcomes |

The PSP Service:

* **produces** → `payment_request`
* **consumes** → `payment_response`


---

## 🧩 Design Decisions

### Asynchronous Processing

* Payments are processed via Kafka
* Improves scalability and fault tolerance

### Polling Model

* Clients poll for transaction status
* Avoids long-lived connections

### Loose Coupling

* Services communicate via events
* Enables independent scaling and failure isolation


---

## 🧠 Redis Usage

Redis is used as a **response store**.

### Purpose:

* Store latest transaction state
* Enable fast polling by clients

### Example:

```id="nx2nyr"
Key: transaction:<txn_id>
Value: {
  "status": "SUCCESS",
}
```
---

## Authentication

All APIs require:

```id="flndy4"
x-api-key: <psp-api-key>
```

* Keys are validated via middleware
* Can be extended with IP allowlisting

---

## Tech Stack

* **Language:** Go
* **Framework:** Gin
* **Messaging:** Kafka
* **Cache:** Redis
* **Database:** PostgreSQL

---

## Project Structure

```id="d7l0u7"
psp-service/
├── cmd/
├── internals/
│   ├── handlers/
│   ├── services/
│   ├── repositories/
│   ├── routes/
│   ├── middleware/
│   ├── kafka/          # producers & consumers
│   └── http_client/
├── migrations/
└── go.mod
```
