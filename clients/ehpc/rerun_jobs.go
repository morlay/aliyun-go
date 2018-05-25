package ehpc

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type RerunJobsRequest struct {
	requests.RpcRequest
	Jobs      string `position:"Query" name:"Jobs"`
	ClusterId string `position:"Query" name:"ClusterId"`
}

func (req *RerunJobsRequest) Invoke(client *sdk.Client) (resp *RerunJobsResponse, err error) {
	req.InitWithApiInfo("EHPC", "2018-04-12", "RerunJobs", "ehs", "")
	resp = &RerunJobsResponse{}
	err = client.DoAction(req, resp)
	return
}

type RerunJobsResponse struct {
	responses.BaseResponse
	RequestId common.String
}
