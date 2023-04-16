package controllers

import (
	"strconv"
	"video/internal/domain"

	"github.com/gofiber/fiber/v2"
)

func Videos(router fiber.Router) {

	router.Get("/", getAllvideos)

	router.Get("/:videoId", getVideoById)

	router.Get("/chann/:channId", getChannelVideos)

}

// Get All video
// @Summary Videos
// @Description get all video
// @Tags Videos
// @Success 200 {Videos} List Videos
// @Failure 404
// @Router /video [get]
func getAllvideos(c *fiber.Ctx) error {
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

	return c.Status(fiber.StatusAccepted).JSON(video.GetAllVideosFromChannel())

}
