package vpc

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ReleaseEipAddressRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	AllocationId         string `position:"Query" name:"AllocationId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (req *ReleaseEipAddressRequest) Invoke(client *sdk.Client) (resp *ReleaseEipAddressResponse, err error) {
	req.InitWithApiInfo("Vpc", "2016-04-28", "ReleaseEipAddress", "vpc", "")
	resp = &ReleaseEipAddressResponse{}
	err = client.DoAction(req, resp)
	return
}

type ReleaseEipAddressResponse struct {
	responses.BaseResponse
	RequestId string
}
