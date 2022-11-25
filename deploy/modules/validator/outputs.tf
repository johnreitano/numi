output "ips" {
  value = [for eip in aws_eip.validator : eip.public_ip]
}
