package vpc

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DeleteVSwitchRequest struct {
	VSwitchId            string `position:"Query" name:"VSwitchId"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (r DeleteVSwitchRequest) Invoke(client *sdk.Client) (response *DeleteVSwitchResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DeleteVSwitchRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Vpc", "2016-04-28", "DeleteVSwitch", "vpc", "")

	resp := struct {
		*responses.BaseResponse
		DeleteVSwitchResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DeleteVSwitchResponse

	err = client.DoAction(&req, &resp)
	return
}

type DeleteVSwitchResponse struct {
	RequestId string
}
