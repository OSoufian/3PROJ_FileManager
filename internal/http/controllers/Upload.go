package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"mime"
	"os"
	"strconv"
	"strings"
	"time"
	"video/internal/domain"
	"path/filepath"

	"github.com/gofiber/fiber/v2"
)

type partialVideo struct {
	Id           uint   `json:"Id"`
	Name          string `json:"Name"`
	Description   string `json:"Description"`
	Icon          string `json:"Icon"`
	VideoURL      string `json:"VideoURL"`
	Views         int    `json:"Views"`
	Size          int64  `json:"Size"`
	CreatedAt     string `json:"CreatedAt"`
	CreationDate  string `json:"CreationDate"`
	IsHide     	  bool   `json:"IsHide"`
	IsBlock       bool   `json:"IsBlock"`
}

type partialCreateVideo struct {
	partialVideo
	ChannelId int64 `json:"channId"`
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

	router.Post("/video", uploadVideo)

	router.Post("/image", uploadImage)

	router.Get("/detail", videoDetail)

	router.Put("/", createWithoutUplaod)

	router.Patch("/", patchVideoByFileName)

	router.Delete("/", deleteVideo)

	router.Get("/files", retrieveAllFile)

	router.Get("/files/:id", getVideoByFileId)
}

// upload image
// @Summary Files
// @Description retrieve a file
// @Tags Files
// @Success 200 {Videos} Videos "video info"
// @MultipartForm video
// @MultipartForm image
// @MultipartForm info
// @Failure 404
// @Router /files [post]
func uploadImage(c *fiber.Ctx) error {

	// c.Request().ContinueReadBodyStream()
	fileType := "image"

	// Get the file from form data
	form, err := c.MultipartForm()
	if err != nil {
		return err
	}

	files := form.File["image"]
	// Change to match your field name
	if len(files) == 0 {
		// Change to match your field name
		files = form.File["image"]
	}

	if len(files) == 0 {
		return c.SendString("no file found")
	}
	file := files[0]

	if fileType == "video" {
		return c.SendString("not an image")
	}

	// Parse the filename parameter from the header
	header := file.Header.Get("Content-Disposition")
	_, params, err := mime.ParseMediaType(header)
	if err != nil {
		log.Println(err)
		return err
	}
	filename, ok := params["filename"]
	if !ok {
		return c.SendString("filename parameter not found")
	}

	filePath := fmt.Sprintf("./data/images/%s", filename)
	if err := c.SaveFile(file, filePath); err != nil {
		return err
	}
	// Return success
	return c.Status(fiber.StatusAccepted).SendString(filename)
}

// // Upload video
// // @Summary Files
// // @Description retrieve a file
// // @Tags Files
// // @Success 200 {Videos} Videos "video info"
// // @MultipartForm video
// // @MultipartForm image
// // @MultipartForm info
// // @Failure 404
// // @Router /files [post]
func uploadVideo(c *fiber.Ctx) error {
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
		c.SendString("not a video")
	}

	if len(files) == 0 {
		return c.SendString("no file found")
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
		return c.SendString("filename parameter not found")
	}

	// filePath := fmt.Sprintf("./data/images/%s", filename)
	if fileType == "video" {

		video := new(domain.Videos)
		video.VideoURL = filename

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
		channel, err := channel.Get()
		if err != nil {
			return c.SendStatus(fiber.ErrBadGateway.Code)
		}

		video.ChannelId = channel.Id

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

		tmpTime := strings.Split(time.Now().UTC().String(), " ")

		video.CreatedAt = strings.Join(tmpTime[:len(tmpTime)-1], " ")

		video.Create()
		filePath := fmt.Sprintf("./data/videos/%s", filename)

		if err := c.SaveFile(file, filePath); err != nil {
			return err
		}
		// Return status accepted
		return c.Status(fiber.StatusAccepted).JSON(fiber.Map{
			"id":       video.Id,
			"filename": filename,
			"status": 200,
		})
	}
	return c.SendStatus(fiber.ErrBadRequest.Code)
}

// Get All video
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

// Get All video
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
	channel, err := channel.Get()
	if err != nil {
		return c.SendStatus(fiber.ErrBadGateway.Code)
	}

	video.ChannelId = channel.Id

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

// Get All video
// @Summary Files
// @Description retrieve a file
// @Tags Files
// @Success 200 {Videos} Videos "video info"
// @Query path
// @Params {partialVideo}
// @Failure 404
// @Router /files [patch]
func patchVideoByFileName(c *fiber.Ctx) error {
	// path := c.Query("path")

	video := new(domain.Videos)

	partial := new(partialVideo)
	if err := partial.Unmarshal(c.Body()); err != nil {
		return c.Status(fiber.ErrInternalServerError.Code).JSON(fiber.Map{
			"err": err.Error(),
		})
	}
	video.Id = partial.Id

	video.GetById()

	if partial.Name != "" {
		video.Name = partial.Name
	}
	if partial.Description != "" {
		video.Description = partial.Description
	}
	if partial.Icon != "" {
		video.Icon = partial.Icon
	}
	video.IsHide = partial.IsHide
	video.IsBlock = partial.IsBlock
	video.Views = partial.Views
	
	video.Update()
	
	return c.Status(fiber.StatusAccepted).JSON(video)

}

// Delete video
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

	if err := os.Remove("./data/videos/" + filename); err != nil {
		log.Println(err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	// Return the file with correct Content-Length header
	return c.SendStatus(201)
}

// Get video by file id
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

// Get All files
// @Summary Files
// @Description retrieve a file
// @Tags Files
// @Success 200 {Videos} Videos "video info"
// @Params id
// @Failure 404
// @Router /files/files [get]
func retrieveAllFile(c *fiber.Ctx) error {

    // Define the directory path
    directoryPath := "./data"

    // Define a slice to store the file names
    var fileNames []string

    // Walk the directory and its subdirectories
    err := filepath.Walk(directoryPath, func(path string, info os.FileInfo, err error) error {
        if err != nil {
            return err
        }
        if !info.Mode().IsRegular() {
            return nil
        }
        fileNames = append(fileNames, info.Name())
        return nil
    })
    if err != nil {
        return err
    }

    // Return the file names as a JSON response
    return c.JSON(fileNames)
}
