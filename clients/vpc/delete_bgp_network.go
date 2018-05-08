package vpc

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DeleteBgpNetworkRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ClientToken          string `position:"Query" name:"ClientToken"`
	RouterId             string `position:"Query" name:"RouterId"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	DstCidrBlock         string `position:"Query" name:"DstCidrBlock"`
}

func (req *DeleteBgpNetworkRequest) Invoke(client *sdk.Client) (resp *DeleteBgpNetworkResponse, err error) {
	req.InitWithApiInfo("Vpc", "2016-04-28", "DeleteBgpNetwork", "vpc", "")
	resp = &DeleteBgpNetworkResponse{}
	err = client.DoAction(req, resp)
	return
}

type DeleteBgpNetworkResponse struct {
	responses.BaseResponse
	RequestId common.String
}
