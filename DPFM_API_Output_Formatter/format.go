package dpfm_api_output_formatter

import (
	"database/sql"
	"fmt"
)

func ConvertToCountry(rows *sql.Rows) (*Country, error) {
	defer rows.Close()
	country := Country{}
	i := 0

	for rows.Next() {
		i++
		err := rows.Scan(
			&country.country,
			&country.IsMarkedForDeletion,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return &country, err
		}

	}
	if i == 0 {
		fmt.Printf("DBに対象のレコードが存在しません。")
		return &country, nil
	}

	return &country, nil
}
