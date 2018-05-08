package emr

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type CreateMainVersionForAdminRequest struct {
	requests.RpcRequest
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	MainVersion     string `position:"Query" name:"MainVersion"`
	Content         string `position:"Query" name:"Content"`
}

func (req *CreateMainVersionForAdminRequest) Invoke(client *sdk.Client) (resp *CreateMainVersionForAdminResponse, err error) {
	req.InitWithApiInfo("Emr", "2016-04-08", "CreateMainVersionForAdmin", "", "")
	resp = &CreateMainVersionForAdminResponse{}
	err = client.DoAction(req, resp)
	return
}

type CreateMainVersionForAdminResponse struct {
	responses.BaseResponse
	RequestId common.String
}
