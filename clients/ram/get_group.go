package ram

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type GetGroupRequest struct {
	GroupName string `position:"Query" name:"GroupName"`
}

func (r GetGroupRequest) Invoke(client *sdk.Client) (response *GetGroupResponse, err error) {
	req := struct {
		*requests.RpcRequest
		GetGroupRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Ram", "2015-05-01", "GetGroup", "", "")

	resp := struct {
		*responses.BaseResponse
		GetGroupResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.GetGroupResponse

	err = client.DoAction(&req, &resp)
	return
}

type GetGroupResponse struct {
	RequestId string
	Group     GetGroupGroup
}

type GetGroupGroup struct {
	GroupName  string
	Comments   string
	CreateDate string
	UpdateDate string
}
