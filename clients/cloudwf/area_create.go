package cloudwf

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type AreaCreateRequest struct {
	requests.RpcRequest
	Name string `position:"Query" name:"Name"`
	Dids string `position:"Query" name:"Dids"`
	Sid  int64  `position:"Query" name:"Sid"`
}

func (req *AreaCreateRequest) Invoke(client *sdk.Client) (resp *AreaCreateResponse, err error) {
	req.InitWithApiInfo("cloudwf", "2017-03-28", "AreaCreate", "", "")
	resp = &AreaCreateResponse{}
	err = client.DoAction(req, resp)
	return
}

type AreaCreateResponse struct {
	responses.BaseResponse
	Success   bool
	Data      string
	ErrorCode int
	ErrorMsg  string
}
