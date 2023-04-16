# jwt-auth

## To Use the Project :- 
1) Install Dependencies
```
go mod tidy
```
2) Run the the project 

```
go run main.go
```


## Admin Account Endpoints

| Endpoint                             | Method | Description                |
|--------------------------------------|--------|----------------------------|
| https://localhost:9000/signup         | POST   | Create a new user account  |
| https://localhost:9000/login          | POST   | Login                      |
| https://localhost:9000/users          | GET    | Get list of all users      |
| https://localhost:9000/users/:user_id | GET    | Get a specific user        |

## User Account Endpoints 

| Endpoint                                 | Method | Description           |
|------------------------------------------|--------|-----------------------|
| https://localhost:9000/login             | POST   | User Login            |
| https://localhost:9000/users/:user_id    | GET    | Get a specific user   |
| https://localhost:9000/users             | GET    | Get list of all users |

## Use POSTMAN or similar app for testing

![image](https://user-images.githubusercontent.com/98258627/226696602-94cb59af-646d-4803-afc2-b671c0cd1b1a.png)
