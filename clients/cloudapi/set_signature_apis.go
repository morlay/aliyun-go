package cloudapi

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type SetSignatureApisRequest struct {
	SignatureId string `position:"Query" name:"SignatureId"`
	GroupId     string `position:"Query" name:"GroupId"`
	ApiIds      string `position:"Query" name:"ApiIds"`
	StageName   string `position:"Query" name:"StageName"`
}

func (r SetSignatureApisRequest) Invoke(client *sdk.Client) (response *SetSignatureApisResponse, err error) {
	req := struct {
		*requests.RpcRequest
		SetSignatureApisRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("CloudAPI", "2016-07-14", "SetSignatureApis", "apigateway", "")

	resp := struct {
		*responses.BaseResponse
		SetSignatureApisResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.SetSignatureApisResponse

	err = client.DoAction(&req, &resp)
	return
}

type SetSignatureApisResponse struct {
	RequestId string
}
