package cms

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ProfileSetRequest struct {
	requests.RpcRequest
	EnableInstallAgentNewECS string `position:"Query" name:"EnableInstallAgentNewECS"`
	EnableActiveAlert        string `position:"Query" name:"EnableActiveAlert"`
	AutoInstall              string `position:"Query" name:"AutoInstall"`
	UserId                   int64  `position:"Query" name:"UserId"`
}

func (req *ProfileSetRequest) Invoke(client *sdk.Client) (resp *ProfileSetResponse, err error) {
	req.InitWithApiInfo("Cms", "2018-03-08", "ProfileSet", "cms", "")
	resp = &ProfileSetResponse{}
	err = client.DoAction(req, resp)
	return
}

type ProfileSetResponse struct {
	responses.BaseResponse
	ErrorCode    int
	ErrorMessage string
	Success      bool
	RequestId    string
}
