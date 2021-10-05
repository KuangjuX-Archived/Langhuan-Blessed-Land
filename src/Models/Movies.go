package Models

import (
	"time"

	orm "github.com/KuangjuX/Lang-Huan-Blessed-Land/DataBases/mysql"
)

// 电影元信息
type Movie struct {
	ID           int64
	Name         string
	Directors    string
	Actors       string
	ScreenWriter string
	Language     string
	Length       int16
	Tags         string
	Description  string
	Score        int32
	ReleaseDate  time.Time
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func CreateMovie(
	name, directors, actors, screen_writer, language string,
	length int16, score int32,
	tags, description string,
	release_date time.Time,
) (string, error) {
	movie := Movie{
		Name:         name,
		Directors:    directors,
		Actors:       actors,
		ScreenWriter: screen_writer,
		Language:     language,
		Length:       length,
		Score:        score,
		Tags:         tags,
		Description:  description,
		ReleaseDate:  release_date,
	}

	res := orm.Db.Create(&movie)
	if res.Error != nil {
		return "Fail to create", res.Error
	} else {
		return "Success to create", nil
	}
}

func FindMovieById(id string) (Movie, error) {
	var movie Movie
	res := orm.Db.Where("id = ?", id).First(&movie)
	return movie, res.Error
}
