package models

import "database/sql"

// Expenditure struct with sched, c_code, count
type Expenditure struct {
  Sched string
  C_code string
  Count int
  Year int
}

// CountSchedulesExp counts the number of schedules
func CountSchedulesExp(db *sql.DB) ([]Expenditure, error) {
  rows, err := db.Query("select sched, count(sched) " +
    "from expenditures " +
    "group by sched " +
    "order by count(sched) desc")

  if err != nil {
		return nil, err
	}

	defer rows.Close()

  amounts := make([]Expenditure, 0, 6)
	for rows.Next() {
		var amount Expenditure
		err = rows.Scan(&amount.Sched, &amount.Count)
		if err == nil {
			amounts = append(amounts, amount)
		}
	}

	return amounts, nil

}

// CountCCodes counts the number of c_codes
func CountCCodes(db *sql.DB) ([]Expenditure, error) {
  rows, err := db.Query("select c_code, count(c_code) " +
    "from expenditures " +
    "group by c_code " +
    "order by count(c_code) desc")

  if err != nil {
		return nil, err
	}

	defer rows.Close()

  amounts := make([]Expenditure, 0, 6)
	for rows.Next() {
		var amount Expenditure
		err = rows.Scan(&amount.C_code, &amount.Count)
		if err == nil {
			amounts = append(amounts, amount)
		}
	}

	return amounts, nil
}

// NumExpendituresEachYear get number of expenditures made each year
func NumExpendituresEachYear(db *sql.DB) ([]Expenditure, error) {
  rows, err := db.Query("select election_year, count(election_year) " +
    " from expenditures " +
    " group by election_year " +
    " order by count(election_year) asc;")

  if err != nil {
		return nil, err
	}

	defer rows.Close()

  amounts := make([]Expenditure, 0, 5)
	for rows.Next() {
		var amount Expenditure
		err = rows.Scan(&amount.Year, &amount.Count)
		if err == nil {
			amounts = append(amounts, amount)
		}
	}

	return amounts, nil
}
