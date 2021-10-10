package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/itp-backend/backend-b-antar-jemput/service"
)

type LocationController interface {
	GetAllProvinces(c *fiber.Ctx) error
    GetAllRegenciesByProvinceId(c *fiber.Ctx) error
}

type locationController struct {
	LocationService service.LocationService
}

func NewLocationController(service service.LocationService) LocationController {
	return locationController{
		LocationService: service,
	}
}

func (lc locationController) GetAllProvinces(c *fiber.Ctx) error {
	provinces, err := lc.LocationService.GetAllLocationProvince()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"code":    fiber.StatusBadRequest,
			"message": err.Error(),
			"data":    nil,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"code":    fiber.StatusOK,
		"message": nil,
		"data":    provinces,
	})
}


func (lc locationController) GetAllRegenciesByProvinceId(c *fiber.Ctx) error {
    provinceId := c.Query("provinceId")
    if provinceId == "" {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "code":    fiber.StatusBadRequest,
            "message": "Please Input the correct of Province Id",
            "data":    nil,
        })
    }
    regencies, err := lc.LocationService.GetAllRegencyByProvince(provinceId)
    if err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "code":    fiber.StatusInternalServerError,
            "message": "There is errors in server",
            "data":    nil,
        })
    }

    return c.Status(fiber.StatusOK).JSON(fiber.Map{
        "code":    fiber.StatusOK,
        "message": nil,
        "data":    regencies,
    })
}
