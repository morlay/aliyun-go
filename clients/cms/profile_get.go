package cms

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type ProfileGetRequest struct {
	requests.RpcRequest
	UserId int64 `position:"Query" name:"UserId"`
}

func (req *ProfileGetRequest) Invoke(client *sdk.Client) (resp *ProfileGetResponse, err error) {
	req.InitWithApiInfo("Cms", "2018-03-08", "ProfileGet", "cms", "")
	resp = &ProfileGetResponse{}
	err = client.DoAction(req, resp)
	return
}

type ProfileGetResponse struct {
	responses.BaseResponse
	ErrorCode                common.Integer
	ErrorMessage             common.String
	Success                  bool
	RequestId                common.String
	UserId                   common.Long
	AutoInstall              bool
	EnableInstallAgentNewECS bool
	EnableActiveAlert        common.String
}
