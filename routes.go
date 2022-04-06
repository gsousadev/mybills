package main

import "github.com/gin-gonic/gin"

func routesHandler() {
	routes := gin.Default()

	routes.POST("/salario-mensal", storeMonthlySalary)

	routes.Run()
}
