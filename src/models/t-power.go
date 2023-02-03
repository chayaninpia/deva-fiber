package models

type TPowerI struct {
}

type TPowerO struct {
	ActivePower int `json:"active_power"`
	PowerInput  int `json:"power_input"`
}

func (TPowerI) GetTableName() string {
	return "t_power"
}
