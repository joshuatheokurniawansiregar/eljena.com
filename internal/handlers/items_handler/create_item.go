package items_handler

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/joshuatheokurniawansiregar/eljena/internal/models/items"
)

func isAllowedExtensions(filename string, allowedExtensions []string) error{
	var ext string = filepath.Ext(filename)
	for _, allowedExt := range allowedExtensions{
		if ext == allowedExt{
			return nil
		}
	}
	return errors.New("the file extension is not allowed")
}

func (h *Handler) CreateItem(ctx *fiber.Ctx)error{
	var itemRequest items.ItemRequest
	
	if err := ctx.BodyParser(&itemRequest); err != nil{
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"code": http.StatusBadRequest,
			"message": err.Error(),
			"data": nil,
		})
	}

	form, err := ctx.MultipartForm()
	if err == nil {
		for key, files := range form.File {
			fmt.Printf("Key: %s, Files: %+v\n", key, files)
		}
	}

	file, err := ctx.FormFile("file")
	if err != nil{
		fmt.Println(err)
		return ctx.Status(http.StatusBadRequest).Status(http.StatusBadRequest).JSON(fiber.Map{
			"code": http.StatusBadRequest,
			"message": err.Error(),
			"data": nil,
		})
	}
	fmt.Println("size")

	


	var allowedExtensions []string = []string{".jpg", ".jpeg", ".png"}
	var filename string= file.Filename
	var extension string = filepath.Ext(filename)
	var directory string= "./public/uploads/"
	var counter int64 = 1
	var basename = strings.TrimSuffix(filename, extension)
	var currentCounterFile string = fmt.Sprintf("%s_%d%s", basename, counter, extension)

	if err = isAllowedExtensions(filename, allowedExtensions); err != nil{
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"code": http.StatusInternalServerError,
			"message": fmt.Sprintf("extension: %s is not allowed", extension),
			"data":nil,
		})
	}


	if _, err:= os.Stat(filepath.Join(directory, filename)); err == nil{
		if _, err := os.Stat(filepath.Join(directory, currentCounterFile));err == nil{
			for{
				var nextCounterFile string = fmt.Sprintf("%s_%d%s", basename, counter, extension)
				currentCounterFile = nextCounterFile
				counter++
				if _, err := os.Stat(filepath.Join(directory, currentCounterFile));os.IsNotExist(err){
					break
				}
			}
			filename = currentCounterFile
		}else{
			filename = currentCounterFile
		}
	}


	ctx.SaveFile(file, fmt.Sprintf("%s%s", directory, filename))
	
	var uploadsDirectoryUrl string = fmt.Sprintf("assets/uploads/%s", filename)

	if err := h.ItemsService.CreateItem(ctx.Context(), itemRequest, uploadsDirectoryUrl);err != nil{
		return ctx.JSON(fiber.Map{
			"code": http.StatusInternalServerError,
			"message": err.Error(),
			"data": nil,
		})
	}

	return ctx.Status(200).JSON(fiber.Map{
		"code": http.StatusOK,
		"message": "Create a new item succeed",
		// "data": map[string]any{
		// 	"itemResult": itemRequest,
		// },
	})
}