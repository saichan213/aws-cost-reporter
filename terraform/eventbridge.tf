resource "aws_cloudwatch_event_rule" "cost_reporter" {
  name                = "${local.service_name}-step-functions-daily-trigger"
  description         = "Triggers Step Functions ${local.service_name} state machine daily"
  schedule_expression = "cron(0 0 * * ? *)"
}

# EventBridge Target
resource "aws_cloudwatch_event_target" "cost_reporter" {
  rule     = aws_cloudwatch_event_rule.cost_reporter.name
  arn      = aws_lambda_function.cost_reporter.arn
  role_arn = aws_iam_role.eventbridge_role.arn
}
