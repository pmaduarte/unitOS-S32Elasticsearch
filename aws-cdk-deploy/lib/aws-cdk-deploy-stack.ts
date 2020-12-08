import * as path from 'path';
import * as lambda from '@aws-cdk/aws-lambda';
import * as cdk from '@aws-cdk/core';
import * as s3 from '@aws-cdk/aws-s3';
import * as log from '@aws-cdk/aws-logs';

import { S3EventSource } from '@aws-cdk/aws-lambda-event-sources';

export class AwsCdkDeployStack extends cdk.Stack {
  
  constructor(scope: cdk.Construct, id: string, props?: cdk.StackProps) {
    super(scope, id, props);

    const bucket = new s3.Bucket(this, 'BucketOutSystemsUnit', {
      blockPublicAccess: s3.BlockPublicAccess.BLOCK_ALL
    });

    const fn = new lambda.Function(this, 'read-log-lambda-handler', {
      code: lambda.Code.fromAsset(path.join(__dirname, '../assets/read-log-lambda-handler')),
      handler: 'read-log-lambda-handler.handler',
      runtime: lambda.Runtime.PYTHON_3_8,
      timeout: cdk.Duration.minutes(3),
      logRetention: log.RetentionDays.ONE_MONTH
    });

    bucket.grantReadWrite(fn);
    
    //Trigger
    fn.addEventSource(new S3EventSource(bucket, {
      events: [ s3.EventType.OBJECT_CREATED]
    }));
    
    // Export BucketName
    new cdk.CfnOutput(this, "BucketNameEnvironmentVariable", {
      value: `export AWS_FILE_UPLOAD_BUCKET=${bucket.bucketName}`,
    });    

    // Export FunctionLogGroup
    new cdk.CfnOutput(this, "FunctionLogGroupEnvironmentVariable", {
      value: `export AWS_CLOUDWATCH_LOG_TRIGGER=/aws/lambda/${fn.functionName}` ,
    });  
  

  }
}
