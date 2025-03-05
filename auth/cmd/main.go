package main

func main() {
	app := NewApplication()
	app.Infologger.Println("starting server on address " + app.HttpServer.Addr)
	app.ErrorLogger.Fatalln(app.HttpServer.ListenAndServe())
}
