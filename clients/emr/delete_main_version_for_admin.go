package emr

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DeleteMainVersionForAdminRequest struct {
	requests.RpcRequest
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	MainVersion     string `position:"Query" name:"MainVersion"`
}

func (req *DeleteMainVersionForAdminRequest) Invoke(client *sdk.Client) (resp *DeleteMainVersionForAdminResponse, err error) {
	req.InitWithApiInfo("Emr", "2016-04-08", "DeleteMainVersionForAdmin", "", "")
	resp = &DeleteMainVersionForAdminResponse{}
	err = client.DoAction(req, resp)
	return
}

type DeleteMainVersionForAdminResponse struct {
	responses.BaseResponse
	RequestId string
}
