package vpc

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeBandwidthPackagePublicIpMonitorDataRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	Period               int    `position:"Query" name:"Period"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	EndTime              string `position:"Query" name:"EndTime"`
	AllocationId         string `position:"Query" name:"AllocationId"`
	StartTime            string `position:"Query" name:"StartTime"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (r DescribeBandwidthPackagePublicIpMonitorDataRequest) Invoke(client *sdk.Client) (response *DescribeBandwidthPackagePublicIpMonitorDataResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DescribeBandwidthPackagePublicIpMonitorDataRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Vpc", "2016-04-28", "DescribeBandwidthPackagePublicIpMonitorData", "vpc", "")

	resp := struct {
		*responses.BaseResponse
		DescribeBandwidthPackagePublicIpMonitorDataResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DescribeBandwidthPackagePublicIpMonitorDataResponse

	err = client.DoAction(&req, &resp)
	return
}

type DescribeBandwidthPackagePublicIpMonitorDataResponse struct {
	RequestId    string
	MonitorDatas DescribeBandwidthPackagePublicIpMonitorDataMonitorDataList
}

type DescribeBandwidthPackagePublicIpMonitorDataMonitorData struct {
	RX                   int64
	TX                   int64
	ReceivedBandwidth    int64
	TransportedBandwidth int64
	Flow                 int64
	Bandwidth            int64
	Packets              int64
	TimeStamp            string
}

type DescribeBandwidthPackagePublicIpMonitorDataMonitorDataList []DescribeBandwidthPackagePublicIpMonitorDataMonitorData

func (list *DescribeBandwidthPackagePublicIpMonitorDataMonitorDataList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeBandwidthPackagePublicIpMonitorDataMonitorData)
	err := json.Unmarshal(data, &m)
	if err != nil {
		return err
	}
	for _, v := range m {
		*list = v
		break
	}
	return nil
}
