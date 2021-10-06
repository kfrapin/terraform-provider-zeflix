package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func init() {
}

func New(version string) func() *schema.Provider {
	return func() *schema.Provider {
		p := &schema.Provider{
			ResourcesMap: map[string]*schema.Resource{
				// no resources yet
			},
			DataSourcesMap: map[string]*schema.Resource{
				"zeflix_catalog": datasourceCatalogRead,
			},
		}
		return p
	}
}
