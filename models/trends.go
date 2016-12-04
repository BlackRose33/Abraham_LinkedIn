package models

import "database/sql"

func GetMostPaid(db *sql.DB) (*Contributor, error) {
  return nil, nil
}

func GetBiggestSpender(db *sql.DB) (*Contributor, error) {
  /*rows, err := db.Query("select e.exp_name, e.exp_amount, e.election_year, " +
    "e.cand_id, c.cand_first, c.cand_last " +
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
  }*/
  return nil, nil
}

func GetBiggestContributor(db *sql.DB) (*Contributor, error) {
  return nil, nil
}

func GetExpenditureChange(db *sql.DB) (*Contributor, error) {
  return nil, nil
}

func GetContributionChange(db *sql.DB) (*Contributor, error) {
  return nil, nil
}
