package emr

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ModifyUserChannelInfoRequest struct {
	requests.RpcRequest
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	ChannelId       string `position:"Query" name:"ChannelId"`
}

func (req *ModifyUserChannelInfoRequest) Invoke(client *sdk.Client) (resp *ModifyUserChannelInfoResponse, err error) {
	req.InitWithApiInfo("Emr", "2016-04-08", "ModifyUserChannelInfo", "", "")
	resp = &ModifyUserChannelInfoResponse{}
	err = client.DoAction(req, resp)
	return
}

type ModifyUserChannelInfoResponse struct {
	responses.BaseResponse
	RequestId string
	Success   bool
}
