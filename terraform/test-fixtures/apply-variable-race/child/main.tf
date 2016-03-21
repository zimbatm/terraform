variable "foo" {}
variable "bar" {}

module "grandchild" {
  source = "./grandchild"

  bar = "${var.bar}"
}

resource "aws_instance" "baz" {}

output "foo" { value = "${var.foo}" }
output "bar" { value = "${var.bar}" }
output "baz" { value = "${aws_instance.baz.id}" }
output "qux" { value = "${module.grandchild.qux}" }
