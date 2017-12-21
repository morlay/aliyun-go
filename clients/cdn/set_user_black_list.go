package cdn

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type SetUserBlackListRequest struct {
	ConfigUrl     string `position:"Query" name:"ConfigUrl"`
	SecurityToken string `position:"Query" name:"SecurityToken"`
	OwnerAccount  string `position:"Query" name:"OwnerAccount"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
}

func (r SetUserBlackListRequest) Invoke(client *sdk.Client) (response *SetUserBlackListResponse, err error) {
	req := struct {
		*requests.RpcRequest
		SetUserBlackListRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Cdn", "2014-11-11", "SetUserBlackList", "", "")

	resp := struct {
		*responses.BaseResponse
		SetUserBlackListResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.SetUserBlackListResponse

	err = client.DoAction(&req, &resp)
	return
}

type SetUserBlackListResponse struct {
	RequestId string
}
