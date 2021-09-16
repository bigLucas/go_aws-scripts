package main

import (
	"context"
	"log"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatchevents"
)

func Disable_all_rules(cfg aws.Config) {
	client := cloudwatchevents.NewFromConfig(cfg)
	rule_names := get_rule_names(*client)
	disable_rules(*client, rule_names)
}

func get_rule_names(client cloudwatchevents.Client) []string {
	result := []string{}
	output, err := client.ListRules(context.TODO(), &cloudwatchevents.ListRulesInput{
		NamePrefix: aws.String("ddl-adiq-"),
	})
	if err != nil {
		log.Fatal(err)
	}
	for _, rule := range output.Rules {
		result = append(result, *rule.Name)
	}
	return result
}

func disable_rules(client cloudwatchevents.Client, rule_names []string) {
	for _, rule_name := range rule_names {
		if !strings.Contains(rule_name, "ari") && !strings.Contains(rule_name, "euler") {
			_, err := client.DisableRule(context.TODO(), &cloudwatchevents.DisableRuleInput{
				Name: aws.String(rule_name),
			})
			if err != nil {
				log.Fatal(err)
			}
		}
	}
}
