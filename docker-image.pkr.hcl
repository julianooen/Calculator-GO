packer {
  required_plugins {
    docker = {
      version = ">= 1.0.1"
      source  = "github.com/hashicorp/docker"
    }
  }
}

source "docker" "golang" {
  image  = "golang:1.19"
  commit = true
  changes = [
    "WORKDIR /calc",
    "EXPOSE 8080",
    "ENTRYPOINT /calc/calc"
  ]
}

build {
  name = "GO-calc"
  sources = [
    "source.docker.golang"
  ]
  provisioner "file" {
    source      = "./calc"
    destination = "/calc"

  }

  provisioner "shell" {
    inline = ["cd /calc", "go mod tidy", "go build calc"]
  }

  post-processor "docker-tag" {
    repository = "go-calc"
    tags       = ["latest"]
    only       = ["docker.golang"]
  }
}