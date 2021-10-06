resource "zeflix_catalog" "top_100_catalog" {
  name = "Top 10 Best Movies Ever"
}

output "my_catalog" {
  value = zeflix_catalog.top_100_catalog
}