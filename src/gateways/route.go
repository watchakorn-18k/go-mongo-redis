package gateways

import "github.com/gofiber/fiber/v2"

func GatewayUsers(gateway HTTPGateway, app *fiber.App) {
	api := app.Group("/api/redis_mongo")

	api.Get("/get_users_redis", gateway.GetAllUserDataRedis)
	api.Get("/set_users_redis", gateway.SetAllUserDataRedis)
	api.Get("/get_users_mongo", gateway.GetAllUserData)
}
