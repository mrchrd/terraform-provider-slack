package slack

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/nlopes/slack"
)

func resourceUser() *schema.Resource {
	return &schema.Resource{
		Create: resourceUserCreate,
		Read:   resourceUserRead,
		Update: resourceUserUpdate,
		Delete: resourceUserDelete,
		Exists: resourceUserExists,

		Schema: map[string]*schema.Schema{
			"email": {
				Type:        schema.TypeString,
				Description: "User's email address",
				Required:    true,
			},

			"first_name": {
				Type:        schema.TypeString,
				Description: "User's first name",
				Required:    true,
			},

			"last_name": {
				Type:        schema.TypeString,
				Description: "User's last name",
				Required:    true,
			},

			"team_name": {
				Type:        schema.TypeString,
				Description: "Team",
				Required:    true,
			},
		},
	}
}

func resourceUserExists(d *schema.ResourceData, meta interface{}) (b bool, e error) {
	api := slack.New(meta.(*Config).Token)

	_, err := api.GetUserByEmail(d.Get("email").(string))
	if err != nil {
		return false, err
	}

	return true, nil
}

func resourceUserCreate(d *schema.ResourceData, meta interface{}) error {
	api := slack.New(meta.(*Config).Token)

	err := api.InviteToTeam(d.Get("team_name").(string), d.Get("first_name").(string), d.Get("last_name").(string), d.Get("email").(string))
	if err != nil {
		return err
	}

	/*
		user, err := api.GetUserByEmail(d.Get("email").(string))
		if err != nil {
			return err
		}
	*/

	d.SetId(d.Get("email").(string))
	return nil
}

func resourceUserRead(d *schema.ResourceData, meta interface{}) error {
	api := slack.New(meta.(*Config).Token)

	_, err := api.GetUserInfo(d.Id())
	if err != nil {
		d.SetId("")
		return nil
	}

	return nil
}

func resourceUserUpdate(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceUserDelete(d *schema.ResourceData, meta interface{}) error {
	api := slack.New(meta.(*Config).Token)

	err := api.DisableUser(d.Get("team_name").(string), d.Id())
	if err != nil {
		return err
	}

	return nil
}
