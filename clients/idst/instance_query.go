package idst

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type InstanceQueryRequest struct {
	requests.RoaRequest
}

func (req *InstanceQueryRequest) Invoke(client *sdk.Client) (resp *InstanceQueryResponse, err error) {
	req.InitWithApiInfo("IDST", "2018-01-08", "InstanceQuery", "/console/instance/query", "", "")
	req.Method = "POST"

	resp = &InstanceQueryResponse{}
	err = client.DoAction(req, resp)
	return
}

type InstanceQueryResponse struct {
	responses.BaseResponse
}
