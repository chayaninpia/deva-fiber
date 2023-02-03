package handler

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/chayaninpia/deva-fiber/src/models"
	"github.com/chayaninpia/deva-fiber/src/utils"
	"github.com/gofiber/fiber/v2"
	"xorm.io/xorm"
)

func Power(c *fiber.Ctx) error {

	column := c.Params("column")
	number, _ := strconv.Atoi(c.Params("numberRecord"))

	engine, err := utils.InitXorm()
	if err != nil {
		return err
	}

	resQuery, err := QueryTPowerByLimit(engine, column, number)
	if err != nil {
		return err
	}

	res := SumTPower(resQuery)

	resMarshal, err := json.Marshal(&res)
	if err != nil {
		return err
	}

	return c.SendString(string(resMarshal))
}

func QueryTPowerByLimit(e *xorm.Engine, column string, number int) ([]models.TPowerO, error) {

	response := make([]models.TPowerO, 0)
	columns := ``

	switch column {
	case "active_power":
		columns = `active_power`
	case "power_input":
		columns = `power_input`
	case "all":
		columns = `active_power, power_input`
	default:
		return nil, fmt.Errorf(`invalid column: "%v" `, column)
	}

	qs := e.Table(models.TPowerI{}.GetTableName()).Select(columns).Limit(number)

	resQuery, err := qs.QueryInterface()
	if err != nil {
		return nil, err
	}

	switch column {
	case "active_power":
		for _, v := range resQuery {
			response = append(response, models.TPowerO{
				ActivePower: int(v["active_power"].(int32)),
			})
		}
	case "power_input":
		for _, v := range resQuery {
			response = append(response, models.TPowerO{
				PowerInput: int(v["power_input"].(int32)),
			})
		}
	case "all":
		for _, v := range resQuery {
			response = append(response, models.TPowerO{
				ActivePower: int(v["active_power"].(int32)),
				PowerInput:  int(v["power_input"].(int32)),
			})
		}
	}
	return response, nil
}

func SumTPower(req []models.TPowerO) *models.TPowerO {

	res := &models.TPowerO{}
	for _, v := range req {
		res.ActivePower += v.ActivePower
		res.PowerInput += v.PowerInput
	}

	return res
}
