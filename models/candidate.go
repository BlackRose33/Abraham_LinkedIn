package models

import "database/sql"

// CandInfo werhiqweriun
type CandInfo struct {
	Candidate Candidate
	Position  []Positions
}

// Positions aweiurnhalewj r
type Positions struct {
	OfficeCd int
	Year     int
}

// Candidate encompasses fields that describe political candidates
type Candidate struct {
	ID        string
	FirstName string
	LastName  string
	CandClass string
	Count     int
}

// GetCandidateByID attempts to find a candidate in the database based on an
// ID
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
		if err == nil {
			return &candidate, nil
		}
	}

	return nil, nil
}

// GetCandidateByIDWithInfo attempts to find a candidate in the database based on
// an ID
func GetCandidateByIDWithInfo(db *sql.DB, id string) (*CandInfo, error) {
	if id == "" {
		var candidate CandInfo
		candidate.Candidate.ID = " "
		candidate.Candidate.FirstName = " "
		candidate.Candidate.LastName = " "
		candidate.Position = []Positions{}
		return &candidate, nil
	}

	rows, err := db.Query("SELECT cand_id, cand_first, cand_last FROM "+
		"candidate WHERE cand_id = ?", id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	if rows.Next() {
		var candidate CandInfo
		err = rows.Scan(&candidate.Candidate.ID, &candidate.Candidate.FirstName, &candidate.Candidate.LastName)
		if err != nil {
			return nil, err
		}

		rows2, err := db.Query("SELECT office_code, election_year FROM candidacy WHERE cand_id = ?", id)
		if err != nil {
			return nil, err
		}
		defer rows2.Close()

		positions := make([]Positions, 0, 4)
		for rows2.Next() {
			var pos Positions
			err = rows2.Scan(&pos.OfficeCd, &pos.Year)
			if err == nil {
				positions = append(positions, pos)
			}
		}
		candidate.Position = positions

		return &candidate, nil
	}
	return nil, nil
}

// GetCandidates gets all candidates in the database
func GetCandidates(db *sql.DB) ([]Candidate, error) {
	rows, err := db.Query("SELECT cand_id, cand_first, cand_last FROM " +
		"candidate ORDER BY cand_last, cand_first ASC")
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
		err = rows.Scan(&contribution.ID, &contribution.EntityName,
			&contribution.Amount)
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

// GetContributeDates gets the dates of this candidate's contributions
func GetContributeDates(db *sql.DB, candID string) ([]CandidateAmount, error) {
	rows, err := db.Query("select cand_id, str_to_date(con_date, '%m/%d/%Y %k:%i:%s') as dt " +
		" from contributes " +
		" where cand_id = ? " +
		" group by dt " +
		" order by dt asc")

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	amounts := make([]CandidateAmount, 0, 110)
	for rows.Next() {
		var amount CandidateAmount
		err = rows.Scan(&amount.ID, &amount.EntityName)
		if err == nil {
			amounts = append(amounts, amount)
		}
	}

	return amounts, nil
}

// GetMostCandClass finds the most common candidate class for candidates
func GetMostCandClass(db *sql.DB) ([]Candidate, error) {
	rows, err := db.Query("select cand_class, count(cand_class) " +
		"from candidate " +
		"group by cand_class " +
		"order by count(cand_class) desc")

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	amounts := make([]Candidate, 0, 3)
	for rows.Next() {
		var amount Candidate
		err = rows.Scan(&amount.CandClass, &amount.Count)
		if err == nil {
			amounts = append(amounts, amount)
		}
	}

	return amounts, nil
}

// CandidateAmountSummary describes the contribution / expenditure distribution
// for a candidate per month for all campaigns, used for generating a graph
type CandidateAmountSummary struct {
	Year            int
	Month           int
	IsContributions bool
	Amount          float64
}

// GetSummariesForCandidate gets an aggregate summary for a candidate per month
// for expenditures and contributions
func GetSummariesForCandidate(db *sql.DB, candID string) (
	[]CandidateAmountSummary, error) {

	summaries, err := GetContributionSummaryForCandidate(db, candID)
	if err != nil {
		return nil, err
	}

	summaries2, err := GetExpenditureSummaryForCandidate(db, candID)
	if err != nil {
		return nil, err
	}

	for _, summary := range summaries2 {
		summaries = append(summaries, summary)
	}

	return summaries, nil
}

// GetContributionSummaryForCandidate gets a list of contribution totals per
// month for a candidate
func GetContributionSummaryForCandidate(db *sql.DB, candID string) (
	[]CandidateAmountSummary, error) {

	rows, err := db.Query("SELECT election_year, month(STR_TO_DATE(con_date, "+
		"'%m/%d/%Y %k:%i:%s')) mnth, sum(con_amount) FROM contributes WHERE "+
		"cand_id = ? GROUP BY mnth, election_year", candID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	summaries := make([]CandidateAmountSummary, 0, 5)
	for rows.Next() {
		summary := CandidateAmountSummary{IsContributions: true}
		err = rows.Scan(&summary.Year, &summary.Month, &summary.Amount)
		if err == nil {
			summaries = append(summaries, summary)
		}
	}

	return summaries, nil
}

// GetExpenditureSummaryForCandidate gets a list of expenditure totals per
// month for a candidate
func GetExpenditureSummaryForCandidate(db *sql.DB, candID string) (
	[]CandidateAmountSummary, error) {

	rows, err := db.Query("SELECT election_year, month(STR_TO_DATE(exp_date, "+
		"'%m/%d/%Y %k:%i:%s')) mnth, sum(exp_amount) FROM expenditures WHERE "+
		"cand_id = ? GROUP BY mnth, election_year", candID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	summaries := make([]CandidateAmountSummary, 0, 5)
	for rows.Next() {
		summary := CandidateAmountSummary{IsContributions: false}
		err = rows.Scan(&summary.Year, &summary.Month, &summary.Amount)
		if err == nil {
			summaries = append(summaries, summary)
		}
	}

	return summaries, nil
}

// Candidacy describes a candidate's run for elected office
type Candidacy struct {
	Year       int
	OfficeCode int
}

// GetCandidacyHistory gets a list of candidacy records for a candidate
func GetCandidacyHistory(db *sql.DB, candID string) ([]Candidacy, error) {
	rows, err := db.Query("SELECT election_year, office_code FROM candidacy "+
		"WHERE cand_id = ?", candID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	history := make([]Candidacy, 0, 10)
	for rows.Next() {
		var entry Candidacy
		err = rows.Scan(&entry.Year, &entry.OfficeCode)
		if err == nil {
			history = append(history, entry)
		}
	}

	return history, nil
}
