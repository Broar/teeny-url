package main

import (
	"crypto/sha256"
	"encoding/base64"
	"log"
	"net/http"

	"github.com/emicklei/go-restful"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var db *gorm.DB

type ShortURL struct {
	gorm.Model
	Key   string
	Value string
}

type CreateShortURLRequest struct {
	URL string
}

type URLResource struct{}

func (resource *URLResource) RegisterTo(container *restful.Container) {
	service := new(restful.WebService)
	service.Path("/url").
		Consumes(restful.MIME_XML, restful.MIME_JSON).
		Produces(restful.MIME_XML, restful.MIME_JSON)

	service.Route(service.GET("/{key}").To(resource.getShortURL))
	service.Route(service.PUT("").To(resource.createShortURL))
	service.Route(service.POST("").To(resource.createShortURL))

	container.Add(service)
}

func (resource *URLResource) createShortURL(request *restful.Request, response *restful.Response) {
	url := new(CreateShortURLRequest)
	err := request.ReadEntity(&url)

	if err == nil {
		hash := sha256.New()
		hash.Write([]byte(url.URL))
		key := base64.URLEncoding.EncodeToString(hash.Sum(nil))[:12]
		short := ShortURL{Key: key, Value: url.URL}
		db.Create(&short)
		response.WriteEntity(short)
	} else {
		response.WriteError(http.StatusInternalServerError, err)
	}
}

func (resource *URLResource) getShortURL(request *restful.Request, response *restful.Response) {
	key := request.PathParameter("key")

	var url ShortURL
	err := db.First(&url, "key = ?", key).Error

	if err == nil {
		response.WriteEntity(url)
	} else {
		response.WriteError(http.StatusNotFound, err)
	}
}

func main() {
	var err error
	db, err = gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("Failed to connect database")
	}

	defer db.Close()

	db.AutoMigrate(&ShortURL{})

	container := restful.NewContainer()
	resource := URLResource{}
	resource.RegisterTo(container)

	cors := restful.CrossOriginResourceSharing{
		ExposeHeaders:  []string{"X-My-Header"},
		AllowedHeaders: []string{"Content-Type", "Accept"},
		AllowedMethods: []string{"GET", "POST"},
		CookiesAllowed: false,
		Container:      container}

	container.Filter(cors.Filter)
	container.Filter(container.OPTIONSFilter)

	log.Fatal(http.ListenAndServe(":8080", container))
}
