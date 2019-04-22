provider "aws" {
  region                  = "${var.aws_region}"
  shared_credentials_file = "/home/vimal/.aws"
}

variable "aws_region" {
  default = "us-west-2"
}

module "terraform_complete" {
  source = "./modules"
}