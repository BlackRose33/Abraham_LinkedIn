package models

// Contribution struct with year, amount, first name, last name, occupation
type Contribution struct {
	Year      int
	Amount    float64
	FirstName string
	LastName  string
	Occupation string
	Count int
}

// CountOccupation counts the number of each occupation
func CountOccupation(db *sql.DB) (*Contribution, error) {
	rows, err := db.Query("select occupation, count(occupation) " +
		"from contributes " +
		"group by occupation " +
		"order by count(occupation) desc")

  if err != nil {
		return nil, err
	}

	defer rows.Close()

  amounts := make([]Contribution, 0, 25)
	for rows.Next() {
		var amount Contribution
		err = rows.Scan(&amount.Occupation, &amount.Count)
		if err == nil {
			amounts = append(amounts, amount)
		}
	}

	return amounts, nil
}
