module "child" {
  source = "./child"
  foo = "bar"
  bar = "baz"
}
output "final" {
  value = "${module.child.foo}${module.child.bar}${module.child.baz}${module.child.qux}"
}
