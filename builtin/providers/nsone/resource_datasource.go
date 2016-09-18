package nsone

import (
	"github.com/hashicorp/terraform/helper/schema"

	nsone "gopkg.in/ns1/ns1-go.v2/rest"
	"gopkg.in/ns1/ns1-go.v2/rest/model/data"
)

func dataSourceResource() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"sourcetype": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
		Create: DataSourceCreate,
		Read:   DataSourceRead,
		Update: DataSourceUpdate,
		Delete: DataSourceDelete,
	}
}

func dataSourceToResourceData(d *schema.ResourceData, ds *data.Source) {
	d.SetId(ds.ID)
	d.Set("name", ds.Name)
	d.Set("sourcetype", ds.Type)
}

// DataSourceCreate creates an ns1 datasource
func DataSourceCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*nsone.Client)
	ds := data.NewSource(d.Get("name").(string), d.Get("sourcetype").(string))
	if _, err := client.DataSources.Create(ds); err != nil {
		return err
	}
	dataSourceToResourceData(d, ds)
	return nil
}

// DataSourceRead fetches info for the given datasource from ns1
func DataSourceRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*nsone.Client)
	ds, _, err := client.DataSources.Get(d.Id())
	if err != nil {
		return err
	}
	dataSourceToResourceData(d, ds)
	return nil
}

// DataSourceDelete deteltes the given datasource from ns1
func DataSourceDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*nsone.Client)
	_, err := client.DataSources.Delete(d.Id())
	d.SetId("")
	return err
}

// DataSourceUpdate updates the datasource with given parameters
func DataSourceUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*nsone.Client)
	ds := data.NewSource(d.Get("name").(string), d.Get("sourcetype").(string))
	ds.ID = d.Id()
	if _, err := client.DataSources.Update(ds); err != nil {
		return err
	}
	dataSourceToResourceData(d, ds)
	return nil
}
