/*
 * Copyright 2022 XXIV
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
package cocktaildb

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	httpClient "net/http"
	"net/url"
)

type cocktails struct {
	Drinks []Cocktail  `json:"drinks"`
}

type Cocktail struct {
	DateModified                string      `json:"dateModified"`
	IDDrink                     string      `json:"idDrink"`
	StrAlcoholic                string      `json:"strAlcoholic"`
	StrCategory                 string      `json:"strCategory"`
	StrCreativeCommonsConfirmed string      `json:"strCreativeCommonsConfirmed"`
	StrDrink                    string      `json:"strDrink"`
	StrDrinkAlternate           string 		`json:"strDrinkAlternate"`
	StrDrinkThumb               string      `json:"strDrinkThumb"`
	StrGlass                    string      `json:"strGlass"`
	StrIBA                      string      `json:"strIBA"`
	StrImageAttribution         string      `json:"strImageAttribution"`
	StrImageSource              string      `json:"strImageSource"`
	StrIngredient1              string      `json:"strIngredient1"`
	StrIngredient10             string 		`json:"strIngredient10"`
	StrIngredient11             string 		`json:"strIngredient11"`
	StrIngredient12             string 		`json:"strIngredient12"`
	StrIngredient13             string 		`json:"strIngredient13"`
	StrIngredient14             string 		`json:"strIngredient14"`
	StrIngredient15             string 		`json:"strIngredient15"`
	StrIngredient2              string      `json:"strIngredient2"`
	StrIngredient3              string      `json:"strIngredient3"`
	StrIngredient4              string      `json:"strIngredient4"`
	StrIngredient5              string      `json:"strIngredient5"`
	StrIngredient6              string      `json:"strIngredient6"`
	StrIngredient7              string      `json:"strIngredient7"`
	StrIngredient8              string 		`json:"strIngredient8"`
	StrIngredient9              string 		`json:"strIngredient9"`
	StrInstructions             string      `json:"strInstructions"`
	StrInstructionsDE           string      `json:"strInstructionsDE"`
	StrInstructionsES           string 		`json:"strInstructionsES"`
	StrInstructionsFR           string 		`json:"strInstructionsFR"`
	StrInstructionsIT           string      `json:"strInstructionsIT"`
	StrInstructionsZH_HANS      string 		`json:"strInstructionsZH-HANS"`
	StrInstructionsZH_HANT      string 		`json:"strInstructionsZH-HANT"`
	StrMeasure1                 string      `json:"strMeasure1"`
	StrMeasure10                string 		`json:"strMeasure10"`
	StrMeasure11               	string 		`json:"strMeasure11"`
	StrMeasure12                string 		`json:"strMeasure12"`
	StrMeasure13                string 		`json:"strMeasure13"`
	StrMeasure14                string 		`json:"strMeasure14"`
	StrMeasure15                string 		`json:"strMeasure15"`
	StrMeasure2                 string      `json:"strMeasure2"`
	StrMeasure3                 string      `json:"strMeasure3"`
	StrMeasure4                 string      `json:"strMeasure4"`
	StrMeasure5                 string      `json:"strMeasure5"`
	StrMeasure6                 string      `json:"strMeasure6"`
	StrMeasure7                 string      `json:"strMeasure7"`
	StrMeasure8                 string 		`json:"strMeasure8"`
	StrMeasure9                 string 		`json:"strMeasure9"`
	StrTags                     string      `json:"strTags"`
	StrVideo                    string 		`json:"strVideo"`
}

type ingredients struct {
	Ingredients []Ingredient `json:"ingredients"`
}

type Ingredient struct {
	IDIngredient   string `json:"idIngredient"`
	StrABV         string `json:"strABV"`
	StrAlcohol     string `json:"strAlcohol"`
	StrDescription string `json:"strDescription"`
	StrIngredient  string `json:"strIngredient"`
	StrType        string `json:"strType"`
}

type filters struct {
	Drinks []Filter `json:"drinks"`
}

type Filter struct {
	IDDrink       string `json:"idDrink"`
	StrDrink      string `json:"strDrink"`
	StrDrinkThumb string `json:"strDrinkThumb"`
}

type categoriesFilter struct {
	Drinks []struct {
		StrCategory string `json:"strCategory"`
	} `json:"drinks"`
}

type glassesFilter struct {
	Drinks []struct {
		StrGlass string `json:"strGlass"`
	} `json:"drinks"`
}

type ingredientsFilter struct {
	Drinks []struct {
		StrIngredient1 string `json:"strIngredient1"`
	} `json:"drinks"`
}

type alcoholicFilter struct {
	Drinks []struct {
		StrAlcoholic string `json:"strAlcoholic"`
	} `json:"drinks"`
}

func http(endpoint string) (string, error) {
	response, err := httpClient.Get(fmt.Sprintf("https://thecocktaildb.com/api/json/v1/1/%s", endpoint))
	if err != nil {
		return "", errors.New("ERROR")
	}
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", errors.New("ERROR")
	}
	return string(body), nil
}

// Search  cocktail by name
func Search(s string) ([]Cocktail, error) {
	response, err := http(fmt.Sprintf("search.php?s=%s",url.QueryEscape(s)))
	if err != nil {
		return []Cocktail{}, newError("error")
	}
	if len(response) == 0 {
		return []Cocktail{}, newError("error")
	}
	var drink cocktails
	jsonError := json.Unmarshal([]byte(response), &drink)
	if jsonError != nil {
		return []Cocktail{}, newError("error")
	}
	if len(drink.Drinks) == 0 {
		return []Cocktail{}, newError("error")
	}
	var responseSlice []Cocktail
	for _, v := range drink.Drinks {
		responseSlice= append(responseSlice,v)
	}
	return responseSlice, nil
}

// SearchByLetter search cocktails by first letter
func SearchByLetter(b byte) ([]Cocktail, error) {
	response, err := http(fmt.Sprintf("search.php?f=%c", b))
	if err != nil {
		return []Cocktail{}, newError("error")
	}
	if len(response) == 0 {
		return []Cocktail{}, newError("error")
	}
	var drink cocktails
	jsonError := json.Unmarshal([]byte(response), &drink)
	if jsonError != nil {
		return []Cocktail{}, newError("error")
	}
	if len(drink.Drinks) == 0 {
		return []Cocktail{}, newError("error")
	}
	var responseSlice []Cocktail
	for _, v := range drink.Drinks {
		responseSlice= append(responseSlice,v)
	}
	return responseSlice, nil
}

// SearchIngredient search ingredient by name
func SearchIngredient(s string) (Ingredient, error) {
	response, err := http(fmt.Sprintf("search.php?i=%s",url.QueryEscape(s)))
	if err != nil {
		return Ingredient{}, newError("error")
	}
	if len(response) == 0 {
		return Ingredient{}, newError("error")
	}
	var drink ingredients
	jsonError := json.Unmarshal([]byte(response), &drink)
	if jsonError != nil {
		return Ingredient{}, newError("error")
	}
	if len(drink.Ingredients) == 0 {
		return Ingredient{}, newError("error")
	}
	ingredient := drink.Ingredients[0]
	return ingredient, nil
}

// SearchByID search cocktail details by id
func SearchByID(i int64) (Cocktail, error) {
	response, err := http(fmt.Sprintf("lookup.php?i=%d", i))
	if err != nil {
		return Cocktail{}, newError("error")
	}
	if len(response) == 0 {
		return Cocktail{}, newError("error")
	}
	var drink cocktails
	jsonError := json.Unmarshal([]byte(response), &drink)
	if jsonError != nil {
		return Cocktail{}, newError("error")
	}
	if len(drink.Drinks) == 0 {
		return Cocktail{}, newError("error")
	}
	ingredient := drink.Drinks[0]
	return ingredient, nil
}

// SearchIngredientByID search ingredient by ID
func SearchIngredientByID(i int64) (Ingredient, error) {
	response, err := http(fmt.Sprintf("lookup.php?iid=%d", i))
	if err != nil {
		return Ingredient{}, newError("error")
	}
	if len(response) == 0 {
		return Ingredient{}, newError("error")
	}
	var drink ingredients
	jsonError := json.Unmarshal([]byte(response), &drink)
	if jsonError != nil {
		return Ingredient{}, newError("error")
	}
	if len(drink.Ingredients) == 0 {
		return Ingredient{}, newError("error")
	}
	ingredient := drink.Ingredients[0]
	return ingredient, nil
}

// Random cocktail
func Random() (Cocktail, error) {
	response, err := http("random.php")
	if err != nil {
		return Cocktail{}, newError("error")
	}
	if len(response) == 0 {
		return Cocktail{}, newError("error")
	}
	var drink cocktails
	jsonError := json.Unmarshal([]byte(response), &drink)
	if jsonError != nil {
		return Cocktail{}, newError("error")
	}
	if len(drink.Drinks) == 0 {
		return Cocktail{}, newError("error")
	}
	ingredient := drink.Drinks[0]
	return ingredient, nil
}

// FilterByIngredient filter by ingredient
func FilterByIngredient(s string) ([]Filter, error) {
	response, err := http(fmt.Sprintf("filter.php?i=%s",url.QueryEscape(s)))
	if err != nil {
		return []Filter{}, newError("error")
	}
	if len(response) == 0 {
		return []Filter{}, newError("error")
	}
	var drink filters
	jsonError := json.Unmarshal([]byte(response), &drink)
	if jsonError != nil {
		return []Filter{}, newError("error")
	}
	if len(drink.Drinks) == 0 {
		return []Filter{}, newError("error")
	}
	var responseSlice []Filter
	for _, v := range drink.Drinks {
		responseSlice= append(responseSlice,v)
	}
	return responseSlice, nil
}

// FilterByAlcoholic filter by alcoholic
func FilterByAlcoholic(s string) ([]Filter, error) {
	response, err := http(fmt.Sprintf("filter.php?a=%s",url.QueryEscape(s)))
	if err != nil {
		return []Filter{}, newError("error")
	}
	if len(response) == 0 {
		return []Filter{}, newError("error")
	}
	var drink filters
	jsonError := json.Unmarshal([]byte(response), &drink)
	if jsonError != nil {
		return []Filter{}, newError("error")
	}
	if len(drink.Drinks) == 0 {
		return []Filter{}, newError("error")
	}
	var responseSlice []Filter
	for _, v := range drink.Drinks {
		responseSlice= append(responseSlice,v)
	}
	return responseSlice, nil
}

// FilterByCategory filter by category
func FilterByCategory(s string) ([]Filter, error) {
	response, err := http(fmt.Sprintf("filter.php?c=%s",url.QueryEscape(s)))
	if err != nil {
		return []Filter{}, newError("error")
	}
	if len(response) == 0 {
		return []Filter{}, newError("error")
	}
	var drink filters
	jsonError := json.Unmarshal([]byte(response), &drink)
	if jsonError != nil {
		return []Filter{}, newError("error")
	}
	if len(drink.Drinks) == 0 {
		return []Filter{}, newError("error")
	}
	var responseSlice []Filter
	for _, v := range drink.Drinks {
		responseSlice= append(responseSlice,v)
	}
	return responseSlice, nil
}

// FilterByGlass filter by glass
func FilterByGlass(s string) ([]Filter, error) {
	response, err := http(fmt.Sprintf("filter.php?g=%s",url.QueryEscape(s)))
	if err != nil {
		return []Filter{}, newError("error")
	}
	if len(response) == 0 {
		return []Filter{}, newError("error")
	}
	var drink filters
	jsonError := json.Unmarshal([]byte(response), &drink)
	if jsonError != nil {
		return []Filter{}, newError("error")
	}
	if len(drink.Drinks) == 0 {
		return []Filter{}, newError("error")
	}
	var responseSlice []Filter
	for _, v := range drink.Drinks {
		responseSlice= append(responseSlice,v)
	}
	return responseSlice, nil
}

// CategoriesFilter List the categories filter
func CategoriesFilter() ([]string, error) {
	response, err := http("list.php?c=list")
	if err != nil {
		return []string{}, newError("error")
	}
	if len(response) == 0 {
		return []string{}, newError("error")
	}
	var drink categoriesFilter
	jsonError := json.Unmarshal([]byte(response), &drink)
	if jsonError != nil {
		return []string{}, newError("error")
	}
	if len(drink.Drinks) == 0 {
		return []string{}, newError("error")
	}
	var responseSlice []string
	for _, v := range drink.Drinks {
		responseSlice= append(responseSlice,v.StrCategory)
	}
	return responseSlice, nil
}

// GlassesFilter List the glasses filter
func GlassesFilter() ([]string, error) {
	response, err := http("list.php?g=list")
	if err != nil {
		return []string{}, newError("error")
	}
	if len(response) == 0 {
		return []string{}, newError("error")
	}
	var drink glassesFilter
	jsonError := json.Unmarshal([]byte(response), &drink)
	if jsonError != nil {
		return []string{}, newError("error")
	}
	if len(drink.Drinks) == 0 {
		return []string{}, newError("error")
	}
	var responseSlice []string
	for _, v := range drink.Drinks {
		responseSlice= append(responseSlice,v.StrGlass)
	}
	return responseSlice, nil
}

// IngredientsFilter List the ingredients filter
func IngredientsFilter() ([]string, error) {
	response, err := http("list.php?i=list")
	if err != nil {
		return []string{}, newError("error")
	}
	if len(response) == 0 {
		return []string{}, newError("error")
	}
	var drink ingredientsFilter
	jsonError := json.Unmarshal([]byte(response), &drink)
	if jsonError != nil {
		return []string{}, newError("error")
	}
	if len(drink.Drinks) == 0 {
		return []string{}, newError("error")
	}
	var responseSlice []string
	for _, v := range drink.Drinks {
		responseSlice= append(responseSlice,v.StrIngredient1)
	}
	return responseSlice, nil
}

// AlcoholicFilter List the alcoholic filter
func AlcoholicFilter() ([]string, error) {
	response, err := http("list.php?a=list")
	if err != nil {
		return []string{}, newError("error")
	}
	if len(response) == 0 {
		return []string{}, newError("error")
	}
	var drink alcoholicFilter
	jsonError := json.Unmarshal([]byte(response), &drink)
	if jsonError != nil {
		return []string{}, newError("error")
	}
	if len(drink.Drinks) == 0 {
		return []string{}, newError("error")
	}
	var responseSlice []string
	for _, v := range drink.Drinks {
		responseSlice= append(responseSlice,v.StrAlcoholic)
	}
	return responseSlice, nil
}