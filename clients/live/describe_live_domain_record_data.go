package live

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeLiveDomainRecordDataRequest struct {
	requests.RpcRequest
	RecordType string `position:"Query" name:"RecordType"`
	DomainName string `position:"Query" name:"DomainName"`
	EndTime    string `position:"Query" name:"EndTime"`
	StartTime  string `position:"Query" name:"StartTime"`
	OwnerId    int64  `position:"Query" name:"OwnerId"`
}

func (req *DescribeLiveDomainRecordDataRequest) Invoke(client *sdk.Client) (resp *DescribeLiveDomainRecordDataResponse, err error) {
	req.InitWithApiInfo("live", "2016-11-01", "DescribeLiveDomainRecordData", "live", "")
	resp = &DescribeLiveDomainRecordDataResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeLiveDomainRecordDataResponse struct {
	responses.BaseResponse
	RequestId       string
	RecordDataInfos DescribeLiveDomainRecordDataRecordDataInfoList
}

type DescribeLiveDomainRecordDataRecordDataInfo struct {
	Date   string
	Total  int
	Detail DescribeLiveDomainRecordDataDetail
}

type DescribeLiveDomainRecordDataDetail struct {
	MP4 int
	FLV int
	TS  int
}

type DescribeLiveDomainRecordDataRecordDataInfoList []DescribeLiveDomainRecordDataRecordDataInfo

func (list *DescribeLiveDomainRecordDataRecordDataInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeLiveDomainRecordDataRecordDataInfo)
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
