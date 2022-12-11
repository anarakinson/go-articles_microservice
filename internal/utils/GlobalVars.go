package utils

import (
    "internal/entities"
)

var SigningKey = []byte("signing_key12345") // for HS256

var trueUser = entities.User{
    Username: "1",
    Password: "1",
}
