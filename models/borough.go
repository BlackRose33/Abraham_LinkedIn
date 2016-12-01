package models

import "database/sql"

const (
	bronxCode        = "X"
	brooklynCode     = "K"
	manhattanCode    = "M"
	queensCode       = "Q"
	statenIslandCode = "S"
	unassignedCode   = ""
	multiBoroughCode = "Z"
)

// BoroughPercentages contains counts for multiple boroughs to help create
// a heatmap
type BoroughPercentages struct {
	Bronx             float64 `json:"bronx"`
	Brooklyn          float64 `json:"brooklyn"`
	Manhattan         float64 `json:"manhattan"`
	Queens            float64 `json:"queens"`
	StatenIsland      float64 `json:"staten_island"`
	UnassignedCount   float64 `json:"unassigned_count"`
	MultiBoroughCount float64 `json:"multi_borough_count"`
	IsAmountMode      bool    `json:"is_amount_mode"`
}

// GetBoroughContribAmountPercentagesForCandidate generates heatmap
// data for a candidate, but by *amount*, and not by number of contribs
func GetBoroughContribAmountPercentagesForCandidate(db *sql.DB,
	candidateID string) (*BoroughPercentages, error) {

	amounts := map[string]float64{
		bronxCode:        0,
		brooklynCode:     0,
		manhattanCode:    0,
		queensCode:       0,
		statenIslandCode: 0,
		unassignedCode:   0,
		multiBoroughCode: 0,
	}

	query := "SELECT borough, sum(con_amount) from contributor c1, contributes " +
		"c2 where c1.con_name = c2.con_name "
	args := []interface{}{}
	if len(candidateID) > 0 {
		query += "and c2.cand_id = ? "
		args = []interface{}{candidateID}
	}
	query += "group by borough"

	rows, err := db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var boroughCode string
		var amount float64
		err = rows.Scan(&boroughCode, &amount)

		if err != nil {
			continue
		}

		if _, ok := amounts[boroughCode]; ok {
			amounts[boroughCode] = amount
		}
	}

	max := float64(0.0001)
	for _, amt := range amounts {
		if amt > max {
			max = amt
		}
	}

	return &BoroughPercentages{
		Bronx:             .2 + amounts[bronxCode]/max*.8,
		Brooklyn:          .2 + amounts[brooklynCode]/max*.8,
		Manhattan:         .2 + amounts[manhattanCode]/max*.8,
		Queens:            .2 + amounts[queensCode]/max*.8,
		StatenIsland:      .2 + amounts[statenIslandCode]/max*.8,
		UnassignedCount:   amounts[unassignedCode],
		MultiBoroughCount: amounts[multiBoroughCode],
		IsAmountMode:      true,
	}, nil
}

// GetBoroughContribPercentagesForCandidate attempts to build heat-map
// data for contributions for a candidate
func GetBoroughContribPercentagesForCandidate(db *sql.DB, candidateID string) (
	*BoroughPercentages, error) {

	counts := map[string]float64{
		bronxCode:        0,
		brooklynCode:     0,
		manhattanCode:    0,
		queensCode:       0,
		statenIslandCode: 0,
		unassignedCode:   0,
		multiBoroughCode: 0,
	}

	query := "SELECT borough, count(*) from contributor c1"
	args := []interface{}{}
	if len(candidateID) > 0 {
		query += ", contributes c2 where c1.con_name = c2.con_name and " +
			"c2.cand_id = ?"
		args = []interface{}{candidateID}
	}
	query += " group by borough"

	rows, err := db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var boroughCode string
		var count int
		err = rows.Scan(&boroughCode, &count)
		if _, ok := counts[boroughCode]; ok {
			counts[boroughCode] = float64(count)
		}
	}

	max := float64(0.0001)
	for _, amt := range counts {
		if amt > max {
			max = amt
		}
	}

	return &BoroughPercentages{
		Bronx:             .2 + counts[bronxCode]/max*.8,
		Brooklyn:          .2 + counts[brooklynCode]/max*.8,
		Manhattan:         .2 + counts[manhattanCode]/max*.8,
		Queens:            .2 + counts[queensCode]/max*.8,
		StatenIsland:      .2 + counts[statenIslandCode]/max*.8,
		UnassignedCount:   counts[unassignedCode],
		MultiBoroughCount: counts[multiBoroughCode],
		IsAmountMode:      false,
	}, nil
}
