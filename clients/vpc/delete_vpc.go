package vpc

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DeleteVpcRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	VpcId                string `position:"Query" name:"VpcId"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (req *DeleteVpcRequest) Invoke(client *sdk.Client) (resp *DeleteVpcResponse, err error) {
	req.InitWithApiInfo("Vpc", "2016-04-28", "DeleteVpc", "vpc", "")
	resp = &DeleteVpcResponse{}
	err = client.DoAction(req, resp)
	return
}

type DeleteVpcResponse struct {
	responses.BaseResponse
	RequestId common.String
}
