resource "aws_instance" "seed" {
  depends_on = [aws_subnet.seed, aws_security_group.seed]

  count                       = var.num_instances
  ami                         = var.ami
  instance_type               = "t2.micro"
  subnet_id                   = aws_subnet.seed.id
  key_name                    = "numi-key"
  vpc_security_group_ids      = [aws_security_group.seed.id]
  associate_public_ip_address = false

  lifecycle {
    ignore_changes = [associate_public_ip_address]
  }

  tags = {
    Environment = var.env
    Project     = var.project
    Name        = "${var.project}-${var.env}-seed-${count.index}"
  }
}

resource "aws_eip" "seed" {
  count    = var.num_instances
  instance = aws_instance.seed[count.index].id
  vpc      = true
  tags = {
    Environment = var.env
    Project     = var.project
    Name        = "${var.project}-${var.env}-seed-eip-${count.index}"
  }
}

resource "aws_route53_record" "seed_api_a_record" {
  depends_on = [aws_eip.seed]
  count      = var.num_instances

  zone_id = var.dns_zone_id
  name    = "${var.domain_prefix}seed-${count.index}-api"
  type    = "A"
  ttl     = 600
  records = [aws_eip.seed[count.index].public_ip]
}

resource "aws_route53_record" "seed_rpc_a_record" {
  depends_on = [aws_eip.seed]
  count      = var.num_instances

  zone_id = var.dns_zone_id
  name    = "${var.domain_prefix}seed-${count.index}-rpc"
  type    = "A"
  ttl     = 600
  records = [aws_eip.seed[count.index].public_ip]
}
