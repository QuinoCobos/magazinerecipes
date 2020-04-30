package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/tealeg/xlsx"

	l "magazinerecipes/logger"
	s "magazinerecipes/sql"
	st "magazinerecipes/structs"
	u "magazinerecipes/utils"

	_ "github.com/go-sql-driver/mysql"
)

var (
	user, password string
	level          int
)

func main() {
	//Getting arguments
	level = 2
	//args()
	//Setting loggers
	l.Init(level)
	defer recoverFatal()
	l.Info("Inicializando")
	fmt.Println("Introduzca usuario:")
	fmt.Scanln(&user)
	fmt.Println("Introduzca contraseña:")
	fmt.Scanln(&password)
	l.Info(fmt.Sprintf("Conectando como: %s", user))
	//Setting database connection
	db, err := s.Connect(user, password, "thermomix_dbvorwerk_tm")
	u.Check(err)
	defer db.Close()
	l.Info("Conexion establecida")

	bookSlice, err := readBook()
	u.Check(err)
	productMap := st.PopulatePropertyMap()
	recipes := listRecipes(bookSlice, productMap, true)

	u.Check(s.Duplicate(st.RecipeTable()))
	_, err3 := s.Truncate(st.RecipeTable())
	u.Check(err3)

	l.Info("Cargando recetas. Espere...")
	for _, recipe := range recipes {
		i, err := recipe.Insert()
		u.Check(err)
		recipe.ID = i
	}
	l.Info(fmt.Sprintf("Se han cargado %s recetas\r\n", strconv.Itoa(len(recipes))))
	close(0)
}

func recoverFatal() {
	e := recover()
	l.Error(fmt.Sprintf("%s", e))
	close(2)
}

func args() {
	args := os.Args
	for i := 0; i < len(args); i++ {
		if i == 0 {
			continue
		}
		arg := args[i]
		fmt.Println(arg + "=" + args[i+1])
		switch arg {
		case "-u":
			user = args[i+1]
			i++
		case "-p":
			password = args[i+1]
			i++
		case "-l":
			var err error
			level, err = strconv.Atoi(args[i+1])
			u.Check(err)
			i++
		default:
			err := errors.New(fmt.Sprintf("Unknown argument %s", arg))
			panic(err)
		}
	}
}

func close(code int) {
	l.Info(fmt.Sprintf(u.ExitMessage, code))
	for i := 5; i > 0; i-- {
		fmt.Println(fmt.Sprintf("%d...", i))
		time.Sleep(time.Second)
	}
	os.Exit(code)
}

func readBook() ([][][]string, error) {
	dir, err := os.UserHomeDir()
	u.Check(err)
	dateS := time.Now().Format(u.DateFormat)
	bookName := dir + u.FilesPath + u.FilesName + dateS + u.XlsxExt
	return xlsx.FileToSliceUnmerged(bookName)
}

func listRecipes(file [][][]string, prodMap map[int]int, headers bool) []st.Recipe {
	var recipes []st.Recipe
	for _, sheet := range file {
		for i, row := range sheet {
			if (headers && i == 0) || row[1] == "" {
				continue
			}
			var recipe st.Recipe
			recipe.Name = strings.ReplaceAll(row[0], "'", "\\'")
			recipe.Magazine = row[1]
			magazine, _ := strconv.Atoi(recipe.Magazine)
			recipe.ProductID = prodMap[magazine]
			recipe.Page = row[2]
			recipe.TimeText = row[3]
			recipe.KcalText = row[4]
			recipes = append(recipes, recipe)
			l.Debug(recipe.String())
		}
	}
	return recipes
}
