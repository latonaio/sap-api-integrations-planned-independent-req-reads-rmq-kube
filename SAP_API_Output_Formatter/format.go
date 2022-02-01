package sap_api_output_formatter

import (
	"encoding/json"
	"sap-api-integrations-planned-independent-req-reads-rmq-kube/SAP_API_Caller/responses"

	"github.com/latonaio/golang-logging-library-for-sap/logger"
	"golang.org/x/xerrors"
)

func ConvertToHeader(raw []byte, l *logger.Logger) ([]Header, error) {
	pm := &responses.Header{}

	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to Header. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 10 {
		l.Info("raw data has too many Results. %d Results exist. show the first 10 of Results array", len(pm.D.Results))
	}

	header := make([]Header, 0, 10)
	for i := 0; i < 10 && i < len(pm.D.Results); i++ {
		data := pm.D.Results[i]
		header = append(header, Header{
	UpdateMc:                      data.UpdateMc,
	ToPlndIndepRqmtItemOc:         data.ToPlndIndepRqmtItemOc,
	Product:                       data.Product,
	Plant:                         data.Plant,
	MRPArea:                       data.MRPArea,
	PlndIndepRqmtType:             data.PlndIndepRqmtType,
	PlndIndepRqmtVersion:          data.PlndIndepRqmtVersion,
	RequirementPlan:               data.RequirementPlan,
	RequirementSegment:            data.RequirementSegment,
	RequirementPlanIsExternal:     data.RequirementPlanIsExternal,
	PlndIndepRqmtInternalID:       data.PlndIndepRqmtInternalID,
	PlndIndepRqmtIsActive:         data.PlndIndepRqmtIsActive,
	WBSElement:                    data.WBSElement,
	PlndIndepRqmtAcctAssgmtCat:    data.PlndIndepRqmtAcctAssgmtCat,
	PlndIndepRqmtLastChgdDateTime: data.PlndIndepRqmtLastChgdDateTime,
    ToItem:                        data.ToItem.Deferred.URI,
		})
	}

	return header, nil
}

func ConvertToToItem(raw []byte, l *logger.Logger) ([]ToItem, error) {
	pm := &responses.ToItem{}

	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to ToItem. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 10 {
		l.Info("raw data has too many Results. %d Results exist. expected only 1 Result. Use the first of Results array", len(pm.D.Results))
	}
	toItem := make([]ToItem, 0, 10)
	for i := 0; i < 10 && i < len(pm.D.Results); i++ {
		data := pm.D.Results[i]
		toItem = append(toItem, ToItem{
	UpdateMc:                      data.UpdateMc,
	Product:                       data.Product,
	Plant:                         data.Plant,
	MRPArea:                       data.MRPArea,
	PlndIndepRqmtType:             data.PlndIndepRqmtType,
	PlndIndepRqmtVersion:          data.PlndIndepRqmtVersion,
	RequirementPlan:               data.RequirementPlan,
	RequirementSegment:            data.RequirementSegment,
	PlndIndepRqmtPeriod:           data.PlndIndepRqmtPeriod,
	PeriodType:                    data.PeriodType,
	PlndIndepRqmtPeriodStartDate:  data.PlndIndepRqmtPeriodStartDate,
	PlndIndepRqmtInternalID:       data.PlndIndepRqmtInternalID,
	WorkingDayDate:                data.WorkingDayDate,
	PlannedQuantity:               data.PlannedQuantity,
	WithdrawalQuantity:            data.WithdrawalQuantity,
	UnitOfMeasure:                 data.UnitOfMeasure,
	LastChangedByUser:             data.LastChangedByUser,
	LastChangeDate:                data.LastChangeDate,
	PlndIndepRqmtLastChgdDateTime :data.PlndIndepRqmtLastChgdDateTime,
		})
	}

	return toItem, nil
}
