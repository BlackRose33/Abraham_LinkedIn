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
	ID        string
	FirstName string
	LastName  string
	Year      int
	Amount    float64
	EntityName   string
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
