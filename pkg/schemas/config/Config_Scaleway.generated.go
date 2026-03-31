package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/terraform-kops/terraform-provider-kops/pkg/api/config"
	. "github.com/terraform-kops/terraform-provider-kops/pkg/schemas"
)

var _ = Schema

func ConfigScaleway() *schema.Resource {
	res := &schema.Resource{
		Schema: map[string]*schema.Schema{
			"access_key":            Sensitive(OptionalString()),
			"secret_key":            Sensitive(OptionalString()),
			"project_id":            OptionalString(),
			"region":                OptionalString(),
			"zone":                  OptionalString(),
			"profile":               OptionalString(),
			"kops_controller_image": OptionalString(),
		},
	}

	return res
}

func ExpandConfigScaleway(in map[string]interface{}) config.Scaleway {
	if in == nil {
		panic("expand Scaleway failure, in is nil")
	}
	return config.Scaleway{
		AccessKey: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["access_key"]),
		SecretKey: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["secret_key"]),
		ProjectId: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["project_id"]),
		Region: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["region"]),
		Zone: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["zone"]),
		Profile: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["profile"]),
		KopsControllerImage: func(in interface{}) string {
			return string(ExpandString(in))
		}(in["kops_controller_image"]),
	}
}

func FlattenConfigScalewayInto(in config.Scaleway, out map[string]interface{}) {
	out["access_key"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.AccessKey)
	out["secret_key"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.SecretKey)
	out["project_id"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.ProjectId)
	out["region"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.Region)
	out["zone"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.Zone)
	out["profile"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.Profile)
	out["kops_controller_image"] = func(in string) interface{} {
		return FlattenString(string(in))
	}(in.KopsControllerImage)
}

func FlattenConfigScaleway(in config.Scaleway) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenConfigScalewayInto(in, out)
	return out
}
