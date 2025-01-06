package utils

import (
	"net/url"
)

func UnescapeQueryValues(valuesMap url.Values) error {
	for _, ls := range valuesMap {
		for i, _ := range ls {
			v, err := url.QueryUnescape(ls[i])
			if err != nil {
				return err
			}
			ls[i] = v
		}
	}
	return nil
}
