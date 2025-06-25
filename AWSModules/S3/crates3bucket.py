import boto3

def create_s3_bucket(bucket_name, region=None):
    s3 = boto3.client('s3', region_name=region)
    
    if region is None:
        response = s3.create_bucket(Bucket=bucket_name)
    else:
        response = s3.create_bucket(
            Bucket=bucket_name,
            CreateBucketConfiguration={'LocationConstraint': region}
        )
    
    return response

print("Creating S3 bucket...")
bucket_name = 'my-unique-bucket-name-shekhar'  # Replace with your
# unique bucket name
region = 'us-west-2'  # Replace with your desired region
response = create_s3_bucket(bucket_name, region)
print("Bucket created successfully:", response)




