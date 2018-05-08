package cms

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type CreateMyGroupsRequest struct {
	requests.RpcRequest
	ContactGroups string `position:"Query" name:"ContactGroups"`
	Options       string `position:"Query" name:"Options"`
	Type          string `position:"Query" name:"Type"`
	ServiceId     int64  `position:"Query" name:"ServiceId"`
	GroupName     string `position:"Query" name:"GroupName"`
	BindUrl       string `position:"Query" name:"BindUrl"`
}

func (req *CreateMyGroupsRequest) Invoke(client *sdk.Client) (resp *CreateMyGroupsResponse, err error) {
	req.InitWithApiInfo("Cms", "2018-03-08", "CreateMyGroups", "cms", "")
	resp = &CreateMyGroupsResponse{}
	err = client.DoAction(req, resp)
	return
}

type CreateMyGroupsResponse struct {
	responses.BaseResponse
	RequestId    common.String
	Success      bool
	ErrorCode    common.Integer
	ErrorMessage common.String
	GroupId      common.Long
}
