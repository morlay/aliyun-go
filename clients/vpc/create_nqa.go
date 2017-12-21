package vpc

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type CreateNqaRequest struct {
	DestinationIp        string `position:"Query" name:"DestinationIp"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ClientToken          string `position:"Query" name:"ClientToken"`
	RouterId             string `position:"Query" name:"RouterId"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (r CreateNqaRequest) Invoke(client *sdk.Client) (response *CreateNqaResponse, err error) {
	req := struct {
		*requests.RpcRequest
		CreateNqaRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Vpc", "2016-04-28", "CreateNqa", "vpc", "")

	resp := struct {
		*responses.BaseResponse
		CreateNqaResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.CreateNqaResponse

	err = client.DoAction(&req, &resp)
	return
}

type CreateNqaResponse struct {
	RequestId string
	NqaId     string
}
