package cms

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type AccessKeyGetRequest struct {
	UserId int64 `position:"Query" name:"UserId"`
}

func (r AccessKeyGetRequest) Invoke(client *sdk.Client) (response *AccessKeyGetResponse, err error) {
	req := struct {
		*requests.RpcRequest
		AccessKeyGetRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Cms", "2017-03-01", "AccessKeyGet", "cms", "")

	resp := struct {
		*responses.BaseResponse
		AccessKeyGetResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.AccessKeyGetResponse

	err = client.DoAction(&req, &resp)
	return
}

type AccessKeyGetResponse struct {
	ErrorCode    int
	ErrorMessage string
	Success      bool
	RequestId    string
	UserId       int64
	AccessKey    string
	SecretKey    string
}
