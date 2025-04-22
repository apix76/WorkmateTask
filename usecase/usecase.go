package usecase

import (
	"WorkmateTask/db"
	"encoding/json"
	"log"
)

type Usecase struct {
	DbAccess db.DbAccess
}

func New(dbAccess db.DbAccess) *Usecase {
	return &Usecase{dbAccess}
}

func (uc *Usecase) Workmate() {
	uc.sum()
}

func (uc *Usecase) sum() {

	for {
		var data struct {
			id     string
			value1 int
			value2 int
		}
		rawData, err := uc.DbAccess.GetTask("sum")

		err = json.Unmarshal(rawData, &data)
		if err != nil {
			log.Println(err)
			continue
		}
		sum := data.value1 + data.value2

		sumJson, _ := json.Marshal(sum)
		uc.DbAccess.AddResult(data.id, sumJson)
	}
}
