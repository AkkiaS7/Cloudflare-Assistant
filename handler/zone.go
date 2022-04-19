package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/thedevsaddam/gojsonq/v2"
)

// GetZoneIdByDomain 从域名获取zoneId, force为true则绕过本地缓存
func GetZoneIdByDomain(token string, domain string, force bool) (string, error) {
	if !force {
		if ZoneIdList[domain] != "" {
			return ZoneIdList[domain], nil
		}
	}
	path := "/zones?name=" + domain
	headerMap := map[string]string{
		"Authorization": "Bearer " + token,
	}
	if body, err := sendReq("GET", path, headerMap, ""); err == nil {
		fmt.Println(body)
		if gojsonq.New().FromString(body).Find("success").(bool) == true {
			if res := gojsonq.New().FromString(body).Find("result.[0].id"); res != nil {
				ZoneIdList[domain] = res.(string)
				return res.(string), nil
			} else {
				return "", errors.New("{\"message\":\"domain not found\",\"success\":false}")
			}
		} else {
			return "", errors.New(body)
		}
	} else {
		return "", errors.New("{\"message\":\"" + err.Error() + "\",\"success\":false}")
	}
}

func SetZoneAccessRule(token string, zoneId string, mode string, value string, note string) (string, error) {
	// 先查询是否存在匹配当前IP或note的规则
	res, _ := GetZoneAccessRule(token, zoneId, mode, value, note)
	if res == "" {
		// 不存在则新增
		var err error
		res, err = AddZoneAccessRule(token, zoneId, mode, value, note)
		if err != nil {
			return "", err
		} else {
			return res, nil
		}
	}
	// 存在则修改现有规则
	return UpdateZoneAccessRule(token, zoneId, res, mode, value, note)
}

// GetZoneAccessRule 查询该ip或note在当前zone的访问规则
func GetZoneAccessRule(token string, zoneId string, mode string, value string, note string) (string, error) {
	path := "/zones/" + zoneId + "/firewall/access_rules/rules"
	// 先尝试根据IP查询
	headerMap := map[string]string{
		"Authorization":       "Bearer " + token,
		"configuration.value": value,
	}
	if body, err := sendReq("GET", path, headerMap, ""); err == nil {
		fmt.Println(body)
		if gojsonq.New().FromString(body).Find("success").(bool) == true {
			if res := gojsonq.New().FromString(body).Find("result.[0].id"); res != nil {
				return res.(string), nil
			} else {
				// 如果没有查询到，则尝试根据note查询
				headerMap = map[string]string{
					"Authorization": "Bearer " + token,
					"notes":         note,
				}
				if body, err := sendReq("GET", path, headerMap, ""); err == nil {
					fmt.Println(body)
					if gojsonq.New().FromString(body).Find("success").(bool) == true {
						if res := gojsonq.New().FromString(body).Find("result.[0].id"); res != nil {
							return res.(string), nil
						} else {
							return "", errors.New("{\"message\":\"rule not found\",\"success\":false}")
						}
					} else {
						return "", errors.New(body)
					}
				} else {
					return "", errors.New("{\"message\":\"" + err.Error() + "\",\"success\":false}")
				}
			}
		} else {
			return "", errors.New(body)
		}
	} else {
		return "", errors.New("{\"message\":\"" + err.Error() + "\",\"success\":false}")
	}
}

// AddZoneAccessRule
func AddZoneAccessRule(token string, zoneId string, mode string, value string, note string) (string, error) {
	return "", nil
}

// UpdateZoneAccessRule
func UpdateZoneAccessRule(token string, zoneId string, ruleID string, mode string, value string, note string) (string, error) {
	path := "/zones/" + zoneId + "/firewall/access_rules/rules/" + ruleID
	headerMap := map[string]string{
		"Authorization": "Bearer " + token,
		"Content-Type":  "application/json",
	}
	type bodyBuilder struct {
		Mode string `json:"mode"`
		Note string `json:"note"`
	}
	bodyB, _ := json.Marshal(bodyBuilder{Mode: mode, Note: note})
	body := string(bodyB)
	fmt.Println(body)
	if body, err := sendReq("PATCH", path, headerMap, body); err == nil {
		fmt.Println(body)
		if gojsonq.New().FromString(body).Find("success").(bool) == true {
			return "", nil
		} else {
			return "", errors.New(body)
		}
	} else {
		return "", errors.New("{\"message\":\"" + err.Error() + "\",\"success\":false}")
	}
}
