package zeflix

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func datasourceCatalog() *schema.Resource {
	return &schema.Resource{
		ReadContext: datasourceCatalogRead,
		Schema: map[string]*schema.Schema{
			"id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"name": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func datasourceCatalogRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := &http.Client{Timeout: 10 * time.Second}
	var diags diag.Diagnostics
	id := d.Get("id").(string)
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/catalog/%s", "http://localhost:8080", id), nil)
	if err != nil {
		return diag.FromErr(err)
	}

	r, err := client.Do(req)
	if err != nil {
		return diag.FromErr(err)
	}
	defer r.Body.Close()

	catalog := make(map[string]interface{})
	err = json.NewDecoder(r.Body).Decode(&catalog)
	if err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("name", catalog["Name"]); err != nil {
		return diag.FromErr(err)
	}

	d.SetId(catalog["Id"].(string))

	return diags
}
