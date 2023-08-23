package main

import (
"net/http"
"coding4women/handlers"

)

func main(){

http.HandleFunc("/", handlers.HomePage)
http.ListenAndServe(":8000",nil)
}