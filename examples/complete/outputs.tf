output "id" {
  value = module.servicebus_topic.id
}

output "name" {
  value = module.servicebus_topic.name
}

output "endpoint" {
  value = module.servicebus_namespace.endpoint
}

output "resource_group_name" {
  value = module.resource_group.name
}

output "namespace_name" {
  value = module.servicebus_namespace.name
}
