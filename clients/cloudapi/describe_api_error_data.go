package cloudapi

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DescribeApiErrorDataRequest struct {
	requests.RpcRequest
	ApiId     string `position:"Query" name:"ApiId"`
	GroupId   string `position:"Query" name:"GroupId"`
	StartTime string `position:"Query" name:"StartTime"`
	EndTime   string `position:"Query" name:"EndTime"`
}

func (req *DescribeApiErrorDataRequest) Invoke(client *sdk.Client) (resp *DescribeApiErrorDataResponse, err error) {
	req.InitWithApiInfo("CloudAPI", "2016-07-14", "DescribeApiErrorData", "apigateway", "")
	resp = &DescribeApiErrorDataResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeApiErrorDataResponse struct {
	responses.BaseResponse
	RequestId    common.String
	ClientErrors DescribeApiErrorDataMonitorItemList
	ServerErrors DescribeApiErrorDataMonitorItemList
}

type DescribeApiErrorDataMonitorItem struct {
	ItemTime  common.String
	ItemValue common.String
}

type DescribeApiErrorDataMonitorItemList []DescribeApiErrorDataMonitorItem

func (list *DescribeApiErrorDataMonitorItemList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeApiErrorDataMonitorItem)
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
