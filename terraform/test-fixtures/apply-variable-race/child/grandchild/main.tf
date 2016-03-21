variable "bar" {}

resource "aws_instance" "qux" {}

output "qux" { value = "${var.bar}${aws_instance.qux.id}" }
