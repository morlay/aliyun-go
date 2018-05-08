package cdn

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DescribeLiveStreamRecordContentRequest struct {
	requests.RpcRequest
	AppName       string `position:"Query" name:"AppName"`
	SecurityToken string `position:"Query" name:"SecurityToken"`
	DomainName    string `position:"Query" name:"DomainName"`
	EndTime       string `position:"Query" name:"EndTime"`
	StartTime     string `position:"Query" name:"StartTime"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
	StreamName    string `position:"Query" name:"StreamName"`
}

func (req *DescribeLiveStreamRecordContentRequest) Invoke(client *sdk.Client) (resp *DescribeLiveStreamRecordContentResponse, err error) {
	req.InitWithApiInfo("Cdn", "2014-11-11", "DescribeLiveStreamRecordContent", "", "")
	resp = &DescribeLiveStreamRecordContentResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeLiveStreamRecordContentResponse struct {
	responses.BaseResponse
	RequestId             common.String
	RecordContentInfoList DescribeLiveStreamRecordContentRecordContentInfoList
}

type DescribeLiveStreamRecordContentRecordContentInfo struct {
	OssEndpoint     common.String
	OssBucket       common.String
	OssObjectPrefix common.String
	StartTime       common.String
	EndTime         common.String
	Duration        common.Float
}

type DescribeLiveStreamRecordContentRecordContentInfoList []DescribeLiveStreamRecordContentRecordContentInfo

func (list *DescribeLiveStreamRecordContentRecordContentInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeLiveStreamRecordContentRecordContentInfo)
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
