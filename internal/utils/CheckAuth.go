package utils

import (
    "fmt"
    "net/http"

    "github.com/dgrijalva/jwt-go/v4"
)


func CheckAuth(endpoint func(http.ResponseWriter, *http.Request)) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

        // defer r.Body.Close()
        if r.Header["Token"] != nil {
            token, err := jwt.Parse(r.Header["Token"][0], func(token *jwt.Token) (interface{}, error) {
                _, ok := token.Method.(*jwt.SigningMethodHMAC)
                if !ok {
                    fmt.Println("not ok!")
                    return nil, fmt.Errorf("There was an error")
                }
                fmt.Println("return SigningKey, nil")
                return SigningKey, nil
            })

            if err != nil {
                defer func() {
                    if panicMessage := recover(); panicMessage != nil {
                        fmt.Printf("Panic '%v' captured", panicMessage)
                    }
                }()
                // panic("Oh no!")
                fmt.Println("token err != nil")
                // fmt.Println(err)
                fmt.Fprintf(w, err.Error())
            }

            if token.Valid {
                fmt.Println("token.Valid")
                endpoint(w, r)
            }
        } else {

            fmt.Println("Not Authorized")
            fmt.Fprintf(w, "Not Authorized")
        }
    })
}
