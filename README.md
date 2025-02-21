# Shoes Store Project

## Contributors:
- **Nurislam Toktarov**
- **Kuanysh Aitzhanov**

This project is a Shoes Store application featuring a **Go-based backend** and a **HTML, CSS, and JS frontend**.

---

## Prerequisites
- **Git**: Make sure Git is installed.
- **Go**: Download and install Go.

---

## Setup Instructions

### 1. Clone the Repository
Clone the project to your desktop:
```bash

git clone https://github.com/KUANAIT/Shoes_Store.git
```

## 2. Setup the Backend

### Step 1: Install Dependencies
Run the following command to download all required Go dependencies:
```bash
go mod tidy
```
## Install MongoDB Dependencies
To interact with MongoDB, you need to install the MongoDB Go driver. Run the following command to install the required dependency:

```bash
go get go.mongodb.org/mongo-driver/mongo
```
## 3. Run the Backend Server
Start the backend by executing the following command:
```bash
go run main.go
```
Run in browser http://localhost:3000/
## Final Notes
Once both the backend and frontend servers are running, you should be able to access the application via your browser.  
For additional details or troubleshooting, please consult the project documentation or reach out to the maintainers.

## API Endpoints

### User Operations:
- **Create a user**: `/user/create`
- **Get a user by ID**: `/user/getbyid?id={id}`
- **Update a user**: `/user/update?id={id}`
- **Delete a user**: `/user/delete?id={id}`

### Shoe Operations:
- **Create a shoe**: `/create`
- **Get all shoes**: `/getall`
- **Get a shoe by ID**: `/getbyid`
- **Delete a shoe**: `/delete?id={id}`
- **Update a shoe**: `/update?id={id}`

### Order Operations:
- **Create an order**: `/order/create`
- **Get orders by user ID**: `/order/getbyuser?id={id}`
- **Create order by user ID (as user)**: `/order/createforuser?id={id}`
