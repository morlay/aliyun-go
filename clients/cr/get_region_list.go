package cr

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type GetRegionListRequest struct {
	requests.RoaRequest
}

func (req *GetRegionListRequest) Invoke(client *sdk.Client) (resp *GetRegionListResponse, err error) {
	req.InitWithApiInfo("cr", "2016-06-07", "GetRegionList", "/regions", "", "")
	req.Method = "GET"

	resp = &GetRegionListResponse{}
	err = client.DoAction(req, resp)
	return
}

type GetRegionListResponse struct {
	responses.BaseResponse
}
