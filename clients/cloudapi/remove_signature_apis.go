package cloudapi

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type RemoveSignatureApisRequest struct {
	SignatureId string `position:"Query" name:"SignatureId"`
	GroupId     string `position:"Query" name:"GroupId"`
	ApiIds      string `position:"Query" name:"ApiIds"`
	StageName   string `position:"Query" name:"StageName"`
}

func (r RemoveSignatureApisRequest) Invoke(client *sdk.Client) (response *RemoveSignatureApisResponse, err error) {
	req := struct {
		*requests.RpcRequest
		RemoveSignatureApisRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("CloudAPI", "2016-07-14", "RemoveSignatureApis", "apigateway", "")

	resp := struct {
		*responses.BaseResponse
		RemoveSignatureApisResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.RemoveSignatureApisResponse

	err = client.DoAction(&req, &resp)
	return
}

type RemoveSignatureApisResponse struct {
	RequestId string
}
