package cloudwf

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type BusinessShowListRequest struct {
	requests.RpcRequest
	Page int `position:"Query" name:"Page"`
	Per  int `position:"Query" name:"Per"`
}

func (req *BusinessShowListRequest) Invoke(client *sdk.Client) (resp *BusinessShowListResponse, err error) {
	req.InitWithApiInfo("cloudwf", "2017-03-28", "BusinessShowList", "", "")
	resp = &BusinessShowListResponse{}
	err = client.DoAction(req, resp)
	return
}

type BusinessShowListResponse struct {
	responses.BaseResponse
	Success   bool
	Data      string
	ErrorCode int
	ErrorMsg  string
}
