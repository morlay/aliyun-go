package arms4finance

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ARMSQueryDataSetRequest struct {
	Measuress     *ARMSQueryDataSetMeasuresList   `position:"Query" type:"Repeated" name:"Measures"`
	IntervalInSec int                             `position:"Query" name:"IntervalInSec"`
	DateStr       string                          `position:"Query" name:"DateStr"`
	IsDrillDown   string                          `position:"Query" name:"IsDrillDown"`
	MinTime       int64                           `position:"Query" name:"MinTime"`
	DatasetId     int64                           `position:"Query" name:"DatasetId"`
	MaxTime       int64                           `position:"Query" name:"MaxTime"`
	Dimensionss   *ARMSQueryDataSetDimensionsList `position:"Query" type:"Repeated" name:"Dimensions"`
}

func (r ARMSQueryDataSetRequest) Invoke(client *sdk.Client) (response *ARMSQueryDataSetResponse, err error) {
	req := struct {
		*requests.RpcRequest
		ARMSQueryDataSetRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("ARMS4FINANCE", "2017-11-30", "ARMSQueryDataSet", "", "")

	resp := struct {
		*responses.BaseResponse
		ARMSQueryDataSetResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.ARMSQueryDataSetResponse

	err = client.DoAction(&req, &resp)
	return
}

type ARMSQueryDataSetDimensions struct {
	Key   string `name:"Key"`
	Value string `name:"Value"`
}

type ARMSQueryDataSetResponse struct {
	Data string
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
