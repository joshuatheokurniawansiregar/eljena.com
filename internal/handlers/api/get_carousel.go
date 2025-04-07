package api

import "github.com/gofiber/fiber/v2"

type CarouselData struct{
	ImageUrl string `json:"imageUrl"`
	ItemName string `json:"itemName"`
	Description string `json:"description"`
	SubCategory string `json:"subCategory"`
	Tags []string `json:"tags"`
}

func(h *Handler) GetCarousel(ctx *fiber.Ctx) error {
	var carouselData []CarouselData = []CarouselData{
		{
			ImageUrl: "/assets/Aero-Glider.jpg",
			ItemName: "Aero Glider",
			Description: "Goes To School adalah service kegiatan belajar-mengajar yang dilakukan tim kami.",
			SubCategory: "Goes To School",
			Tags: []string{"program","technology", "science"},
		},
		{
			ImageUrl: "/assets/boneka_potty.jpg",
			ItemName: "Boneka Potty",
			Description: "lorem ipsum 2",
			SubCategory: "Test 1",
			Tags: []string{"program","math", "test 1", "test 2", "test 3"},
		},
		{
			ImageUrl: "/assets/boneka_potty.jpg",
			ItemName: "Layang-Layang",
			Description: "this is description",
			SubCategory: "Test 2",
			Tags: []string{"program","technolgy", "creativity"},
		},
		{
			ImageUrl: "/assets/boneka_potty.jpg",
			ItemName: "Ondel-Ondel",
			Description: "this is description",
			SubCategory: "Test 2",
			Tags: []string{"program","technolgy", "math"},
		},
	}

	return ctx.Status(200).JSON(carouselData)
}