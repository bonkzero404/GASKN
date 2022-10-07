package utils

import (
	"encoding/json"

	"github.com/gofiber/fiber/v2"
)

type RouteFeature struct {
	route_group string
	route_name  string
	description string
	only_admin  bool
}

type FeatureLists struct {
	Group       string   `json:"group"`
	Method      string   `json:"method"`
	Endpoint    string   `json:"endpoint"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Params      []string `json:"params"`
	OnlyAdmin   bool     `json:"only_admin"`
}

type FeatureUnderGroup struct {
	Method      string   `json:"method"`
	Endpoint    string   `json:"endpoint"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Params      []string `json:"params"`
	OnlyAdmin   bool     `json:"only_admin"`
}

type FeatureGroup struct {
	Group string              `json:"group"`
	Items []FeatureUnderGroup `json:"items"`
}

func (f *RouteFeature) SetGroup(str string) *RouteFeature {
	f.route_group = str
	return f
}

func (f *RouteFeature) SetName(str string) *RouteFeature {
	f.route_name = str
	return f
}

func (f *RouteFeature) SetDescription(str string) *RouteFeature {
	f.description = str
	return f
}

func (f *RouteFeature) SetOnlyAdmin(a bool) *RouteFeature {
	f.only_admin = a
	return f
}

func (f *RouteFeature) Exec() string {
	var iface = make(map[string]interface{})

	if f.route_group != "" {
		iface["group"] = f.route_group
	} else {
		iface["group"] = ""
	}

	if f.route_name != "" {
		iface["name"] = f.route_name
	} else {
		iface["name"] = ""
	}

	if f.description != "" {
		iface["description"] = f.description
	} else {
		iface["description"] = ""
	}

	iface["only_admin"] = f.only_admin

	res, _ := json.Marshal(iface)

	f.cleanup()

	iface["only_admin"] = f.only_admin

	return string(res)
}

func (f *RouteFeature) cleanup() {
	if f.route_group != "" {
		f.route_group = ""
	}

	if f.route_name != "" {
		f.route_name = ""
	}

	if f.description != "" {
		f.description = ""
	}

	if f.only_admin {
		f.only_admin = false
	}

}

func IsJSON(str string) bool {
	var js json.RawMessage
	return json.Unmarshal([]byte(str), &js) == nil
}

func ExtractRouteAsFeatures(app *fiber.App) []FeatureLists {
	var resp []FeatureLists

	for _, items := range app.Stack() {
		for _, item := range items {

			if item.Name != "" && IsJSON(item.Name) {
				var nameInfo = make(map[string]interface{})

				json.Unmarshal([]byte(item.Name), &nameInfo)

				resp = append(resp, FeatureLists{
					Group:       nameInfo["group"].(string),
					Method:      item.Method,
					Endpoint:    item.Path,
					Name:        nameInfo["name"].(string),
					Description: nameInfo["description"].(string),
					Params:      item.Params,
					OnlyAdmin:   nameInfo["only_admin"].(bool),
				})
			}
		}
	}
	return resp
}

func FeaturesGroupLists(app *fiber.App) []FeatureGroup {
	var list = ExtractRouteAsFeatures(app)
	m := make(map[string]bool)
	var a = []string{}
	var resp []FeatureGroup

	for _, item := range list {
		if !m[item.Group] {
			a = append(a, item.Group)
			m[item.Group] = true
		}
	}

	for idx, val := range a {
		resp = append(resp, FeatureGroup{
			Group: val,
			Items: []FeatureUnderGroup{},
		})

		for _, item := range list {
			if val == item.Group {
				resp[idx].Items = append(resp[idx].Items, FeatureUnderGroup{
					Method:      item.Method,
					Endpoint:    item.Endpoint,
					Name:        item.Name,
					Description: item.Description,
					Params:      item.Params,
					OnlyAdmin:   item.OnlyAdmin,
				})
			}
		}
	}

	return resp
}
