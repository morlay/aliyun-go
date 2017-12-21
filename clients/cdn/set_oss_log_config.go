package cdn

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type SetOssLogConfigRequest struct {
	Bucket        string `position:"Query" name:"Bucket"`
	SecurityToken string `position:"Query" name:"SecurityToken"`
	Enable        string `position:"Query" name:"Enable"`
	Prefix        string `position:"Query" name:"Prefix"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
}

func (r SetOssLogConfigRequest) Invoke(client *sdk.Client) (response *SetOssLogConfigResponse, err error) {
	req := struct {
		*requests.RpcRequest
		SetOssLogConfigRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Cdn", "2014-11-11", "SetOssLogConfig", "", "")

	resp := struct {
		*responses.BaseResponse
		SetOssLogConfigResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.SetOssLogConfigResponse

	err = client.DoAction(&req, &resp)
	return
}

type SetOssLogConfigResponse struct {
	RequestId string
}
