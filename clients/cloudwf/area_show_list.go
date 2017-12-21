package cloudwf

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type AreaShowListRequest struct {
	requests.RpcRequest
	Page int   `position:"Query" name:"Page"`
	Per  int   `position:"Query" name:"Per"`
	Sid  int64 `position:"Query" name:"Sid"`
}

func (req *AreaShowListRequest) Invoke(client *sdk.Client) (resp *AreaShowListResponse, err error) {
	req.InitWithApiInfo("cloudwf", "2017-03-28", "AreaShowList", "", "")
	resp = &AreaShowListResponse{}
	err = client.DoAction(req, resp)
	return
}

type AreaShowListResponse struct {
	responses.BaseResponse
	Success   bool
	Data      string
	ErrorCode int
	ErrorMsg  string
}
