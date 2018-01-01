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

func main() {
	var err error
	db, err = gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("Failed to connect database")
	}

	defer db.Close()

	db.AutoMigrate(&ShortURL{})

	ws := new(restful.WebService)
	ws.Path("/url").
		Consumes(restful.MIME_XML, restful.MIME_JSON).
		Produces(restful.MIME_XML, restful.MIME_JSON)

	ws.Route(ws.GET("/{key}").To(getShortURL))
	ws.Route(ws.PUT("").To(createShortURL))
	ws.Route(ws.POST("").To(createShortURL))

	restful.Add(ws)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func createShortURL(request *restful.Request, response *restful.Response) {
	url := new(CreateShortURLRequest)
	err := request.ReadEntity(&url)

	if err == nil {
		hash := sha256.New()
		hash.Write([]byte(url.URL))
		key := base64.URLEncoding.EncodeToString(hash.Sum(nil))
		short := ShortURL{Key: key, Value: url.URL}
		db.Create(&short)
		response.WriteEntity(short)
	} else {
		response.WriteError(http.StatusInternalServerError, err)
	}
}

func getShortURL(request *restful.Request, response *restful.Response) {
	key := request.PathParameter("key")

	var url ShortURL
	err := db.First(&url, "key = ?", key).Error

	if err == nil {
		response.WriteEntity(url)
	} else {
		response.WriteError(http.StatusNotFound, err)
	}
}
