package sap_api_output_formatter

type PlannedIndependentRequirementReads struct {
	ConnectionKey string `json:"connection_key"`
	Result        bool   `json:"result"`
	RedisKey      string `json:"redis_key"`
	Filepath      string `json:"filepath"`
	APISchema     string `json:"api_schema"`
	MaterialCode  string `json:"material_code"`
	Plant         string `json:"plant"`
	Deleted       bool   `json:"deleted"`
}

type Header struct {
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
    ToItem                        string      `json:"to_Item"`	
}

type ToItem struct {
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
}
