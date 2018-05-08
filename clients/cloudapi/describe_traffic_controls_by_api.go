package cloudapi

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DescribeTrafficControlsByApiRequest struct {
	requests.RpcRequest
	GroupId   string `position:"Query" name:"GroupId"`
	ApiId     string `position:"Query" name:"ApiId"`
	StageName string `position:"Query" name:"StageName"`
}

func (req *DescribeTrafficControlsByApiRequest) Invoke(client *sdk.Client) (resp *DescribeTrafficControlsByApiResponse, err error) {
	req.InitWithApiInfo("CloudAPI", "2016-07-14", "DescribeTrafficControlsByApi", "apigateway", "")
	resp = &DescribeTrafficControlsByApiResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeTrafficControlsByApiResponse struct {
	responses.BaseResponse
	RequestId           common.String
	TrafficControlItems DescribeTrafficControlsByApiTrafficControlItemList
}

type DescribeTrafficControlsByApiTrafficControlItem struct {
	TrafficControlItemId   common.String
	TrafficControlItemName common.String
	BoundTime              common.String
}

type DescribeTrafficControlsByApiTrafficControlItemList []DescribeTrafficControlsByApiTrafficControlItem

func (list *DescribeTrafficControlsByApiTrafficControlItemList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeTrafficControlsByApiTrafficControlItem)
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
