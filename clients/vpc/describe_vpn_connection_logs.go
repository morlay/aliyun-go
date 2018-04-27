package vpc

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeVpnConnectionLogsRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	MinutePeriod         int    `position:"Query" name:"MinutePeriod"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	VpnConnectionId      string `position:"Query" name:"VpnConnectionId"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	PageSize             int    `position:"Query" name:"PageSize"`
	From                 int    `position:"Query" name:"From"`
	To                   int    `position:"Query" name:"To"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	PageNumber           int    `position:"Query" name:"PageNumber"`
}

func (req *DescribeVpnConnectionLogsRequest) Invoke(client *sdk.Client) (resp *DescribeVpnConnectionLogsResponse, err error) {
	req.InitWithApiInfo("Vpc", "2016-04-28", "DescribeVpnConnectionLogs", "vpc", "")
	resp = &DescribeVpnConnectionLogsResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeVpnConnectionLogsResponse struct {
	responses.BaseResponse
	RequestId   string
	Count       int
	IsCompleted bool
	PageNumber  int
	PageSize    int
	Data        DescribeVpnConnectionLogsDatumList
}

type DescribeVpnConnectionLogsDatumList []string

func (list *DescribeVpnConnectionLogsDatumList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]string)
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
