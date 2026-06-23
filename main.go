package main

import (
	"errors"
	"fmt"
)

// sub 5
type invalidRadiusError struct {
}

func (d invalidRadiusError) Error() string {
	return "Радиус не может быть отрицательным"
}

// sub6
func calculateCircleArea(d int) (float64, error) {
	result := float64(d)
	if d < 0 {
		result = 0
	} else {
		return result, nil
	}

	return result, invalidRadiusError{}

}

//sub 7

func findUser(s []string, d string) (int, error) {
	for i, v := range s {
		if v == d {
			return i, nil
		}

	}
	return -1, errors.New("пользователь не найден")

}

//sub 2

func divide(d, s int) (int, interface{}) {

	if s == 0 {
		return 0, errors.New("Ошибка: деление на ноль невозможно")
	}
	return d / s, nil

}

//sub 3

func checkAge(d int) error {

	if d <= 0 || d >= 120 {
		return errors.New("Указан некоректный возраст")
	} else {
		return nil
	}

}

// sub 4

func validatePassword(d string) (bool, error) {
	if len(d) < 6 {
		return false, errors.New("Пароль слишком короткий")
	}
	return true, nil
}

func main() {
	//sub 1

	m := map[string]int{"Pen": 2, "Clock": 5}

	result, ok := m["Pen"]

	if ok {
		fmt.Println(result)

	} else {
		fmt.Println("Товар отсутствует")
	}

	// sub 3

	fmt.Println(checkAge(12))

	// sub 4

	fmt.Println(validatePassword("ass"))
	fmt.Println(validatePassword("assasdas"))

	//sub 5
	err := invalidRadiusError{}

	fmt.Println(err)

	//sub 6

	fmt.Println(calculateCircleArea(-1))

	// sub 7

	u := []string{"asd", "dsa"}

	fmt.Println(findUser(u, "adsa"))

}
