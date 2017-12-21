package vpc

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DeleteBgpGroupRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ClientToken          string `position:"Query" name:"ClientToken"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	BgpGroupId           string `position:"Query" name:"BgpGroupId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (r DeleteBgpGroupRequest) Invoke(client *sdk.Client) (response *DeleteBgpGroupResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DeleteBgpGroupRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Vpc", "2016-04-28", "DeleteBgpGroup", "vpc", "")

	resp := struct {
		*responses.BaseResponse
		DeleteBgpGroupResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DeleteBgpGroupResponse

	err = client.DoAction(&req, &resp)
	return
}

type DeleteBgpGroupResponse struct {
	RequestId string
}
