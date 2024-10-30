include .env
plan:
    terraform plan -var "sns_email_adress=$(MAIL_ADDRESS)"

apply:
    terraform apply -var "sns_email_adress=$(MAIL_ADDRESS)"
