package mts

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DeleteTemplateRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	TemplateId           string `position:"Query" name:"TemplateId"`
}

func (r DeleteTemplateRequest) Invoke(client *sdk.Client) (response *DeleteTemplateResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DeleteTemplateRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Mts", "2014-06-18", "DeleteTemplate", "", "")

	resp := struct {
		*responses.BaseResponse
		DeleteTemplateResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DeleteTemplateResponse

	err = client.DoAction(&req, &resp)
	return
}

type DeleteTemplateResponse struct {
	RequestId  string
	TemplateId string
}
