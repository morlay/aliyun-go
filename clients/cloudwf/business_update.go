package cloudwf

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type BusinessUpdateRequest struct {
	requests.RpcRequest
	Warn             int    `position:"Query" name:"Warn"`
	BusinessCity     string `position:"Query" name:"BusinessCity"`
	WarnEmail        string `position:"Query" name:"WarnEmail"`
	BusinessAddress  string `position:"Query" name:"BusinessAddress"`
	Bid              int64  `position:"Query" name:"Bid"`
	BusinessManager  string `position:"Query" name:"BusinessManager"`
	BusinessProvince string `position:"Query" name:"BusinessProvince"`
}

func (req *BusinessUpdateRequest) Invoke(client *sdk.Client) (resp *BusinessUpdateResponse, err error) {
	req.InitWithApiInfo("cloudwf", "2017-03-28", "BusinessUpdate", "", "")
	resp = &BusinessUpdateResponse{}
	err = client.DoAction(req, resp)
	return
}

type BusinessUpdateResponse struct {
	responses.BaseResponse
	Success   bool
	Data      string
	ErrorCode int
	ErrorMsg  string
}
