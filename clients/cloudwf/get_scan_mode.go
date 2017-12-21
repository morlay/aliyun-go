package cloudwf

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type GetScanModeRequest struct {
	requests.RpcRequest
	MacLists *GetScanModeMacListList `position:"Query" type:"Repeated" name:"MacList"`
}

func (req *GetScanModeRequest) Invoke(client *sdk.Client) (resp *GetScanModeResponse, err error) {
	req.InitWithApiInfo("cloudwf", "2017-03-28", "GetScanMode", "", "")
	resp = &GetScanModeResponse{}
	err = client.DoAction(req, resp)
	return
}

type GetScanModeResponse struct {
	responses.BaseResponse
	RequestId string
	Success   bool
	Message   string
	Data      string
	ErrorCode int
	ErrorMsg  string
}

type GetScanModeMacListList []string

func (list *GetScanModeMacListList) UnmarshalJSON(data []byte) error {
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
