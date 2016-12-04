package models

import (
	"database/sql"
	"fmt"
)

// Contribution struct with year, amount, first name, last name, occupation
type Contribution struct {
	Year       int
	Amount     float64
	FirstName  string
	LastName   string
	Occupation string
	Count      int
	Sched      string
}

// CountOccupation counts the number of each occupation
func CountOccupation(db *sql.DB) ([]Contribution, error) {
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

// CountSchedulesCon counts the number of schedules of each (ex. ABC, M, D)
func CountSchedulesCon(db *sql.DB) ([]Contribution, error) {
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
func NumContributionsEachYear(db *sql.DB) ([]Contribution, error) {
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

// Contributor contains information about the amount a specific contributor
// gave to a candidate
type Contributor struct {
	Name   string
	Amount float64
	Year   int
}

// GetTopTenContributorsForCandidate gets a list of the top contributors
// for a candidate in any candidacy
func GetTopTenContributorsForCandidate(db *sql.DB, candID string) ([]Contributor,
	error) {

	rows, err := db.Query("SELECT con_name, sum(con_amount) total_amt, "+
		"election_year FROM contributes WHERE cand_id = ? GROUP BY con_name, "+
		"election_year ORDER BY total_amt DESC LIMIT 10", candID)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	defer rows.Close()

	contributors := make([]Contributor, 0, 10)
	for rows.Next() {
		var contributor Contributor
		err = rows.Scan(&contributor.Name, &contributor.Amount, &contributor.Year)
		if err == nil {
			contributors = append(contributors, contributor)
		}
	}

	return contributors, nil
}

// Expense contains information about where a candidate's money went most
// frequently in a campaign.
type Expense struct {
	Name   string
	Amount float64
	Year   int
}

// GetTopTenExpensesForCandidate gets a list of the top expenses
// for a candidate in any candidacy
func GetTopTenExpensesForCandidate(db *sql.DB, candID string) ([]Expense,
	error) {

	rows, err := db.Query("SELECT exp_name, sum(exp_amount) total_amt, "+
		"election_year FROM expenditures WHERE cand_id = ? GROUP BY exp_name, "+
		"election_year ORDER BY total_amt DESC LIMIT 10", candID)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	defer rows.Close()

	expenditures := make([]Expense, 0, 10)
	for rows.Next() {
		var expenditure Expense
		err = rows.Scan(&expenditure.Name, &expenditure.Amount, &expenditure.Year)
		if err == nil {
			expenditures = append(expenditures, expenditure)
		}
	}

	return expenditures, nil
}

// Reason holds info about an expenditure reason and its frequency
type Reason struct {
	Str       string
	Frequency int
}

// GetTopReasonsForCandidate gets top explanations for expenditures for a
// campaign
func GetTopReasonsForCandidate(db *sql.DB, candID string) ([]Reason, error) {
	rows, err := db.Query("SELECT explanation, count(1) cnt FROM expenditures "+
		"WHERE cand_id = ? GROUP BY explanation ORDER BY cnt DESC LIMIT 10", candID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	reasons := make([]Reason, 0, 10)
	for rows.Next() {
		var reason Reason
		err = rows.Scan(&reason.Str, &reason.Frequency)
		if err == nil {
			reasons = append(reasons, reason)
		}
	}

	return reasons, nil
}

// DoorPrize has information about where a candidate has received door prizes
type DoorPrize struct {
	From   string
	Amount float64
}

// GetTopDoorPrizesGiven ahnreaiweorjno
func GetTopDoorPrizesGiven(db *sql.DB, candID string) ([]DoorPrize, error) {
	rows, err := db.Query("SELECT con_name, sum(con_amount) FROM contributes "+
		"WHERE sched = 'd' AND cand_id = ? GROUP BY con_name ORDER BY sum(con_amount) "+
		"DESC LIMIT 10", candID)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	defer rows.Close()

	prizes := make([]DoorPrize, 0, 10)
	for rows.Next() {
		var prize DoorPrize
		err = rows.Scan(&prize.From, &prize.Amount)
		if err == nil {
			prizes = append(prizes, prize)
		}
	}

	return prizes, nil
}

// GetTopDoorPrizesPurchased ahnreaiweorjno
func GetTopDoorPrizesPurchased(db *sql.DB, candID string) ([]DoorPrize, error) {
	rows, err := db.Query("SELECT exp_name, sum(exp_amount) FROM expenditures "+
		"WHERE sched = 'd' AND cand_id = ? GROUP BY exp_name ORDER BY sum(exp_amount) "+
		"DESC LIMIT 10", candID)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	defer rows.Close()

	prizes := make([]DoorPrize, 0, 10)
	for rows.Next() {
		var prize DoorPrize
		err = rows.Scan(&prize.From, &prize.Amount)
		if err == nil {
			prizes = append(prizes, prize)
		}
	}

	return prizes, nil
}
