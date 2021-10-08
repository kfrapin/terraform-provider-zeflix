resource "zeflix_movie" "gladiator" {
  name = "Gladiator"
}

output "my_movie" {
  value = zeflix_movie.gladiator
}