package arms4finance

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type ARMSQueryDataSetRequest struct {
	requests.RpcRequest
	Measuress     *ARMSQueryDataSetMeasuresList   `position:"Query" type:"Repeated" name:"Measures"`
	IntervalInSec int                             `position:"Query" name:"IntervalInSec"`
	DateStr       string                          `position:"Query" name:"DateStr"`
	IsDrillDown   string                          `position:"Query" name:"IsDrillDown"`
	MinTime       int64                           `position:"Query" name:"MinTime"`
	DatasetId     int64                           `position:"Query" name:"DatasetId"`
	MaxTime       int64                           `position:"Query" name:"MaxTime"`
	Dimensionss   *ARMSQueryDataSetDimensionsList `position:"Query" type:"Repeated" name:"Dimensions"`
}

func (req *ARMSQueryDataSetRequest) Invoke(client *sdk.Client) (resp *ARMSQueryDataSetResponse, err error) {
	req.InitWithApiInfo("ARMS4FINANCE", "2017-11-30", "ARMSQueryDataSet", "", "")
	resp = &ARMSQueryDataSetResponse{}
	err = client.DoAction(req, resp)
	return
}

type ARMSQueryDataSetDimensions struct {
	Key   string `name:"Key"`
	Value string `name:"Value"`
}

type ARMSQueryDataSetResponse struct {
	responses.BaseResponse
	Data common.String
}

type ARMSQueryDataSetMeasuresList []string

func (list *ARMSQueryDataSetMeasuresList) UnmarshalJSON(data []byte) error {
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

type ARMSQueryDataSetDimensionsList []ARMSQueryDataSetDimensions

func (list *ARMSQueryDataSetDimensionsList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ARMSQueryDataSetDimensions)
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
