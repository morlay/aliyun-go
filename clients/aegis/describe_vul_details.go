package aegis

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeVulDetailsRequest struct {
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

func (req *DescribeVulDetailsRequest) Invoke(client *sdk.Client) (resp *DescribeVulDetailsResponse, err error) {
	req.InitWithApiInfo("aegis", "2016-11-11", "DescribeVulDetails", "vipaegis", "")
	resp = &DescribeVulDetailsResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeVulDetailsResponse struct {
	responses.BaseResponse
	RequestId      string
	Name           string
	AliasName      string
	Level          string
	VulPublishTs   int64
	Type           string
	Product        string
	HasPatch       bool
	PatchPublishTs int64
	PatchSource    string
	Cvss           string
	CveIds         string
	Advice         string
	Description    string
	PendingCount   int
	HandledCount   int
	CveLists       DescribeVulDetailsCveListList
}

type DescribeVulDetailsCveListList []string

func (list *DescribeVulDetailsCveListList) UnmarshalJSON(data []byte) error {
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
