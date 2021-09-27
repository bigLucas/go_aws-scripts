package main

import (
	"context"
	"flag"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
)

func main() {
	cfg := config_aws()

	sub_cmd_names := []string{
		"disable_rules",
		"clean_envs",
		"clean_s3",
	}

	// sub-commands
	disableRulesCmd := flag.NewFlagSet(sub_cmd_names[0], flag.ExitOnError)
	cleanEnvsCmd := flag.NewFlagSet(sub_cmd_names[1], flag.ExitOnError)
	cleanS3Cmd := flag.NewFlagSet(sub_cmd_names[2], flag.ExitOnError)

	// flag of the sub-commands
	cleanS3Term := cleanS3Cmd.String("term", "", "Pass a term string to be used as a prefix of the buckets to be deleted (required)")

	if len(os.Args) < 2 {
		fmt.Println("You must pass at least one sub-command (disable_rules, clean_envs, clean_s3)")
		os.Exit(1)
	}

	pos_sub_cmds := 1
	switch os.Args[pos_sub_cmds] {
	case sub_cmd_names[0]:
		disableRulesCmd.Parse(os.Args[2:])
	case sub_cmd_names[1]:
		cleanEnvsCmd.Parse(os.Args[2:])
	case sub_cmd_names[2]:
		cleanS3Cmd.Parse(os.Args[2:])
	default:
		flag.PrintDefaults()
		os.Exit(1)
	}

	if disableRulesCmd.Parsed() {
		Disable_all_rules(cfg)
	}

	if cleanEnvsCmd.Parsed() {
		Clean_all_envs(cfg)
	}

	if cleanS3Cmd.Parsed() {
		if *cleanS3Term == "" {
			cleanEnvsCmd.PrintDefaults()
			os.Exit(1)
		}
		DeleteBucketsByTerm(cfg, *cleanS3Term)
	}
}

func config_aws() aws.Config {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return cfg
}
