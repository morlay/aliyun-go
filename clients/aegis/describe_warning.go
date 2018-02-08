package aegis

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeWarningRequest struct {
	requests.RpcRequest
	TypeNames       string `position:"Query" name:"TypeNames"`
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	RiskName        string `position:"Query" name:"RiskName"`
	StatusList      string `position:"Query" name:"StatusList"`
	SourceIp        string `position:"Query" name:"SourceIp"`
	RiskLevels      string `position:"Query" name:"RiskLevels"`
	PageSize        int    `position:"Query" name:"PageSize"`
	CurrentPage     int    `position:"Query" name:"CurrentPage"`
	Dealed          string `position:"Query" name:"Dealed"`
	SubTypeNames    string `position:"Query" name:"SubTypeNames"`
	Uuids           string `position:"Query" name:"Uuids"`
}

func (req *DescribeWarningRequest) Invoke(client *sdk.Client) (resp *DescribeWarningResponse, err error) {
	req.InitWithApiInfo("aegis", "2016-11-11", "DescribeWarning", "vipaegis", "")
	resp = &DescribeWarningResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeWarningResponse struct {
	responses.BaseResponse
	RequestId   string
	Count       int
	PageSize    int
	TotalCount  int
	CurrentPage int
	Warnings    DescribeWarningWarningList
}

type DescribeWarningWarningList []string

func (list *DescribeWarningWarningList) UnmarshalJSON(data []byte) error {
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
