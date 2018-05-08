package cloudapi

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DeleteApiGroupRequest struct {
	requests.RpcRequest
	GroupId string `position:"Query" name:"GroupId"`
}

func (req *DeleteApiGroupRequest) Invoke(client *sdk.Client) (resp *DeleteApiGroupResponse, err error) {
	req.InitWithApiInfo("CloudAPI", "2016-07-14", "DeleteApiGroup", "apigateway", "")
	resp = &DeleteApiGroupResponse{}
	err = client.DoAction(req, resp)
	return
}

type DeleteApiGroupResponse struct {
	responses.BaseResponse
	RequestId common.String
}
