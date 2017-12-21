package cdn

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type SetRemoveQueryStringConfigRequest struct {
	KeepOssArgs   string `position:"Query" name:"KeepOssArgs"`
	SecurityToken string `position:"Query" name:"SecurityToken"`
	DomainName    string `position:"Query" name:"DomainName"`
	AliRemoveArgs string `position:"Query" name:"AliRemoveArgs"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
}

func (r SetRemoveQueryStringConfigRequest) Invoke(client *sdk.Client) (response *SetRemoveQueryStringConfigResponse, err error) {
	req := struct {
		*requests.RpcRequest
		SetRemoveQueryStringConfigRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Cdn", "2014-11-11", "SetRemoveQueryStringConfig", "", "")

	resp := struct {
		*responses.BaseResponse
		SetRemoveQueryStringConfigResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.SetRemoveQueryStringConfigResponse

	err = client.DoAction(&req, &resp)
	return
}

type SetRemoveQueryStringConfigResponse struct {
	RequestId string
}
