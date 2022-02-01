package responses

type Header struct {
	D struct {
		Results []struct {
			Metadata struct {
				ID   string `json:"id"`
				URI  string `json:"uri"`
				Type string `json:"type"`
				Etag string `json:"etag"`
			} `json:"__metadata"`
			UpdateMc                      bool        `json:"Update_mc"`
			ToPlndIndepRqmtItemOc         bool        `json:"to_PlndIndepRqmtItem_oc"`
			Product                       string      `json:"Product"`
			Plant                         string      `json:"Plant"`
			MRPArea                       string      `json:"MRPArea"`
			PlndIndepRqmtType             string      `json:"PlndIndepRqmtType"`
			PlndIndepRqmtVersion          string      `json:"PlndIndepRqmtVersion"`
			RequirementPlan               string      `json:"RequirementPlan"`
			RequirementSegment            string      `json:"RequirementSegment"`
			RequirementPlanIsExternal     bool        `json:"RequirementPlanIsExternal"`
			PlndIndepRqmtInternalID       string      `json:"PlndIndepRqmtInternalID"`
			PlndIndepRqmtIsActive         string      `json:"PlndIndepRqmtIsActive"`
			WBSElement                    string      `json:"WBSElement"`
			PlndIndepRqmtAcctAssgmtCat    string      `json:"PlndIndepRqmtAcctAssgmtCat"`
			PlndIndepRqmtLastChgdDateTime string      `json:"PlndIndepRqmtLastChgdDateTime"`
			ToItem           struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"to_PlndIndepRqmtItem"`
		} `json:"results"`
	} `json:"d"`
}
