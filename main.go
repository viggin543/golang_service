package main

import "github.com/viggin543/go_http_server/bootstrap"

func main() {
    r := bootstrap.SetupServer()
    _ = r.Run()
}
