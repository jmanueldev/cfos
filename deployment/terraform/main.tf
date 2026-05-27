provider "aws" {
  region = "us-east-1"
}

resource "aws_eks_cluster" "cfos" {
  name = "cfos"
}