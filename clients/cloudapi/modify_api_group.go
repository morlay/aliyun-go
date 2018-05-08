package cloudapi

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type ModifyApiGroupRequest struct {
	requests.RpcRequest
	GroupId     string `position:"Query" name:"GroupId"`
	GroupName   string `position:"Query" name:"GroupName"`
	Description string `position:"Query" name:"Description"`
}

func (req *ModifyApiGroupRequest) Invoke(client *sdk.Client) (resp *ModifyApiGroupResponse, err error) {
	req.InitWithApiInfo("CloudAPI", "2016-07-14", "ModifyApiGroup", "apigateway", "")
	resp = &ModifyApiGroupResponse{}
	err = client.DoAction(req, resp)
	return
}

type ModifyApiGroupResponse struct {
	responses.BaseResponse
	RequestId   common.String
	GroupId     common.String
	GroupName   common.String
	SubDomain   common.String
	Description common.String
}
