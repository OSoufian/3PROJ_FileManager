package controllers

import (
	"strconv"
	"video/internal/domain"
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func Videos(router fiber.Router) {

	router.Get("/all", getAllvideos)

	router.Get("/:videoId", getVideoById)

	router.Get("/chann/:channId", getChannelVideos)

	router.Delete("/chann/:channId", deleteChannelVideos)

	router.Get("/", retrieveVideo)
}

// Get All video
// @Summary Videos
// @Description get all video
// @Tags Videos
// @Success 200 {Videos} List Videos
// @Failure 404
// @Router /video [get]
func getAllvideos(c *fiber.Ctx) error {
		
	if strings.TrimSpace(c.Query("q")) != "" {
		return getSearchVideos(c)
	}

	videosModels := domain.Videos{}
	video, err := videosModels.GetAll()
	if err != nil {
		return c.SendStatus(fiber.ErrBadRequest.Code)
	}
	return c.Status(200).JSON(video)
}

// Get Video by Id
// @Summary Videos
// @Description get a video by id
// @Tags Videos
// @Success 200 {Videos} Get a Video by id
// @Failure 404
// @Router /video/:videoID [get]
func getVideoById(c *fiber.Ctx) error {
	id := c.Params("videoId")
	videoId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return c.Status(fiber.ErrBadRequest.Code).JSON(err)
	}
	video := domain.Videos{}

	video.Id = uint(videoId)

	return c.Status(fiber.StatusAccepted).JSON(video.GetById())
}

// Get Channel Videos
// @Summary Videos
// @Description get all video from a channel
// @Tags Videos
// @Success 200 {Videos} List of Videos
// @Failure 404
// @Router /video/chann/:channId [get]
func getChannelVideos(c *fiber.Ctx) error {
	id := c.Params("channId")
	channId, err := strconv.ParseInt(id, 10, 64)

	if err != nil {
		return c.Status(fiber.ErrBadRequest.Code).JSON(err)
	}

	video := domain.Videos{}
	video.ChannelId = uint(channId)

	orderByParams := c.Query("orderBy")
	orderBy := strings.Split(orderByParams, ",")
	orderedVideos := video.GetAllVideosFromChannel(orderBy...)
	return c.Status(fiber.StatusAccepted).JSON(orderedVideos)
}

func deleteChannelVideos(c *fiber.Ctx) error {
	id := c.Params("channId")
	channId, err := strconv.ParseInt(id, 10, 64)

	if err != nil {
		return c.Status(fiber.ErrBadRequest.Code).JSON(err)
	}

	video := domain.Videos{}

	video.DeleteAllVideosFromChannel(uint(channId))
	return c.Status(fiber.StatusOK).JSON(video)
}

// Get video
// @Summary Files
// @Description retrieve a file
// @Tags Files
// @Success 200 {Blob} Retrieve a blob file
// @Query videoname
// @Failure 404
// @Router /video?videoname [get]
func retrieveVideo(c *fiber.Ctx) error {
	// Get the videoname from the request parameters
	videoname := c.Query("videoname")

	// Open the file
	file, err := os.Open("./data/videos/" + videoname)
	if err != nil {
		return err
	}
	defer file.Close()

	// Return the file with correct Content-Length header
	return c.SendFile(file.Name())
}

// Get Searched video
// @Summary Videos
// @Description get searched video
// @Tags Videos
// @Success 200 {Videos} List Videos
// @Failure 404
// @Router /video [get]
func getSearchVideos(c *fiber.Ctx) error {
	searchQuery := c.Query("q") // Get the search query from the query parameter "q"
	videosModels := domain.Videos{}
	video, err := videosModels.GetSearch(searchQuery)
	if err != nil {
		return c.SendStatus(fiber.ErrBadRequest.Code)
	}
	return c.Status(fiber.StatusOK).JSON(video)
}

// package controllers

// import (
// 	"strconv"
// 	"video/internal/domain"
// 	"os"
// 	"strings"

// 	"github.com/gofiber/fiber/v2"
// )

// func Videos(router fiber.Router) {

// 	router.Get("/all", getAllvideos)

// 	router.Get("/:videoId", getVideoById)

// 	router.Get("/chann/:channId", getChannelVideos)

// 	router.Get("/", retrieveVideo)
// }

// // Get All video
// // @Summary Videos
// // @Description get all video
// // @Tags Videos
// // @Success 200 {Videos} List Videos
// // @Failure 404
// // @Router /video [get]
// func getAllvideos(c *fiber.Ctx) error {
		
// 	if strings.TrimSpace(c.Query("q")) != "" {
// 		return getSearchVideos(c)
// 	}
	
// 	videosModels := domain.Videos{}

// 	orderByParams := c.Query("orderBy")
// 	orderBy := strings.Split(orderByParams, ",")

// 	orderedVideos, err := videosModels.GetAll(orderBy...)
// 	if err != nil {
// 		return c.SendStatus(fiber.ErrBadRequest.Code)
// 	}
// 	return c.Status(200).JSON(orderedVideos)
// }

// // Get Video by Id
// // @Summary Videos
// // @Description get a video by id
// // @Tags Videos
// // @Success 200 {Videos} Get a Video by id
// // @Failure 404
// // @Router /video/:videoID [get]
// func getVideoById(c *fiber.Ctx) error {
// 	id := c.Params("videoId")
// 	videoId, err := strconv.ParseInt(id, 10, 64)
// 	if err != nil {
// 		return c.Status(fiber.ErrBadRequest.Code).JSON(err)
// 	}
// 	video := domain.Videos{}

// 	video.Id = uint(videoId)

// 	return c.Status(fiber.StatusAccepted).JSON(video.GetById())
// }

// // Get Channel Videos
// // @Summary Videos
// // @Description get all video from a channel
// // @Tags Videos
// // @Success 200 {Videos} List of Videos
// // @Failure 404
// // @Router /video/chann/:channId [get]
// func getChannelVideos(c *fiber.Ctx) error {
// 	id := c.Params("channId")
// 	channId, err := strconv.ParseInt(id, 10, 64)

// 	if err != nil {
// 		return c.Status(fiber.ErrBadRequest.Code).JSON(err)
// 	}

// 	video := domain.Videos{}
// 	video.ChannelId = uint(channId)

// 	orderByParams := c.Query("orderBy")
// 	orderBy := strings.Split(orderByParams, ",")
// 	orderedVideos := video.GetAllVideosFromChannel(orderBy...)
// 	return c.Status(fiber.StatusAccepted).JSON(orderedVideos)
// }

// // Get video
// // @Summary Files
// // @Description retrieve a file
// // @Tags Files
// // @Success 200 {Blob} Retrieve a blob file
// // @Query videoname
// // @Failure 404
// // @Router /video?videoname [get]
// func retrieveVideo(c *fiber.Ctx) error {
// 	// Get the videoname from the request parameters
// 	videoname := c.Query("videoname")

// 	// Open the file
// 	file, err := os.Open("./data/videos/" + videoname)
// 	if err != nil {
// 		return err
// 	}
// 	defer file.Close()

// 	// Return the file with correct Content-Length header
// 	return c.SendFile(file.Name())
// }

// // Get Searched video
// // @Summary Videos
// // @Description get searched video
// // @Tags Videos
// // @Success 200 {Videos} List Videos
// // @Failure 404
// // @Router /video [get]
// func getSearchVideos(c *fiber.Ctx) error {
// 	searchQuery := c.Query("q") // Get the search query from the query parameter "q"

// 	videosModels := domain.Videos{}
// 	orderByParams := c.Query("orderBy")
// 	orderBy := strings.Split(orderByParams, ",")
// 	orderedVideos, err := videosModels.GetSearch(searchQuery, orderBy...)
// 	if err != nil {
// 		return c.SendStatus(fiber.ErrBadRequest.Code)
// 	}
// 	return c.Status(fiber.StatusOK).JSON(orderedVideos)
// }