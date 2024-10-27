locals {
  sns_topic_name = toset(["total-cost", "service-costs"])
}

resource "aws_sns_topic" "cost_reporter" {
  for_each = local.sns_topic_name
  name     = "${local.service_name}-${each.value}-topic"
}

resource "aws_sns_topic_subscription" "cost_reporter" {
  for_each  = aws_sns_topic.cost_reporter
  topic_arn = each.value.arn
  protocol  = "email"
  endpoint  = var.sns_email_adress
}
