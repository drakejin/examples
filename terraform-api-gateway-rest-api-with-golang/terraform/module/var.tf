variable "vpc_id" {
  type = string
}

variable "subnet_ids" {
  type = list(string)
}

variable "sg_ids" {
  type = list(string)
}

variable "record" {
  type = object({
    acm_arn = string
    zone_id = string
    name    = string
  })
}

variable "meta" {
  type = object({
    team    = string
    service = string
    env     = string
  })
}

variable "lambda" {
  type = object({
    version = string
    handler = string
    runtime = string
    memory_size = number
    concurrent = number
    timeout = number
    s3_bucket = string
    s3_object_key = string
  })
}
