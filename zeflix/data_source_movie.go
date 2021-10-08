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

func datasourceMovie() *schema.Resource {
	return &schema.Resource{
		ReadContext: datasourceMovieRead,
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

func datasourceMovieRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := &http.Client{Timeout: 10 * time.Second}
	var diags diag.Diagnostics
	id := d.Get("id").(string)
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/movie/%s", "http://localhost:8080", id), nil)
	if err != nil {
		return diag.FromErr(err)
	}

	r, err := client.Do(req)
	if err != nil {
		return diag.FromErr(err)
	}
	defer r.Body.Close()

	Movie := make(map[string]interface{})
	err = json.NewDecoder(r.Body).Decode(&Movie)
	if err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("name", Movie["Name"]); err != nil {
		return diag.FromErr(err)
	}

	d.SetId(Movie["Id"].(string))

	return diags
}
