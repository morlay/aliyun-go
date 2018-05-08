package sts

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type AssumeRoleRequest struct {
	requests.RpcRequest
	RoleArn         string `position:"Query" name:"RoleArn"`
	RoleSessionName string `position:"Query" name:"RoleSessionName"`
	DurationSeconds int64  `position:"Query" name:"DurationSeconds"`
	Policy          string `position:"Query" name:"Policy"`
}

func (req *AssumeRoleRequest) Invoke(client *sdk.Client) (resp *AssumeRoleResponse, err error) {
	req.InitWithApiInfo("Sts", "2015-04-01", "AssumeRole", "", "")
	resp = &AssumeRoleResponse{}
	err = client.DoAction(req, resp)
	return
}

type AssumeRoleResponse struct {
	responses.BaseResponse
	RequestId       common.String
	Credentials     AssumeRoleCredentials
	AssumedRoleUser AssumeRoleAssumedRoleUser
}

type AssumeRoleCredentials struct {
	SecurityToken   common.String
	AccessKeySecret common.String
	AccessKeyId     common.String
	Expiration      common.String
}

type AssumeRoleAssumedRoleUser struct {
	Arn           common.String
	AssumedRoleId common.String
}
