package cms

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type AccessKeyGetRequest struct {
	requests.RpcRequest
	UserId int64 `position:"Query" name:"UserId"`
}

func (req *AccessKeyGetRequest) Invoke(client *sdk.Client) (resp *AccessKeyGetResponse, err error) {
	req.InitWithApiInfo("Cms", "2017-03-01", "AccessKeyGet", "cms", "")
	resp = &AccessKeyGetResponse{}
	err = client.DoAction(req, resp)
	return
}

type AccessKeyGetResponse struct {
	responses.BaseResponse
	ErrorCode    int
	ErrorMessage string
	Success      bool
	RequestId    string
	UserId       int64
	AccessKey    string
	SecretKey    string
}
