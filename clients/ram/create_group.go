package ram

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type CreateGroupRequest struct {
	Comments  string `position:"Query" name:"Comments"`
	GroupName string `position:"Query" name:"GroupName"`
}

func (r CreateGroupRequest) Invoke(client *sdk.Client) (response *CreateGroupResponse, err error) {
	req := struct {
		*requests.RpcRequest
		CreateGroupRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Ram", "2015-05-01", "CreateGroup", "", "")

	resp := struct {
		*responses.BaseResponse
		CreateGroupResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.CreateGroupResponse

	err = client.DoAction(&req, &resp)
	return
}

type CreateGroupResponse struct {
	RequestId string
	Group     CreateGroupGroup
}

type CreateGroupGroup struct {
	GroupName  string
	Comments   string
	CreateDate string
}
