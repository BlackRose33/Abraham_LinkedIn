package models

import "database/sql"

// Candidate encompasses fields that describe political candidates
type Candidate struct {
	ID        string
	FirstName string
	LastName  string
}

// GetCandidateByID attempts to find a candidate in the database based on
// an ID
func GetCandidateByID(db *sql.DB, id string) (*Candidate, error) {
	rows, err := db.Query("SELECT cand_id, cand_first, cand_last FROM "+
		"candidate WHERE cand_id = ?", id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	if rows.Next() {
		var candidate Candidate
		err = rows.Scan(&candidate.ID, &candidate.FirstName, &candidate.LastName)
		if err != nil {
			return nil, err
		}
		return &candidate, nil
	}
	return nil, nil
}

// GetCandidates gets all candidates in the database
func GetCandidates(db *sql.DB) ([]Candidate, error) {
	rows, err := db.Query("SELECT cand_id, cand_first, cand_last FROM " +
		"candidate")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	candidates := make([]Candidate, 0, 100)
	for rows.Next() {
		var candidate Candidate
		err = rows.Scan(&candidate.ID, &candidate.FirstName, &candidate.LastName)
		if err == nil {
			candidates = append(candidates, candidate)
		}
	}

	return candidates, nil
}

// CandidateAmount holds information about expenditures from a candidate
type CandidateAmount struct {
	ID         string
	FirstName  string
	LastName   string
	Year       int
	Amount     float64
	EntityName string
	OfficeCd   int
}

// GetBiggestSpenderOfAnyYear gets the... biggest spender of any year?
func GetBiggestSpenderOfAnyYear(db *sql.DB) (*CandidateAmount, error) {
	rows, err := db.Query("select e.exp_name, e.exp_amount, e.election_year, e.cand_id, c.cand_first, c.cand_last " +
		"from expenditures e join candidate c on e.cand_id = c.cand_id " +
		"group by exp_amount " +
		"order by exp_amount desc " +
		"limit 1")

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	if rows.Next() {
		var candidate CandidateAmount
		err = rows.Scan(&candidate.EntityName, &candidate.Amount, &candidate.Year,
			&candidate.ID, &candidate.FirstName, &candidate.LastName)
		if err == nil {
			return &candidate, nil
		}
	}
	return nil, nil
}

// GetMostPaidAnyYear returns the candidate who was paid the most in any year
func GetMostPaidAnyYear(db *sql.DB) (*CandidateAmount, error) {
	rows, err := db.Query("select con_name, con_amount, election_year " +
		" from contributes " +
		" group by con_amount " +
		" order by con_amount desc " +
		" limit 1")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if rows.Next() {
		var contributes CandidateAmount
		err = rows.Scan(&contributes.EntityName, &contributes.Amount, &contributes.Year)
		if err == nil {
			return &contributes, nil
		}
	}
	return nil, nil
}

// GetCandidateContributionsForEachYear attempts to find the amount of
// total contributions a candidate received in any year
func GetCandidateContributionsForEachYear(db *sql.DB, candID string) (
	[]CandidateAmount, error) {

	rows, err := db.Query("SELECT c1.con_name, c1.con_amount, c1.election_year, "+
		"c1.cand_id, c2.cand_first, c2.cand_last FROM contributes c1 JOIN "+
		"candidate c2 ON c1.cand_id = c2.cand_id WHERE c1.cand_id = ? "+
		"ORDER BY con_amount DESC", candID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	amounts := make([]CandidateAmount, 0, 5)
	for rows.Next() {
		var amount CandidateAmount
		err = rows.Scan(&amount.EntityName, &amount.Amount, &amount.Year,
			&amount.ID, &amount.FirstName, &amount.LastName)
		if err == nil {
			amounts = append(amounts, amount)
		}
	}

	return amounts, nil
}

// GetTotalAmountThisCandidateContributed get total amount that this candidate (ex. 605) contributed (ex. $999999)
func GetTotalAmountThisCandidateContributed(db *sql.DB, candID string) (*CandidateAmount, error) {
	rows, err := db.Query("select cand_id, sum(con_amount) "+
		"from contributes "+
		"where cand_id = ?", candID)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	if rows.Next() {
		var contributes CandidateAmount
		err = rows.Scan(&contributes.EntityName, &contributes.Amount)
		if err == nil {
			return &contributes, nil
		}
	}
	return nil, nil
}

// GetTotalAmountThisCandidateReceived get total amount this candidate (ex. 605) received
func GetTotalAmountThisCandidateReceived(db *sql.DB, candID string) (*CandidateAmount, error) {
	rows, err := db.Query("select cand_id, exp_amount, sum(exp_amount) "+
		"from expenditures "+
		"where cand_id = ?", candID)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	if rows.Next() {
		var contribution CandidateAmount
		err = rows.Scan(&contribution.EntityName, &contribution.Amount)
		if err == nil {
			return &contribution, nil
		}
	}
	return nil, nil
}

// GetHighestSpenderByPositionForYear ecnj eroiubhcriuqbherioucuebwcriuqriu
func GetHighestSpenderByPositionForYear(db *sql.DB, year int) ([]CandidateAmount,
	error) {

	rows, err := db.Query("SELECT election_year, cand_id, office_code, amt, "+
		"cand_first, cand_last, exp_name FROM (SELECT c2.cand_first, "+
		"c2.cand_last, e.exp_name, c1.election_year, c1.cand_id, c1.office_code, "+
		"SUM(exp_amount) amt FROM expenditures e, candidacy c1, candidate c2 "+
		"WHERE e.cand_id = c2.cand_id AND e.cand_id = c1.cand_id AND "+
		"e.election_year = c1.election_year AND e.election_year = ? GROUP BY "+
		"c1.office_code, e.cand_id ORDER BY amt DESC) result GROUP BY office_code",
		year)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	amounts := make([]CandidateAmount, 0, 6)
	for rows.Next() {
		var amount CandidateAmount
		err = rows.Scan(&amount.Year, &amount.ID, &amount.OfficeCd, &amount.Amount,
			&amount.FirstName, &amount.LastName, &amount.EntityName)
		if err == nil {
			amounts = append(amounts, amount)
		}
	}

	return amounts, nil
}

// GetHighestPaidByPositionForYear hrinachunarih
func GetHighestPaidByPositionForYear(db *sql.DB, year int) ([]CandidateAmount,
	error) {

	rows, err := db.Query("SELECT election_year, cand_id, office_code, amt, "+
		"cand_first, cand_last, con_name FROM (SELECT c3.cand_first, "+
		"c3.cand_last, c1.con_name, c2.election_year, c2.cand_id, c2.office_code, "+
		"SUM(con_amount) amt FROM contributes c1, candidacy c2, candidate c3 "+
		"WHERE c1.cand_id = c3.cand_id AND c1.cand_id = c2.cand_id "+
		"AND c1.election_year = c2.election_year AND c1.election_year = ? "+
		"GROUP BY c2.office_code, c1.cand_id ORDER BY amt DESC) result "+
		"GROUP BY office_code", year)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	amounts := make([]CandidateAmount, 0, 6)
	for rows.Next() {
		var amount CandidateAmount
		err = rows.Scan(&amount.Year, &amount.ID, &amount.OfficeCd, &amount.Amount,
			&amount.FirstName, &amount.LastName, &amount.EntityName)
		if err == nil {
			amounts = append(amounts, amount)
		}
	}

	return amounts, nil
}

// GetTotalAmountThisCandidateContributedEachYear get total amount this candidate (ex. 605) contributed each year
func GetTotalAmountThisCandidateContributedEachYear(db *sql.DB, candID string) ([]CandidateAmount, error) {
	rows, err := db.Query("select cand_id, sum(con_amount), election_year "+
		"from contributes "+
		"where cand_id = ? "+
		"group by election_year "+
		"order by election_year desc", candID)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	amounts := make([]CandidateAmount, 0, 5)
	for rows.Next() {
		var amount CandidateAmount
		err = rows.Scan(&amount.ID, &amount.Amount, &amount.Year)
		if err == nil {
			amounts = append(amounts, amount)
		}
	}

	return amounts, nil
}
