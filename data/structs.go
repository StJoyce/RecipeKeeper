package data

type Recipe struct{
Id string `json:"Id"`
RecipeName string `"json:'Recepe Name"`
Source string `json:"Source"`
PrepTime string `json:"Presentation Time"`
CookTime string `json:"Cook Time"`
ServingSize int `json:"Serving Size"`
Ingredients map[string]string `json:"Ingredients"`
Instructions []string `json:"Directions"`
Tags []string `json:"Tags"`
} 

var AllRecipes []Recipe