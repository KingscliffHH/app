# project-dashboard

## Infra Deploy

Deploy Infra once to create all necessary resources.
Copy `VPCEndpoint` from Cloudformation Outputs and add a Private Endpoint on MongoDB to allow the database connection.
After the endpoint has been created successfully, create a database connection on MongoDB using the Private Endpoint created.
Copy the URI, set the password and update MONGODB_URI on GithubActions.
Deploy the infra again and the environment will be completely configured.

After the infra has been deployed successfully, edit the loadbalancer to add HTTPS redirect to HTTP.
Add the SSL certificate. Edit the security group of the loadbalancer to listen to port 443.

## AWS

Use `export AWS_PROFILE=<profile>` to set the current profile.
Available profiles:

- innov
- dev
- prod

Use `aws sso login` get credentials. This will have to be used everyday as the credentials are temporary.

Use the command below to export credentials to env vars.

```bash
eval "$(aws configure export-credentials --profile innov --format env)"

```

## OBS

The default docker sample from elasticbeanstalk-samples-${AWS::Region} doesn't work. Had to download from https://docs.aws.amazon.com/elasticbeanstalk/latest/dg/tutorials.html and upload manually.
This wont be a problem for our app as We'll build the app ourselves

## Backend

You must be on the folder `backend`

To run the application

```bash
make run
```

To run a stress test on the app

```bash
make stress
```

## MongoDB

- Create Serverless DB on Mongo
- Create VPC Endpoint on AWS
- Create PrivateLink on Mongo

## Local .env Example

Create a `.env` file with the folowing content on the backend folder.

```
MONGODB_URI="mongodb://localhost:27017"
MONGODB_NAME="test"
```

................
