package vpc

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type AssociateEipAddressRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	InstanceId           string `position:"Query" name:"InstanceId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	InstanceType         string `position:"Query" name:"InstanceType"`
	AllocationId         string `position:"Query" name:"AllocationId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (req *AssociateEipAddressRequest) Invoke(client *sdk.Client) (resp *AssociateEipAddressResponse, err error) {
	req.InitWithApiInfo("Vpc", "2016-04-28", "AssociateEipAddress", "vpc", "")
	resp = &AssociateEipAddressResponse{}
	err = client.DoAction(req, resp)
	return
}

type AssociateEipAddressResponse struct {
	responses.BaseResponse
	RequestId string
}
