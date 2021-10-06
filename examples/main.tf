data "zeflix_catalog" "my_catalog" {
    id = "d0416022-a424-4cc1-9ecd-5700b5321df5"
}

output "my_catalog" {
    value = data.zeflix_catalog.my_catalog
}