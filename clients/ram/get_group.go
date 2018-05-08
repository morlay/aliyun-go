package ram

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type GetGroupRequest struct {
	requests.RpcRequest
	GroupName string `position:"Query" name:"GroupName"`
}

func (req *GetGroupRequest) Invoke(client *sdk.Client) (resp *GetGroupResponse, err error) {
	req.InitWithApiInfo("Ram", "2015-05-01", "GetGroup", "", "")
	resp = &GetGroupResponse{}
	err = client.DoAction(req, resp)
	return
}

type GetGroupResponse struct {
	responses.BaseResponse
	RequestId common.String
	Group     GetGroupGroup
}

type GetGroupGroup struct {
	GroupName  common.String
	Comments   common.String
	CreateDate common.String
	UpdateDate common.String
}
