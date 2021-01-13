locals {
  meta = {
    team    = "squaaat"
    service = "squaaat-api"
  }
}

module "alpha" {
  source = "./module"

  meta = {
    team    = local.meta.team
    service = local.meta.service
    env     = "alpha"
  }

  lambda = {
    handler = "sq"
    version = "v1"
    runtime = "go1.x"
    memory_size = 128
    concurrent = 2
    timeout = 1
    s3_bucket = aws_s3_bucket_object.object.bucket
    s3_object_key = aws_s3_bucket_object.object.key
  }

  record = {
    acm_arn = data.aws_acm_certificate.alpha_acm.arn
    zone_id = data.terraform_remote_state.common.outputs.route53_zone.zone_id
    name    = "api.alpha.squaaat.com"
  }

  vpc_id = data.terraform_remote_state.common.outputs.vpc.vpc_id
  subnet_ids = [
    data.terraform_remote_state.common.outputs.vpc.subnet_public_a_id,
    data.terraform_remote_state.common.outputs.vpc.subnet_public_b_id,
  ]
  sg_ids = [
    data.terraform_remote_state.common.outputs.vpc.sg_basic_id,
  ]
}

output "alpha" {
  value = module.alpha
}

data "aws_acm_certificate" "alpha_acm" {
  provider = aws.us_east_1
  domain   = "alpha.squaaat.com"
  statuses = ["ISSUED"]
}

data "aws_acm_certificate" "prod_acm" {
  provider = aws.us_east_1
  domain   = "squaaat.com"
  statuses = ["ISSUED"]
}

data "aws_s3_bucket" "repo" {
  bucket = data.terraform_remote_state.common.outputs.s3_lambda.name
}

data "archive_file" "dist" {
  type        = "zip"
  source_file  = "${path.cwd}/../dist/sq"
  output_path = "${path.cwd}/../dist/sq.zip"
}

resource "aws_s3_bucket_object" "object" {
  bucket = data.aws_s3_bucket.repo.bucket

  source = data.archive_file.dist.output_path
  key    = "${data.aws_s3_bucket.repo.bucket}/${local.meta.service}/${formatdate("YYYYMMDDhhmmss", timestamp())}.zip"

  etag = filemd5(data.archive_file.dist.output_path)
}
