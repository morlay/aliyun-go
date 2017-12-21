package emr

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ModifyClusterNameRequest struct {
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	Id              string `position:"Query" name:"Id"`
	Name            string `position:"Query" name:"Name"`
}

func (r ModifyClusterNameRequest) Invoke(client *sdk.Client) (response *ModifyClusterNameResponse, err error) {
	req := struct {
		*requests.RpcRequest
		ModifyClusterNameRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Emr", "2016-04-08", "ModifyClusterName", "", "")

	resp := struct {
		*responses.BaseResponse
		ModifyClusterNameResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.ModifyClusterNameResponse

	err = client.DoAction(&req, &resp)
	return
}

type ModifyClusterNameResponse struct {
	RequestId string
}
