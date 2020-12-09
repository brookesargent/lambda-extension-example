package main

import (
	"context"
	"os"

	"github.com/aws/aws-lambda-go/lambdacontext"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/honeycombio/beeline-go"
)

func handler(ctx context.Context) error {
	lc, _ := lambdacontext.FromContext(ctx)
	_, rootSpan := beeline.StartSpan(ctx, "handler")
	defer rootSpan.Send()
	rootSpan.AddTraceField("function_name", lambdacontext.FunctionName)
	rootSpan.AddTraceField("function_version", lambdacontext.FunctionVersion)
	rootSpan.AddTraceField("brand", "oralc")
	rootSpan.AddTraceField("region", os.Getenv("AWS_REGION"))
	rootSpan.AddTraceField("lambda.request_id", lc.AwsRequestID)
	return nil
}

func main() {
	beeline.Init(beeline.Config{
		WriteKey: os.Getenv(""),
	})
	lambda.Start(handler)
}
