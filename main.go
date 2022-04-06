package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	routesHandler()
}

func storeMonthlySalary(ginContext *gin.Context) {

	var monthlySalary monthlySalary

	if err := ginContext.BindJSON(&monthlySalary); err == nil {

		db, err := sql.Open("mysql", "root:root@tcp(localhost:3100)/mybills")

		if err != nil {
			panic(err)
		}

		db.SetConnMaxLifetime(time.Minute * 3)
		db.SetMaxOpenConns(10)
		db.SetMaxIdleConns(10)

		_, err = db.Exec(fmt.Sprintf("INSERT INTO monthly_salary(month, salary) values( %d, %.2f)", monthlySalary.Month, monthlySalary.Salary))

		if err == nil {
			ginContext.Status(http.StatusOK)
		} else {
			ginContext.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}

	} else {
		ginContext.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}
}
