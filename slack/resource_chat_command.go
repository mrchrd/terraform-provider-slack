package slack

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/nlopes/slack"
)

func resourceChatCommand() *schema.Resource {
	return &schema.Resource{
		Create: resourceChatCommandCreate,
		Read:   resourceChatCommandRead,
		Update: resourceChatCommandUpdate,
		Delete: resourceChatCommandDelete,
		Exists: resourceChatCommandExists,

		Schema: map[string]*schema.Schema{
			"channel": {
				Type:        schema.TypeString,
				Description: "ID of the public channel to execute the command in",
				Required:    true,
			},

			"command": {
				Type:        schema.TypeString,
				Description: "Slash command to be executed",
				Required:    true,
			},

			"text": {
				Type:        schema.TypeString,
				Description: "Additional parameters provided to the slash command",
				Optional:    true,
			},
		},
	}
}

func resourceChatCommandExists(d *schema.ResourceData, meta interface{}) (b bool, e error) {
	api := slack.New(meta.(*Config).Token)

	_, err := api.GetChatCommandByEmail(d.Get("email").(string))
	if err != nil {
		return false, err
	}

	return true, nil
}

func resourceChatCommandCreate(d *schema.ResourceData, meta interface{}) error {
	api := slack.New(meta.(*Config).Token)

	err := api.InviteToTeam(d.Get("team_name").(string), d.Get("first_name").(string), d.Get("last_name").(string), d.Get("email").(string))
	if err != nil {
		return err
	}

	/*
		user, err := api.GetChatCommandByEmail(d.Get("email").(string))
		if err != nil {
			return err
		}
	*/

	d.SetId(d.Get("email").(string))
	return nil
}

func resourceChatCommandRead(d *schema.ResourceData, meta interface{}) error {
	api := slack.New(meta.(*Config).Token)

	_, err := api.GetChatCommandInfo(d.Id())
	if err != nil {
		d.SetId("")
		return nil
	}

	return nil
}

func resourceChatCommandUpdate(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceChatCommandDelete(d *schema.ResourceData, meta interface{}) error {
	api := slack.New(meta.(*Config).Token)

	err := api.DisableChatCommand(d.Get("team_name").(string), d.Id())
	if err != nil {
		return err
	}

	return nil
}
