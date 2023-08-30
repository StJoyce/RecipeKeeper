package main

import (
"net/http"
"fmt"
"coding4women/handlers"
"coding4women/data"

)


func main(){
data.FetchAllRecipes()
fmt.Println(data.AllRecipes)
http.HandleFunc("/", handlers.HomePage)
http.ListenAndServe(":8000",nil)

}