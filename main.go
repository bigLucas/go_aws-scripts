package main

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/lambda"
	"github.com/aws/aws-sdk-go-v2/service/lambda/types"
)

func main() {

	clean_env()

}

func clean_env() {
	cfg := config_aws()
	client := lambda.NewFromConfig(cfg)
	envs := get_envs(*client)

	delete(envs, "ADIQ_SFTP_HOST")
	delete(envs, "ADIQ_SFTP_USER")
	delete(envs, "ADIQ_SFTP_PASSWORD")

	new_envs := types.Environment{
		Variables: envs,
	}

	update_envs(*client, new_envs)
}

func config_aws() aws.Config {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Fatal(err)
	}
	return cfg
}

func get_envs(client lambda.Client) map[string]string {
	output, err := client.GetFunctionConfiguration(context.TODO(), &lambda.GetFunctionConfigurationInput{
		FunctionName: aws.String("ddl-adiq-lucas2-dates-collector"),
	})
	if err != nil {
		log.Fatal(err)
	}
	return output.Environment.Variables
}

func update_envs(client lambda.Client, envs types.Environment) {
	_, err := client.UpdateFunctionConfiguration(context.TODO(), &lambda.UpdateFunctionConfigurationInput{
		FunctionName: aws.String("ddl-adiq-lucas2-dates-collector"),
		Environment:  &envs,
	})
	if err != nil {
		log.Fatal(err)
	}
}
