package data

import (
	"fmt"
	"math/rand"
	"time"
)

//Website defines the structure for an API website
// swagger:model
type Website struct {
    // the id of the website
    //
    // required: false
    // min: 1
    ID int `json:"id"`

    // the id of the user it belongs to
    //
    // required: true
    // min: 1
    UserID int `json:"userID" validate:"required,gt=0"`
    
    // name of the Website
    //
    // required: true
    // max length: 50
    Name string `json:"name" validate:"required"`
    
    // url of the Website
    //
    // required: true
    URL string `json:"url" validate:"required"`
    
    // link of script for fetching website articles
    //
    // required: true
    ScriptLink string `json:"scriptLink" validate:"required,script"`
    
    CreatedOn string `json:"-"`
    
    UpdatedOn string `json:"-"`
    
    DeletedOn string `json:"-"`
}

type Websites []*Website

func GetWebsites() Websites {
    return websiteList
}

func GetWebsiteByID(id int) (*Website, error) {
    i := findWebsiteByID(id)
    if i == -1 {
        return nil, ErrorWebsiteNotFound
    }

    return websiteList[i], nil
}

func AddWebsite(w Website) {
    w.ID = rand.Int()
    websiteList = append(websiteList, &w)
}

func UpdateWebsite(w Website) error {
    pos := findWebsiteByID(w.ID)
    if pos == -1 {
        return ErrorWebsiteNotFound
    }
    
    websiteList[pos] = &w
    
    return nil
}

func DeleteWebsite(id int) error {
    pos := findWebsiteByID(id)
    if pos == -1 {
        return ErrorWebsiteNotFound
    }

    websiteList = append(websiteList[:pos], websiteList[pos+1])

    return nil
}

var ErrorWebsiteNotFound = fmt.Errorf("Website not found")

func findWebsiteByID(id int) int {
    for i, w := range websiteList {
        if w.ID == id {
            return i
        }
    }

    return -1
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
