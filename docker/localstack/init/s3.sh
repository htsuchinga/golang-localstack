# create bucket
awslocal s3 mb s3://test
awslocal s3 ls

# put sample.json
awslocal s3 cp /json/sample.json s3://test
awslocal s3 ls s3://test

# put sample.csv
awslocal s3 cp /csv/sample.csv s3://test
awslocal s3 ls s3://test
