package vpc

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeVpnGatewaysRequest struct {
	requests.RpcRequest
	BusinessStatus       string `position:"Query" name:"BusinessStatus"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	VpcId                string `position:"Query" name:"VpcId"`
	PageSize             int    `position:"Query" name:"PageSize"`
	VpnGatewayId         string `position:"Query" name:"VpnGatewayId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	PageNumber           int    `position:"Query" name:"PageNumber"`
	Status               string `position:"Query" name:"Status"`
}

func (req *DescribeVpnGatewaysRequest) Invoke(client *sdk.Client) (resp *DescribeVpnGatewaysResponse, err error) {
	req.InitWithApiInfo("Vpc", "2016-04-28", "DescribeVpnGateways", "vpc", "")
	resp = &DescribeVpnGatewaysResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeVpnGatewaysResponse struct {
	responses.BaseResponse
	RequestId   string
	TotalCount  int
	PageNumber  int
	PageSize    int
	VpnGateways DescribeVpnGatewaysVpnGatewayList
}

type DescribeVpnGatewaysVpnGateway struct {
	VpnGatewayId   string
	VpcId          string
	VSwitchId      string
	InternetIp     string
	CreateTime     int64
	EndTime        int64
	Spec           string
	Name           string
	Description    string
	Status         string
	BusinessStatus string
	ChargeType     string
}

type DescribeVpnGatewaysVpnGatewayList []DescribeVpnGatewaysVpnGateway

func (list *DescribeVpnGatewaysVpnGatewayList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeVpnGatewaysVpnGateway)
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
