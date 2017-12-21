package cloudapi

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeAppSecurityRequest struct {
	requests.RpcRequest
	AppId int64 `position:"Query" name:"AppId"`
}

func (req *DescribeAppSecurityRequest) Invoke(client *sdk.Client) (resp *DescribeAppSecurityResponse, err error) {
	req.InitWithApiInfo("CloudAPI", "2016-07-14", "DescribeAppSecurity", "apigateway", "")
	resp = &DescribeAppSecurityResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeAppSecurityResponse struct {
	responses.BaseResponse
	RequestId    string
	AppKey       string
	AppSecret    string
	CreatedTime  string
	ModifiedTime string
}
