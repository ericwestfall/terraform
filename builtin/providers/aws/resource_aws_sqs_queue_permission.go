package aws

import (
	"fmt"
	//"log"

	"github.com/aws/aws-sdk-go/aws"
	//"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceAwsSqsQueuePermission() *schema.Resource {
	return &schema.Resource{
		Create: resourceAwsSqsQueuePermissionCreate,
		Read:   resourceAwsSqsQueuePermissionRead,
		Update: nil,
		Delete: resourceAwsSqsQueuePermissionDelete,

		Schema: map[string]*schema.Schema{
			"queue_url": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"label": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"actions": &schema.Schema{
				Type:     schema.TypeList,
				ForceNew: true,
				Required: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"account_ids": &schema.Schema{
				Type:     schema.TypeList,
				ForceNew: true,
				Required: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

// if v, ok := d.GetOk("allowed_account_ids"); ok {
//   config.AllowedAccountIds = v.(*schema.Set).List()
// }

// "allowed_account_ids": {
//   Type:          schema.TypeSet,
//   Elem:          &schema.Schema{Type: schema.TypeString},
//   Optional:      true,
//   ConflictsWith: []string{"forbidden_account_ids"},
//   Set:           schema.HashString,
// },

func resourceAwsSqsQueuePermissionRead(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceAwsSqsQueuePermissionDelete(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceAwsSqsQueuePermissionCreate(d *schema.ResourceData, meta interface{}) error {
	conn := meta.(*AWSClient).sqsconn

	url := d.Get("queue_url").(string)
	label := d.Get("label").(string)
	accountIds := d.Get("account_ids")
	actions := d.Get("actions")

	params := &sqs.AddPermissionInput{
		AWSAccountIds: expandStringList(accountIds.([]interface{})),
		Actions:       expandStringList(actions.([]interface{})),
		Label:         aws.String(label),
		QueueUrl:      aws.String(url),
	}

	resp, err := conn.AddPermission(params)
	if err != nil {
		return fmt.Errorf("Error adding permission.")
	}

	fmt.Printf("Response: %s", *resp)

	d.SetId("sqs-permission-" + label + "-" + url)

	return nil
}

// func resourceAwsSqsQueuePolicyUpsert(d *schema.ResourceData, meta interface{}) error {
// 	conn := meta.(*AWSClient).sqsconn
// 	url := d.Get("queue_url").(string)
//
// 	_, err := conn.SetQueueAttributes(&sqs.SetQueueAttributesInput{
// 		QueueUrl: aws.String(url),
// 		Attributes: aws.StringMap(map[string]string{
// 			"Policy": d.Get("policy").(string),
// 		}),
// 	})
// 	if err != nil {
// 		return fmt.Errorf("Error updating SQS attributes: %s", err)
// 	}
//
// 	d.SetId("sqs-policy-" + url)
//
// 	return resourceAwsSqsQueuePolicyRead(d, meta)
// }
