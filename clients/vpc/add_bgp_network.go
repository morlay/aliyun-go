package vpc

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type AddBgpNetworkRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ClientToken          string `position:"Query" name:"ClientToken"`
	RouterId             string `position:"Query" name:"RouterId"`
	VpcId                string `position:"Query" name:"VpcId"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	DstCidrBlock         string `position:"Query" name:"DstCidrBlock"`
}

func (req *AddBgpNetworkRequest) Invoke(client *sdk.Client) (resp *AddBgpNetworkResponse, err error) {
	req.InitWithApiInfo("Vpc", "2016-04-28", "AddBgpNetwork", "vpc", "")
	resp = &AddBgpNetworkResponse{}
	err = client.DoAction(req, resp)
	return
}

type AddBgpNetworkResponse struct {
	responses.BaseResponse
	RequestId string
}
