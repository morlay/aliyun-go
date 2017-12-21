package ecs

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ModifyImageSharePermissionRequest struct {
	AddAccount1          string `position:"Query" name:"AddAccount.1"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ImageId              string `position:"Query" name:"ImageId"`
	AddAccount9          string `position:"Query" name:"AddAccount.9"`
	AddAccount8          string `position:"Query" name:"AddAccount.8"`
	AddAccount7          string `position:"Query" name:"AddAccount.7"`
	AddAccount6          string `position:"Query" name:"AddAccount.6"`
	AddAccount5          string `position:"Query" name:"AddAccount.5"`
	AddAccount10         string `position:"Query" name:"AddAccount.10"`
	AddAccount4          string `position:"Query" name:"AddAccount.4"`
	AddAccount3          string `position:"Query" name:"AddAccount.3"`
	AddAccount2          string `position:"Query" name:"AddAccount.2"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	RemoveAccount1       string `position:"Query" name:"RemoveAccount.1"`
	RemoveAccount2       string `position:"Query" name:"RemoveAccount.2"`
	RemoveAccount3       string `position:"Query" name:"RemoveAccount.3"`
	RemoveAccount4       string `position:"Query" name:"RemoveAccount.4"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RemoveAccount9       string `position:"Query" name:"RemoveAccount.9"`
	RemoveAccount5       string `position:"Query" name:"RemoveAccount.5"`
	RemoveAccount6       string `position:"Query" name:"RemoveAccount.6"`
	RemoveAccount7       string `position:"Query" name:"RemoveAccount.7"`
	RemoveAccount8       string `position:"Query" name:"RemoveAccount.8"`
	RemoveAccount10      string `position:"Query" name:"RemoveAccount.10"`
}

func (r ModifyImageSharePermissionRequest) Invoke(client *sdk.Client) (response *ModifyImageSharePermissionResponse, err error) {
	req := struct {
		*requests.RpcRequest
		ModifyImageSharePermissionRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Ecs", "2014-05-26", "ModifyImageSharePermission", "ecs", "")

	resp := struct {
		*responses.BaseResponse
		ModifyImageSharePermissionResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.ModifyImageSharePermissionResponse

	err = client.DoAction(&req, &resp)
	return
}

type ModifyImageSharePermissionResponse struct {
	RequestId string
}
