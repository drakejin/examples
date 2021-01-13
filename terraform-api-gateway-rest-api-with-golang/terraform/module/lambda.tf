data "aws_caller_identity" "current" {}
data "aws_region" "current" {}

# Lambda
resource "aws_lambda_permission" "lambda_permission" {
  statement_id  = "AllowExecutionFromAPIGateway"
  action        = "lambda:InvokeFunction"
  function_name = aws_lambda_function.lambda.function_name
  principal     = "apigateway.amazonaws.com"

  # More: http://docs.aws.amazon.com/apigateway/latest/developerguide/api-gateway-control-access-using-iam-policies-to-invoke-api.html
//  source_arn = "arn:aws:execute-api:${data.aws_region.current.name}:${data.aws_caller_identity.current.account_id}:${aws_api_gateway_rest_api.api.id}/${aws_api_gateway_deployment.deployment.stage_name}/${aws_api_gateway_method.method.http_method}${aws_api_gateway_resource.resource.path}"
  source_arn = "${aws_api_gateway_rest_api.api.execution_arn}/*/*/*"
//  source_arn = aws_api_gateway_rest_api.api.execution_arn
}

resource "aws_lambda_function" "lambda" {
  handler       = var.lambda.handler
  function_name = "${var.meta.team}-${var.meta.service}-${var.meta.env}"
  role          = aws_iam_role.role.arn
  runtime       = var.lambda.runtime
  memory_size      = var.lambda.memory_size
  timeout          = var.lambda.timeout

  s3_bucket = var.lambda.s3_bucket
  s3_key    = var.lambda.s3_object_key

  vpc_config {
    # Every subnet should be able to reach an EFS mount target in the same Availability Zone. Cross-AZ mounts are not permitted.
    subnet_ids         = var.subnet_ids
    security_group_ids = [aws_security_group.sg.id]
  }
}

resource "aws_security_group" "sg" {
  name        = "${var.meta.team}-${var.meta.service}-${var.meta.env}"
  description = "${var.meta.team}-${var.meta.service}-${var.meta.env}"
  vpc_id      = var.vpc_id

  ingress {
    from_port = 0
    to_port   = 0
    protocol  = -1
    self      = true
  }

  ingress {
    security_groups = var.sg_ids
    from_port       = 0
    to_port         = 0
    protocol        = -1
  }


  egress {
    from_port   = 0
    to_port     = 0
    protocol    = -1
    cidr_blocks = ["0.0.0.0/0"]
  }

  # VPC 피어링한 상대의 vpc id로 지정할수 없음.
  # https://docs.aws.amazon.com/vpc/latest/peering/vpc-peering-security-groups.html
  # ingress {
  #   from_port   = 0
  #   to_port     = 0
  #   protocol    = -1
  #   cidr_blocks = ["${var.peering_cidrs}"]
  # }
}

resource "aws_iam_role" "role" {
  name = "${var.meta.team}-${var.meta.service}-${var.meta.env}"

  assume_role_policy = <<POLICY
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Action": "sts:AssumeRole",
      "Principal": {
        "Service": "lambda.amazonaws.com"
      },
      "Effect": "Allow",
      "Sid": ""
    }
  ]
}
POLICY
}

resource "aws_iam_policy" "policy" {
  name = "${var.meta.team}-${var.meta.service}-${var.meta.env}"
  description = "${var.meta.team}-${var.meta.service}-${var.meta.env}"

  policy = <<EOF
{
    "Version": "2012-10-17",
    "Statement": [
        {
            "Effect": "Allow",
            "Action": [
                "logs:CreateLogGroup",
                "logs:CreateLogStream",
                "logs:PutLogEvents",
                "ec2:CreateNetworkInterface",
                "ec2:DescribeNetworkInterfaces",
                "ec2:DeleteNetworkInterface",
                "ec2:AssignPrivateIpAddresses",
                "ec2:UnassignPrivateIpAddresses"
            ],
            "Resource": "*"
        }
    ]
}
EOF
}

resource "aws_iam_role_policy_attachment" "role_attach" {
  role = aws_iam_role.role.name
  policy_arn = aws_iam_policy.policy.arn
}

