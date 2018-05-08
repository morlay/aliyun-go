package vpc

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DeleteBgpGroupRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ClientToken          string `position:"Query" name:"ClientToken"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	BgpGroupId           string `position:"Query" name:"BgpGroupId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (req *DeleteBgpGroupRequest) Invoke(client *sdk.Client) (resp *DeleteBgpGroupResponse, err error) {
	req.InitWithApiInfo("Vpc", "2016-04-28", "DeleteBgpGroup", "vpc", "")
	resp = &DeleteBgpGroupResponse{}
	err = client.DoAction(req, resp)
	return
}

type DeleteBgpGroupResponse struct {
	responses.BaseResponse
	RequestId common.String
}
