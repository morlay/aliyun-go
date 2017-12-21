package drds

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeReadOnlyAccountRequest struct {
	DbName         string `position:"Query" name:"DbName"`
	DrdsInstanceId string `position:"Query" name:"DrdsInstanceId"`
}

func (r DescribeReadOnlyAccountRequest) Invoke(client *sdk.Client) (response *DescribeReadOnlyAccountResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DescribeReadOnlyAccountRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Drds", "2017-10-16", "DescribeReadOnlyAccount", "", "")

	resp := struct {
		*responses.BaseResponse
		DescribeReadOnlyAccountResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DescribeReadOnlyAccountResponse

	err = client.DoAction(&req, &resp)
	return
}

type DescribeReadOnlyAccountResponse struct {
	RequestId string
	Success   bool
	Data      DescribeReadOnlyAccountData
}

type DescribeReadOnlyAccountData struct {
	DbName         string
	DrdsInstanceId string
	AccountName    string
}
