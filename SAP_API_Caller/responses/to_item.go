package responses

type ToItem struct {
	D struct {
		Results []struct {
			Metadata struct {
				ID   string `json:"id"`
				URI  string `json:"uri"`
				Type string `json:"type"`
				Etag string `json:"etag"`
			} `json:"__metadata"`
			UpdateMc                      bool        `json:"Update_mc"`
			Product                       string      `json:"Product"`
			Plant                         string      `json:"Plant"`
			MRPArea                       string      `json:"MRPArea"`
			PlndIndepRqmtType             string      `json:"PlndIndepRqmtType"`
			PlndIndepRqmtVersion          string      `json:"PlndIndepRqmtVersion"`
			RequirementPlan               string      `json:"RequirementPlan"`
			RequirementSegment            string      `json:"RequirementSegment"`
			PlndIndepRqmtPeriod           string      `json:"PlndIndepRqmtPeriod"`
			PeriodType                    string      `json:"PeriodType"`
			PlndIndepRqmtPeriodStartDate  string      `json:"PlndIndepRqmtPeriodStartDate"`
			PlndIndepRqmtInternalID       string      `json:"PlndIndepRqmtInternalID"`
			WorkingDayDate                string      `json:"WorkingDayDate"`
			PlannedQuantity               string      `json:"PlannedQuantity"`
			WithdrawalQuantity            string      `json:"WithdrawalQuantity"`
			UnitOfMeasure                 string      `json:"UnitOfMeasure"`
			LastChangedByUser             string      `json:"LastChangedByUser"`
			LastChangeDate                string      `json:"LastChangeDate"`
			PlndIndepRqmtLastChgdDateTime string      `json:"PlndIndepRqmtLastChgdDateTime"`
		} `json:"results"`
	} `json:"d"`
}
