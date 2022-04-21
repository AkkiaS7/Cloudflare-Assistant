package controller

import (
	"context"
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
	r.GET("/zones/firewall/access_rule/set", ac.setZoneAccessRule)
}

// setZoneAccessRule http://localhost/zones/firewall/access_rule/set?
// zone=<zone_id|domain>
// mode=<mode>
// notes=<notes>
// target=<targetIP>
// cftoken=<cftoken>
func (ac *apiController) setZoneAccessRule(c *gin.Context) {
	cftoken := c.Query("cftoken")
	if cftoken == "" {
		c.JSON(400, gin.H{
			"code":    400,
			"message": "cftoken is required",
		})
		return
	}
	cfapi, err := cloudflare.NewWithAPIToken(cftoken)
	if err != nil {
		c.JSON(400, gin.H{
			"code":    400,
			"error":   "get cloudflare api error",
			"message": err.Error(),
		})
		return
	}
	zoneId, err := cfapi.ZoneIDByName(c.Query("zone"))
	searchBy := cloudflare.AccessRule{
		Notes: c.Query("notes"),
	}
	if err != nil {
		c.JSON(400, gin.H{
			"code":    400,
			"error":   "get zone id error",
			"message": err.Error(),
		})
		return
	}
	res, err := cfapi.ListZoneAccessRules(context.Background(), zoneId, searchBy, 1)
	if err != nil {
		c.JSON(400, gin.H{
			"code":    400,
			"error":   "ErrorListZoneAccessRules",
			"message": err.Error(),
		})
		return
	}
	// delete all existing access rule
	for _, rule := range res.Result {
		if _, err := cfapi.DeleteZoneAccessRule(context.Background(), zoneId, rule.ID); err != nil {
			c.JSON(400, gin.H{
				"code":    400,
				"error":   "ErrorDeleteZoneAccessRule",
				"message": err.Error(),
			})
			return
		}
	}
	// create a new one
	newRule, err := cfapi.CreateZoneAccessRule(context.Background(), zoneId, cloudflare.AccessRule{
		Mode:  c.Query("mode"),
		Notes: c.Query("notes"),
		Configuration: cloudflare.AccessRuleConfiguration{
			Target: "ip",
			Value:  c.Query("target"),
		},
	})
	if err != nil {
		c.JSON(400, gin.H{
			"code":    400,
			"error":   "ErrorCreateZoneAccessRule",
			"message": "E" + err.Error(),
		})
		return
	}
	c.JSON(200, newRule.Result)

}
