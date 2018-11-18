package controllers

import (
	"github.com/astaxie/beego"
	"crawl_movie/models"
	"github.com/astaxie/beego/httplib"
	"fmt"
)

type CrawMovieController struct {
	beego.Controller
}

func (c *CrawMovieController) CrawlMovie() {
	sUrl := "https://movie.douban.com/subject/26147417/"
	rsp := httplib.Get(sUrl)
	sMovieHtml,err := rsp.String()
	if err != nil{
		panic(err)
	}

	var movieInfo models.MovieInfo
	movieInfo.Movie_name = models.GetMovieName(sMovieHtml)
	movieInfo.Movie_director = models.GetMovieDirector(sMovieHtml)
	movieInfo.Movie_main_character = models.GetMovieMainCharacter(sMovieHtml)
	movieInfo.Movie_grade = models.GetMovieGrade(sMovieHtml)
	movieInfo.Movie_on_time = models.GetMovieOnTime(sMovieHtml)
	movieInfo.Movie_span= models.GetMovieRunningTime(sMovieHtml)
	movieInfo.Movie_type= models.GetMovieGenre(sMovieHtml)

	 //id,err := models.AddMovie(&movieInfo)

	// 连接到redis
	models.ConnectRedis("127.0.0.1:6379")
	urls := models.GetMovieUrls(sMovieHtml)

	for _,url := range urls{
		models.PutinQueue(url)
	}


	 c.Ctx.WriteString(fmt.Sprintf("%v", urls))
}
