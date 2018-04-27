package emr

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type CreateFlowProjectUserRequest struct {
	requests.RpcRequest
	UserAccountId string `position:"Query" name:"UserAccountId"`
	ProjectId     string `position:"Query" name:"ProjectId"`
	UserName      string `position:"Query" name:"UserName"`
}

func (req *CreateFlowProjectUserRequest) Invoke(client *sdk.Client) (resp *CreateFlowProjectUserResponse, err error) {
	req.InitWithApiInfo("Emr", "2016-04-08", "CreateFlowProjectUser", "", "")
	resp = &CreateFlowProjectUserResponse{}
	err = client.DoAction(req, resp)
	return
}

type CreateFlowProjectUserResponse struct {
	responses.BaseResponse
	RequestId string
	Data      bool
}
