package models

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/astaxie/beego/orm"
	"regexp"
)

var (
	db orm.Ormer
)

type MovieInfo struct {
	Id                  int64
	Movie_id            int64
	Movie_name          string
	Movie_pic           string
	Movie_director      string
	Movie_writer        string
	Movie_country       string
	Movie_language      string
	Movie_main_character string
	Movie_type          string
	Movie_on_time       string
	Movie_span          string
	Movie_grade         string
	_Create_time        string
}

func init() {
	orm.Debug = true
	orm.RegisterDataBase("default", "mysql", "root:root@tcp(127.0.0.1:3306)/test?charset=utf8", 30)
	orm.RegisterModel(new(MovieInfo))
	db = orm.NewOrm()
}

func AddMovie(movie_info *MovieInfo) (int64, error) {
	id, err := db.Insert(movie_info)
	return id, err
}

func GetMovieDirector(movieHtml string) string {
	regexpStr := `<a.*?rel="v:directedBy">(.*?)</a>`
	result := GetInfoByRegxep(movieHtml, regexpStr)
	return result[0][1]
}

func GetMovieName(movieHtml string) string {
	regexpStr := `<span\s*?property="v:itemreviewed">(.*?)</span>`
	result := GetInfoByRegxep(movieHtml, regexpStr)
	return result[0][1]
}

func GetMovieMainCharacter(movieHtml string) string{
	regexpStr := `<a.*?rel="v:starring">(.*?)</a>`
	result := GetInfoByRegxep(movieHtml, regexpStr)

	mainCharacters := ""
	for _,v := range result{
		mainCharacters += v[1] + "/"
	}
	return mainCharacters
}

func GetMovieGrade(movieHtml string)string{
	reg := regexp.MustCompile(`<strong.*?property="v:average">(.*?)</strong>`)
	result := reg.FindAllStringSubmatch(movieHtml, -1)

	return string(result[0][1])
}

func GetMovieGenre(movieHtml string)string{
	reg := regexp.MustCompile(`<span.*?property="v:genre">(.*?)</span>`)
	result := reg.FindAllStringSubmatch(movieHtml, -1)

	movieGenre := ""
	for _,v := range result{
		movieGenre += v[1] + "/"
	}
	return movieGenre
}

func GetMovieOnTime(movieHtml string) string{
	reg := regexp.MustCompile(`<span.*?property="v:initialReleaseDate".*?>(.*?)\(.*</span>`)
	result := reg.FindAllStringSubmatch(movieHtml, -1)

	return string(result[0][1])
}

func GetMovieRunningTime(movieHtml string) string{
	reg := regexp.MustCompile(`<span.*?property="v:runtime".*?>(.*?)</span>`)
	result := reg.FindAllStringSubmatch(movieHtml, -1)

	return string(result[0][1])
}

func GetInfoByRegxep(content string, regexpString string) [][]string{
	reg := regexp.MustCompile(regexpString)
	result := reg.FindAllStringSubmatch(content, -1)
	return result
}

func GetMovieUrls(movieHtml string)[]string{
	reg := regexp.MustCompile(`<a\s*?href="(https://movie.douban.com/subject/.*?)"`)
	result := reg.FindAllStringSubmatch(movieHtml, -1)

	var movieSets []string
	for _,v := range result{
		movieSets = append(movieSets, v[1])
	}
	return movieSets
}