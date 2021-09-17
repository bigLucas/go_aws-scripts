# [WIP] Go scripts for AWS

## What do we need to run?
- We need to have Go installed. 
- We need to have the right credentials of AWS in the system where the script will run, see more [here](https://aws.github.io/aws-sdk-go-v2/docs/getting-started/) and [here](https://aws.github.io/aws-sdk-go-v2/docs/configuring-sdk/#specifying-credentials) to understand how to set the AWS credentials.

## What does this script do?
- We can use this script to clean the environment variables of the AWS lambdas.
- We can use this script to disable AWS cloud watch events.

## How to run?
Inside the project folder run the below command:
```shell
go run *.go
```
## Notes
- The file `event_watch_rule.go` has the **logic to disable the triggers of some lambdas**.
- The file `lambda_env.go` has the logic to **remove the environment variables**.

## To-do
- [ ] turn the scripts into a CLI tool.
- [ ] get the variables as the function names, the names of environment variables, and the stages as inputs in the environment where the script is running or over the CLI arguments.
