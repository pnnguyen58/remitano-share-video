package controllers

import (
	"errors"
	"fmt"
	beego "github.com/beego/beego/v2/server/web"
	"remitano-share-video/models"
)

type MainController struct {
	beego.Controller
}

type Response struct {
	Data  interface{} `json:"data"`
	Error string      `json:"err"`
}

type Header struct {
	Authorization []string `json:"Authorization"`
}

func (c *MainController) Prepare() {
	token := c.GetSession("token")
	if token != nil {
		if userId, err := checkToken(token.(string)); err != nil {
			c.Redirect("/v1/users/login", 302)
		} else {
			_ = models.SetToken(userId, token.(string))
		}
	} else {
		c.Redirect("/v1/users/login", 302)
	}

}

func (c *MainController) Get() {
	// Get list video
	var limit, offset int
	_ = c.Ctx.Input.Bind(&limit, "limit")
	_ = c.Ctx.Input.Bind(&offset, "offset")
	userId := c.GetSession("userId")
	videos, _ := models.GetVideos(userId.(int64), limit, offset)
	c.Data["videos"] = loadTable(videos)
	c.TplName = "index.html"
}

func checkToken(token string) (userId int64, err error) {
	return models.CheckToken(token)
}

func getVideo(u *VideoController) (video models.Video, err error) {
	video.Title = u.GetString("title")
	video.Link = u.GetString("link")
	video.Description = u.GetString("description")
	fmt.Println(u.GetString("username"))
	if video.Title == "" || video.Link == "" {
		return video, errors.New("missing required input")
	}
	return
}

func getUser(u *UserController) (user models.User, err error) {
	user.Username = u.GetString("username")
	user.Password = u.GetString("password")
	if user.Username == "" || user.Password == "" {
		return user, errors.New("missing required input")
	}
	return
}

func loadTable(videos []models.Video) (videoTable string) {
	videoTable = "<table style=\"width:100%\">\n<tr>\n<th>Title</th>\n<th>Description</th>\n<th>Link</th>\n</tr>\n"
	rowTemplate := "<tr>\n<td>%v</td>\n<td>%v</td>\n<td>%v</td>\n</tr>\n"
	for _, v := range videos {
		videoTable += fmt.Sprintf(rowTemplate, v.Title, v.Description, v.Link)
	}
	videoTable += "</table>"
	return
}
