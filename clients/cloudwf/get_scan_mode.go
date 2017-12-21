package cloudwf

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type GetScanModeRequest struct {
	MacLists *GetScanModeMacListList `position:"Query" type:"Repeated" name:"MacList"`
}

func (r GetScanModeRequest) Invoke(client *sdk.Client) (response *GetScanModeResponse, err error) {
	req := struct {
		*requests.RpcRequest
		GetScanModeRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("cloudwf", "2017-03-28", "GetScanMode", "", "")

	resp := struct {
		*responses.BaseResponse
		GetScanModeResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.GetScanModeResponse

	err = client.DoAction(&req, &resp)
	return
}

type GetScanModeResponse struct {
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
