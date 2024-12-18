output "id" {
  value = module.servicebus_topic.id
}

output "name" {
  value = module.servicebus_topic.name
}

output "endpoint" {
  value = module.servicebus_namespace.endpoint

}
