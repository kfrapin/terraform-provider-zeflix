data "zeflix_catalog" "my_catalog" {
    id = "12517246-293a-44c5-860d-7b658b2818cf"
}

output "my_catalog" {
    value = data.zeflix_catalog.my_catalog
}