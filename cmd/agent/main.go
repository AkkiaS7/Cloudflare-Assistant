package main

import (
	"Cloudflare-Assistant/service"
	"Cloudflare-Assistant/tool"
	"github.com/cloudflare/cloudflare-go"
	"log"
)

var (
	Conf  *service.AgentConfig
	CFAPI *service.CFAPI
)

func init() {
	Conf = &service.AgentConfig{}
	err := Conf.Read("config/agent.yaml")
	if err != nil {
		panic(err)
	}
	CFAPI, err = service.NewCFAPIFromToken(Conf.Token)
	if err != nil {
		panic(err)
	}
}

func main() {
	setZoneAccessWhiteList()
}

// setZoneAccessWhiteList sets the access whitelist for all zone given by config.
func setZoneAccessWhiteList() {
	for zone, notes := range Conf.ZoneAccessWhitelist {
		zoneID, err := getZoneID(zone)
		if err != nil {
			log.Printf("Error getting zoneID for %s: %s", zone, err)
			continue
		}
		ip, err := tool.GetIPFromAPI()
		if err != nil {
			log.Printf("Error getting IP from API: %s", err)
			continue
		}
		rule := cloudflare.AccessRule{
			Notes: notes,
			Mode:  "whitelist",
			Configuration: cloudflare.AccessRuleConfiguration{
				Target: "ip",
				Value:  ip,
			},
		}
		res, err := CFAPI.SetZoneAccessRuleByNotes(zoneID, rule)
		if err == nil {
			log.Printf("Set access rule for %s to %s, ruleID:%s", zone, ip, res)
		} else {
			log.Printf("Error setting access rule for %s: %s", zone, err.Error())
		}
	}
}

// getZoneID returns the zoneID of the given zone name.
func getZoneID(domain string) (string, error) {
	if _, ok := Conf.ZoneIDList[domain]; !ok {
		zoneID, err := CFAPI.GetZoneIDFromDomain(domain)
		if err != nil {
			return "", err
		}
		Conf.ZoneIDList[domain] = zoneID
		return zoneID, nil
	}
	return Conf.ZoneIDList[domain], nil

}
