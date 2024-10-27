resource "aws_sns_topic" "cost_reporter" {
  name = "${local.service_name}-topic"
}

resource "aws_sns_topic_subscription" "cost_reporter" {
  topic_arn = aws_sns_topic.cost_reporter.arn
  protocol  = "email"
  endpoint  = var.sns_email_adress
}
