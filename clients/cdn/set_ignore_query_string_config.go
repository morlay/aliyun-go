package cdn

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type SetIgnoreQueryStringConfigRequest struct {
	KeepOssArgs   string `position:"Query" name:"KeepOssArgs"`
	HashKeyArgs   string `position:"Query" name:"HashKeyArgs"`
	SecurityToken string `position:"Query" name:"SecurityToken"`
	Enable        string `position:"Query" name:"Enable"`
	DomainName    string `position:"Query" name:"DomainName"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
}

func (r SetIgnoreQueryStringConfigRequest) Invoke(client *sdk.Client) (response *SetIgnoreQueryStringConfigResponse, err error) {
	req := struct {
		*requests.RpcRequest
		SetIgnoreQueryStringConfigRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Cdn", "2014-11-11", "SetIgnoreQueryStringConfig", "", "")

	resp := struct {
		*responses.BaseResponse
		SetIgnoreQueryStringConfigResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.SetIgnoreQueryStringConfigResponse

	err = client.DoAction(&req, &resp)
	return
}

type SetIgnoreQueryStringConfigResponse struct {
	RequestId string
}
