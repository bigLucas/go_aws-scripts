package main

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func DeleteBucketsByTerm(cfg aws.Config, term string) {
	client := s3.NewFromConfig(cfg)
	bucket_names := getBucketNames(*client)
	filtered_names := filterBucketsByTerm(bucket_names, term)
	deleteAllBucketsByName(*client, filtered_names)
}

func getBucketNames(client s3.Client) []string {
	result := []string{}
	out, err := client.ListBuckets(context.TODO(), &s3.ListBucketsInput{})
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	for _, bucket := range out.Buckets {
		result = append(result, *bucket.Name)
	}
	return result
}

func filterBucketsByTerm(buckets []string, term string) []string {
	result := []string{}
	for _, bucket := range buckets {
		if strings.HasPrefix(bucket, term) {
			result = append(result, bucket)
		}
	}
	return result
}

func deleteAllBucketsByName(client s3.Client, bucket_names []string) {
	for _, bucket_name := range bucket_names {
		_, err := client.DeleteBucket(context.TODO(), &s3.DeleteBucketInput{
			Bucket: aws.String(bucket_name),
		})
		if err != nil {
			fmt.Println(err)
		}
	}
}
