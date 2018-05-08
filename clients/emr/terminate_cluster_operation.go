package emr

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type TerminateClusterOperationRequest struct {
	requests.RpcRequest
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	OperationId     string `position:"Query" name:"OperationId"`
	ClusterId       string `position:"Query" name:"ClusterId"`
}

func (req *TerminateClusterOperationRequest) Invoke(client *sdk.Client) (resp *TerminateClusterOperationResponse, err error) {
	req.InitWithApiInfo("Emr", "2016-04-08", "TerminateClusterOperation", "", "")
	resp = &TerminateClusterOperationResponse{}
	err = client.DoAction(req, resp)
	return
}

type TerminateClusterOperationResponse struct {
	responses.BaseResponse
	RequestId common.String
}
