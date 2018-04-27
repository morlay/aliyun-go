package emr

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DeleteResourcePoolRequest struct {
	requests.RpcRequest
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	ResourcePoolId  string `position:"Query" name:"ResourcePoolId"`
	ClusterId       string `position:"Query" name:"ClusterId"`
}

func (req *DeleteResourcePoolRequest) Invoke(client *sdk.Client) (resp *DeleteResourcePoolResponse, err error) {
	req.InitWithApiInfo("Emr", "2016-04-08", "DeleteResourcePool", "", "")
	resp = &DeleteResourcePoolResponse{}
	err = client.DoAction(req, resp)
	return
}

type DeleteResourcePoolResponse struct {
	responses.BaseResponse
	RequestId string
}
