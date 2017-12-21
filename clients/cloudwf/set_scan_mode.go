package cloudwf

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type SetScanModeRequest struct {
	Operation int                     `position:"Query" name:"Operation"`
	MacLists  *SetScanModeMacListList `position:"Query" type:"Repeated" name:"MacList"`
}

func (r SetScanModeRequest) Invoke(client *sdk.Client) (response *SetScanModeResponse, err error) {
	req := struct {
		*requests.RpcRequest
		SetScanModeRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("cloudwf", "2017-03-28", "SetScanMode", "", "")

	resp := struct {
		*responses.BaseResponse
		SetScanModeResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.SetScanModeResponse

	err = client.DoAction(&req, &resp)
	return
}

type SetScanModeResponse struct {
	RequestId string
	Success   bool
	Message   string
	Data      string
	ErrorCode int
	ErrorMsg  string
}

type SetScanModeMacListList []string

func (list *SetScanModeMacListList) UnmarshalJSON(data []byte) error {
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
