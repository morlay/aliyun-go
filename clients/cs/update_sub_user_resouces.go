package cs

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type UpdateSubUserResoucesRequest struct {
	requests.RoaRequest
}

func (req *UpdateSubUserResoucesRequest) Invoke(client *sdk.Client) (resp *UpdateSubUserResoucesResponse, err error) {
	req.InitWithApiInfo("CS", "2015-12-15", "UpdateSubUserResouces", "/ram/resources", "", "")
	req.Method = "POST"

	resp = &UpdateSubUserResoucesResponse{}
	err = client.DoAction(req, resp)
	return
}

type UpdateSubUserResoucesResponse struct {
	responses.BaseResponse
}
