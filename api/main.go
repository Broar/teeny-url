package main

import (
	"crypto/sha256"
	"encoding/base64"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/emicklei/go-restful"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"gopkg.in/yaml.v2"
)

type Config struct {
	DB struct {
		Name string
		Host string
		SSL  string
	}
}

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

var environment string
var username string
var password string
var config Config
var db *gorm.DB

func init() {
	flag.StringVar(&environment, "environment", "dev", "")
	flag.StringVar(&username, "username", "", "PostgreSQL username")
	flag.StringVar(&password, "password", "", "PostgreSQL password")
}

func main() {
	flag.Parse()

	file, err := ioutil.ReadFile(fmt.Sprintf("config/%s.yml", environment))
	if err != nil {
		panic(fmt.Sprintf("Failed to read YAML file: %v", err))
	}

	config := Config{}
	err = yaml.Unmarshal([]byte(file), &config)
	if err != nil {
		panic(fmt.Sprintf("Failed to parse YAML file: %v", err))
	}

	fmt.Printf("host=%s user=%s dbname=%s sslmode=%s password=%s\n", config.DB.Host, username, config.DB.Name, config.DB.SSL, password)
	db, err = gorm.Open("postgres", fmt.Sprintf("host=%s user=%s dbname=%s sslmode=%s password=%s", config.DB.Host, username, config.DB.Name, config.DB.SSL, password))
	if err != nil {
		panic("Failed to connect to database")
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
