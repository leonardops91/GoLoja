package Product

import (
	"github.com/leonardops91/projetospessoais/loja/database"
)

type Product struct {
	Id       int
	Name     string
	Price    float64
	Quantity int
}

func GetProduct(id string) Product {
	db := database.DbConnect()

	query, err := db.Query("select * from product where id=$1", id)
	if err != nil {
		panic(err.Error())
	}
	product := Product{}
	for query.Next() {
		var Id, Quantity int
		var Name string
		var Price float64

		err = query.Scan(&Id, &Name, &Price, &Quantity)
		if err != nil {
			panic(err.Error())
		}
		product.Id = Id
		product.Name = Name
		product.Price = Price
		product.Quantity = Quantity
	}
	defer db.Close()

	return product
}

func GetProducts() []Product {
	db := database.DbConnect()

	query, err := db.Query("select * from product order by id;")
	if err != nil {
		panic(err.Error())
	}
	p := Product{}
	products := []Product{}
	for query.Next() {
		var id, quantity int
		var name string
		var price float64

		err := query.Scan(&id, &name, &price, &quantity)
		if err != nil {
			panic(err.Error())
		}

		p.Id = id
		p.Name = name
		p.Price = price
		p.Quantity = quantity

		products = append(products, p)
	}

	defer db.Close()

	return products

}

func InsertProduct(name string, price float64, quantity int) {
	db := database.DbConnect()
	insertQuery, err := db.Prepare("insert into product (name, price, quantity) values ($1, $2, $3);")
	if err != nil {
		panic(err.Error())
	}
	insertQuery.Exec(name, price, quantity)

	defer db.Close()

}

func UpdateProduct(id int, name string, price float64, quantity int) {
	db := database.DbConnect()

	query, err := db.Prepare("Update product set name=$1, price=$2, quantity=$3 where id=$4")
	if err != nil {
		panic(err.Error())
	}
	query.Exec(name, price, quantity, id)

	defer db.Close()
}

func DeleteProduct(id string) {
	db := database.DbConnect()

	deleteQuery, err := db.Prepare("Delete from product where id=$1")

	if err != nil {
		panic(err.Error())
	}

	deleteQuery.Exec(id)

	defer db.Close()

}
