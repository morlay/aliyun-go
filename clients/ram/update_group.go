package ram

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type UpdateGroupRequest struct {
	NewGroupName string `position:"Query" name:"NewGroupName"`
	NewComments  string `position:"Query" name:"NewComments"`
	GroupName    string `position:"Query" name:"GroupName"`
}

func (r UpdateGroupRequest) Invoke(client *sdk.Client) (response *UpdateGroupResponse, err error) {
	req := struct {
		*requests.RpcRequest
		UpdateGroupRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Ram", "2015-05-01", "UpdateGroup", "", "")

	resp := struct {
		*responses.BaseResponse
		UpdateGroupResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.UpdateGroupResponse

	err = client.DoAction(&req, &resp)
	return
}

type UpdateGroupResponse struct {
	RequestId string
	Group     UpdateGroupGroup
}

type UpdateGroupGroup struct {
	GroupName  string
	Comments   string
	CreateDate string
	UpdateDate string
}
