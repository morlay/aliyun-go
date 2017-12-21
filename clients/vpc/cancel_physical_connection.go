package vpc

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type CancelPhysicalConnectionRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ClientToken          string `position:"Query" name:"ClientToken"`
	PhysicalConnectionId string `position:"Query" name:"PhysicalConnectionId"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	UserCidr             string `position:"Query" name:"UserCidr"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (req *CancelPhysicalConnectionRequest) Invoke(client *sdk.Client) (resp *CancelPhysicalConnectionResponse, err error) {
	req.InitWithApiInfo("Vpc", "2016-04-28", "CancelPhysicalConnection", "vpc", "")
	resp = &CancelPhysicalConnectionResponse{}
	err = client.DoAction(req, resp)
	return
}

type CancelPhysicalConnectionResponse struct {
	responses.BaseResponse
	RequestId string
}
