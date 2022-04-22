package service

import (
	"context"
	"errors"
	"github.com/cloudflare/cloudflare-go"
)

type CFAPI struct {
	Client *cloudflare.API
}

// NewCFAPI returns a new instance of CFAPI From Token
func NewCFAPIFromToken(token string) (*CFAPI, error) {
	api, err := cloudflare.NewWithAPIToken(token)
	if err != nil {
		return nil, err
	}
	return &CFAPI{
		Client: api,
	}, nil
}

func (cfapi *CFAPI) SetZoneAccessRuleByNotes(zoneId string, rule cloudflare.AccessRule) (string, error) {
	client := cfapi.Client
	rules, err := client.ListZoneAccessRules(context.Background(), zoneId, rule, 1)
	if err != nil {
		return "errorListZoneAccessRules", err
	}
	if len(rules.Result) == 0 {
		res, err := client.CreateZoneAccessRule(context.Background(), zoneId, rule)
		if err != nil {
			return "ErrorCreateZoneAccessRule", err
		}
		return res.Result.ID, nil
	}
	// delete all rules
	for _, rule := range rules.Result {
		_, err := client.DeleteZoneAccessRule(context.Background(), zoneId, rule.ID)
		if err != nil {
			return "ErrorDeleteZoneAccessRule", err
		}
	}
	// create new rule
	res, err := client.CreateZoneAccessRule(context.Background(), zoneId, rule)
	if err != nil {
		return "ErrorCreateZoneAccessRule", err
	}
	return res.Result.ID, nil

}

func (cfapi *CFAPI) GetZoneIDFromDomain(domain string) (string, error) {
	client := cfapi.Client
	zones, err := client.ListZones(context.Background(), domain)
	if err != nil {
		return "", err
	}
	for _, zone := range zones {
		if zone.Name == domain {
			return zone.ID, nil
		}
	}
	return "", errors.New("zone not found")
}
