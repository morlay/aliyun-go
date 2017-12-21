package mts

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type QueryAuthConfigRequest struct {
	ResourceOwnerId      string `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              string `position:"Query" name:"OwnerId"`
}

func (r QueryAuthConfigRequest) Invoke(client *sdk.Client) (response *QueryAuthConfigResponse, err error) {
	req := struct {
		*requests.RpcRequest
		QueryAuthConfigRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Mts", "2014-06-18", "QueryAuthConfig", "", "")

	resp := struct {
		*responses.BaseResponse
		QueryAuthConfigResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.QueryAuthConfigResponse

	err = client.DoAction(&req, &resp)
	return
}

type QueryAuthConfigResponse struct {
	RequestId string
	Key1      string
	Key2      string
}
