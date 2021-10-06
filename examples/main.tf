resource "zeflix_catalog" "top_100_catalog" {
    name = "Top 100 Best Movies Ever"
}

output "my_catalog" {
    value = data.zeflix_catalog.top_100_catalog
}