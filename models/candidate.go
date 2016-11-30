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

// CandidateSpender holds information about expenditures from a candidate
type CandidateSpender struct {
	ID        string
	FirstName string
	LastName  string
	Year      int
	Amount    float64
	ExpName   string
}

// GetBiggestSpenderOfAnyYear gets the... biggest spender of any year?
func GetBiggestSpenderOfAnyYear(db *sql.DB) (*CandidateSpender, error) {
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
		var candidate CandidateSpender
		err = rows.Scan(&candidate.ExpName, &candidate.Amount, &candidate.Year,
			&candidate.ID, &candidate.FirstName, &candidate.LastName)
		if err == nil {
			return &candidate, nil
		}
	}
	return nil, nil
}
