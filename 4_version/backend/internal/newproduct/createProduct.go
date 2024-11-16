package newproduct

import (
	"encoding/json"
	"errors"
	entity "helloapp"
	"log"
	"unicode"
)

func MakeProductPattern(data_entry []byte) entity.NewProduct {
	product := entity.NewProduct{}

	err := json.Unmarshal(data_entry, &product)
	if err != nil {
		log.Fatal("Failed unmarshal")
	}

	if product.Product_param == nil {
		product.Product_param = make(map[string]string)
	}

	return product
}

// . - means int; * - means Letter
func checkMask(val string, mask string) error {
	if len(val) != len(mask) {
		return errors.New("not a pattern")
	}

	for i := 0; i < len(mask); i++ {
		switch mask[i] {
		case '.':
			// Проверяем, что символ в input - цифра
			if !unicode.IsDigit(rune(val[i])) {
				return errors.New("not a pattern")
			}
		case '*':
			// Проверяем, что символ в input - буква
			if !unicode.IsLetter(rune(val[i])) {
				return errors.New("not a pattern")
			}
		default:
			// Проверяем, что символы совпадают
			if val[i] != mask[i] {
				return errors.New("not a pattern")
			}
		}
	}

	return nil
}

func CheckProductPattern(pattern map[string]string, product map[string]string) error {
	for param, mask := range pattern {
		val, ok := product[param]
		if !ok {
			return errors.New("not a pattern")
		}
		err := checkMask(val, mask)
		if err != nil {
			return err
		}
	}

	return nil
}