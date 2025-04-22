package rest

import (
	"WorkmateTask/conf"
	"WorkmateTask/db"
	"WorkmateTask/usecase"
	"encoding/json"
	"github.com/labstack/echo/v4"
	"github.com/redis/go-redis/v9"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

func Server(con conf.Conf, db db.DbAccess) {
	use := usecase.New(db)
	use.Workmate()
	
	e := echo.New()
	e.GET("/sum", func(c echo.Context) error {
		var sum struct {
			id     string
			value1 int
			value2 int
		}
		sum.id = strconv.Itoa(rand.Int())
		sumJson, _ := json.Marshal(sum)

		db.AddTask("sum", sumJson)

		result := 0
		for {

			resultByte, err := db.GetResult(sum.id)
			if err == redis.Nil {
				time.Sleep(1 * time.Second)
				continue
			}

			json.Unmarshal(resultByte, &result)
			break
		}
		return c.JSON(http.StatusOK, result)
	})

	e.Logger.Fatal(e.Start(con.HttpPort))
}
