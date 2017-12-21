package cr

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type GetRegionRequest struct {
	requests.RoaRequest
	Domain string `position:"Query" name:"Domain"`
}

func (req *GetRegionRequest) Invoke(client *sdk.Client) (resp *GetRegionResponse, err error) {
	req.InitWithApiInfo("cr", "2016-06-07", "GetRegion", "/regions", "", "")
	req.Method = "GET"

	resp = &GetRegionResponse{}
	err = client.DoAction(req, resp)
	return
}

type GetRegionResponse struct {
	responses.BaseResponse
}
