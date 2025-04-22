package conf

import (
	"encoding/json"
	"fmt"
	"os"
)

type Conf struct {
	HttpPort string
	RdbUrl   string
}

func NewConf() (Conf, error) {
	var con Conf

	file, err := os.Open("Conf.cfg")
	if err != nil {
		return Conf{}, fmt.Errorf("open Conf.cfg: %w", err)
	}
	defer file.Close()

	err = json.NewDecoder(file).Decode(&con)

	if err != nil {
		return Conf{}, fmt.Errorf("decode Conf.cfg: %w", err)
	}

	return con, err
}
