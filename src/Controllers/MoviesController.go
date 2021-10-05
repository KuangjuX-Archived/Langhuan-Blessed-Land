package Controllers

import (
	"strconv"
	"time"

	"github.com/KuangjuX/Lang-Huan-Blessed-Land/Help/json"
	"github.com/KuangjuX/Lang-Huan-Blessed-Land/Models"
	"github.com/gin-gonic/gin"
)

func CreateMovie(c *gin.Context) {
	name := c.PostForm("name")
	directors := c.PostForm("directors")
	actors := c.PostForm("actors")
	screen_writer := c.PostForm("screen_writer")
	language := c.PostForm("language")
	length, _ := strconv.ParseInt(c.PostForm("length"), 10, 16)
	score, _ := strconv.ParseInt(c.PostForm("score"), 10, 32)
	tags := c.PostForm("tags")
	description := c.PostForm("description")
	release_date := c.PostForm("release_date")
	date, err := time.Parse("2006-01-02T15:04:05.000Z", release_date)
	msg, err := Models.CreateMovie(
		name, directors, actors, screen_writer,
		language, int16(length), int32(score),
		tags, description, date,
	)

	if err != nil {
		json.JsonMsgWithError(c, msg, err)
	} else {
		json.JsonMsgWithSuccess(c, msg)
	}
}

func FindMovieById(c *gin.Context) {
	movie_id := c.Query("movie_id")
	movie, err := Models.FindMovieById(movie_id)
	if err != nil {
		json.JsonMsgWithError(c, "Fail to find movie", err)
	} else {
		json.JsonDataWithSuccess(c, movie)
	}
}
