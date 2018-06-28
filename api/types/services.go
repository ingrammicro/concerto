package types

type Service struct {
	ID          string   `json:"id" header:"ID"`
	Name        string   `json:"name" header:"NAME"`
	Description string   `json:"description" header:"DESCRIPTION"`
	Public      bool     `json:"public" header:"PUBLIC"`
	License     string   `json:"license" header:"LICENSE"`
	Recipes     []string `json:"recipes"  header:"RECIPES" show:"nolist"`
}
