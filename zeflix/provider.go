package zeflix

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func Provider() func() *schema.Provider {
	return func() *schema.Provider {
		p := &schema.Provider{
			ResourcesMap: map[string]*schema.Resource{
				"zeflix_catalog": ressourceCatalog(),
			},
			DataSourcesMap: map[string]*schema.Resource{
				"zeflix_catalog": datasourceCatalog(),
			},
			// provider configuration
			Schema: map[string]*schema.Schema{
				"api_endpoint": {
					Type: schema.TypeString,
					Default:  "http://localhost:8080",
					Optional: true,
				},
				"api_token": {
					Type: schema.TypeString,
					Required: true,
				},
			},
			// add provider configuration in contect
			ConfigureContextFunc: providerConfigure,
		}
		return p
	}
}

func providerConfigure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	return map[string]string{
		"api_endpoint": d.Get("api_endpoint").(string),
		"api_token":    d.Get("api_token").(string),
	}, nil
}
