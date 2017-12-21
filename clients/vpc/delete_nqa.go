package vpc

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DeleteNqaRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ClientToken          string `position:"Query" name:"ClientToken"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	NqaId                string `position:"Query" name:"NqaId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (r DeleteNqaRequest) Invoke(client *sdk.Client) (response *DeleteNqaResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DeleteNqaRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Vpc", "2016-04-28", "DeleteNqa", "vpc", "")

	resp := struct {
		*responses.BaseResponse
		DeleteNqaResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DeleteNqaResponse

	err = client.DoAction(&req, &resp)
	return
}

type DeleteNqaResponse struct {
	RequestId string
}
