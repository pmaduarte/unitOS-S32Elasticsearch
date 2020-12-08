#!/usr/bin/env node
import 'source-map-support/register';
import * as cdk from '@aws-cdk/core';
import { AwsCdkDeployStack } from '../lib/aws-cdk-deploy-stack';

const app = new cdk.App();
new AwsCdkDeployStack(app, 'AwsCdkDeployStack');
