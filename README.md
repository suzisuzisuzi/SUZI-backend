![backend-header](https://github.com/suzisuzisuzi/SUZI-backend/assets/81961406/de9497a4-ff37-4b01-aae9-9e682a91b936)
# SUZI Backend

This is the backend for the SUZI Project for the Google Solution Challenge 2024 by Team TBD.

## Instructions to Run
Create a .env file at the root of the project and add the following environment variables for database. Make sure you have a PostgreSQL database running on your machine. Replace the values with your own database credentials:
```
DBHOST=<HOST>
DBUSER=<USERNAME>
DBPASSWORD=<PASSWORD>
DB=postgres
DBPORT=5432
```
Run the following commands to start the server:
```
go run cmd/main.go
```
