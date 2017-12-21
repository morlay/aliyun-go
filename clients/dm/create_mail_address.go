package dm

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type CreateMailAddressRequest struct {
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	AccountName          string `position:"Query" name:"AccountName"`
	ReplyAddress         string `position:"Query" name:"ReplyAddress"`
	Sendtype             string `position:"Query" name:"Sendtype"`
}

func (r CreateMailAddressRequest) Invoke(client *sdk.Client) (response *CreateMailAddressResponse, err error) {
	req := struct {
		*requests.RpcRequest
		CreateMailAddressRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Dm", "2015-11-23", "CreateMailAddress", "", "")

	resp := struct {
		*responses.BaseResponse
		CreateMailAddressResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.CreateMailAddressResponse

	err = client.DoAction(&req, &resp)
	return
}

type CreateMailAddressResponse struct {
	RequestId string
}
