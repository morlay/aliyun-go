package drds

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeCreateDrdsInstanceStatusRequest struct {
	DrdsInstanceId string `position:"Query" name:"DrdsInstanceId"`
}

func (r DescribeCreateDrdsInstanceStatusRequest) Invoke(client *sdk.Client) (response *DescribeCreateDrdsInstanceStatusResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DescribeCreateDrdsInstanceStatusRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Drds", "2017-10-16", "DescribeCreateDrdsInstanceStatus", "", "")

	resp := struct {
		*responses.BaseResponse
		DescribeCreateDrdsInstanceStatusResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DescribeCreateDrdsInstanceStatusResponse

	err = client.DoAction(&req, &resp)
	return
}

type DescribeCreateDrdsInstanceStatusResponse struct {
	RequestId string
	Success   bool
	Data      DescribeCreateDrdsInstanceStatusData
}

type DescribeCreateDrdsInstanceStatusData struct {
	Status string
}
