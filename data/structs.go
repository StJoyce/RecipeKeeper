package data

type Receipe struct{
Id string
RecipeName string
Source string
PrepTime string
CookTime string
ServingSize int
Ingredients map[string]string
Instructions []string
Tags []string
} 
