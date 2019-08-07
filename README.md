# Go Contacts API
A Go Rest API in with MySQL Database and JWT

Postman Collection : [PostMan API Docs And Collection](https://documenter.getpostman.com/view/5756370/SVYruKML?version=latest)

Database Configuration 
1. Open your .env file
2. Change the following variables accordingly

```
db_name = go_contacts

db_pass = password

db_user = root
```

#### Running the project:
1. Clone to your go workspace folder and copy to src directory
2. Download dependencies:
  
    `go get -u github/gorilla/mux`
     
    `go get -u github/jinzhu/gorm`
     
    `go get -u github/dgrijalva/jwt-go`
    
    `go get -u github/joho/godotenv`
    
3. Cd into project
4. `go run main.go`
     