package onepassword

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceUser() *schema.Resource {
	return &schema.Resource{
		Read:   resourceUserRead,
		Create: resourceUserCreate,
		Delete: resourceUserDelete,
		Importer: &schema.ResourceImporter{
			State: func(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
				err := resourceUserRead(d, meta)
				return []*schema.ResourceData{d}, err
			},
		},
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"email": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceUserRead(d *schema.ResourceData, meta interface{}) error {
	m := meta.(*Meta)
	v, err := m.onePassClient.ReadUser(getID(d))
	if err != nil {
		return err
	}
	d.SetId(v.UUID)
	if err := d.Set("name", v.Name); err != nil {
		return err
	}

	return d.Set("state", v.State)
}

func resourceUserCreate(d *schema.ResourceData, meta interface{}) error {
	m := meta.(*Meta)
	_, err := m.onePassClient.CreateUser(&User{
		Name: d.Get("name").(string),
		Email: d.Get("email").(string),

	})
	if err != nil {
		return err
	}
	return resourceGroupRead(d, meta)
}

func resourceUserDelete(d *schema.ResourceData, meta interface{}) error {
	m := meta.(*Meta)
	err := m.onePassClient.DeleteUser(getID(d))
	if err == nil {
		d.SetId("")
	}
	return err
}
