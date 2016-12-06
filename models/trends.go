package models

import (
	"database/sql"
	"strings"
)

//GetMostPaid asdfjnaoifjnoij
func GetMostPaid(db *sql.DB) (*CandidateAmount, error) {
	rows, err := db.Query("SELECT c2.cand_first, c2.cand_last, c1.election_year," +
		" c1.cand_id, c1.office_code, SUM(con_amount) amt FROM contributes c," +
		" candidacy c1, candidate c2 WHERE c.cand_id = c2.cand_id AND" +
		" c.cand_id = c1.cand_id AND c.election_year = c1.election_year AND" +
		" c.sched <> 'M' GROUP BY c.cand_id, c.election_year ORDER BY amt DESC" +
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
		" c1.sched <> 'M' GROUP BY c2.con_name, c1.election_year ORDER BY amt" +
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
		"FROM contributes WHERE sched <> 'M' GROUP BY election_year " +
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

// Refund contains information about a refund given by a candidate
type Refund struct {
	CanFName    string
	CanLName    string
	ContribName string
	Amount      float64
	Year        int
	DaysPassed  int
}

// GetBiggestRefunds gets information about the refunds with the largest
// monetary value from the database
func GetBiggestRefunds(db *sql.DB) ([]Refund, error) {
	rows, err := db.Query("SELECT c1.cand_first, c1.cand_last, c2.con_name, " +
		"ABS(c2.con_amount) amt, DATEDIFF(" +
		"STR_TO_DATE(c2.refund_date, '%m/%d/%Y %k:%i:%s'), " +
		"STR_TO_DATE(c2.con_date, '%m/%d/%Y %k:%i:%s')) time_gap, " +
		"c2.election_year FROM candidate c1, contributes c2 WHERE " +
		"c1.cand_id = c2.cand_id AND sched = 'M' ORDER BY amt DESC LIMIT 10")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	refunds := make([]Refund, 0, 10)
	for rows.Next() {
		var refund Refund
		err = rows.Scan(&refund.CanFName, &refund.CanLName, &refund.ContribName,
			&refund.Amount, &refund.DaysPassed, &refund.Year)
		if err == nil {
			refunds = append(refunds, refund)
		}
	}
	return refunds, nil
}

// GetMostDelayedRefunds gets information about the refunds with the largest
// time gap from contribution to refund from the database
func GetMostDelayedRefunds(db *sql.DB) ([]Refund, error) {
	rows, err := db.Query("SELECT c1.cand_first, c1.cand_last, c2.con_name, " +
		"ABS(c2.con_amount) amt, DATEDIFF(" +
		"STR_TO_DATE(c2.refund_date, '%m/%d/%Y %k:%i:%s'), " +
		"STR_TO_DATE(c2.con_date, '%m/%d/%Y %k:%i:%s')) time_gap, " +
		"c2.election_year FROM candidate c1, contributes c2 WHERE " +
		"c1.cand_id = c2.cand_id AND sched = 'M' ORDER BY time_gap DESC LIMIT 10")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	refunds := make([]Refund, 0, 10)
	for rows.Next() {
		var refund Refund
		err = rows.Scan(&refund.CanFName, &refund.CanLName, &refund.ContribName,
			&refund.Amount, &refund.DaysPassed, &refund.Year)
		if err == nil {
			refunds = append(refunds, refund)
		}
	}
	return refunds, nil
}

// AmountAndCount is used for CFP data
type AmountAndCount struct {
	Amount float64
	Count  int
	Year   int
}

// GetMatchAmountAndParticipantCount gets the number of participants and
// amount matched from CFP program
func GetMatchAmountAndParticipantCount(db *sql.DB) ([]AmountAndCount, error) {
	rows, err := db.Query(`SELECT q1.match_amount, q2.num_particip,
		q1.election_year FROM (SELECT election_year, SUM(match_amt) match_amount
		FROM contributes WHERE match_amt > 0 GROUP BY election_year) q1,
		(SELECT c1.election_year, COUNT(1) num_particip FROM candidacy c1,
		candidate c2 WHERE c1.cand_id = c2.cand_id AND c2.cand_class = 'P'
		GROUP BY c1.election_year) q2 WHERE q1.election_year = q2.election_year`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	res := make([]AmountAndCount, 0, 5)
	for rows.Next() {
		var item AmountAndCount
		err = rows.Scan(&item.Amount, &item.Count, &item.Year)
		if err == nil {
			res = append(res, item)
		}
	}

	return res, nil
}

// GetAvgMatchAmountAndParticipantCount gets the number of participants and
// amount matched from CFP program
func GetAvgMatchAmountAndParticipantCount(db *sql.DB) ([]AmountAndCount, error) {
	rows, err := db.Query(`SELECT q1.match_amount, q2.num_particip,
		q1.election_year FROM (SELECT election_year, AVG(match_amt) match_amount
		FROM contributes WHERE match_amt > 0 GROUP BY election_year) q1,
		(SELECT c1.election_year, COUNT(1) num_particip FROM candidacy c1,
		candidate c2 WHERE c1.cand_id = c2.cand_id AND c2.cand_class = 'P'
		GROUP BY c1.election_year) q2 WHERE q1.election_year = q2.election_year`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	res := make([]AmountAndCount, 0, 5)
	for rows.Next() {
		var item AmountAndCount
		err = rows.Scan(&item.Amount, &item.Count, &item.Year)
		if err == nil {
			res = append(res, item)
		}
	}

	return res, nil
}

// StringCount has a string and a count
type StringCount struct {
	Str    string
	Count  int
	Amount float64
}

// GetMostCommonOccupations adsfiasf
func GetMostCommonOccupations(db *sql.DB) ([]StringCount, error) {
	rows, err := db.Query(`SELECT occupation, COUNT(occupation), AVG(con_amount)
		FROM (SELECT occupation, con_amount FROM contributes GROUP BY con_name)
		occupations WHERE occupation <> ''
    GROUP BY occupation
    ORDER BY COUNT(occupation) DESC LIMIT 10`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	res := make([]StringCount, 0, 10)
	for rows.Next() {
		var item StringCount
		err = rows.Scan(&item.Str, &item.Count, &item.Amount)
		if err == nil {
			item.Str = strings.Title(strings.ToLower(item.Str))
			res = append(res, item)
		}
	}

	return res, nil
}

// GetHighestContribOccupations adsfiasf
func GetHighestContribOccupations(db *sql.DB) ([]StringCount, error) {
	rows, err := db.Query(`SELECT occupation, COUNT(occupation), AVG(con_amount)
		FROM (SELECT occupation, con_amount FROM contributes GROUP BY con_name)
		occupations WHERE occupation <> ''
    GROUP BY occupation
		HAVING COUNT(occupation) > 10
    ORDER BY AVG(con_amount) DESC LIMIT 10`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	res := make([]StringCount, 0, 10)
	for rows.Next() {
		var item StringCount
		err = rows.Scan(&item.Str, &item.Count, &item.Amount)
		if err == nil {
			item.Str = strings.Title(strings.ToLower(item.Str))
			res = append(res, item)
		}
	}

	return res, nil
}
