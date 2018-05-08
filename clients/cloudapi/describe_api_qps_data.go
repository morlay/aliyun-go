package cloudapi

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DescribeApiQpsDataRequest struct {
	requests.RpcRequest
	ApiId     string `position:"Query" name:"ApiId"`
	GroupId   string `position:"Query" name:"GroupId"`
	StartTime string `position:"Query" name:"StartTime"`
	EndTime   string `position:"Query" name:"EndTime"`
}

func (req *DescribeApiQpsDataRequest) Invoke(client *sdk.Client) (resp *DescribeApiQpsDataResponse, err error) {
	req.InitWithApiInfo("CloudAPI", "2016-07-14", "DescribeApiQpsData", "apigateway", "")
	resp = &DescribeApiQpsDataResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeApiQpsDataResponse struct {
	responses.BaseResponse
	RequestId     common.String
	CallSuccesses DescribeApiQpsDataMonitorItemList
	CallFails     DescribeApiQpsDataMonitorItemList
}

type DescribeApiQpsDataMonitorItem struct {
	ItemTime  common.String
	ItemValue common.String
}

type DescribeApiQpsDataMonitorItemList []DescribeApiQpsDataMonitorItem

func (list *DescribeApiQpsDataMonitorItemList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeApiQpsDataMonitorItem)
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
