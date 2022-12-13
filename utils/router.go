package utils

import (
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
)

type RouteFeature struct {
	RouteGroup     string `json:"route_group"`
	RouteName      string `json:"route_name"`
	Description    string `json:"description"`
	Tenant         bool   `json:"group"`
	DescriptionKey string
}

type FeatureLists struct {
	Group       string   `json:"group"`
	Method      string   `json:"method"`
	Endpoint    string   `json:"endpoint"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Params      []string `json:"params"`
	Tenant      bool     `json:"tenant"`
}

type FeatureUnderGroup struct {
	Method      string   `json:"method"`
	Endpoint    string   `json:"endpoint"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Params      []string `json:"params"`
	Tenant      bool     `json:"tenant"`
}

type FeatureGroup struct {
	Group string              `json:"group"`
	Items []FeatureUnderGroup `json:"items"`
}

type GasknRouter struct {
	Router        fiber.Router
	SetRoute      fiber.Router
	GroupName     string
	Ctx           *fiber.Ctx
	RouterOptions RouteFeature
}

func (f *GasknRouter) Set(app fiber.Router) *GasknRouter {
	f.Router = app
	f.RouterOptions.Tenant = false
	return f
}

func (f *GasknRouter) Group(prefix string, handlers ...fiber.Handler) *GasknRouter {
	f.Router = f.Router.Group(prefix, handlers...)
	return f
}

func (f *GasknRouter) SetGroupName(name string) *GasknRouter {
	f.GroupName = name
	return f
}

func (f *GasknRouter) Get(path string, handlers ...fiber.Handler) *GasknRouter {
	f.SetRoute = f.Router.Get(path, handlers...)
	return f
}

func (f *GasknRouter) Post(path string, handlers ...fiber.Handler) *GasknRouter {
	f.SetRoute = f.Router.Post(path, handlers...)
	return f
}

func (f *GasknRouter) Patch(path string, handlers ...fiber.Handler) *GasknRouter {
	f.SetRoute = f.Router.Patch(path, handlers...)
	return f
}

func (f *GasknRouter) Put(path string, handlers ...fiber.Handler) *GasknRouter {
	f.SetRoute = f.Router.Put(path, handlers...)
	return f
}

func (f *GasknRouter) Delete(path string, handlers ...fiber.Handler) *GasknRouter {
	f.SetRoute = f.Router.Delete(path, handlers...)
	return f
}

func (f *GasknRouter) Options(path string, handlers ...fiber.Handler) *GasknRouter {
	f.SetRoute = f.Router.Options(path, handlers...)
	return f
}

func (f *GasknRouter) Head(path string, handlers ...fiber.Handler) *GasknRouter {
	f.SetRoute = f.Router.Head(path, handlers...)
	return f
}

func (f *GasknRouter) Trace(path string, handlers ...fiber.Handler) *GasknRouter {
	f.SetRoute = f.Router.Trace(path, handlers...)
	return f
}

func (f *GasknRouter) SetRouteName(name string) *GasknRouter {
	f.RouterOptions.RouteName = name
	return f
}

func (f *GasknRouter) SetRouteDescription(desc string) *GasknRouter {
	f.RouterOptions.Description = desc
	return f
}

func (f *GasknRouter) SetRouteDescriptionKeyLang(keyLang string) *GasknRouter {
	f.RouterOptions.DescriptionKey = keyLang
	return f
}

func (f *GasknRouter) ImplementDescriptionLang() func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		fmt.Print("AIAIA")
		f.RouterOptions.Description = Lang(c, f.RouterOptions.DescriptionKey)
		return c.Next()
	}
}

func (f *GasknRouter) SetRouteTenant(v bool) *GasknRouter {
	f.RouterOptions.Tenant = v
	return f
}

func (f *GasknRouter) Execute() fiber.Router {
	var iface = make(map[string]interface{})

	f.RouterOptions.RouteGroup = f.GroupName

	if f.RouterOptions.RouteGroup != "" {
		iface["group"] = f.RouterOptions.RouteGroup
	} else {
		iface["group"] = ""
	}

	if f.RouterOptions.RouteName != "" {
		iface["name"] = f.RouterOptions.RouteName
	} else {
		iface["name"] = ""
	}

	if f.RouterOptions.DescriptionKey != "" {
		iface["description_key"] = f.RouterOptions.DescriptionKey
	} else {
		iface["description_key"] = ""
	}

	if f.RouterOptions.Description != "" {
		iface["description"] = f.RouterOptions.Description
	} else {
		iface["description"] = ""
	}

	iface["tenant"] = f.RouterOptions.Tenant

	res, _ := json.Marshal(iface)

	f.cleanup()

	iface["tenant"] = f.RouterOptions.Tenant

	f.SetRoute.Name(string(res))

	return f.SetRoute
}

func (f *GasknRouter) cleanup() {
	if f.RouterOptions.RouteGroup != "" {
		f.RouterOptions.RouteGroup = ""
	}

	if f.RouterOptions.RouteName != "" {
		f.RouterOptions.RouteName = ""
	}

	if f.RouterOptions.Description != "" {
		f.RouterOptions.Description = ""
	}

	if f.RouterOptions.DescriptionKey != "" {
		f.RouterOptions.DescriptionKey = ""
	}

	if f.RouterOptions.Tenant {
		f.RouterOptions.Tenant = false
	}

}

func IsJSON(str string) bool {
	var js json.RawMessage
	return json.Unmarshal([]byte(str), &js) == nil
}

func ExtractRouteAsFeatures(c *fiber.Ctx, isTenant bool) []FeatureLists {
	var resp []FeatureLists

	for _, items := range c.App().Stack() {
		for _, item := range items {

			if item.Name != "" && IsJSON(item.Name) {
				var nameInfo = make(map[string]interface{})

				err := json.Unmarshal([]byte(item.Name), &nameInfo)
				if err != nil {
					return nil
				}

				var descLang = nameInfo["description_key"].(string)
				var desc = ""

				if descLang != "" {
					desc = Lang(c, nameInfo["description_key"].(string))
				} else {
					desc = nameInfo["description"].(string)
				}

				if nameInfo["tenant"].(bool) == true && isTenant == true {
					resp = append(resp, FeatureLists{
						Group:       nameInfo["group"].(string),
						Method:      item.Method,
						Endpoint:    item.Path,
						Name:        nameInfo["name"].(string),
						Description: desc,
						Params:      item.Params,
						Tenant:      nameInfo["tenant"].(bool),
					})
				} else if isTenant == false {
					resp = append(resp, FeatureLists{
						Group:       nameInfo["group"].(string),
						Method:      item.Method,
						Endpoint:    item.Path,
						Name:        nameInfo["name"].(string),
						Description: desc,
						Params:      item.Params,
						Tenant:      nameInfo["tenant"].(bool),
					})
				}
			}
		}
	}
	return resp
}

func FeaturesGroupLists(c *fiber.Ctx, isTenant bool) []FeatureGroup {
	var list = ExtractRouteAsFeatures(c, isTenant)
	m := make(map[string]bool)
	var a []string
	var resp []FeatureGroup

	for _, item := range list {
		if !m[item.Group] && item.Tenant == true && isTenant == true {
			a = append(a, item.Group)
			m[item.Group] = true
		} else if !m[item.Group] && isTenant == false {
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
			if val == item.Group && item.Tenant == true && isTenant == true {
				resp[idx].Items = append(resp[idx].Items, FeatureUnderGroup{
					Method:      item.Method,
					Endpoint:    item.Endpoint,
					Name:        item.Name,
					Description: item.Description,
					Params:      item.Params,
					Tenant:      item.Tenant,
				})
			} else if val == item.Group && isTenant == false {
				resp[idx].Items = append(resp[idx].Items, FeatureUnderGroup{
					Method:      item.Method,
					Endpoint:    item.Endpoint,
					Name:        item.Name,
					Description: item.Description,
					Params:      item.Params,
					Tenant:      item.Tenant,
				})
			}
		}
	}

	return resp
}
