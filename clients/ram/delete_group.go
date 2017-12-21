package ram

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DeleteGroupRequest struct {
	GroupName string `position:"Query" name:"GroupName"`
}

func (r DeleteGroupRequest) Invoke(client *sdk.Client) (response *DeleteGroupResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DeleteGroupRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Ram", "2015-05-01", "DeleteGroup", "", "")

	resp := struct {
		*responses.BaseResponse
		DeleteGroupResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DeleteGroupResponse

	err = client.DoAction(&req, &resp)
	return
}

type DeleteGroupResponse struct {
	RequestId string
}
