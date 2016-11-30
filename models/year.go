package models

import "database/sql"

// YearAmount keeps track of the amount of expenditures/contributions in
// a year
type YearAmount struct {
	Year   int
	Amount float64
}

// GetYearExpenditures returns the amount of expenditures in each year
func GetYearExpenditures(db *sql.DB) ([]YearAmount, error) {
	var expenditures []YearAmount

	rows, err := db.Query("select election_year, avg(exp_amount) " +
		"from expenditures group by election_year order by election_year desc")

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var expenditure YearAmount
		err = rows.Scan(&expenditure.Year, &expenditure.Amount)
		if err == nil {
			expenditures = append(expenditures, expenditure)
		}
	}

	return expenditures, nil
}

// GetYearContributions returns the amount of contributions in each year
func GetYearContributions(db *sql.DB) ([]YearAmount, error) {
	var contributions []YearAmount

	rows, err := db.Query("select election_year, avg(con_amount) " +
		"from contributes group by election_year order by election_year desc")

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var contribution YearAmount
		err = rows.Scan(&contribution.Year, &contribution.Amount)
		if err == nil {
			contributions = append(contributions, contribution)
		}
	}

	return contributions, nil
}
