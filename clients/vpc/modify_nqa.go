package vpc

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ModifyNqaRequest struct {
	DestinationIp        string `position:"Query" name:"DestinationIp"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ClientToken          string `position:"Query" name:"ClientToken"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	NqaId                string `position:"Query" name:"NqaId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (r ModifyNqaRequest) Invoke(client *sdk.Client) (response *ModifyNqaResponse, err error) {
	req := struct {
		*requests.RpcRequest
		ModifyNqaRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Vpc", "2016-04-28", "ModifyNqa", "vpc", "")

	resp := struct {
		*responses.BaseResponse
		ModifyNqaResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.ModifyNqaResponse

	err = client.DoAction(&req, &resp)
	return
}

type ModifyNqaResponse struct {
	RequestId string
}
