package structs

import (
	"database/sql"
	"errors"
	"fmt"
	s "magazinerecipes/sql"
	u "magazinerecipes/utils"
)

const (
	ProductPropertyTable = "sylius_product_property"
)

//ProductProperty struct
type ProductProperty struct {
	ID         int    `json:"id"`
	ProductId  int    `json:"product_id"`
	PropertyId int    `json:"property_id"`
	Value      string `json:"value"`
}

//MaxMagazineNumber Retrieves magazines max number
func MaxMagazineNumber() int {
	var i ProductProperty
	err := s.CountDistinct(ProductPropertyTable, "value", "property_id=23 AND value NOT LIKE '-%%' AND value<>0").Scan(&i.ID)
	if err == sql.ErrNoRows {
		return 0
	}
	u.Check(err)
	return i.ID
}

//PopulatePropertyMap creates a MagazineNumb by ProductId map
func PopulatePropertyMap() map[int]int {
	query := fmt.Sprintf("SELECT MAX(product_id) AS product_id, value FROM %s WHERE property_id=23 AND value NOT LIKE '-%%' AND value<>0 GROUP BY value", ProductPropertyTable)
	rows, err := s.SelectQuery(query)
	u.Check(err)
	num := MaxMagazineNumber()
	if num == 0 {
		numErr := errors.New("No rows")
		panic(numErr)
	}
	propMap := make(map[int]int, num)
	for rows.Next() {
		var key int
		var value int
		rows.Scan(&value, &key)
		propMap[key] = value
	}
	return propMap
}
