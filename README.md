
# E-Commerce Microservices System


## Overview

This project is an implementation of an e-commerce platform using a microservice architecture. Each service in the system is designed to perform a specific function and communicates with other services through message brokers (RabbitMQ). The system is built with Go and uses PostgreSQL for database management.

The main services are:

- Shop Service: Manages shops.
- Product Service: Manages product catalogs.
- Order Service: Handles order creation, payment confirmation, and order cancellation.
- Warehouse Service: Manages stock across different warehouses.
- Customer Service: Manages user information and authentication.

## Running the Project


   ```bash
   docker compose -f docker-compose.yml up --build -d
   ```
## Access the Services

- RabbitMQ Management UI: http://localhost:15672 Username: user, Password: password

- Prometheus: http://localhost:9090

- Grafana: http://localhost:9000 Username: admin, Password: admin


- Product Service: http://127.0.0.1:3000

- User Service: http://127.0.0.1:3001

- Shop Service: http://127.0.0.1:3002

- Warehouse Service: http://127.0.0.1:3003

- Order Service: http://127.0.0.1:3004

## API Endpoints

## User Service
- POST /register - Register a new customer.
- POST /login - Login a customer.
- GET /me - Get customer profile.

## Product Service
- GET /products - Get all products.

## Shop Service
- GET /shops - Get all shops.
- POST /shops - Create new shop.
- PUT /shops/:id - Update shop data by id.
- DELETE /shops/:id - Delete shop data by id.
- GET /shops/:id - Delete shop data by id.

## Warehouse Service
- POST /warehouses - Create new warehouse.
- POST /warehouses/:warehouseID/activate - Activate warehouse.
- POST /warehouses/:warehouseID/shop/:shopID - Set shop's warehouses.
- POST /warehouses/:warehouseID/product/:productID/adjust - Adjust warehouse's stock.
- POST /warehouses/transfer - Transfer stock to another warehouse.

## Warehouse Service
- POST /orders/checkout - Checkout Order.
- POST /orders/payment/confirm - Payment Confirmation.
- POST /orders/cancel - Cancel Order.