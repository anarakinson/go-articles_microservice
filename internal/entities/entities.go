package entities

import (

)

type Article struct {
    Id int `json:id`
    Title string `json:titel`
    Announce string `json:announce`
    Text string `json:text`
}
