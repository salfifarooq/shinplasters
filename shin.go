package shinplasters

import (
	"encoding/json"
	"errors"
	"strings"
)

/**
 * Created by  ♅ Salfi Farooq on 28/09/18.
 */

func GetAll() (map[string]int32, error) {

	subunits, err := parse()

	if err != nil {
		return nil, err
	}

	return subunits, err

}

func GetSubunit(currencyCode string )(int32,error){

	subunits, err := parse()

	if err != nil {
		return 0, err
	}

	if subunits[strings.ToUpper(currencyCode)] == 0 {
		return 0, errors.New(currencyCode + " Does not exist.")
	}

	return subunits[strings.ToUpper(currencyCode)] , err

}


func parse() (map[string]int32, error) {

	var subunit map[string]int32

	err := json.Unmarshal([]byte(Data), &subunit)
	if err != nil {
		return nil, err
	}

	return subunit, nil
}
