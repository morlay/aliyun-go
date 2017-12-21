package cdn

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeLiveStreamBpsDataRequest struct {
	requests.RpcRequest
	AppName       string `position:"Query" name:"AppName"`
	SecurityToken string `position:"Query" name:"SecurityToken"`
	DomainName    string `position:"Query" name:"DomainName"`
	EndTime       string `position:"Query" name:"EndTime"`
	StartTime     string `position:"Query" name:"StartTime"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
	StreamName    string `position:"Query" name:"StreamName"`
}

func (req *DescribeLiveStreamBpsDataRequest) Invoke(client *sdk.Client) (resp *DescribeLiveStreamBpsDataResponse, err error) {
	req.InitWithApiInfo("Cdn", "2014-11-11", "DescribeLiveStreamBpsData", "", "")
	resp = &DescribeLiveStreamBpsDataResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeLiveStreamBpsDataResponse struct {
	responses.BaseResponse
	RequestId string
	BpsDatas  DescribeLiveStreamBpsDataDomainBpsModelList
}

type DescribeLiveStreamBpsDataDomainBpsModel struct {
	Time string
	Bps  float32
}

type DescribeLiveStreamBpsDataDomainBpsModelList []DescribeLiveStreamBpsDataDomainBpsModel

func (list *DescribeLiveStreamBpsDataDomainBpsModelList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeLiveStreamBpsDataDomainBpsModel)
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
