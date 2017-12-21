package cloudwf

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type AreaUpdateRequest struct {
	requests.RpcRequest
	Name string `position:"Query" name:"Name"`
	Dids string `position:"Query" name:"Dids"`
	Aid  int64  `position:"Query" name:"Aid"`
	Sid  int64  `position:"Query" name:"Sid"`
}

func (req *AreaUpdateRequest) Invoke(client *sdk.Client) (resp *AreaUpdateResponse, err error) {
	req.InitWithApiInfo("cloudwf", "2017-03-28", "AreaUpdate", "", "")
	resp = &AreaUpdateResponse{}
	err = client.DoAction(req, resp)
	return
}

type AreaUpdateResponse struct {
	responses.BaseResponse
	Success   bool
	Data      string
	ErrorCode int
	ErrorMsg  string
}
