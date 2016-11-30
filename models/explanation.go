package models

import "database/sql"

// ExplanationCount represents the amount of times an explanation was used
// for an expenditure
type ExplanationCount struct {
	Explanation string
	Count       int
}

// GetExplanationCounts gets the frequency of each explanation
func GetExplanationCounts(db *sql.DB) ([]ExplanationCount, error) {
	rows, err := db.Query("SELECT explanation, COUNT(explanation) " +
		"FROM expenditures GROUP BY explanation ORDER BY COUNT(explanation) DESC")
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
