module "child1" {
  source = "./child"
  foo = "bar"
  bar = "baz"
}
output "final1" {
  value = "${module.child1.foo}${module.child1.bar}"
}
module "child2" {
  source = "./child"
  foo = "bar"
  bar = "baz"
}
output "final1" {
  value = "${module.child2.foo}${module.child2.bar}"
}
module "child3" {
  source = "./child"
  foo = "bar"
  bar = "baz"
}
output "final3" {
  value = "${module.child3.foo}${module.child3.bar}"
}
