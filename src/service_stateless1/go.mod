module service_stateless1

go 1.14

import (
    "github.com/gin-gonic/gin"
    "github.com/badoux/checkmail"
    "github.com/jinzhu/gorm"
    "github.com/dgrijalva/jwt-go"
    "github.com/gorilla/mux"
    "github.com/joho/godotenv"
    "github.com/jinzhu/gorm/dialects/mysql"
    "github.com/jinzhu/gorm/dialects/postgres"
    "github.com/go-sql-driver/mysql"
    "github.com/graphql-go/graphql"
    "github.com/streadway/amqp"
    "go.mongodb.org/mongo-driver/mongo"
    "golang.org/x/crypto/bcrypt"
    "gopkg.in/go-playground/assert.v1"
)