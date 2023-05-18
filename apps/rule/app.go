package rule

import (
	"github.com/go-playground/validator/v10"
)

const (
	AppName = "rule"
)

var (
	validate = validator.New()
)

func NewRules() *Rules {
	return &Rules{}
}

func (req *Rules) GroupSet() []*Group {
	gps := make([]*Group, 0, len(req.Data.Groups))
	for i := range req.Data.Groups {
		gp := &Group{
			Id:       req.Data.Groups[i].Id,
			Name:     req.Data.Groups[i].Name,
			File:     req.Data.Groups[i].File,
			Interval: req.Data.Groups[i].Interval,
		}

		gps = append(gps, gp)
	}
	return gps
}

func (req *Rules) RuleSet(group *Group) []*Rule {
	rus := make([]*Rule, 0, len(group.Rules))
	for i := range group.Rules {
		ru := &Rule{
			Id:          group.Rules[i].Id,
			Name:        group.Rules[i].Name,
			Query:       group.Rules[i].Query,
			GroupId:     group.Rules[i].GroupId,
			Labels:      group.Rules[i].Labels,
			Annotations: group.Rules[i].Annotations,
			GroupName:   group.Name,
			Level:       group.Rules[i].Labels["level"],
			Service:     group.Rules[i].Labels["service"],
		}
		rus = append(rus, ru)
	}

	return rus
}
