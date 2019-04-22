# Archive a single file.

data "archive_file" "zip-lambda-go" {
  type        = "zip"
  source_dir = "/home/vimal/go/src/nike-aws-go/aws"
  output_path = "${path.cwd}/files/init.zip"
}



resource "aws_lambda_function" "go_lambda" {
  filename         = "${path.cwd}/files/init.zip"
  function_name    = "lambda_go"
  role             = "${aws_iam_role.lambda_exec_role.arn}"
  handler          = "main"
  source_code_hash = "${base64sha256(data.archive_file.zip-lambda-go.output_path)}"
  runtime          = "go1.x"

  environment {
    variables = {
      go = "go"
    }
  }
}

resource "aws_iam_role" "lambda_exec_role" {
  name = "lambda_exec_role"
  assume_role_policy = "${data.local_file.policy_lambda_file.content}"
}

//resource "aws_iam_policy" "policy" {
//  name        = "lambda-policy"
//  description = "A lambda policy"
//  policy      = "${file(data.local_file.policy_lambda_file)}"
//}

data "local_file" "policy_lambda_file" {
  filename = "${path.module}/lambdapolicy.json"
}