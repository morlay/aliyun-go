package cms

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type AccessKeyGetRequest struct {
	requests.RpcRequest
	UserId int64 `position:"Query" name:"UserId"`
}

func (req *AccessKeyGetRequest) Invoke(client *sdk.Client) (resp *AccessKeyGetResponse, err error) {
	req.InitWithApiInfo("Cms", "2018-03-08", "AccessKeyGet", "cms", "")
	resp = &AccessKeyGetResponse{}
	err = client.DoAction(req, resp)
	return
}

type AccessKeyGetResponse struct {
	responses.BaseResponse
	ErrorCode    common.Integer
	ErrorMessage common.String
	Success      bool
	RequestId    common.String
	UserId       common.Long
	AccessKey    common.String
	SecretKey    common.String
}
