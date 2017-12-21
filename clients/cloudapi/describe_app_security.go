package cloudapi

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeAppSecurityRequest struct {
	AppId int64 `position:"Query" name:"AppId"`
}

func (r DescribeAppSecurityRequest) Invoke(client *sdk.Client) (response *DescribeAppSecurityResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DescribeAppSecurityRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("CloudAPI", "2016-07-14", "DescribeAppSecurity", "apigateway", "")

	resp := struct {
		*responses.BaseResponse
		DescribeAppSecurityResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DescribeAppSecurityResponse

	err = client.DoAction(&req, &resp)
	return
}

type DescribeAppSecurityResponse struct {
	RequestId    string
	AppKey       string
	AppSecret    string
	CreatedTime  string
	ModifiedTime string
}
