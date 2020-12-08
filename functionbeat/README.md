# Functionbeat 7.10.0

## Getting Started

To get started with Functionbeat, you need to set up Elasticsearch on
your localhost first. After that, start Functionbeat with:

     ./functionbeat -c functionbeat.yml -e

This will start Functionbeat and send the data to your Elasticsearch
instance. To load the dashboards for Functionbeat into Kibana, run:

    ./functionbeat setup -e

For further steps visit the
[Quick start](https://www.elastic.co/guide/en/beats/functionbeat/7.10/functionbeat-installation-configuration.html) guide.

## Documentation

Visit [Elastic.co Docs](https://www.elastic.co/guide/en/beats/functionbeat/7.10/index.html)
for the full Functionbeat documentation.

## Release notes

https://www.elastic.co/guide/en/beats/libbeat/7.10/release-notes-7.10.0.html


# Challenge

## Deploy solution

    export AWS_ACCESS_KEY_ID=<credentials.txt>
    export AWS_SECRET_ACCESS_KEY=<credentials.txt>
    export AWS_DEFAULT_REGION=<credentials.txt>
    export AWS_DEPLOY_BUCKET=<credentials.txt>
    export AWS_CLOUDWATCH_LOG_TRIGGER=<credentials.txt>

    export CLOUD_ID=<credentials.txt>
    export CLOUD_AUTH=<credentials.txt>
    export CLOUD_ELASTICSEARCH_URL=<credentials.txt>

    curl --user $CLOUD_AUTH -XPUT -H 'Content-Type: application/json' $CLOUD_ELASTICSEARCH_URL/_template/iislogs -d@functionbeat.template.json
    ./functionbeat setup -e
    ./functionbeat deploy stream-cloudwatch-logs-elasticsearch -e

## Remove solution

    curl --user $CLOUD_AUTH -XDELETE $CLOUD_ELASTICSEARCH_URL'/iislogs-*'
    ./functionbeat remove stream-cloudwatch-logs-elasticsearch -e