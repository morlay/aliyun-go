package cloudwf

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ListAccountConfigRequest struct {
	OrderCol    string `position:"Query" name:"OrderCol"`
	Length      int    `position:"Query" name:"Length"`
	SearchEmail string `position:"Query" name:"SearchEmail"`
	PageIndex   int    `position:"Query" name:"PageIndex"`
	OrderDir    string `position:"Query" name:"OrderDir"`
}

func (r ListAccountConfigRequest) Invoke(client *sdk.Client) (response *ListAccountConfigResponse, err error) {
	req := struct {
		*requests.RpcRequest
		ListAccountConfigRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("cloudwf", "2017-03-28", "ListAccountConfig", "", "")

	resp := struct {
		*responses.BaseResponse
		ListAccountConfigResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.ListAccountConfigResponse

	err = client.DoAction(&req, &resp)
	return
}

type ListAccountConfigResponse struct {
	RequestId string
	Success   bool
	Message   string
	Data      string
	ErrorCode int
	ErrorMsg  string
}
