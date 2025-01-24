# NATS Messaging Examples

This repository demonstrates four core NATS messaging models:

1. **Pub/Sub (Publish/Subscribe)**
2. **Queue Groups**
3. **Request-Reply**
4. **JetStream**

Each model is explained below with its purpose, use cases, and instructions for running the examples.

---

## Table of Contents

1. [Prerequisites](#prerequisites)

---

## Prerequisites

Before running the examples, ensure you have the following installed:

-   **Go** (version 1.20 or higher)
-   **NATS Server** (run `nats-server -js` to start the server with JetStream enabled)
-   **NATS CLI** (optional, for monitoring streams and consumers)
