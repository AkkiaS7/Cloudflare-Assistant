package controller

import (
	"Cloudflare-Assistant/service"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"
)

type oauth2Controller struct {
	r                 gin.IRoutes
	githubOAuthConfig *oauth2.Config
}

type githubUser struct {
	Login     string `json:"login"`
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	AvatarURL string `json:"avatar_url"`
}

func InitOauth2Controller(r gin.IRoutes) {
	oc := &oauth2Controller{r: r}
	oc.registerRouter()
	oc.setGithubOAuthConfig()
}

// setGithubOAuthConfig sets the oauth config for github login
func (oc *oauth2Controller) setGithubOAuthConfig() {
	oc.githubOAuthConfig = &oauth2.Config{
		ClientID:     service.Conf.Oauth2.Github.ClientID,
		ClientSecret: service.Conf.Oauth2.Github.ClientSecret,
		RedirectURL:  service.Conf.Oauth2.Github.RedirectURL,
		Scopes:       service.Conf.Oauth2.Github.Scopes,
		Endpoint:     github.Endpoint,
	}
}

func (oc *oauth2Controller) registerRouter() {
	r := oc.r
	flag := false
	if service.Conf.Oauth2.Github.Enable {
		flag = true
		r.GET("/github/callback", oc.githubCallback)
		r.GET("/github/login", oc.githubLogin)
	}
	if !flag {
		panic("no oauth2 enabled")
	}
}

func (oc *oauth2Controller) githubLogin(c *gin.Context) {

	session := sessions.Default(c)
	session.Set("oauth2_state", service.GetRandomString(32))
	if err := session.Save(); err != nil {
		_ = c.AbortWithError(500, err)
		return
	}
	c.Redirect(302, oc.githubOAuthConfig.RedirectURL)
}

func (oc *oauth2Controller) githubCallback(c *gin.Context) {
	c.HTML(200, "oauth2/callback.html", nil)
}
