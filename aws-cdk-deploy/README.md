# CDK TypeScript project

The `cdk.json` file tells the CDK Toolkit how to execute your app.

## Useful commands

 * `npm run build`   compile typescript to js
 * `npm run watch`   watch for changes and compile
 * `npm run test`    perform the jest unit tests
 * `cdk deploy`      deploy this stack to your default AWS account/region
 * `cdk diff`        compare deployed stack with current state
 * `cdk synth`       emits the synthesized CloudFormation template

## Reference

[AWS CDK Reference Documentation](https://docs.aws.amazon.com/cdk/api/latest)

[The Typescript workshop](https://cdkworkshop.com/20-typescript/20-create-project.html)


# Challenge

## Deploy solution

    export AWS_ACCESS_KEY_ID=<credentials.txt>
    export AWS_SECRET_ACCESS_KEY=<credentials.txt>
    export AWS_DEFAULT_REGION=<credentials.txt>

    cdk deploy
