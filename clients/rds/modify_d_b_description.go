package rds

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ModifyDBDescriptionRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	DBName               string `position:"Query" name:"DBName"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
	DBDescription        string `position:"Query" name:"DBDescription"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (r ModifyDBDescriptionRequest) Invoke(client *sdk.Client) (response *ModifyDBDescriptionResponse, err error) {
	req := struct {
		*requests.RpcRequest
		ModifyDBDescriptionRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Rds", "2014-08-15", "ModifyDBDescription", "rds", "")

	resp := struct {
		*responses.BaseResponse
		ModifyDBDescriptionResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.ModifyDBDescriptionResponse

	err = client.DoAction(&req, &resp)
	return
}

type ModifyDBDescriptionResponse struct {
	RequestId string
}
