package ehpc

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DeleteJobsRequest struct {
	requests.RpcRequest
	Jobs      string `position:"Query" name:"Jobs"`
	ClusterId string `position:"Query" name:"ClusterId"`
}

func (req *DeleteJobsRequest) Invoke(client *sdk.Client) (resp *DeleteJobsResponse, err error) {
	req.InitWithApiInfo("EHPC", "2018-04-12", "DeleteJobs", "ehs", "")
	resp = &DeleteJobsResponse{}
	err = client.DoAction(req, resp)
	return
}

type DeleteJobsResponse struct {
	responses.BaseResponse
	RequestId common.String
}
