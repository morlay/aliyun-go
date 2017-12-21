package polardb

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ModifyParameterRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
	Forcerestart         string `position:"Query" name:"Forcerestart"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	Parameters           string `position:"Query" name:"Parameters"`
}

func (r ModifyParameterRequest) Invoke(client *sdk.Client) (response *ModifyParameterResponse, err error) {
	req := struct {
		*requests.RpcRequest
		ModifyParameterRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("polardb", "2017-08-01", "ModifyParameter", "polardb", "")

	resp := struct {
		*responses.BaseResponse
		ModifyParameterResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.ModifyParameterResponse

	err = client.DoAction(&req, &resp)
	return
}

type ModifyParameterResponse struct {
	RequestId string
}
