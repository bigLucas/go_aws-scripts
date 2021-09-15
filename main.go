package main

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/lambda"
	"github.com/aws/aws-sdk-go-v2/service/lambda/types"
)

func main() {
	// exe_clean_env()
}

func exe_clean_env() {
	stages := []string{"ari", "lucas2", "gabriel", "euler", "igor", "beta", "frontend", "qa", "pagflow", "dev", "paypi", "feeh"}

	for _, stage := range stages {
		func_names := get_func_names(stage)
		for _, func_name := range func_names {
			clean_env(stage, func_name)
		}
	}
}

func get_func_names(stage string) []string {
	return []string{
		fmt.Sprintf("ddl-adiq-%s-dates-collector", stage),
		fmt.Sprintf("ddl-adiq-%s-edi-storage", stage),
		fmt.Sprintf("ddl-adiq-%s-edi-collector-and-parser", stage),
	}
}

func clean_env(stage string, func_name string) {
	cfg := config_aws()
	client := lambda.NewFromConfig(cfg)
	envs, err := get_envs(*client, stage, func_name)
	if err == nil {
		delete(envs, "ADIQ_SFTP_HOST")
		delete(envs, "ADIQ_SFTP_USER")
		delete(envs, "ADIQ_SFTP_PASSWORD")

		new_envs := types.Environment{
			Variables: envs,
		}

		update_envs(*client, new_envs, stage, func_name)
	}
}

func config_aws() aws.Config {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Fatal(err)
	}
	return cfg
}

func get_envs(client lambda.Client, stage string, func_name string) (map[string]string, error) {
	output, err := client.GetFunctionConfiguration(context.TODO(), &lambda.GetFunctionConfigurationInput{
		FunctionName: aws.String(func_name),
	})
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return output.Environment.Variables, nil
}

func update_envs(client lambda.Client, envs types.Environment, stage string, func_name string) {
	_, err := client.UpdateFunctionConfiguration(context.TODO(), &lambda.UpdateFunctionConfigurationInput{
		FunctionName: aws.String(func_name),
		Environment:  &envs,
	})
	if err != nil {
		log.Fatal(err)
	}
}
