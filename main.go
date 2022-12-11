package main

import (
    "fmt"
    "log"
    "net/http"

    // "internal/entities"
    "internal/utils"

    "github.com/gorilla/mux"
    "github.com/spf13/viper"
    "github.com/joho/godotenv"
)



func main() {
    fmt.Println("Hello")

    // parse configs
    err := initConfig()
    if err != nil {
        log.Fatal("[!] Error when parsing configs: %s", err.Error())
    }
    // parse variables
    err = godotenv.Load()
    if err != nil {
        log.Fatal("[!] Error when parsing environment variables: %s", err.Error())
    }

    // Run server
    RunServer()
}

func initConfig() error {
    viper.AddConfigPath("configs")
    viper.SetConfigName("config")
    return viper.ReadInConfig()
}

func RunServer() {
    router := mux.NewRouter()
    port := viper.GetString("app.port")
    fmt.Println("port:", port)

    // router.HandleFunc("/main/", ShowArticle).Methods("GET")
    router.Handle("/main/", utils.CheckAuth(utils.ShowArticle)).Methods("GET")
    router.HandleFunc("/login/", utils.LogIn).Methods("POST")

    // http.Handle("/", router)
    http.ListenAndServe(fmt.Sprintf(":%s", port), router)
}
