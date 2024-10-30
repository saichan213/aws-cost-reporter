# aws-cost-reporter

Report AWS account cost per day and per services everyday with e-mail.

## how to create aws resources

- modify .env-sample
  - input your mail address to MAIL_ADDRESS
  - rename .env-sample to .env

```
$ basename `pwd`
terraform
$ make <plan or apply>
```

## how to build lambda function

if you modify lamda function, you can rebuild in this way

```
$ basename `pwd`
lambda
$ make build
```

## CI/CD

You can use Github Actions to CI(terraform plan) and CD(terraform apply)

You must add two secrets on your repository

- ACTIONS_ROLE
  - AWS IAM Role to OIDC
- MAIL_ADDRESS
  - Email address with SNS
