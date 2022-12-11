package utils

import (
    "fmt"
    "time"
    "net/http"
    "encoding/json"

    "internal/entities"

    "github.com/dgrijalva/jwt-go/v4"
)


func LogIn(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Content-Type", "application/json")

    var user entities.User
    json.NewDecoder(r.Body).Decode(&user)
    fmt.Println("user:", user)

    token, err := checkLogin(user)
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println("token:")
    fmt.Println(token)
}

func GenerateJWT() (string, error) {
    token := jwt.New(jwt.SigningMethodHS256)

    claims := token.Claims.(jwt.MapClaims)

    claims["exp"] = time.Now().Add(time.Minute * 10).Unix()
    claims["user"] = "WTF"
    claims["authorized"] = true

    token_str, err := token.SignedString(SigningKey)
    if err != nil {
        return "", err
    }
    return token_str, nil
}

func checkLogin(u entities.User) (string, error) {
    if u.Username != trueUser.Username {
        fmt.Println("Username incorrect")
        return "", fmt.Errorf("incorrect username")
    }
    if u.Password != trueUser.Password {
        fmt.Println("Username incorrect")
        return "", fmt.Errorf("incorrect username")
    }

    token, err := GenerateJWT()
    if err != nil {
        return "", err
    }
    fmt.Println("token:", token)
    return token, nil
}
