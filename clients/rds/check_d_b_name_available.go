package rds

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type CheckDBNameAvailableRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	DBName               string `position:"Query" name:"DBName"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ClientToken          string `position:"Query" name:"ClientToken"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (r CheckDBNameAvailableRequest) Invoke(client *sdk.Client) (response *CheckDBNameAvailableResponse, err error) {
	req := struct {
		*requests.RpcRequest
		CheckDBNameAvailableRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Rds", "2014-08-15", "CheckDBNameAvailable", "rds", "")

	resp := struct {
		*responses.BaseResponse
		CheckDBNameAvailableResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.CheckDBNameAvailableResponse

	err = client.DoAction(&req, &resp)
	return
}

type CheckDBNameAvailableResponse struct {
	RequestId string
}
