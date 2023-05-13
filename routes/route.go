package routes

import (
	"ukprakerja/controllers"

	"github.com/labstack/echo/v4"
)

// Untuk routing ke controller API
func InitRoute() *echo.Echo {
	e := echo.New()
	e.GET("/items", controllers.GetItemsController)
	e.POST("/items", controllers.InsertItemController)
	e.GET("/items/:id", controllers.GetSingleItemController)
	e.PUT("/items/:id", controllers.UpdateItemController)
	e.DELETE("/items/:id", controllers.DeleteItemController)
	e.GET("/employees", controllers.GetEmployeesController)
	e.POST("/employees", controllers.InsertEmployeeController)
	e.GET("/employees/:id", controllers.GetSingleEmployeeController)
	e.PUT("/employees/:id", controllers.UpdateEmployeeController)
	e.DELETE("/employees/:id", controllers.DeleteEmployeeController)
	return e
}
