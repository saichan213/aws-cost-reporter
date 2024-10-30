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
