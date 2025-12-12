![CI Status](https://github.com/k4rldoherty/bridge-backend/actions/workflows/ci.yml/badge.svg)
# Bridge

## Overview

### Tech Stack
- This backend is built using **Go** and **PostgreSQL**.

### Business Reasoning
- The ordering system is designed for a **club shop model**, where external clubs (clients of the parent company) collect orders from their members and transmit them to the company for processing, where the orders are then manufactured and ordered in bulk. This system streamlines the ordering process, ensuring efficient management of customer and order details without handling payments.

- This system is beneficial for companies that regularly do batch order, like when a football team releases a new kit, and a large quantity of orders will come in in bulk, so instead of processing hundreds of individual transactions, and the company can handle them all at once through an invoice.

### Key Features

1. **Order Entry**
   - **Point-of-Sale Integration**: Cashiers enter customer and order information directly into the system during the transaction.
   - **Real-Time Updates**: Orders are updated in real-time, providing immediate visibility to the HQ staff.

2. **Order Management**
   - **Centralized Dashboard**: HQ staff can view all incoming orders, track their status, and manage fulfillment.
   - **Status Updates**: Staff can easily mark orders as pending or shipped, providing insight into the order’s lifecycle.

3. **Customer Records**
   - **Database of Customers**: Store detailed information about customers for easy reference and future ordering.
   - **Privacy Management**: Ensure that customer information is handled securely and in compliance with applicable regulations.

4. **Bulk Ordering System**
   - **Consolidated Handling**: Since payments are processed in bulk, the system focuses purely on order management and tracking.
   - **Reduction of Redundancy**: Removes the need for multiple individual transactions, streamlining the overall process.

### Operational Flow

1. **Customer Order Placement**:
   - Customers place orders through the cashier at club shops.

2. **Order Submission**:
   - Cashiers enter order details and customer information into the system.

3. **Order Visibility at HQ**:
   - Orders are immediately accessible to headquarters staff for processing.

4. **Order Processing**:
   - HQ staff reviews and manages orders, marking them accordingly as they move through the fulfillment process.

5. **Shipping Notifications**:
   - Once an order is fulfilled, HQ staff can update its status, indicating that it has been shipped to the club.


