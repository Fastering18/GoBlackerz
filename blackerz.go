package blackerz

import (
   "io/ioutil"
   "log"
   "net/http"
   "encoding/json"
)

type Bot struct {
    LongDescription string `json:"LongDescription"`
    ShortDescription string `json:"ShortDescription"`
    Avatar string `json:"avatar"`
    InviteLink string `json:"inviteLink"`
    Name string `json:"name"`
    Owner BotOwner `json:"owner"`
    Prefix string `json:"prefix"`
    ServerCount int `json:"serverCount"`
    Tag string `json:"tag"`
    Upvotes int `json:"upvotes"`
    Verified bool `json:"verified"`
    ID string `json:"id"`
}
type BotOwner struct {
    ID string `json:"id"`
    Name string `json:"name"`
}

// get bot data
func GetBotData(id string) Bot {
    resp, err := http.Get("https://blackerz.herokuapp.com/api/v1/bots/"+id)
    if err != nil {
      log.Fatalln(err)
    }
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
      log.Fatalln(err)
    }
    var data Bot;
    json.Unmarshal(body, &data)
    return data
}
