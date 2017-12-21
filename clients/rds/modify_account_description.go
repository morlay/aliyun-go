package rds

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ModifyAccountDescriptionRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	AccountName          string `position:"Query" name:"AccountName"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	AccountDescription   string `position:"Query" name:"AccountDescription"`
}

func (r ModifyAccountDescriptionRequest) Invoke(client *sdk.Client) (response *ModifyAccountDescriptionResponse, err error) {
	req := struct {
		*requests.RpcRequest
		ModifyAccountDescriptionRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Rds", "2014-08-15", "ModifyAccountDescription", "rds", "")

	resp := struct {
		*responses.BaseResponse
		ModifyAccountDescriptionResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.ModifyAccountDescriptionResponse

	err = client.DoAction(&req, &resp)
	return
}

type ModifyAccountDescriptionResponse struct {
	RequestId string
}
