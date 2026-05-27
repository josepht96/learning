output "vpc_id" {
  description = "Output for VPC"
  value       = data.terraform_remote_state.common.outputs.vpc_id
}