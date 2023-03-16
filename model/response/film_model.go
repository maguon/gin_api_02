package response

import (
	"gin_api_02/global"
	"time"

	pq "github.com/lib/pq"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type FilmInfo struct {
	global.EXTEND_MODEL
	UniqueId    string         `json:"uniqueId" form:"uniqueId" gorm:"column:unique_id" `
	Url         string         `json:"url" form:"url" gorm:"column:url" `
	Studio      string         `json:"studio" form:"studio" `
	Duration    int            `json:"duration,string" form:"duration" gorm:"column:duration" `
	Thumbnail   string         `json:"thumbnail" form:"thumbnail" `
	Img         string         `json:"img" form:"img" `
	Title       string         `json:"title" form:"title" `
	Content     string         `json:"content" form:"content" `
	PublishDate time.Time      `json:"publishDate" form:"publishDate" gorm:"column:publish_date" `
	Category    pq.StringArray `json:"category" form:"category"  gorm:"column:category" `
	Casts       pq.StringArray `json:"casts" form:"casts"  gorm:"column:casts"`
	Remarks     string         `json:"remarks" form:"remarks" `
	Status      int            `json:"status" form:"status" `
}

func (FilmInfo) TableName() string {
	return "film_info"
}

type FilmQuery struct {
	global.EXTEND_SEARCH
	FilmInfo
	UniqueId     string    `json:"uniqueId" form:"uniqueId" `
	CategoryId   string    `json:"categoryId" form:"categoryId" `
	ActorId      string    `json:"actorId" form:"actorId" `
	ActressId    string    `json:"actressId" form:"actressId" `
	DirectorId   string    `json:"directorId" form:"directorId" `
	StudioId     string    `json:"studioId" form:"studioId" `
	PublishStart time.Time `json:"publishStart" form:"publishStart" time_format:"2006-01-02"`
	PublishEnd   time.Time `json:"publishEnd" form:"publishEnd" time_format:"2006-01-02"`
}

type Actor struct {
	ID        primitive.ObjectID `json:"id"           bson:"_id"`
	Name      string             `json:"name"         bson:"name"`
	Url       string             `json:"url"         bson:"url"`
	FilmCount int                `json:"filmCount"         bson:"film_count"`
}
type Actress struct {
	ID        primitive.ObjectID `json:"id"           bson:"_id"`
	Name      string             `json:"name"         bson:"name"`
	Url       string             `json:"url"         bson:"url"`
	Avatar    string             `json:"avatar"         bson:"avtar"`
	Unit      string             `json:"unit"         bson:"unit"`
	Birth     time.Time          `json:"birth"         bson:"birth"`
	Figure    []string           `json:"figure"         bson:"figure"`
	FilmCount int                `json:"filmCount"         bson:"film_count"`
}

type ActressQuery struct {
	global.EXTEND_SEARCH
	Actress
}
type Category struct {
	ID        primitive.ObjectID `json:"id"           bson:"_id"`
	Category  string             `json:"category"         bson:"category"`
	Url       string             `json:"url"         bson:"url"`
	FilmCount int                `json:"filmCount"         bson:"film_count"`
}

type Director struct {
	ID        primitive.ObjectID `json:"id"           bson:"_id"`
	Name      string             `json:"name"         bson:"name"`
	Url       string             `json:"url"         bson:"url"`
	FilmCount int                `json:"filmCount"         bson:"film_count"`
}

type Studio struct {
	ID        primitive.ObjectID `json:"id"          bson:"_id"`
	Studio    string             `json:"studio"      bson:"studio"`
	Url       string             `json:"url"         bson:"url"`
	FilmCount int                `json:"filmCount"  bson:"film_count"`
}

type Mfilm struct {
	ID          primitive.ObjectID `json:"id"           bson:"_id"`
	UniqueId    string             `json:"uniqueId"     bson:"unique_id"`
	Title       string             `json:"title"     bson:"title"`
	Content     string             `json:"content"     bson:"content"`
	Url         string             `json:"url"     bson:"url"`
	Thumbnail   string             `json:"thumbnail"     bson:"thumbnail"`
	Image       string             `json:"image"     bson:"image"`
	Time        string             `json:"time"     bson:"time"`
	PublishDate primitive.DateTime `json:"publishDate" bson:"publish_date"`
	Actor       []Actor            `json:"actor"         bson:"actor"`
	Actress     []Actress          `json:"actress"         bson:"actress"`
	Category    []Category         `json:"category"  bson:"category"`
	Director    Director           `json:"director"  bson:"director"`
	Studio      Studio             `json:"studio"  bson:"studio"`
}
