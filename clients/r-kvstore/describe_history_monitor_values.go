
package r-kvstore

import (
    "github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
)


type DescribeHistoryMonitorValuesRequest struct {
ResourceOwnerId int64 `position:"Query" name:"ResourceOwnerId"`
ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
OwnerAccount string `position:"Query" name:"OwnerAccount"`
EndTime string `position:"Query" name:"EndTime"`
StartTime string `position:"Query" name:"StartTime"`
OwnerId int64 `position:"Query" name:"OwnerId"`
InstanceId string `position:"Query" name:"InstanceId"`
SecurityToken string `position:"Query" name:"SecurityToken"`
IntervalForHistory string `position:"Query" name:"IntervalForHistory"`
MonitorKeys string `position:"Query" name:"MonitorKeys"`
}

func (r DescribeHistoryMonitorValuesRequest) Invoke(client *sdk.Client) (response *DescribeHistoryMonitorValuesResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DescribeHistoryMonitorValuesRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("R-kvstore", "2015-01-01", "DescribeHistoryMonitorValues", "redisa", "")

	resp := struct {
		*responses.BaseResponse
		DescribeHistoryMonitorValuesResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
    response = &resp.DescribeHistoryMonitorValuesResponse

	err = client.DoAction(&req, &resp)
	return
}

type DescribeHistoryMonitorValuesResponse struct {
RequestId string
MonitorHistory string
}

