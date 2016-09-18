package nsone

import (
	"github.com/hashicorp/terraform/helper/schema"

	nsone "gopkg.in/ns1/ns1-go.v2/rest"
	"gopkg.in/ns1/ns1-go.v2/rest/model/account"
)

func teamResource() *schema.Resource {
	s := map[string]*schema.Schema{
		"id": &schema.Schema{
			Type:     schema.TypeString,
			Computed: true,
		},
		"name": &schema.Schema{
			Type:     schema.TypeString,
			Required: true,
		},
	}
	s = addPermsSchema(s)
	return &schema.Resource{
		Schema: s,
		Create: TeamCreate,
		Read:   TeamRead,
		Update: TeamUpdate,
		Delete: TeamDelete,
	}
}

func teamToResourceData(d *schema.ResourceData, t *account.Team) error {
	d.SetId(t.ID)
	d.Set("name", t.Name)
	permissionsToResourceData(d, t.Permissions)
	return nil
}

func resourceDataToTeam(u *account.Team, d *schema.ResourceData) error {
	u.ID = d.Id()
	u.Name = d.Get("name").(string)
	u.Permissions = resourceDataToPermissions(d)
	return nil
}

// TeamCreate creates the given team in ns1
func TeamCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*nsone.Client)
	mj := account.Team{}
	if err := resourceDataToTeam(&mj, d); err != nil {
		return err
	}
	if _, err := client.Teams.Create(&mj); err != nil {
		return err
	}
	return teamToResourceData(d, &mj)
}

// TeamRead reads the team data from ns1
func TeamRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*nsone.Client)
	mj, _, err := client.Teams.Get(d.Id())
	if err != nil {
		return err
	}
	teamToResourceData(d, mj)
	return nil
}

// TeamDelete deletes the given team from ns1
func TeamDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*nsone.Client)
	_, err := client.Teams.Delete(d.Id())
	d.SetId("")
	return err
}

// TeamUpdate updates the given team in ns1
func TeamUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*nsone.Client)
	mj := account.Team{
		ID: d.Id(),
	}
	if err := resourceDataToTeam(&mj, d); err != nil {
		return err
	}
	if _, err := client.Teams.Update(&mj); err != nil {
		return err
	}
	teamToResourceData(d, &mj)
	return nil
}
