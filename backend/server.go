package main

import (
	"context"
	"github.com/go-macaron/cors"
	"google.golang.org/api/option"
	"google.golang.org/api/youtube/v3"
	"gopkg.in/macaron.v1"
	"log"
	"os"

	"github.com/joho/godotenv"
)


func main()  {
	m := macaron.Classic()
	m.Use(macaron.Renderer())
	m.Use(cors.CORS(cors.Options{}))

	m.Use(func(c*macaron.Context, log *log.Logger) {

		yts, err := youtube.NewService(ctx, option.WithAPIKey(os.Getenv("API_KEY")))

		if err != nil {
			panic("Fail to Youtube service !!!")
		}

		c.Map(yts)
		c.Next()
	})

	m.Get("/", func() string {
		return "Hello World!"
	})

	m.Get("/api/popular", fetchMostPopularVideos)
	m.Get("/api/video/:id", fetchVideoId)

	m.Run(8000)

}

var (
	ctx = context.Background()
)


func init() {
	log.Println("init dotenv")
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
}

func fetchMostPopularVideos(c *macaron.Context, logger *log.Logger, yts *youtube.Service) error {
	logger.Println("fetch most popular videos")

	pageToken := c.Query("pageToken")

	res, err := yts.Videos.
		List([]string{"id,snippet"}).
		Chart("mostPopular").
		MaxResults(3).
		PageToken(pageToken).
		Do()

	if err != nil {
		return err
	}

	c.JSON(200, res)
	return nil
}

func fetchVideoId(c *macaron.Context, logger *log.Logger, yts *youtube.Service) error {

	id := c.Params(":id")
	logger.Printf("fetch most popular videos (%s)\n", id)

	res, err := yts.Videos.
		List([]string{"id,snippet"}).
		Id(id).
		Do()

	if err != nil {
		return err
	}

	c.JSON(200, res)
	return nil
}