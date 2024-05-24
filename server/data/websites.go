package data

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math/rand"
	"regexp"
	"time"

	"github.com/go-playground/validator"
)

type Website struct {
    ID int
    UserID int `json:"userID" validate:"required"`
    Name string `json:"name" validate:"required"`
    URL string `json:"url" validate:"required"`
    ScriptLink string `json:"scriptLink" validate:"required,script"`
    CreatedOn string `json:"-"`
    UpdatedOn string `json:"-"`
    DeletedOn string `json:"-"`
}

func (w *Website) FromJSON(r io.Reader) error {
    e := json.NewDecoder(r)
    return e.Decode(w)
}

func (w *Website) Validate() error {
    validate := validator.New()
    err := validate.RegisterValidation("script", validateScript)
    if err != nil {
        return err    
    }
    return validate.Struct(w)
}

func validateScript(fl validator.FieldLevel) bool {
    re := regexp.MustCompile(`^https:\/\/utfs\.io\/f\/[a-z0-9-.]+$`)
    matches := re.FindAllString(fl.Field().String(), -1)

    return len(matches) == 1 
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
