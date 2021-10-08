resource "zeflix_movie" "terminator_3" {
  name = "Terminator 3 - Le soulevement"
}

output "my_movie" {
  value = zeflix_movie.terminator_3
}