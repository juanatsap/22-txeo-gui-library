package models

import (
	"database/sql"
	"fmt"
	"os"
	"strings"

	_ "github.com/mattn/go-sqlite3"
	"github.com/olekukonko/tablewriter"
)

// type Category struct {
// 	ID         primitive.ObjectID `bson:"_id,omitempty"`
// 	Count      int                `bson:"count"`
// 	Deleted    bool               `bson:"deleted"`
// 	Name       string             `bson:"name"`
// 	Traduction Traduction         `bson:"traduction"`
// 	// Otros campos según sea necesario
// }

// Category model adapted for SQLite
type Category struct {
	ID          int            // SQLite uses int for primary keys by default
	Name        string         // The name of the category
	ShortName   string         // The short name of the category
	Count       int            // A count of how many times this category is used
	Deleted     bool           // Whether the category is marked as deleted
	Traduction  sql.NullString // To handle cases where a translation might be optional or NULL
	Icon        string
	Color       string
	Tags        Tags
	Subcategory string
	Concepts    Concepts
}
type Tags []Tag

type Tag struct {
	Name string
	Slug string
}
type Categories []Category
type Concepts []Concept
type Concept struct {
	Name              string
	Icon              string
	ShortName         string
	Tags              Tags
	CategoryShortName string
}

// Print
func (categories *Categories) PrintCategories() {
	for _, category := range *categories {
		fmt.Println(category.Name)
	}
}
func (category Category) GetUnknownCategory(b Block) Category {

	return Category{
		Name:        "?",
		Icon:        "❓",
		Color:       "",
		Subcategory: "",
		Tags:        Tags{},
		Concepts:    []Concept{b.Concept},
		ShortName:   "Desconocido",
	}
}
func (category Category) TryToAssignCategory(b Block, categories []Category) Category {

	var categoryName string
	var categoryShortName string
	var categoryIcon string
	var categoryColor string
	var categorySubcategory string
	var categoryTags Tags

CATEGORIES_LABEL:
	for _, category := range categories {
		for _, storedConcept := range category.Concepts {
			if storedConcept.Name == b.Concept.Name {

				// Assign the category to the block
				categoryName = category.Name
				categoryShortName = category.ShortName
				categoryIcon = category.Icon
				categoryColor = category.Color
				categorySubcategory = category.Subcategory
				categoryTags = category.Tags
				break CATEGORIES_LABEL
			}
		}
	}

	return Category{
		Name:        categoryName,
		ShortName:   categoryShortName,
		Icon:        categoryIcon,
		Color:       categoryColor,
		Subcategory: categorySubcategory,
		Tags:        categoryTags,
	}
}
func (categories Categories) GetCategories() Categories {
	return categories
}
func (categories *Categories) AssignCategoryToSelectedConcept(row int, category Category, Blocks Blocks) Concept {

	// Get The slected block
	selectedBlock := Blocks[row]

	fmt.Printf("Category %s selected for row %d\n", category.Name, row+1)
	fmt.Printf("Concept %s selected for row %d\n", selectedBlock.Concept.Name, row+1)

	// Get the selected category
	// selectedCategory := selectedBlock.Category

	// Create a concept object
	concept := Concept{
		Name:              selectedBlock.Concept.Name,
		Icon:              category.Icon,
		ShortName:         selectedBlock.Concept.Name,
		Tags:              selectedBlock.Concept.Tags,
		CategoryShortName: category.ShortName,
	}
	return concept
}

func NewConceptFromString(conceptAsString string) Concept {

	return Concept{
		Name: conceptAsString,
	}
}

func (categories *Categories) SortByShortName() {
	// Sort the categories by short name
	for i := 0; i < len(*categories)-1; i++ {
		for j := i + 1; j < len(*categories); j++ {
			if (*categories)[i].ShortName > (*categories)[j].ShortName {
				(*categories)[i], (*categories)[j] = (*categories)[j], (*categories)[i]
			}
		}
	}
}

func PrintCategoriesAndConcepts(cats Categories) {
	// Print a master table for categories
	categoryTable := tablewriter.NewWriter(os.Stdout)
	categoryTable.SetHeader([]string{"Category Name", "Short Name", "Icon", "Subcategory"})

	for _, c := range cats {
		categoryTable.Append([]string{c.Name, c.ShortName, c.Icon, c.Subcategory})
	}
	categoryTable.SetBorder(true)
	categoryTable.SetAutoWrapText(false)
	fmt.Println("=== Categories ===")
	categoryTable.Render()

	// For each category, print a concepts table
	for _, c := range cats {
		if len(c.Concepts) == 0 {
			continue
		}

		fmt.Printf("\n=== Concepts for Category: %s (%s) ===\n", c.Name, c.ShortName)

		conceptTable := tablewriter.NewWriter(os.Stdout)
		conceptTable.SetHeader([]string{"Concept Name", "Short Name", "Icon", "Tags"})

		for _, co := range c.Concepts {
			// Gather tags into a string
			var tagList string
			for i, t := range co.Tags {
				if i > 0 {
					tagList += ", "
				}
				tagList += fmt.Sprintf("%s(%s)", t.Name, t.Slug)
			}

			conceptTable.Append([]string{co.Name, co.ShortName, co.Icon, tagList})
		}

		conceptTable.SetBorder(true)
		conceptTable.SetAutoWrapText(false)
		conceptTable.Render()
	}
}
func (cats Categories) PrintCombinedTable() {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Cat ID", "Category Name", "Cat ShortName", "Cat Icon", "Concept Name", "Concept ShortName", "Concept Icon", "Tags"})
	table.SetAutoWrapText(false)
	table.SetBorder(true)

	// If you want to show categories even if they have no concepts, set this to true:
	showEmptyCategories := false

	for _, c := range cats {
		if len(c.Concepts) == 0 && showEmptyCategories {
			// Print a row with only category info and empty concepts
			table.Append([]string{fmt.Sprintf("%d", c.ID), c.Name, c.ShortName, c.Icon, "", "", "", ""})
			continue
		}

		for _, co := range c.Concepts {
			// Gather tags into a string
			var tagList []string
			for _, t := range co.Tags {
				tagList = append(tagList, fmt.Sprintf("%s(%s)", t.Name, t.Slug))
			}

			table.Append([]string{
				fmt.Sprintf("%d", c.ID), // Category ID
				c.Name,
				c.ShortName,
				c.Icon,
				co.Name,
				co.ShortName,
				co.Icon,
				strings.Join(tagList, ", "),
			})
		}
	}

	table.Render()
}
