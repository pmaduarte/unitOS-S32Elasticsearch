import logging
from urllib.parse import unquote_plus

import boto3
import botocore

logger = logging.getLogger()
logging.getLogger().setLevel(logging.INFO)

def handler(event, context):

    records = event.get('Records', [])
    if not records:
        logger.warning("no records found")
        return
    
    logger.info(records)
        
    logger.info(f"processing (len(records)) files (s)")
    s3 = boto3.client('s3')
    for record in records:
        _dump_file(s3, record)

 
def _dump_file(s3, record):
    """ Dumps an S3 file to cloudwatch logs """

    try:
        bucket = record['s3']['bucket']['name']
        key = unquote_plus(record['s3']['object']['key'])
        
        logger.info(f"handling bucket {bucket}")
        
        logger.info(f"handling file {key}") 
        s3_object = s3.get_object(Bucket=bucket, Key=key) 
        
        # using iter_lines allow to stream large files in s3
        iter_lines = s3_object['Body'].iter_lines()
        for line in iter_lines:
            _dump_line(line)
        
        s3.delete_object(Bucket=bucket, Key=key)
        logger.info(f"file {key} deleted")

    except botocore.exceptions.ClientError as e:
        if e.response['Error']['Code'] ==  'NoSuchKey':
            logger.error(f"file not found: {key}")
        else:
            raise # something else has gone wrong

def _dump_line(line):
    logger.info(line.decode("utf-8"))