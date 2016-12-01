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

// HighestMatchAmount find highest match amount in table
func HighestMatchAmount(db *sql.DB) (*Contribution, error) {
	rows, err := db.Query("select match_amt " +
		"from contributes " +
		"group by match_amt " +
		"order by match_amt desc " +
		"limit 1")

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	if rows.Next() {
		var contributes Contribution
		err = rows.Scan(&contributes.Amount)
		if err == nil {
			return &contributes, nil
		}
	}
	return nil, nil
}

// CountSchedules counts the number of schedules of each (ex. ABC, M, D)
func CountSchedules(db *sql.DB) (*Contribution, error) {
  rows, err := db.Query("select sched, count(sched) " +
    "from contributes " +
    "group by sched " +
    "order by count(sched) desc")

  if err != nil {
		return nil, err
	}

	defer rows.Close()

  amounts := make([]Contribution, 0, 6)
	for rows.Next() {
		var amount Contribution
		err = rows.Scan(&amount.Sched, &amount.Count)
		if err == nil {
			amounts = append(amounts, amount)
		}
	}

	return amounts, nil

}

// NumContributionsEachYear get number of contributions each year
func NumContributionsEachYear(db *sql.DB) (*Contribution, error) {
  rows, err := db.Query("select election_year, count(election_year) " +
    " from contributes " +
    " group by election_year " +
    " order by count(election_year) asc;")

  if err != nil {
		return nil, err
	}

	defer rows.Close()

  amounts := make([]Contribution, 0, 5)
	for rows.Next() {
		var amount Contribution
		err = rows.Scan(&amount.Year, &amount.Count)
		if err == nil {
			amounts = append(amounts, amount)
		}
	}

	return amounts, nil
}
