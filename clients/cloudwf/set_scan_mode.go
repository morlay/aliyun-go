package cloudwf

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type SetScanModeRequest struct {
	requests.RpcRequest
	Operation int                     `position:"Query" name:"Operation"`
	MacLists  *SetScanModeMacListList `position:"Query" type:"Repeated" name:"MacList"`
}

func (req *SetScanModeRequest) Invoke(client *sdk.Client) (resp *SetScanModeResponse, err error) {
	req.InitWithApiInfo("cloudwf", "2017-03-28", "SetScanMode", "", "")
	resp = &SetScanModeResponse{}
	err = client.DoAction(req, resp)
	return
}

type SetScanModeResponse struct {
	responses.BaseResponse
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
