package main

import (
	config "shopifun/Config"
	model "shopifun/Model"
	"shopifun/router"
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/log"
	"github.com/gofiber/fiber/v3/middleware/cors"
	"github.com/gofiber/fiber/v3/middleware/limiter"
)

func main() {

	config.ConnectDB()

	db := config.DB

	// fmt.Printf("succes connect db %v", db)

	db.AutoMigrate(&model.User{})
	app := fiber.New()
	app.Use(cors.New())

	app.Use(limiter.New(limiter.Config{
		Max:        100,
		Expiration: 30 * time.Second,
	}))

	router.SetupRoutes(app)

	// type Person struct {
	// 	Name    string `json:"name"`
	// 	Age     uint8
	// 	Address string `json:"address"`
	// }

	// type ResponseMessage struct {
	// 	Status  int
	// 	Message string
	// 	Data    interface{}
	// }

	// app.Get("/", func(c fiber.Ctx) error {
	// 	data := Person{
	// 		Name:    "Grame",
	// 		Age:     20,
	// 		Address: "Jakarta",
	// 	}
	// 	return c.JSON(data)
	// })

	// // pendekatan params
	// app.Get("/person/:name", func(c fiber.Ctx) error {
	// 	namePerson := c.Params("name")

	// 	sendData := Person{
	// 		Name:    namePerson,
	// 		Age:     20,
	// 		Address: "Jakarta",
	// 	}

	// 	return c.JSON(sendData)
	// })

	// // query
	// app.Get("/personaja/", func(c fiber.Ctx) error {
	// 	name := c.Query("name")
	// 	ageInt, _ := strconv.Atoi(c.Query("age"))

	// 	age := uint8(ageInt)

	// 	sendData := Person{
	// 		Name:    name,
	// 		Age:     age,
	// 		Address: "Jakarta",
	// 	}
	// 	return c.JSON(sendData)
	// })

	// // body json
	// app.Post("/register", func(c fiber.Ctx) error {
	// 	// Get raw body from POST request:

	// 	auth := c.Get("X-code")

	// 	body := c.Body()

	// 	var person Person
	// 	if err := json.Unmarshal(body, &person); err != nil {
	// 		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
	// 			"error": "Invalid JSON",
	// 		})
	// 	}
	// 	name := person.Name
	// 	age := person.Age

	// 	response := ResponseMessage{
	// 		Status:  200,
	// 		Message: "Success get data",
	// 		Data: fiber.Map{
	// 			"name": name,
	// 			"age":  age,
	// 			"auth": auth,
	// 		},
	// 	}

	// 	return c.Status(fiber.StatusAccepted).JSON(response)
	// })

	// body url encoded

	log.Fatal(app.Listen(":3000"))
}
