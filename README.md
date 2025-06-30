# Go E-Wallet API

A simple digital wallet RESTful API built with Go (Gin), PostgreSQL, and JWT authentication. This project includes features like user registration, login, profile editing, wallet top-up, peer-to-peer wallet transfer, and viewing transfer history with search functionality. It follows a modular structure using the pgx PostgreSQL driver and middleware for token verification.

## Features

✅ User Registration & Login

✅ JWT-based Authentication

✅ Profile Update

✅ Top-Up Wallet

✅ Wallet-to-Wallet Transfer

✅ Transfer History with Search (by name or phone)

✅ Modular Project Structure

✅ PostgreSQL Integration (using pgx)


## Entity Relationship Diagram


```mermaid
erDiagram
direction LR

  users ||--o{ sessions : creates
  users ||--|| wallets : owns
  wallets ||--o{ topups : has
  topups }o--|| payment_methods : uses
  wallets ||--o{ transfers : sends
  wallets ||--o{ transfers : receives

  users {
        int id_user PK 
        string email
        string password
        string pin
        string username 
        string phone
        string profile_picture
    }

  sessions {
      int id_session PK
      int id_user FK
      date issued_at 
  }

  wallets {
      int id_wallet PK
      int id_user FK
      decimal balance
  }

  topups {
    int id_topup PK
    int id_wallet FK
    decimal amount
    int payment_method FK
    decimal admin_fee
    decimal tax
    date created_at
  }

  payment_methods{
    int id_payment_method PK
    string payment_method
  }

  transfers {
    int id_transfer PK
    int id_sender_wallet FK
    int id_receiver_wallet FK
    decimal amount
    string notes
    date created_at
    date deleted_at
  }

```

## How to Clone and Use

Make sure you have Golang installed on your device.

#### 1. Clone the repository
```
git clone https://github.com/ranandasatria/fgo24-be-weeklytask.git
```

#### 2. Navigate into the project directory
```
cd fgo24-be-weeklytask
```

#### 3. Install the dependencies
```
go mod tidy
```

#### 4. Setup .env 
Create a .env file in the root folder with the following variables:
```
APP_SECRET=your_jwt_secret_key
DATABASE_URL=postgres://username:password@localhost:5433/ewallet
```

#### 5. Run the program
```
go run main.go
```

## 📫 API Endpoints

| Method | Endpoint             | Description                        | Auth Required |
|--------|----------------------|------------------------------------|---------------|
| POST   | `/register`          | Register a new user                | No            |
| POST   | `/login`             | Login and get JWT token            | No            |
| PATCH  | `/profile/:id`       | Edit profile by user ID            | Yes           |
| POST   | `/topup`             | Top up wallet balance              | Yes           |
| POST   | `/transfer`          | Transfer balance to another user   | Yes           |
| GET    | `/transfer/history`  | View or search transfer history    | Yes           |

## 📄 License

This project is licensed under the **MIT License**.  

## ©️ Copyright

&copy; 2025 Kodacademy
