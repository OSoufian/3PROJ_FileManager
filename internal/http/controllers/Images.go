package controllers

import (
	"os"

	"github.com/gofiber/fiber/v2"
)

func Images(router fiber.Router) {

	router.Get("/", retrieveImage)

}

// Get image
// @Summary Files
// @Description retrieve a file
// @Tags Files
// @Success 200 {Blob} Retrieve a blob file
// @Query imagename
// @Failure 404
// @Router /files [get]
func retrieveImage(c *fiber.Ctx) error {
	// Get the imagename from the request parameters
	imagename := c.Query("imagename")

	// Open the file
	file, err := os.Open("./data/images/" + imagename)
	if err != nil {
		return err
	}
	defer file.Close()

	// Return the file with correct Content-Length header
	return c.SendFile(file.Name())
}