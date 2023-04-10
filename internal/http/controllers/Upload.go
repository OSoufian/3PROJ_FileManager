package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"mime"
	"os"
	"strconv"
	"time"
	"video/internal/domain"

	"github.com/gofiber/fiber/v2"
)

type partialVideo struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
}

type partialCreateVideo struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
	ChannelId   int64  `json:"channelId"`
}

func (p *partialVideo) Unmarshal(body []byte) error {
	return json.Unmarshal(body, &p)
}
func (p *partialCreateVideo) Unmarshal(body []byte) error {
	return json.Unmarshal(body, &p)
}

func (p *partialCreateVideo) UnmarshalString(body string) error {
	return json.Unmarshal([]byte(body), &p)
}

func Uploader(router fiber.Router) {

	router.Post("/:type", fileUpload)

	router.Get("/", retriveFile)

	router.Get("/detail", videoDetail)

	router.Put("/", createWithoutUplaod)

	router.Patch("/", patchVideoByFileName)

	router.Delete("/", deleteVideo)

	router.Get("/files", retriveAllFile)

	router.Get("/files/:id", getVideoByFileId)
}

// Get All videos
// @Summary Files
// @Description retrieve a file
// @Tags Files
// @Success 200 {Videos} Videos "video info"
// @MultipartForm video
// @MultipartForm image
// @MultipartForm info
// @Failure 404
// @Router /files [post]
func fileUpload(c *fiber.Ctx) error {

	// c.Request().ContinueReadBodyStream()

	fileType := "video"

	// Get the file from form data
	form, err := c.MultipartForm()
	if err != nil {
		return err
	}

	files := form.File["video"]
	// Change to match your field name
	if len(files) == 0 {
		// Change to match your field name
		files = form.File["image"]
		fileType = "image"
	}

	if len(files) == 0 {
		return fmt.Errorf("no file found")
	}
	file := files[0]

	// Parse the filename parameter from the header
	header := file.Header.Get("Content-Disposition")
	_, params, err := mime.ParseMediaType(header)
	if err != nil {
		log.Println(err)
		return err
	}
	filename, ok := params["filename"]
	if !ok {
		return fmt.Errorf("filename parameter not found")
	}

	if fileType == "video" {

		video := new(domain.Videos)
		video.VideoURL = filename
		video.CreationAt = time.Now()

		channel := new(domain.Channel)

		partial := new(partialCreateVideo)

		videosProperties, ok := form.Value["info"]
		if !ok {
			return fmt.Errorf("info parameter not found")
		}

		if err := partial.UnmarshalString(videosProperties[0]); err != nil {
			return c.Status(fiber.ErrInternalServerError.Code).JSON(fiber.Map{
				"err": err.Error(),
			})
		}

		channel.Id = uint(partial.ChannelId)

		if channel.Get() == nil {
			return c.SendStatus(fiber.ErrBadGateway.Code)
		}

		video.ChannelId = channel.Get().Id

		if partial.Name != "" {
			video.Name = partial.Name
		}
		if partial.Description != "" {
			video.Description = partial.Description
		}
		if partial.Icon != "" {
			video.Icon = partial.Icon
		}

		video.Size = file.Size

		video.Create()
	}

	// Return success
	return c.SaveFile(file, fmt.Sprintf("./data/%s", filename))
}

// Get All videos
// @Summary Files
// @Description retrieve a file
// @Tags Files
// @Success 200 {Blob} Retrieve a blob file
// @Query filename
// @Failure 404
// @Router /files [get]
func retriveFile(c *fiber.Ctx) error {
	// Get the filename from the request parameters
	filename := c.Query("filename")

	// Open the file
	file, err := os.Open("./data/" + filename)
	if err != nil {
		return err
	}
	defer file.Close()

	// Return the file with correct Content-Length header
	return c.SendFile(file.Name())
}

// Get All videos
// @Summary Files
// @Description retrieve a file
// @Tags Files
// @Success 200 {Videos} Videos "video info"
// @Query path
// @Failure 404
// @Router /files/detail [get]
func videoDetail(c *fiber.Ctx) error {
	path := c.Query("path")
	video := new(domain.Videos)
	video.VideoURL = path

	return c.JSON(video.Get())
}

// Get All videos
// @Summary Files
// @Description retrieve a file
// @Tags Files
// @Success 200 {Videos} Videos "video info"
// @Params {partialCreateVideo}
// @Failure 404
// @Router /files [put]
func createWithoutUplaod(c *fiber.Ctx) error {
	path := c.Query("path")

	video := new(domain.Videos)
	video.VideoURL = path
	video = video.Get()

	channel := new(domain.Channel)

	partial := new(partialCreateVideo)
	if err := partial.Unmarshal(c.Body()); err != nil {
		return c.Status(fiber.ErrInternalServerError.Code).JSON(fiber.Map{
			"err": err.Error(),
		})
	}

	channel.Id = uint(partial.ChannelId)

	if channel.Get() == nil {
		return c.SendStatus(fiber.ErrBadGateway.Code)
	}

	video.ChannelId = channel.Get().Id

	if partial.Name != "" {
		video.Name = partial.Name
	}
	if partial.Description != "" {
		video.Description = partial.Description
	}
	if partial.Icon != "" {
		video.Icon = partial.Icon
	}

	video.Create()

	return c.Status(fiber.StatusAccepted).JSON(video)
}

// Get All videos
// @Summary Files
// @Description retrieve a file
// @Tags Files
// @Success 200 {Videos} Videos "video info"
// @Query path
// @Params {partialVideo}
// @Failure 404
// @Router /files [patch]
func patchVideoByFileName(c *fiber.Ctx) error {
	path := c.Query("path")

	video := new(domain.Videos)
	video.VideoURL = path
	video = video.Get()

	partial := new(partialVideo)
	if err := partial.Unmarshal(c.Body()); err != nil {
		return c.Status(fiber.ErrInternalServerError.Code).JSON(fiber.Map{
			"err": err.Error(),
		})
	}

	if partial.Name != "" {
		video.Name = partial.Name
	}
	if partial.Description != "" {
		video.Description = partial.Description
	}
	if partial.Icon != "" {
		video.Icon = partial.Icon
	}

	video.Update()

	return c.Status(fiber.StatusAccepted).JSON(video)

}

// Get All videos
// @Summary Files
// @Description retrieve a file
// @Tags Files
// @Success 201
// @Query filename
// @Failure 404
// @Router /files [delete]
func deleteVideo(c *fiber.Ctx) error {
	// Get the filename from the request parameters
	filename := c.Query("filename")

	video := domain.Videos{}
	video.VideoURL = filename

	if video.Find() {
		video.Delete()
	}

	if err := os.Remove("./data/" + filename); err != nil {
		log.Println(err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	// Return the file with correct Content-Length header
	return c.SendStatus(201)
}

// Get All videos
// @Summary Files
// @Description retrieve a file
// @Tags Files
// @Success 200 {Videos} Videos "video info"
// @Params id
// @Failure 404
// @Router /files/:id [get]
func getVideoByFileId(c *fiber.Ctx) error {
	video := domain.Videos{}
	videoParams := c.Params("id")

	videoId, err := strconv.ParseUint(videoParams, 10, len(videoParams))

	if err != nil {
		return c.Status(fiber.ErrBadGateway.Code).JSON(err)
	}

	video.Id = uint(videoId)

	return c.SendFile(video.VideoURL)
}

// Get All videos
// @Summary Files
// @Description retrieve a file
// @Tags Files
// @Success 200 {Videos} Videos "video info"
// @Params id
// @Failure 404
// @Router /files/files [get]
func retriveAllFile(c *fiber.Ctx) error {

	// quering := c.Query("quering")

	// Define the directory path
	directoryPath := "./data"

	// Open the directory
	directory, err := os.Open(directoryPath)
	if err != nil {
		return err
	}
	defer directory.Close()

	// Read the directory contents
	files, err := directory.Readdir(0)
	if err != nil {
		return err
	}

	// Define a slice to store the file names
	var fileNames []string

	// Loop through the files and add their names to the slice
	for _, file := range files {
		if file.Mode().IsRegular() {
			fileNames = append(fileNames, file.Name())
		}
	}

	// Return the file names as a JSON response
	return c.JSON(fileNames)
}
