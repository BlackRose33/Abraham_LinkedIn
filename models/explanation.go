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
}

// GetExplanationCounts gets the frequency of each explanation
func GetExplanationCounts(db *sql.DB) ([]ExplanationCount, error) {
	rows, err := db.Query("SELECT explanation, COUNT(1) FROM expenditures " +
		"GROUP BY explanation HAVING COUNT(1) > 2 ORDER BY COUNT(1) DESC LIMIT 40")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	counts := make([]ExplanationCount, 0, 20)
	for rows.Next() {
		var count ExplanationCount
		err = rows.Scan(&count.Explanation, &count.Count)
		if err == nil {
			count.Explanation = strings.Title(strings.ToLower(count.Explanation))
			counts = append(counts, count)
		}
	}

	return counts, nil
}

// GetPurposeCounts gets the purposes in order of most common and their counts
func GetPurposeCounts(db *sql.DB) ([]ExplanationCount, error) {
	rows, err := db.Query("select purpose, count(purpose) " +
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
		err = rows.Scan(&count.Explanation, &count.Count)
		if err == nil {
			counts = append(counts, count)
		}
	}

	return counts, nil
}
