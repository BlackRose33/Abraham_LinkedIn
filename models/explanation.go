package models

import (
	"database/sql"
	"strings"
)

// ExplanationCount represents the amount of times an explanation was used
// for an expenditure
type ExplanationCount struct {
	Explanation string
	Count       int
	Amount      float64
}

// GetExpensiveExplanations asdfasdfajodsin
func GetExpensiveExplanations(db *sql.DB) ([]ExplanationCount, error) {
	rows, err := db.Query("SELECT explanation, COUNT(1), SUM(exp_amount) amt " +
		"FROM expenditures WHERE sched <> 'L' GROUP BY explanation HAVING " +
		"COUNT(1) > 2 ORDER BY amt DESC LIMIT 40")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	counts := make([]ExplanationCount, 0, 20)
	for rows.Next() {
		var count ExplanationCount
		err = rows.Scan(&count.Explanation, &count.Count, &count.Amount)
		if err == nil {
			count.Explanation = strings.Title(strings.ToLower(count.Explanation))
			counts = append(counts, count)
		}
	}

	return counts, nil
}

// GetExplanationCounts gets the frequency of each explanation
func GetExplanationCounts(db *sql.DB) ([]ExplanationCount, error) {
	rows, err := db.Query("SELECT explanation, COUNT(1), SUM(exp_amount) FROM " +
		"expenditures WHERE sched <> 'L' GROUP BY explanation HAVING " +
		"COUNT(1) > 2 ORDER BY COUNT(1) DESC LIMIT 40")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	counts := make([]ExplanationCount, 0, 20)
	for rows.Next() {
		var count ExplanationCount
		err = rows.Scan(&count.Explanation, &count.Count, &count.Amount)
		if err == nil {
			count.Explanation = strings.Title(strings.ToLower(count.Explanation))
			counts = append(counts, count)
		}
	}

	return counts, nil
}

// GetPurposeCounts gets the purposes in order of most common and their counts
func GetPurposeCounts(db *sql.DB) ([]ExplanationCount, error) {
	rows, err := db.Query("select purpose, count(purpose), sum(exp_amount) " +
		"from expenditures " +
		"group by purpose " +
		"order by count(purpose) desc")

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	counts := make([]ExplanationCount, 0, 20)
	for rows.Next() {
		var count ExplanationCount
		err = rows.Scan(&count.Explanation, &count.Count, &count.Amount)
		if err == nil {
			counts = append(counts, count)
		}
	}

	return counts, nil
}
