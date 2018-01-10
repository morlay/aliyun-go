package idst

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type InstanceDetailRequest struct {
	requests.RoaRequest
}

func (req *InstanceDetailRequest) Invoke(client *sdk.Client) (resp *InstanceDetailResponse, err error) {
	req.InitWithApiInfo("IDST", "2018-01-08", "InstanceDetail", "/console/instance/detail", "", "")
	req.Method = "POST"

	resp = &InstanceDetailResponse{}
	err = client.DoAction(req, resp)
	return
}

type InstanceDetailResponse struct {
	responses.BaseResponse
}
