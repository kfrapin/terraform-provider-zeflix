data "zeflix_catalog" "my_catalog" {
    # required parameter:
    # id = "<catalog-id>"
}

output "my_catalog" {
    value = data.zeflix_catalog.my_catalog
}