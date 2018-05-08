package vpc

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type UnassociateEipAddressRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	InstanceId           string `position:"Query" name:"InstanceId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	InstanceType         string `position:"Query" name:"InstanceType"`
	Force                string `position:"Query" name:"Force"`
	AllocationId         string `position:"Query" name:"AllocationId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (req *UnassociateEipAddressRequest) Invoke(client *sdk.Client) (resp *UnassociateEipAddressResponse, err error) {
	req.InitWithApiInfo("Vpc", "2016-04-28", "UnassociateEipAddress", "vpc", "")
	resp = &UnassociateEipAddressResponse{}
	err = client.DoAction(req, resp)
	return
}

type UnassociateEipAddressResponse struct {
	responses.BaseResponse
	RequestId common.String
}
