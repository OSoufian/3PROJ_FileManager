package controllers

import (
    "strconv"
    "video/internal/domain"

    "github.com/gofiber/fiber/v2"
)

func Videos(router fiber.Router) {

    router.Get("/", func(c *fiber.Ctx) error {
        videosModels := domain.Videos{}
        videos, err := videosModels.GetAll()
        if err != nil {
            return c.SendStatus(fiber.ErrBadRequest.Code)
        }
        return c.Status(200).JSON(videos)
    })

    router.Get("/:videoId", func(c *fiber.Ctx) error {
        id := c.Params("videoId")
        videoId, err := strconv.ParseInt(id, 10, 64)
        if err != nil {
            return c.Status(fiber.ErrBadRequest.Code).JSON(err)
        }
        video := domain.Videos{}

        video.Id = uint(videoId)

        return c.Status(fiber.StatusAccepted).JSON(video.GetById())
    })

    router.Get("/chann/:channId", func(c *fiber.Ctx) error {
        id := c.Params("channId")
        channId, err := strconv.ParseInt(id, 10, 64)

        if err != nil {
            return c.Status(fiber.ErrBadRequest.Code).JSON(err)
        }

        video := domain.Videos{}

        video.ChannelId = uint(channId)

        return c.Status(fiber.StatusAccepted).JSON(video.GetAllVideosFromChannel())

    })

}