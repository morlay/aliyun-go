package emr

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type RefreshClusterResourcePoolRequest struct {
	requests.RpcRequest
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	ResourcePoolId  int64  `position:"Query" name:"ResourcePoolId"`
	ClusterId       string `position:"Query" name:"ClusterId"`
}

func (req *RefreshClusterResourcePoolRequest) Invoke(client *sdk.Client) (resp *RefreshClusterResourcePoolResponse, err error) {
	req.InitWithApiInfo("Emr", "2016-04-08", "RefreshClusterResourcePool", "", "")
	resp = &RefreshClusterResourcePoolResponse{}
	err = client.DoAction(req, resp)
	return
}

type RefreshClusterResourcePoolResponse struct {
	responses.BaseResponse
	RequestId          string
	WorkFlowInstanceId string
	OperationId        string
}
