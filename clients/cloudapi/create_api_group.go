package cloudapi

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type CreateApiGroupRequest struct {
	requests.RpcRequest
	GroupName   string `position:"Query" name:"GroupName"`
	Description string `position:"Query" name:"Description"`
}

func (req *CreateApiGroupRequest) Invoke(client *sdk.Client) (resp *CreateApiGroupResponse, err error) {
	req.InitWithApiInfo("CloudAPI", "2016-07-14", "CreateApiGroup", "apigateway", "")
	resp = &CreateApiGroupResponse{}
	err = client.DoAction(req, resp)
	return
}

type CreateApiGroupResponse struct {
	responses.BaseResponse
	RequestId   common.String
	GroupId     common.String
	GroupName   common.String
	SubDomain   common.String
	Description common.String
}
