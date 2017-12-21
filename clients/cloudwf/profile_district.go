package cloudwf

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ProfileDistrictRequest struct {
	requests.RpcRequest
	Gsid int64 `position:"Query" name:"Gsid"`
}

func (req *ProfileDistrictRequest) Invoke(client *sdk.Client) (resp *ProfileDistrictResponse, err error) {
	req.InitWithApiInfo("cloudwf", "2017-03-28", "ProfileDistrict", "", "")
	resp = &ProfileDistrictResponse{}
	err = client.DoAction(req, resp)
	return
}

type ProfileDistrictResponse struct {
	responses.BaseResponse
	Success   bool
	Data      string
	ErrorCode int
	ErrorMsg  string
}
