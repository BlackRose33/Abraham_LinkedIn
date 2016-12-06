package models

import "database/sql"

//GetMostPaid asdfjnaoifjnoij
func GetMostPaid(db *sql.DB) (*CandidateAmount, error) {
	rows, err := db.Query("SELECT c2.cand_first, c2.cand_last, c1.election_year," +
		" c1.cand_id, c1.office_code, SUM(con_amount) amt FROM contributes c," +
		" candidacy c1, candidate c2 WHERE c.cand_id = c2.cand_id AND" +
		" c.cand_id = c1.cand_id AND c.election_year = c1.election_year AND" +
		" c.refund_date = \"\" GROUP BY c.cand_id, c.election_year ORDER BY amt DESC" +
		" limit 1")

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var amount CandidateAmount
		err = rows.Scan(&amount.FirstName, &amount.LastName, &amount.Year, &amount.ID,
			&amount.OfficeCd, &amount.Amount)
		if err == nil {
			return &amount, nil
		}
	}

	return nil, nil
}

// GetBiggestSpender asdfasdfoajidsfoij
func GetBiggestSpender(db *sql.DB) (*CandidateAmount, error) {
	rows, err := db.Query("SELECT c2.cand_first, c2.cand_last, c1.election_year," +
		" c1.cand_id, c1.office_code, SUM(exp_amount) amt FROM expenditures e," +
		" candidacy c1, candidate c2 WHERE e.cand_id = c2.cand_id AND" +
		" e.cand_id = c1.cand_id AND e.election_year = c1.election_year GROUP BY" +
		" e.cand_id, e.election_year ORDER BY amt DESC limit 1")

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var amount CandidateAmount
		err = rows.Scan(&amount.FirstName, &amount.LastName, &amount.Year, &amount.ID,
			&amount.OfficeCd, &amount.Amount)
		if err == nil {
			return &amount, nil
		}
	}

	return nil, nil
}

// GetBiggestContributor asdfjasdfojpoij
func GetBiggestContributor(db *sql.DB) (*Contributor, error) {
	rows, err := db.Query("SELECT c1.con_name, c1.election_year, SUM(c1.con_amount)" +
		" amt FROM contributes c1, contributor c2 WHERE c1.con_name = c2.con_name AND" +
		" c1.refund_date = \"\" GROUP BY c2.con_name, c1.election_year ORDER BY amt" +
		" DESC limit 1")

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var amount Contributor
		err = rows.Scan(&amount.Name, &amount.Year, &amount.Amount)
		if err == nil {
			return &amount, nil
		}
	}

	return nil, nil
}

//Trend asdfjanodsfij
type Trend struct {
	Year   int
	Amount float64
}

// GetExpenditureChange aidsfjnasdfoijn
func GetExpenditureChange(db *sql.DB) ([]Trend, error) {
	rows, err := db.Query("SELECT election_year, AVG(exp_amount) amt " +
		"FROM expenditures GROUP BY election_year ORDER BY election_year ASC;")

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	amounts := make([]Trend, 0, 5)
	for rows.Next() {
		var amount Trend
		err = rows.Scan(&amount.Year, &amount.Amount)
		if err == nil {
			amounts = append(amounts, amount)
		}
	}

	return amounts, nil
}

// GetContributionChange asdfadsfoij
func GetContributionChange(db *sql.DB) ([]Trend, error) {
	rows, err := db.Query("SELECT election_year, AVG(con_amount) amt " +
		"FROM contributes WHERE refund_date = \"\" GROUP BY election_year " +
		"ORDER BY election_year ASC")

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	amounts := make([]Trend, 0, 5)
	for rows.Next() {
		var amount Trend
		err = rows.Scan(&amount.Year, &amount.Amount)
		if err == nil {
			amounts = append(amounts, amount)
		}
	}

	return amounts, nil
}