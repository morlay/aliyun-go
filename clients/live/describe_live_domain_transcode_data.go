package live

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeLiveDomainTranscodeDataRequest struct {
	requests.RpcRequest
	DomainName string `position:"Query" name:"DomainName"`
	EndTime    string `position:"Query" name:"EndTime"`
	StartTime  string `position:"Query" name:"StartTime"`
	OwnerId    int64  `position:"Query" name:"OwnerId"`
}

func (req *DescribeLiveDomainTranscodeDataRequest) Invoke(client *sdk.Client) (resp *DescribeLiveDomainTranscodeDataResponse, err error) {
	req.InitWithApiInfo("live", "2016-11-01", "DescribeLiveDomainTranscodeData", "live", "")
	resp = &DescribeLiveDomainTranscodeDataResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeLiveDomainTranscodeDataResponse struct {
	responses.BaseResponse
	RequestId          string
	TranscodeDataInfos DescribeLiveDomainTranscodeDataTranscodeDataInfoList
}

type DescribeLiveDomainTranscodeDataTranscodeDataInfo struct {
	Date   string
	Total  int
	Detail string
}

type DescribeLiveDomainTranscodeDataTranscodeDataInfoList []DescribeLiveDomainTranscodeDataTranscodeDataInfo

func (list *DescribeLiveDomainTranscodeDataTranscodeDataInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeLiveDomainTranscodeDataTranscodeDataInfo)
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
