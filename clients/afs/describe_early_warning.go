package afs

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeEarlyWarningRequest struct {
	requests.RpcRequest
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	SourceIp        string `position:"Query" name:"SourceIp"`
}

func (req *DescribeEarlyWarningRequest) Invoke(client *sdk.Client) (resp *DescribeEarlyWarningResponse, err error) {
	req.InitWithApiInfo("afs", "2018-01-12", "DescribeEarlyWarning", "", "")
	resp = &DescribeEarlyWarningResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeEarlyWarningResponse struct {
	responses.BaseResponse
	RequestId     string
	HasWarning    bool
	BizCode       string
	EarlyWarnings DescribeEarlyWarningEarlyWarningList
}

type DescribeEarlyWarningEarlyWarning struct {
	WarnOpen  bool
	Title     string
	Content   string
	Frequency string
	TimeOpen  bool
	TimeBegin string
	TimeEnd   string
	Channel   string
}

type DescribeEarlyWarningEarlyWarningList []DescribeEarlyWarningEarlyWarning

func (list *DescribeEarlyWarningEarlyWarningList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeEarlyWarningEarlyWarning)
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
