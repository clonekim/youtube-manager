package main

import (
	"context"
	"github.com/go-macaron/cors"
	"google.golang.org/api/option"
	"google.golang.org/api/youtube/v3"
	"gopkg.in/macaron.v1"
	"log"
	"os"
	"strconv"

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
	m.Get("/api/related/:id", fetchRelatedVideo)
	m.Get("/api/videoCategory", fetchVideoCategory)

	m.Run(8000)

}

var (
	ctx = context.Background()
	region = ""
	hl = ""
)


func init() {
	log.Println("init dotenv")
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	region  = os.Getenv("REGION")
	hl = os.Getenv("HL")
}

func fetchMostPopularVideos(c *macaron.Context, logger *log.Logger, yts *youtube.Service) error {
	logger.Println("fetch most popular videos")

	pageToken := c.Query("pageToken")

	var itemCount int64 = 5
	itemCount, _ = strconv.ParseInt(c.Query("count"), 10, 64)

	if itemCount == 0 {
		itemCount = 3
	}

	res, err := yts.Videos.
		List([]string{"id,snippet"}).
		Chart("mostPopular").
		MaxResults(itemCount).
		PageToken(pageToken).
		RegionCode(region).
		Hl(hl).
		//VideoCategoryId("20").
		Do()

	if err != nil {
		return err
	}

	c.JSON(200, res)
	return nil
}

func fetchVideoId(c *macaron.Context, logger *log.Logger, yts *youtube.Service) error {

	id := c.Params(":id")
	logger.Printf("fetch video by id(%s)\n", id)

	res, err := yts.Videos.
		List([]string{"id,snippet"}).
		Id(id).
		RegionCode(region).
		Hl(hl).
		Do()

	if err != nil {
		return err
	}

	c.JSON(200, res)
	return nil
}

func fetchRelatedVideo(c *macaron.Context, logger *log.Logger, yts *youtube.Service) error {

	id := c.Params(":id")
	pageToken := c.Query("pageToken")

	res, err := yts.Search.
		List([]string{"id,snippet"}).
		RelatedToVideoId(id).
		PageToken(pageToken).
		Type("video").
		RegionCode(region).
		Do()

	if err != nil {
		return err
	}

	c.JSON(200, res)
	return nil
}

func fetchVideoCategory(c *macaron.Context, logger *log.Logger, yts *youtube.Service) error {

	res, err := yts.VideoCategories.
		List([]string{"id","snippet"}).
		RegionCode(region).
		Hl(hl).
		Do()

	if err != nil {
		return err
	}

	c.JSON(200, res)
	return nil
}