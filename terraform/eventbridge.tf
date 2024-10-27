data "aws_sfn_state_machine" "cost_reporter" {
  name = "cost-reporter-state-machine"
}

resource "aws_cloudwatch_event_rule" "cost_reporter" {
  name                = "${local.service_name}-step-functions-daily-trigger"
  description         = "Triggers Step Functions ${local.service_name} state machine daily"
  schedule_expression = "cron(0 0 * * ? *)"
}

# EventBridge Target
resource "aws_cloudwatch_event_target" "cost_reporter" {
  rule     = aws_cloudwatch_event_rule.cost_reporter.name
  arn      = data.aws_sfn_state_machine.cost_reporter.arn
  role_arn = aws_iam_role.eventbridge_role.arn
}

# IAM Role for EventBridge
resource "aws_iam_role" "eventbridge_role" {
  name = "${local.service_name}-eventbridge-step-functions-role"

  assume_role_policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Action = "sts:AssumeRole"
        Effect = "Allow"
        Principal = {
          Service = "events.amazonaws.com"
        }
      }
    ]
  })
}

# IAM Policy for EventBridge to invoke Step Functions
resource "aws_iam_role_policy" "eventbridge_policy" {
  name = "${local.service_name}-eventbridge-step-functions-policy"
  role = aws_iam_role.eventbridge_role.id

  policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Effect = "Allow"
        Action = [
          "states:StartExecution"
        ]
        Resource = data.aws_sfn_state_machine.cost_reporter.arn
      }
    ]
  })
}
