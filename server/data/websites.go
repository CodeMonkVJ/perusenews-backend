package data

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math/rand"
	"time"
)

type Website struct {
    ID int
    UserID int `json:"userID"`
    Name string `json:"name"`
    URL string `json:"url"`
    ScriptLink string `json:"scriptLink"`
    CreatedOn string `json:"-"`
    UpdatedOn string `json:"-"`
    DeletedOn string `json:"-"`
}

func (w *Website) FromJSON(r io.Reader) error {
    e := json.NewDecoder(r)
    return e.Decode(w)
}

type Websites []*Website

func (w *Websites) ToJSON(wr io.Writer) error {
    e := json.NewEncoder(wr)
    return e.Encode(w)
}

func GetWebsites() Websites {
    return websiteList
}

func AddWebsite(w *Website) {
    w.ID = rand.Int()
    websiteList = append(websiteList, w)
}

func UpdateWebsite(id int, w *Website) error {
    pos, err := findWebsite(id)
    if err != nil {
        return err
    }
    
    log.Println("found website with id ", id)
    w.ID = id
    websiteList[pos] = w
    return nil
}

var ErrorWebsiteNotFound = fmt.Errorf("Website not found")

func findWebsite(id int) (int, error) {
    for i, w := range websiteList {
        if w.ID == id {
            return i, nil
        }
    }

    return -1, ErrorWebsiteNotFound
}

var websiteList = []*Website{
    {
        ID: 13423423,
        UserID: 23423423423,
        Name: "Zomato",
        URL: "https://blog.zomato.com/category/technology",
        ScriptLink: "https://utfs.io/f/89ac14f4-3c6a-460e-bcab-54c4c52156ad-tjg0p2.py",
        CreatedOn: time.Now().UTC().String(),
        UpdatedOn: time.Now().UTC().String(),
    },
    {
        ID: 13423453453,
        UserID: 23423423423,
        Name: "Zomato",
        URL: "https://blog.zomato.com/category/technology",
        ScriptLink: "https://utfs.io/f/89ac14f4-3c6a-460e-bcab-54c4c52156ad-tjg0p2.py",
        CreatedOn: time.Now().UTC().String(),
        UpdatedOn: time.Now().UTC().String(),
    },
}
