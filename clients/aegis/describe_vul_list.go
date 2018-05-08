package aegis

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DescribeVulListRequest struct {
	requests.RpcRequest
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	Product         string `position:"Query" name:"Product"`
	StatusList      string `position:"Query" name:"StatusList"`
	Level           string `position:"Query" name:"Level"`
	Resource        string `position:"Query" name:"Resource"`
	OrderBy         string `position:"Query" name:"OrderBy"`
	Dealed          string `position:"Query" name:"Dealed"`
	CurrentPage     int    `position:"Query" name:"CurrentPage"`
	Type            string `position:"Query" name:"Type"`
	LastTsEnd       int64  `position:"Query" name:"LastTsEnd"`
	BatchName       string `position:"Query" name:"BatchName"`
	PatchId         int64  `position:"Query" name:"PatchId"`
	AliasName       string `position:"Query" name:"AliasName"`
	SourceIp        string `position:"Query" name:"SourceIp"`
	Name            string `position:"Query" name:"Name"`
	PageSize        int    `position:"Query" name:"PageSize"`
	Lang            string `position:"Query" name:"Lang"`
	LastTsStart     int64  `position:"Query" name:"LastTsStart"`
	Necessity       string `position:"Query" name:"Necessity"`
	Uuids           string `position:"Query" name:"Uuids"`
	Direction       string `position:"Query" name:"Direction"`
}

func (req *DescribeVulListRequest) Invoke(client *sdk.Client) (resp *DescribeVulListResponse, err error) {
	req.InitWithApiInfo("aegis", "2016-11-11", "DescribeVulList", "vipaegis", "")
	resp = &DescribeVulListResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeVulListResponse struct {
	responses.BaseResponse
	RequestId   common.String
	Count       common.Integer
	PageSize    common.Integer
	CurrentPage common.Integer
	TotalCount  common.Integer
	VulRecords  DescribeVulListVulRecordList
}

type DescribeVulListVulRecordList []common.String

func (list *DescribeVulListVulRecordList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]common.String)
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
