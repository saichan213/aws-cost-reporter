terraform {
  # terraformのバージョン条件
  required_version = "~> 1.9.5"

  # 実行するproviderの条件
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 5.42.0"
    }
  }
}
