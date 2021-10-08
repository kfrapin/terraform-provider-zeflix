resource "zeflix_movie" "back_future" {
  name = "Back to the future 2"
}

output "my_movie" {
  value = zeflix_movie.back_future
}