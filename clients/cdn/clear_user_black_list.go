package cdn

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ClearUserBlackListRequest struct {
	SecurityToken string `position:"Query" name:"SecurityToken"`
	OwnerAccount  string `position:"Query" name:"OwnerAccount"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
}

func (r ClearUserBlackListRequest) Invoke(client *sdk.Client) (response *ClearUserBlackListResponse, err error) {
	req := struct {
		*requests.RpcRequest
		ClearUserBlackListRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Cdn", "2014-11-11", "ClearUserBlackList", "", "")

	resp := struct {
		*responses.BaseResponse
		ClearUserBlackListResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.ClearUserBlackListResponse

	err = client.DoAction(&req, &resp)
	return
}

type ClearUserBlackListResponse struct {
	RequestId string
}
