package afs

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
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
	RequestId     common.String
	HasWarning    bool
	BizCode       common.String
	EarlyWarnings DescribeEarlyWarningEarlyWarningList
}

type DescribeEarlyWarningEarlyWarning struct {
	WarnOpen  bool
	Title     common.String
	Content   common.String
	Frequency common.String
	TimeOpen  bool
	TimeBegin common.String
	TimeEnd   common.String
	Channel   common.String
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
