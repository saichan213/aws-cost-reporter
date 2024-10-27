data "archive_file" "test_terraform" {
  type        = "zip"
  source_dir  = "../lambda"
  output_path = "cost-reporter.zip"
}

# AWSへ作るlambda function
resource "aws_lambda_function" "cost_reporter" {
  function_name    = "${local.service_name}-function"
  filename         = data.archive_file.test_terraform.output_path
  source_code_hash = data.archive_file.test_terraform.output_base64sha256
  runtime          = "provided.al2"
  role             = aws_iam_role.lambda_role.arn
  handler          = "main"

  environment {
    variables = {
      SNS_TOPIC_ARN = aws_sns_topic.cost_reporter.arn
    }
  }
}
