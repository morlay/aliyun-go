package cdn

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DeleteHttpHeaderConfigRequest struct {
	SecurityToken string `position:"Query" name:"SecurityToken"`
	ConfigID      string `position:"Query" name:"ConfigID"`
	DomainName    string `position:"Query" name:"DomainName"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
}

func (r DeleteHttpHeaderConfigRequest) Invoke(client *sdk.Client) (response *DeleteHttpHeaderConfigResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DeleteHttpHeaderConfigRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Cdn", "2014-11-11", "DeleteHttpHeaderConfig", "", "")

	resp := struct {
		*responses.BaseResponse
		DeleteHttpHeaderConfigResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DeleteHttpHeaderConfigResponse

	err = client.DoAction(&req, &resp)
	return
}

type DeleteHttpHeaderConfigResponse struct {
	RequestId string
}
