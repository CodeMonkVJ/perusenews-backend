package data

type Article struct {
    ID int
    userID int
    websiteID int
    title string
    imageUrl string
    url string
    CreatedOn string
    UpdatedOn string
    DeletedOn string
}

var articleList = []*Article{}
