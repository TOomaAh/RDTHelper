![logo without text](https://github.com/user-attachments/assets/ae43d019-96a4-44df-8cd3-48f64c6d7780)

# RDTHelper

RDTHelper is an improved real-debrid interface. This tool will allow you to upload several '.torrent' in one go. It will also allow you to debride multiple torrents at once.

This tool is developed as part of the Spring framework training.


## Docker

The easiest way to use this tool is to use Docker.

My repo: https://hub.docker.com/repository/docker/bagul/rdthelper

First you must pull the image.

`docker pull bagul/rdthelper:latest`

and just run:

`docker run -d -p 8080:8080 bagul/rdthelper:latest`

you can change the listening port of the server like this:

`docker run -d -p <YOUR-PORT>:8080 bagul/rdthelper:latest`

example:

`docker run -d -p 6549:8080 bagul/rdthelper:latest`

To connect to the server I would do like this:

http://localhost:6549
