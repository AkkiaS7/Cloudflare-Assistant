package api

import (
	"Cloudflare-Assistant/handler"
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/gojsonq/v2"
	"strings"
)

func SetZoneRouter(r *gin.RouterGroup) {
	r.GET("/:zone", GetZoneIdByDomain)
	r.GET("/:zone/firewall/access_rules", SetZoneAccessRule)
}

// SetZoneAccessRule 设置IP访问规则
// mode=block|whitelist|challenge|js_challenge|managed_challenge
// value=1.1.1.1
// notes=my_note
func SetZoneAccessRule(c *gin.Context) {
	zoneID := c.Param("zone")
	token := c.Query("token")
	var force bool
	if c.DefaultQuery("force", "false") == "true" {
		force = true
	}
	// 获取zoneID
	if strings.Contains(zoneID, ".") {
		// 当前传入的是域名
		if ID, err := handler.GetZoneIdByDomain(token, zoneID, force); err != nil {
			c.JSON(200, gojsonq.New().FromString(err.Error()).Get())
			return
		} else {
			zoneID = ID
		}
	} else {
		zoneID = c.Param("zone")
	}

	mode := c.Query("mode")
	value := c.Query("value")
	notes := c.Query("notes")
	var ruleID string
	if id, err := handler.SetZoneAccessRule(token, zoneID, mode, value, notes); err != nil {
		c.JSON(200, gojsonq.New().FromString(err.Error()).Get())
		return
	} else {
		ruleID = id
	}
	c.JSON(200, gin.H{
		"success": true,
		"data":    ruleID,
	})
}

// GetZoneIdByDomain 获取zoneID
// https://example.com/api/zones/example.com?token=12345
func GetZoneIdByDomain(c *gin.Context) {
	var zoneID string
	var force bool
	if c.DefaultQuery("force", "false") == "true" {
		force = true
	}
	if strings.Contains(c.Param("zone"), ".") {
		if ID, err := handler.GetZoneIdByDomain(c.Query("token"), c.Param("zone"), force); err != nil {
			c.JSON(200, gojsonq.New().FromString(err.Error()).Get())
			return
		} else {
			zoneID = ID
		}
	} else {
		zoneID = c.Param("zone")
	}
	c.JSON(200, gin.H{
		"success": true,
		"id":      zoneID,
	})
}
