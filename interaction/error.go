package interaction

import "errors"

func Errk(answer string) (string, error) {
	if answer != "хищные" && answer != "птицы" && answer != "млекопитающие" {
		return "", errors.New("Неверный ввод типа животного")
	}
	return answer, nil
}


