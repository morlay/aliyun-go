package ehpc

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type StopJobsRequest struct {
	requests.RpcRequest
	Jobs      string `position:"Query" name:"Jobs"`
	ClusterId string `position:"Query" name:"ClusterId"`
}

func (req *StopJobsRequest) Invoke(client *sdk.Client) (resp *StopJobsResponse, err error) {
	req.InitWithApiInfo("EHPC", "2017-07-14", "StopJobs", "ehs", "")
	resp = &StopJobsResponse{}
	err = client.DoAction(req, resp)
	return
}

type StopJobsResponse struct {
	responses.BaseResponse
	RequestId common.String
}
