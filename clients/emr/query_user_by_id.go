package emr

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type QueryUserByIdRequest struct {
	requests.RpcRequest
	ResourceOwnerId int64 `position:"Query" name:"ResourceOwnerId"`
}

func (req *QueryUserByIdRequest) Invoke(client *sdk.Client) (resp *QueryUserByIdResponse, err error) {
	req.InitWithApiInfo("Emr", "2016-04-08", "QueryUserById", "", "")
	resp = &QueryUserByIdResponse{}
	err = client.DoAction(req, resp)
	return
}

type QueryUserByIdResponse struct {
	responses.BaseResponse
	RequestId string
	User      QueryUserByIdUser
}

type QueryUserByIdUser struct {
	Id              string
	AliyunId        string
	AliyunOmtId     string
	UserId          string
	Email           string
	Status          string
	DefaultSecGroup string
	RegionId        string
	ChannelId       string
}
