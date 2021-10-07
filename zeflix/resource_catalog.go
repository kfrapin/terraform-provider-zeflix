package zeflix

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ressourceCatalog() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceCatalogCreate,
		ReadContext:   resourceCatalogRead,
		UpdateContext: resourceCatalogUpdate,
		DeleteContext: resourceCatalogDelete,
		Schema: map[string]*schema.Schema{
			"id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceCatalogCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := &http.Client{Timeout: 10 * time.Second}

	// get provider configuration (host, token)
	provider := m.(map[string]string)
	endpoint := provider["api_endpoint"]
	token := provider["api_token"]

	var diags diag.Diagnostics

	// create json body
	catalog := make(map[string]interface{})
	catalog["name"] = d.Get("name").(string)
	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(&catalog)

	// create catalog
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/catalog", endpoint), b)
	req.Header.Add("X-Session-Token", token)
	if err != nil {
		return diag.FromErr(err)
	}
	r, err := client.Do(req)
	if err != nil {
		return diag.FromErr(err)
	}
	defer r.Body.Close()

	// retrieved created id and set it
	created := make(map[string]interface{})
	err = json.NewDecoder(r.Body).Decode(&created)
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(created["Id"].(string))

	// populate terraform state after creation
	resourceCatalogRead(ctx, d, m)

	return diags
}

func resourceCatalogRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := &http.Client{Timeout: 10 * time.Second}
	var diags diag.Diagnostics

	// get provider configuration (host, token)
	provider := m.(map[string]string)
	endpoint := provider["api_endpoint"]
	token := provider["api_token"]

	// use already defined id to retrieve catalog
	id := d.Id()
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/catalog/%s", endpoint, id), nil)
	req.Header.Add("X-Session-Token", token)
	if err != nil {
		return diag.FromErr(err)
	}

	// retrieve catalog request
	r, err := client.Do(req)
	if err != nil {
		return diag.FromErr(err)
	}
	defer r.Body.Close()

	// set catalog informations into state
	catalog := make(map[string]interface{})
	err = json.NewDecoder(r.Body).Decode(&catalog)
	if err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("name", catalog["Name"]); err != nil {
		return diag.FromErr(err)
	}

	return diags
}

func resourceCatalogUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := &http.Client{Timeout: 10 * time.Second}

	// update if name has changed
	if d.HasChange("name") {
		// get provider configuration (host, token)
		provider := m.(map[string]string)
		endpoint := provider["api_endpoint"]
		token := provider["api_token"]

		// prepare request body
		catalog := make(map[string]interface{})
		catalog["name"] = d.Get("name").(string)
		b := new(bytes.Buffer)
		json.NewEncoder(b).Encode(&catalog)

		// update catalog
		req, err := http.NewRequest("PATCH", fmt.Sprintf("%s/catalog/%s", endpoint, d.Id()), b)
		req.Header.Add("X-Session-Token", token)
		if err != nil {
			return diag.FromErr(err)
		}
		r, err := client.Do(req)
		if err != nil {
			return diag.FromErr(err)
		}
		defer r.Body.Close()
	}

	// populate terraform state after creation
	return resourceCatalogRead(ctx, d, m)
}

func resourceCatalogDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := &http.Client{Timeout: 10 * time.Second}
	var diags diag.Diagnostics

	// get provider configuration (host, token)
	provider := m.(map[string]string)
	endpoint := provider["api_endpoint"]
	token := provider["api_token"]

	// use already defined id to retrieve catalog
	id := d.Id()
	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/catalog/%s", endpoint, id), nil)
	req.Header.Add("X-Session-Token", token)
	if err != nil {
		return diag.FromErr(err)
	}

	// delete catalog request
	r, err := client.Do(req)
	if err != nil {
		return diag.FromErr(err)
	}
	defer r.Body.Close()

	// set catalog id to nil for the state
	d.SetId("")

	return diags
}
