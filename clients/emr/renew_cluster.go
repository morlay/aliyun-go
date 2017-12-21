package emr

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type RenewClusterRequest struct {
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	Id              string `position:"Query" name:"Id"`
	ECSIds          string `position:"Query" name:"ECSIds"`
	Period          string `position:"Query" name:"Period"`
}

func (r RenewClusterRequest) Invoke(client *sdk.Client) (response *RenewClusterResponse, err error) {
	req := struct {
		*requests.RpcRequest
		RenewClusterRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Emr", "2016-04-08", "RenewCluster", "", "")

	resp := struct {
		*responses.BaseResponse
		RenewClusterResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.RenewClusterResponse

	err = client.DoAction(&req, &resp)
	return
}

type RenewClusterResponse struct {
	RequestId string
	ClusterId string
}
