package cdn

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type SetL2OssKeyConfigRequest struct {
	SecurityToken  string `position:"Query" name:"SecurityToken"`
	DomainName     string `position:"Query" name:"DomainName"`
	OwnerId        int64  `position:"Query" name:"OwnerId"`
	PrivateOssAuth string `position:"Query" name:"PrivateOssAuth"`
}

func (r SetL2OssKeyConfigRequest) Invoke(client *sdk.Client) (response *SetL2OssKeyConfigResponse, err error) {
	req := struct {
		*requests.RpcRequest
		SetL2OssKeyConfigRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Cdn", "2014-11-11", "SetL2OssKeyConfig", "", "")

	resp := struct {
		*responses.BaseResponse
		SetL2OssKeyConfigResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.SetL2OssKeyConfigResponse

	err = client.DoAction(&req, &resp)
	return
}

type SetL2OssKeyConfigResponse struct {
	RequestId string
}
