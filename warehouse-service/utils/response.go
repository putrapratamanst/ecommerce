package utils

import "github.com/gofiber/fiber/v2"

type APIResponse struct {
    Status  string      `json:"status"`
    Message string      `json:"message"`
    Data    interface{} `json:"data,omitempty"`
}

func SendResponse(c *fiber.Ctx, status int, message string, data interface{}) error {
    response := APIResponse{
        Status:  "success",
        Message: message,
        Data:    data,
    }

    // Menyesuaikan response jika status adalah error
    if status >= 400 {
        response.Status = "error"
    }

    return c.Status(status).JSON(response)
}
