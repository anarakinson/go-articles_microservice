package utils

import (
    "fmt"
    "log"
    "database/sql"

    _ "github.com/go-sql-driver/mysql"

    "internal/database"
    "internal/entities"
)


func GetArticles() ([]entities.Article, error) {
    // parse db configs
    config, err := database.ParseConfig()
    if err != nil {
        log.Println("[!] Error when parsing db configs:", err.Error())
        return nil, err
    }

    // Connect to db
    db, err := sql.Open(
        "mysql",
        fmt.Sprintf(
            "%s:%s@tcp(%s:%s)/%s",
            config.Login,
            config.Passwrd,
            config.Address,
            config.Port,
            config.Name,
        ),
    )
    if err != nil {
        log.Println("[!] Error when connrcting to db:", err.Error())
        return nil, err
    }
    defer db.Close()

    query := fmt.Sprintf("SELECT `id`, `title`, `announce`, `text` FROM `articles`")
    // query := fmt.Sprintf("SELECT `id`, `title`, `announce`, `text` FROM `articles` WHERE `id` = %s", article_id)
    fmt.Println(query)
    res, err := db.Query(query)
    if err != nil {
        log.Println("[!] Error when loading article:", err.Error())
        return nil, err
    }
    defer res.Close()

    // parse result
    var articles []entities.Article
    for res.Next() {
        var post entities.Article
        err = res.Scan(&post.Id, &post.Title, &post.Announce, &post.Text)
        if err != nil {
            log.Println("[!] Error when loading article:", err.Error())
            return nil, err
        }
        articles = append(articles, post)
    }

    return articles, nil
}
