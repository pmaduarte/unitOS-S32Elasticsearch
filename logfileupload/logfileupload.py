import os
import boto3

AWS_ACCESS_KEY_ID = os.getenv('AWS_ACCESS_KEY_ID')
AWS_SECRET_ACCESS_KEY = os.getenv('AWS_SECRET_ACCESS_KEY')
AWS_FILE_UPLOAD_BUCKET = os.getenv('AWS_FILE_UPLOAD_BUCKET')
file = 'u_ex190720.log'

session = boto3.Session(AWS_ACCESS_KEY_ID, AWS_SECRET_ACCESS_KEY)
s3 = session.resource('s3')
s3.Bucket(AWS_FILE_UPLOAD_BUCKET).upload_file(file, file)