data "zeflix_movie" "my_first_movie" {
    # required parameter:
    # id = "<movie-id>"
}

output "my_movie" {
    value = data.zeflix_movie.my_first_movie
}