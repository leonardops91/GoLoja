package controllers

import (
	"html/template"
	"log"
	"net/http"
	"strconv"

	models "github.com/leonardops91/projetospessoais/loja/models"
)

var temp = template.Must(template.ParseGlob("templates/*html"))

const httpForRedirect = 301

func Index(wr http.ResponseWriter, r *http.Request) {
	products := models.GetProducts()
	temp.ExecuteTemplate(wr, "Index", products)
}

func New(wr http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(wr, "NewProduct", nil)
}

func Insert(wr http.ResponseWriter, r *http.Request) {

	name := r.FormValue("name")
	price, err := strconv.ParseFloat(r.FormValue("price"), 64)
	if err != nil {
		log.Println("Erro na conversão de string para float64")
	}
	quantity, err := strconv.Atoi(r.FormValue("quantity"))
	if err != nil {
		log.Println("Erro na conversão de string para inteiro")
	}

	models.InsertProduct(name, price, quantity)

	http.Redirect(wr, r, "/", httpForRedirect)
}

func Edit(wr http.ResponseWriter, r *http.Request) {
	product := models.GetProduct(r.URL.Query().Get("id"))

	temp.ExecuteTemplate(wr, "EditProduct", product)
}

func Update(wr http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.FormValue("id"))
	if err != nil {
		panic(err.Error())
	}
	name := r.FormValue("name")
	price, err := strconv.ParseFloat(r.FormValue("price"), 64)
	if err != nil {
		panic(err.Error())
	}
	quantity, err := strconv.Atoi(r.FormValue("quantity"))
	if err != nil {
		panic(err.Error())
	}
	models.UpdateProduct(id, name, price, quantity)

	http.Redirect(wr, r, "/", httpForRedirect)
}

func Delete(wr http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	models.DeleteProduct(id)

	http.Redirect(wr, r, "/", httpForRedirect)
}
