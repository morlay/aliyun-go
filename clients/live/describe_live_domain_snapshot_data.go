package live

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DescribeLiveDomainSnapshotDataRequest struct {
	requests.RpcRequest
	DomainName string `position:"Query" name:"DomainName"`
	EndTime    string `position:"Query" name:"EndTime"`
	StartTime  string `position:"Query" name:"StartTime"`
	OwnerId    int64  `position:"Query" name:"OwnerId"`
}

func (req *DescribeLiveDomainSnapshotDataRequest) Invoke(client *sdk.Client) (resp *DescribeLiveDomainSnapshotDataResponse, err error) {
	req.InitWithApiInfo("live", "2016-11-01", "DescribeLiveDomainSnapshotData", "live", "")
	resp = &DescribeLiveDomainSnapshotDataResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeLiveDomainSnapshotDataResponse struct {
	responses.BaseResponse
	RequestId         common.String
	SnapshotDataInfos DescribeLiveDomainSnapshotDataSnapshotDataInfoList
}

type DescribeLiveDomainSnapshotDataSnapshotDataInfo struct {
	Date  common.String
	Total common.Integer
}

type DescribeLiveDomainSnapshotDataSnapshotDataInfoList []DescribeLiveDomainSnapshotDataSnapshotDataInfo

func (list *DescribeLiveDomainSnapshotDataSnapshotDataInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeLiveDomainSnapshotDataSnapshotDataInfo)
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
