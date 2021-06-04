package blackerz

import (
   "io/ioutil"
   "log"
   "net/http"
)

// get bot data
func GetBotData(id string) string {
    resp, err := http.Get("https://blackerz.herokuapp.com/api/v1/bots/"+id)
    if err != nil {
      log.Fatalln(err)
    }
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
      log.Fatalln(err)
    }
    return string(body)
}
