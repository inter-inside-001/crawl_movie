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

func GetOneInfoByRegxep(content string, regexpString string) string{
	if content == ""{
		return ""
	}
	reg := regexp.MustCompile(regexpString)
	result := reg.FindAllStringSubmatch(content, -1)

	if len(result) == 0{
		return ""
	}

	return string(result[0][1])
}

func GetMovieDirector(movieHtml string) string {
	regexpStr := `<a.*?rel="v:directedBy">(.*?)</a>`
	result := GetOneInfoByRegxep(movieHtml, regexpStr)
	return result
}

func GetMovieName(movieHtml string) string {
	regexpStr := `<span\s*?property="v:itemreviewed">(.*?)</span>`
	result := GetOneInfoByRegxep(movieHtml, regexpStr)
	return result
}

func GetMovieGrade(movieHtml string)string{
	regexpStr := `<strong.*?property="v:average">(.*?)</strong>`
	result := GetOneInfoByRegxep(movieHtml, regexpStr)
	return result
}

func GetMovieOnTime(movieHtml string) string{
	regexpStr := `<span.*?property="v:initialReleaseDate".*?>(.*?)\(.*</span>`
	result := GetOneInfoByRegxep(movieHtml, regexpStr)
	return result
}

func GetMovieRunningTime(movieHtml string) string{
	regexpStr := `<span.*?property="v:runtime".*?>(.*?)</span>`
	result := GetOneInfoByRegxep(movieHtml, regexpStr)
	return result
}

func GetMulInfoByRegxep(content string, regexpString string) string{
	if content == ""{
		return ""
	}

	reg := regexp.MustCompile(regexpString)
	result := reg.FindAllStringSubmatch(content, -1)

	if len(result) == 0{
		return ""
	}

	res := ""
	for _,v := range result{
		res += v[1] + "/"
	}

	return res
}

func GetMovieMainCharacter(movieHtml string) string{
	regexpStr := `<a.*?rel="v:starring">(.*?)</a>`
	result := GetMulInfoByRegxep(movieHtml, regexpStr)
	return result
}

func GetMovieGenre(movieHtml string)string{
	regexpStr := `<span.*?property="v:genre">(.*?)</span>`
	result := GetMulInfoByRegxep(movieHtml, regexpStr)
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