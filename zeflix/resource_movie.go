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

func ressourceMovie() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceMovieCreate,
		ReadContext:   resourceMovieRead,
		UpdateContext: resourceMovieUpdate,
		DeleteContext: resourceMovieDelete,
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

func resourceMovieCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := &http.Client{Timeout: 10 * time.Second}
	var diags diag.Diagnostics

	// create json body
	movie := make(map[string]interface{})
	movie["name"] = d.Get("name").(string)
	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(&movie)

	// create movie
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/movie", "http://localhost:8080"), b)
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
	resourceMovieRead(ctx, d, m)

	return diags
}

func resourceMovieRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := &http.Client{Timeout: 10 * time.Second}
	var diags diag.Diagnostics

	// use already defined id to retrieve movie
	id := d.Id()
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/movie/%s", "http://localhost:8080", id), nil)
	if err != nil {
		return diag.FromErr(err)
	}

	// retrieve movie request
	r, err := client.Do(req)
	if err != nil {
		return diag.FromErr(err)
	}
	defer r.Body.Close()

	// set movie informations into state
	movie := make(map[string]interface{})
	err = json.NewDecoder(r.Body).Decode(&movie)
	if err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("name", movie["Name"]); err != nil {
		return diag.FromErr(err)
	}

	return diags
}

func resourceMovieUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := &http.Client{Timeout: 10 * time.Second}

	// update if name has changed
	if d.HasChange("name") {
		// prepare request body
		movie := make(map[string]interface{})
		movie["name"] = d.Get("name").(string)
		b := new(bytes.Buffer)
		json.NewEncoder(b).Encode(&movie)
	
		// update movie
		req, err := http.NewRequest("PATCH", fmt.Sprintf("%s/movie/%s", "http://localhost:8080", d.Id()), b)
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
	return resourceMovieRead(ctx, d, m)
}

func resourceMovieDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := &http.Client{Timeout: 10 * time.Second}
	var diags diag.Diagnostics

	// use already defined id to retrieve movie
	id := d.Id()
	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/movie/%s", "http://localhost:8080", id), nil)
	if err != nil {
		return diag.FromErr(err)
	}

	// delete movie request
	r, err := client.Do(req)
	if err != nil {
		return diag.FromErr(err)
	}
	defer r.Body.Close()

	// set movie id to nil for the state
	d.SetId("")
	

	return diags
}
