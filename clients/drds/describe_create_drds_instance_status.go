package drds

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DescribeCreateDrdsInstanceStatusRequest struct {
	requests.RpcRequest
	DrdsInstanceId string `position:"Query" name:"DrdsInstanceId"`
}

func (req *DescribeCreateDrdsInstanceStatusRequest) Invoke(client *sdk.Client) (resp *DescribeCreateDrdsInstanceStatusResponse, err error) {
	req.InitWithApiInfo("Drds", "2017-10-16", "DescribeCreateDrdsInstanceStatus", "", "")
	resp = &DescribeCreateDrdsInstanceStatusResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeCreateDrdsInstanceStatusResponse struct {
	responses.BaseResponse
	RequestId common.String
	Success   bool
	Data      DescribeCreateDrdsInstanceStatusData
}

type DescribeCreateDrdsInstanceStatusData struct {
	Status common.String
}
