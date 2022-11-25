resource "aws_instance" "validator" {
  count                       = var.num_instances
  ami                         = var.ami
  instance_type               = "t2.micro"
  subnet_id                   = aws_subnet.validator.id
  key_name                    = "numi-key"
  vpc_security_group_ids      = [aws_security_group.validator.id]
  associate_public_ip_address = false

  lifecycle {
    ignore_changes = [associate_public_ip_address]
  }

  tags = {
    Environment = var.env
    Project     = var.project
    Name        = "${var.project}-${var.env}-validator-${count.index}"
  }
}

resource "aws_eip" "validator" {
  count    = var.num_instances
  instance = aws_instance.validator[count.index].id
  vpc      = true
  tags = {
    Environment = var.env
    Project     = var.project
    Name        = "${var.project}-${var.env}-validator-eip-${count.index}"
  }
}

resource "aws_route53_record" "validator_api_a_record" {
  depends_on = [aws_eip.validator]
  count      = var.num_instances

  zone_id = var.dns_zone_id
  name    = "${var.domain_prefix}validator-${count.index}-api"
  type    = "A"
  ttl     = 600
  records = [aws_eip.validator[count.index].public_ip]
}

resource "aws_route53_record" "validator_rpc_a_record" {
  depends_on = [aws_eip.validator]
  count      = var.num_instances

  zone_id = var.dns_zone_id
  name    = "${var.domain_prefix}validator-${count.index}-rpc"
  type    = "A"
  ttl     = 600
  records = [aws_eip.validator[count.index].public_ip]
}
