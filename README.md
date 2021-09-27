# [WIP] Go scripts for AWS

## What do we need to run?
- We need to have Go installed. 
- We need to have the right credentials of AWS in the system where the script will run, see more [here](https://aws.github.io/aws-sdk-go-v2/docs/getting-started/) and [here](https://aws.github.io/aws-sdk-go-v2/docs/configuring-sdk/#specifying-credentials) to understand how to set the AWS credentials.

## What does this script do?
- We can use this script to clean the environment variables of the AWS lambdas.
- We can use this script to disable AWS cloud watch events.
- We can use this script to delete S3 buckets by a term used as a prefix.

## How to run?
Inside the project folder run the below command:
```shell
go run *.go
# OR
go build *.go
./aws_utils {disable_rules | clean_envs | clean_s3}
```
## Notes
- The file `event_watch_rule.go` has the **logic to disable the triggers of some lambdas**.
- The file `lambda_env.go` has the logic to **remove the environment variables**.
- The file `clean_s3_buckets.go` has the logic to **delete S3 buckets by a term**.

## To-do
- [X] turn the scripts into a CLI tool.
- [ ] [clean_s3_buckets] make the buckets empty before the delete operation.
- [ ] [lambda_env | event_watch_rule] get the variables as the function names, the names of environment variables, and the stages as inputs in the environment where the script is running or over the CLI arguments.
