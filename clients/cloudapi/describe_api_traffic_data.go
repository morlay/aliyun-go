package cloudapi

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DescribeApiTrafficDataRequest struct {
	requests.RpcRequest
	ApiId     string `position:"Query" name:"ApiId"`
	GroupId   string `position:"Query" name:"GroupId"`
	StartTime string `position:"Query" name:"StartTime"`
	EndTime   string `position:"Query" name:"EndTime"`
}

func (req *DescribeApiTrafficDataRequest) Invoke(client *sdk.Client) (resp *DescribeApiTrafficDataResponse, err error) {
	req.InitWithApiInfo("CloudAPI", "2016-07-14", "DescribeApiTrafficData", "apigateway", "")
	resp = &DescribeApiTrafficDataResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeApiTrafficDataResponse struct {
	responses.BaseResponse
	RequestId     common.String
	CallUploads   DescribeApiTrafficDataMonitorItemList
	CallDownloads DescribeApiTrafficDataMonitorItemList
}

type DescribeApiTrafficDataMonitorItem struct {
	ItemTime  common.String
	ItemValue common.String
}

type DescribeApiTrafficDataMonitorItemList []DescribeApiTrafficDataMonitorItem

func (list *DescribeApiTrafficDataMonitorItemList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeApiTrafficDataMonitorItem)
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
