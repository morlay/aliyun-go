package ehpc

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type RerunJobsRequest struct {
	requests.RpcRequest
	Jobs      string `position:"Query" name:"Jobs"`
	ClusterId string `position:"Query" name:"ClusterId"`
}

func (req *RerunJobsRequest) Invoke(client *sdk.Client) (resp *RerunJobsResponse, err error) {
	req.InitWithApiInfo("EHPC", "2017-07-14", "RerunJobs", "ehs", "")
	resp = &RerunJobsResponse{}
	err = client.DoAction(req, resp)
	return
}

type RerunJobsResponse struct {
	responses.BaseResponse
	RequestId string
}
