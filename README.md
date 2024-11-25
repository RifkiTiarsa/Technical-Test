# Marketplace API

## Feature
- Topup E Wallet

## Program Language
- GO

## Library
- `godotenv`  : `go get github.com/joho/godotenv`
- `jwt`       : `go get -u github.com/golang-jwt/jwt/v5`
- `logger`    : `go get github.com/sirupsen/logrus`

## Framework
- `gin`       : `go get github.com/gin-gonic/gin`

## Database
- `postgres`  : `go get github.com/lib/pq`

## Caching
- `redis` : `go get github.com/redis/go-redis/v9`

---

## Installation
1. Clone Repository
   ```bash
   https://github.com/RifkiTiarsa/Technical-Test.git
2. Run DDL.sql and DML.sql in the assets directory for initial data
3. Create a .env file and copy the example that is in the env_example
4. Install necessary dependencies
   ```bash
   go mod tidy
5. Run the application
   ```bash
   go run .

## API Spesification
- Register
Request :
    - Method    : POST
    - Endpoint  : api/v1/customer/register
    - Header    : 
        - Content-Type  : application/json
        - Accept        : application/json
    - Body      :
        {
            "name"       : "Rifki",
            "email"      : "tiarsarifki@gmail.com",
            "password"   : "rahasia"
        }
Response  :
    - Status    : 201 Created
    - Body      :
      {
          {"status":{"code":201,"message":"Created"},"data":{"id":4,"name":"Tiarsa","email":"tiarsa@gmail.com","password":"$2a$10$0YwoFfWoH2GJNKV8jWiileZA/BRSEYbzaaU4woywtcb4FsRb9AlIe","created_at":"2024-11-25T23:00:03.326005691+07:00","updated_at":"2024-11-25T23:00:03.326006186+07:00"}}
      }

- Login
Request :
    - Method    : POST
    - Endpoint  : api/v1/customer/login
    - Header    : 
        - Content-Type  : application/json
        - Accept        : application/json
    - Body      :
    {
        {
            "email" : "dummy@gmail.com",
            "password" : "rahasia"
        }
    }
Response  :
    - Status    : 200 OK
    - Body      :
      {
          {"status":{"code":200,"message":"OK"},"data":{"token":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJuYW1lIjoiUmlma2kiLCJlbWFpbCI6InRpYXJzYXJpZmtpQGdtYWlsLmNvbSIsImlzcyI6InRpYXJzYXJpZmtpIiwiZXhwIjoxNzMyNTU0MjYzLCJpYXQiOjE3MzI1NTA2NjN9.0PDXs2GO9bpf5BohY0kC4xYR4hZMRL1DiYhBhfdnPHE"}}
      }

- Topup
Request :
    - Method    : POST
    - Endpoint  : api/v1/customer/topup
    - Header    : 
        - Content-Type  : application/json
        - Accept        : application/json
        - Authorization : Bearer Token
    - Body      :
      {
            "customer_id" : 3,
            "merchant_id" : 3,
            "product_id" : 1,
            "payment_method" : "BCA"
      }
Response    :
    - Status    : 201 Created
    - Body      : 
        {
            {"status":{"code":201,"message":"Created"},"data":"Silahkan transfer sebesar 11000.00 ke bank BCA, rekening 123456789 a/n PT EMONEY INDONESIA. Silahkan konfirmasi : {topup_id : 14, amount : 10000.00, price : 11000.00, payment_method : BCA, payment_status : Done} pada link : 'http://localhost:8080/api/v1/topup/callback' jika sudah melakukan transfer"}
        }

- ConfirmTopup
Request :
    - Method    : POST
    - Endpoint  : api/v1/customer/topup/callback
    - Header    : 
        - Content-Type  : application/json
        - Accept        : application/json
        - Authorization : Bearer Token
    Body      :
        {
            "topup_id" : 14, 
            "amount" : 10000,
            "price" : 11000,
            "payment_method" : "BCA", 
            "payment_status" : "Done"
        }
Response    :
    - Status    : 200 OK
    - Body      : 
        {
            {"status":{"code":200,"message":"Topup berhasil"},"data":{"topup_id":14,"amount":10000,"price":11000,"payment_method":"BCA","payment_status":"Done"}}
        }

- Logout
Request :
    - Method    : POST
    - Endpoint  : api/v1/customer/topup/callback
    - Header    : 
        - Content-Type  : application/json
        - Accept        : application/json
        - Authorization : Bearer Token
Response    :
    - Status    : 200 OK
    - Body      :
        {
            {"status":{"code":200,"message":"Logged out successfully"},"data":null}
        }