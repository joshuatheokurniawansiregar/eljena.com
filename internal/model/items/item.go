package items

type Item struct {
	ItemName    string
	Tags        []string
	Description string
	Url         string
}

type SubCategory struct {
	SubCategoryName string
	SubCategoryTag  string
	Items           []Item
}

type Category struct {
	CategoryName  string
	SubCategories []SubCategory
}

func GetItemsByCategories() []Category {
	var goesToSchoolItems []Item = []Item{
		{
			ItemName:    "Aero",
			Tags:        []string{"program", "technology", "science", "science", "science", "science", "science"},
			Description: "This is description of Aero Glider",
			Url:         "/assets/Aero-Glider.jpg",
		},
		{
			ItemName:    "Aero",
			Tags:        []string{"program", "technology", "science", "science", "science", "science", "science"},
			Description: "This is description of Aero Glider",
			Url:         "/assets/Aero-Glider.jpg",
		},
		{
			ItemName:    "Aero",
			Tags:        []string{"program", "technology", "science", "science", "science", "science", "science"},
			Description: "This is description of Aero Glider",
			Url:         "/assets/Aero-Glider.jpg",
		},
		{
			ItemName:    "Aero",
			Tags:        []string{"program", "technology", "science", "science", "science", "science", "science"},
			Description: "This is description of Aero Glider",
			Url:         "/assets/Aero-Glider.jpg",
		},
		{
			ItemName:    "Aero",
			Tags:        []string{"program", "technology", "science", "science", "science", "science", "science"},
			Description: "This is description of Aero Glider",
			Url:         "/assets/Aero-Glider.jpg",
		},
	}

	var digitalMarketingItems []Item = []Item{
		{
			ItemName:    "Aero",
			Tags:        []string{"program", "technology", "science", "science", "science", "science", "science"},
			Description: "This is description of Aero Glider",
			Url:         "/assets/Aero-Glider.jpg",
		},
		{
			ItemName:    "Aero",
			Tags:        []string{"program", "technology", "science", "science", "science", "science", "science"},
			Description: "This is description of Aero Glider",
			Url:         "/assets/Aero-Glider.jpg",
		},
	}

	var categories []Category = []Category{
		{
			CategoryName: "Program",
			SubCategories: []SubCategory{
				{
					SubCategoryName: "Goes To School",
					SubCategoryTag:  "goes_to_school",
					Items:           goesToSchoolItems,
				},
			},
		},
		{
			CategoryName: "Service",
			SubCategories: []SubCategory{
				{
					SubCategoryName: "Digital Marketing",
					SubCategoryTag:  "digital_marketing",
					Items:           digitalMarketingItems,
				},
			},
		},
	}
	return categories
}

func GetItemsByCategoriesForProgram() []Category {
	// var goesToSchoolItems []Item = []Item{
	// 	{
	// 		ItemName:    "Aero",
	// 		Tags:        []string{"program", "technology", "science", "science", "science", "science", "science"},
	// 		Description: "This is description of Aero Glider",
	// 		Url:         "/assets/Aero-Glider.jpg",
	// 	},
	// 	{
	// 		ItemName:    "Aero",
	// 		Tags:        []string{"program", "technology", "science", "science", "science", "science", "science"},
	// 		Description: "This is description of Aero Glider",
	// 		Url:         "/assets/Aero-Glider.jpg",
	// 	},
	// 	{
	// 		ItemName:    "Aero",
	// 		Tags:        []string{"program", "technology", "science", "science", "science", "science", "science"},
	// 		Description: "This is description of Aero Glider",
	// 		Url:         "/assets/Aero-Glider.jpg",
	// 	},
	// 	{
	// 		ItemName:    "Aero",
	// 		Tags:        []string{"program", "technology", "science", "science", "science", "science", "science"},
	// 		Description: "This is description of Aero Glider",
	// 		Url:         "/assets/Aero-Glider.jpg",
	// 	},
	// 	{
	// 		ItemName:    "Aero",
	// 		Tags:        []string{"program", "technology", "science", "science", "science", "science", "science"},
	// 		Description: "This is description of Aero Glider",
	// 		Url:         "/assets/Aero-Glider.jpg",
	// 	},
	// }

	var categories []Category = []Category{
		{
			CategoryName: "Program",
			SubCategories: []SubCategory{
				{
					SubCategoryName: "Goes To School",
					SubCategoryTag:  "goes_to_school",
					// Items:           goesToSchoolItems,
				},
			},
		},
	}
	return categories
}

func GetItemsByCategoriesForServices() []Category {

	// var digitalMarketingItems []Item = []Item{
	// 	{
	// 		ItemName:    "SEO Management",
	// 		Tags:        []string{"program", "technology", "science", "science", "science", "science", "science"},
	// 		Description: "This is description of Aero Glider",
	// 		Url:         "/assets/Aero-Glider.jpg",
	// 	},
	// 	{
	// 		ItemName:    "Landing Page Development",
	// 		Tags:        []string{"program", "technology", "science", "science", "science", "science", "science"},
	// 		Description: "This is description of Aero Glider",
	// 		Url:         "/assets/Aero-Glider.jpg",
	// 	},
	// }

	// var gardenItems []Item = []Item{
	// 	{
	// 		ItemName:    "Vertical Garden",
	// 		Tags:        []string{"program", "technology", "science", "science", "science", "science", "science"},
	// 		Description: "This is description of Aero Glider",
	// 		Url:         "/assets/Aero-Glider.jpg",
	// 	},
	// 	{
	// 		ItemName:    "Horizontal Garden",
	// 		Tags:        []string{"program", "technology", "science", "science", "science", "science", "science"},
	// 		Description: "This is description of Aero Glider",
	// 		Url:         "/assets/Aero-Glider.jpg",
	// 	},
	// }

	var categories []Category = []Category{
		{
			CategoryName: "Service",
			SubCategories: []SubCategory{
				{
					SubCategoryName: "Digital Marketing",
					SubCategoryTag:  "digital_marketing",
					// Items:           digitalMarketingItems,
				},

				{
					SubCategoryName: "Garden",
					SubCategoryTag:  "garden",
					// Items:           gardenItems,
				},
			},
		},
	}
	return categories
}