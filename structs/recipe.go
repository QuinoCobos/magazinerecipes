package structs

import (
	"errors"
	"strconv"

	s "magazinerecipes/sql"
	u "magazinerecipes/utils"
)

var (
	recipeTable  = "sylius_product_recipe"
	recipeFields = []string{"name", "magazine", "product_id", "page", "time_text", "kcal_text"}
)

//Recipe struct
type Recipe struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Magazine  string `json:"magazine"`
	ProductID int    `json:"product_id"`
	Page      string `json:"page"`
	TimeText  string `json:"time_text"`
	KcalText  string `json:"kcal_text"`
}

func (r *Recipe) String() string {
	return strconv.Itoa(r.ID) + ";" +
		r.Name + ";" +
		r.Magazine + ";" +
		strconv.Itoa(r.ProductID) + ";" +
		r.Page + ";" +
		r.TimeText + ";" +
		r.KcalText

}

//Insert a recipe into db
func (r *Recipe) Insert() (int, error) {
	if u.Nil(r.Name) || u.Nil(r.Magazine) || u.Nil(r.ProductID) || u.Nil(r.Page) {
		return 0, errors.New("Error, empty required fields")
	}
	res, err := s.Insert(recipeTable, recipeFields, []string{r.Name, r.Magazine, strconv.Itoa(r.ProductID), r.Page, r.TimeText, r.KcalText})
	id, _ := res.LastInsertId()
	num, _ := res.RowsAffected()
	r.ID = int(id)
	return int(num), err
}

func RecipeTable() string {
	return recipeTable
}
