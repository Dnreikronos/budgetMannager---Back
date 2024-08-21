package main

func main() {
	server := NewAPIServer(":9000")
	server.Run()
}
