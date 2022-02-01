package sap_api_caller

import (
	"fmt"
	"io/ioutil"
	"net/http"
	sap_api_output_formatter "sap-api-integrations-planned-independent-req-reads-rmq-kube/SAP_API_Output_Formatter"
	"strings"
	"sync"

	"github.com/latonaio/golang-logging-library-for-sap/logger"
	"golang.org/x/xerrors"
)

type RMQOutputter interface {
	Send(sendQueue string, payload map[string]interface{}) error
}

type SAPAPICaller struct {
	baseURL string
	apiKey  string  
	outputQueues []string
	outputter    RMQOutputter
	log     *logger.Logger
}

func NewSAPAPICaller(baseUrl string, outputQueueTo []string, outputter RMQOutputter, l *logger.Logger) *SAPAPICaller {
	return &SAPAPICaller{
		baseURL: baseUrl,
		apiKey:  GetApiKey(),  
		outputQueues: outputQueueTo,
		outputter:    outputter,
		log:     l,
	}
}

func (c *SAPAPICaller) AsyncGetPlannedIndependentRequirement(product, plant, mRPArea, plndIndepRqmtType, plndIndepRqmtVersion, requirementPlan, requirementSegment string, accepter []string) {
	wg := &sync.WaitGroup{}
	wg.Add(len(accepter))
	for _, fn := range accepter {
		switch fn {
		case "Header":
			func() {
				c.Header(product, plant, mRPArea, plndIndepRqmtType, plndIndepRqmtVersion, requirementPlan, requirementSegment)
				wg.Done()
			}()
		default:
			wg.Done()
		}
	}

	wg.Wait()
}

func (c *SAPAPICaller) Header(product, plant, mRPArea, plndIndepRqmtType, plndIndepRqmtVersion, requirementPlan, requirementSegment string) {
	headerData, err := c.callPlannedIndependentRequirementSrvAPIRequirementHeader("PlannedIndepRqmt", product, plant, mRPArea, plndIndepRqmtType, plndIndepRqmtVersion, requirementPlan, requirementSegment)
	if err != nil {
		c.log.Error(err)
		return
	}
	err = c.outputter.Send(c.outputQueues[0], map[string]interface{}{"message": headerData, "function": "PlannedIndependentRequirementHeader"})
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(headerData)

	itemData, err := c.callToItem(headerData[0].ToItem)
	if err != nil {
		c.log.Error(err)
		return
	}
	err = c.outputter.Send(c.outputQueues[0], map[string]interface{}{"message": itemData, "function": "PlannedIndependentRequirementItem"})
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(itemData)

}

func (c *SAPAPICaller) callPlannedIndependentRequirementSrvAPIRequirementHeader(api, product, plant, mRPArea, plndIndepRqmtType, plndIndepRqmtVersion, requirementPlan, requirementSegment string) ([]sap_api_output_formatter.Header, error) {
	url := strings.Join([]string{c.baseURL, "API_PLND_INDEP_RQMT_SRV", api}, "/")
	req, _ := http.NewRequest("GET", url, nil)

	c.setHeaderAPIKeyAccept(req)
	c.getQueryWithHeader(req, product, plant, mRPArea, plndIndepRqmtType, plndIndepRqmtVersion, requirementPlan, requirementSegment)

	resp, err := new(http.Client).Do(req)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToHeader(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) callToItem(url string) ([]sap_api_output_formatter.ToItem, error) {
	req, _ := http.NewRequest("GET", url, nil)
	c.setHeaderAPIKeyAccept(req)

	resp, err := new(http.Client).Do(req)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToToItem(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) setHeaderAPIKeyAccept(req *http.Request) {
	req.Header.Set("APIKey", c.apiKey)
	req.Header.Set("Accept", "application/json")
}

func (c *SAPAPICaller) getQueryWithHeader(req *http.Request, product, plant, mRPArea, plndIndepRqmtType, plndIndepRqmtVersion, requirementPlan, requirementSegment string) {
	params := req.URL.Query()
	params.Add("$filter", fmt.Sprintf("Product eq '%s' and Plant eq '%s' and MRPArea eq '%s' and PlndIndepRqmtType eq '%s' and PlndIndepRqmtVersion eq '%s' and RequirementPlan eq '%s' and RequirementSegment eq '%s'", product, plant, mRPArea, plndIndepRqmtType, plndIndepRqmtVersion, requirementPlan, requirementSegment))
	req.URL.RawQuery = params.Encode()
}
