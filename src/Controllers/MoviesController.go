package Controllers

import (
	"strconv"

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
	msg, err := Models.CreateMovie(
		name, directors, actors, screen_writer,
		language, int16(length), int32(score),
		tags, description,
	)

	if err != nil {
		json.JsonMsgWithError(c, msg, err)
	} else {
		json.JsonMsgWithSuccess(c, msg)
	}
}
