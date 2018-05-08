package emr

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
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
	RequestId common.String
	User      QueryUserByIdUser
}

type QueryUserByIdUser struct {
	Id              common.String
	AliyunId        common.String
	AliyunOmtId     common.String
	UserId          common.String
	Email           common.String
	Status          common.String
	DefaultSecGroup common.String
	RegionId        common.String
	ChannelId       common.String
}
