package zeflix

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func Provider() func() *schema.Provider {
	return func() *schema.Provider {
		p := &schema.Provider{
			ResourcesMap: map[string]*schema.Resource{
				// no resources yet
			},
			DataSourcesMap: map[string]*schema.Resource{
				"zeflix_movie": datasourceMovie(),
			},
		}
		return p
	}
}
