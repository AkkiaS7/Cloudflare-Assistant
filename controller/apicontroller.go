package controller

import (
	"Cloudflare-Assistant/service"
	"github.com/cloudflare/cloudflare-go"
	"github.com/gin-gonic/gin"
)

type apiController struct {
	r gin.IRoutes
}

func InitApiController(r gin.IRoutes) {
	ac := &apiController{r: r}
	ac.registerRouter()
}

func (ac *apiController) registerRouter() {
	r := ac.r
	r.GET("/zones/:zone/firewall/access_rule/set", ac.setZoneAccessRule)
}

// setZoneAccessRule http://localhost/zones/123456/firewall/access_rule/set?
// notes=<notes>
// ip=<targetIP>
// cftoken=<cftoken>
func (ac *apiController) setZoneAccessRule(c *gin.Context) {
	api, err := cloudflare.NewWithAPIToken(c.Query("cftoken"))
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
	}
	notes := c.Query("notes")
	zoneid := c.Param("zone")
	ip := c.Query("ip")
	rule := cloudflare.AccessRule{
		Notes: notes,
		Mode:  "whitelist",
		Configuration: cloudflare.AccessRuleConfiguration{
			Target: "ip",
			Value:  ip,
		},
	}
	CFAPI := &service.CFAPI{
		Client: api,
	}
	res, err := CFAPI.SetZoneAccessRuleByNotes(zoneid, rule)
	c.JSON(200, gin.H{
		"success": err == nil,
		"message": res,
	})

}
