package arms4finance

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type WhereInDimQueryRequest struct {
	requests.RpcRequest
	WhereInKey     string                            `position:"Query" name:"WhereInKey"`
	Measuress      *WhereInDimQueryMeasuresList      `position:"Query" type:"Repeated" name:"Measures"`
	IntervalInSec  int                               `position:"Query" name:"IntervalInSec"`
	DateStr        string                            `position:"Query" name:"DateStr"`
	IsDrillDown    string                            `position:"Query" name:"IsDrillDown"`
	MinTime        int64                             `position:"Query" name:"MinTime"`
	DatasetId      int64                             `position:"Query" name:"DatasetId"`
	WhereInValuess *WhereInDimQueryWhereInValuesList `position:"Query" type:"Repeated" name:"WhereInValues"`
	MaxTime        int64                             `position:"Query" name:"MaxTime"`
	Dimensionss    *WhereInDimQueryDimensionsList    `position:"Query" type:"Repeated" name:"Dimensions"`
}

func (req *WhereInDimQueryRequest) Invoke(client *sdk.Client) (resp *WhereInDimQueryResponse, err error) {
	req.InitWithApiInfo("ARMS4FINANCE", "2017-11-30", "WhereInDimQuery", "", "")
	resp = &WhereInDimQueryResponse{}
	err = client.DoAction(req, resp)
	return
}

type WhereInDimQueryDimensions struct {
	Key   string `name:"Key"`
	Value string `name:"Value"`
}

type WhereInDimQueryResponse struct {
	responses.BaseResponse
	Data common.String
}

type WhereInDimQueryMeasuresList []string

func (list *WhereInDimQueryMeasuresList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]string)
	err := json.Unmarshal(data, &m)
	if err != nil {
		return err
	}
	for _, v := range m {
		*list = v
		break
	}
	return nil
}

type WhereInDimQueryWhereInValuesList []string

func (list *WhereInDimQueryWhereInValuesList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]string)
	err := json.Unmarshal(data, &m)
	if err != nil {
		return err
	}
	for _, v := range m {
		*list = v
		break
	}
	return nil
}

type WhereInDimQueryDimensionsList []WhereInDimQueryDimensions

func (list *WhereInDimQueryDimensionsList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]WhereInDimQueryDimensions)
	err := json.Unmarshal(data, &m)
	if err != nil {
		return err
	}
	for _, v := range m {
		*list = v
		break
	}
	return nil
}
