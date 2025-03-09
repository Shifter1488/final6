package actioninfo

import (
	"fmt"
)

type DataParser interface {
	Parse(datastring string) error
	ActionInfo() string
}

func Info(dataset []string, dp DataParser) {
	for _, v := range dataset {
		err := dp.Parse(v)
		if err != nil {
			fmt.Println("Ошибка при парсинге данных:", err)
			continue
		}
		fmt.Println(dp.ActionInfo())
	}
}
