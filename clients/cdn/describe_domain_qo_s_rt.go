package cdn

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeDomainQoSRtRequest struct {
	requests.RpcRequest
	Ip            string `position:"Query" name:"Ip"`
	DomainName    string `position:"Query" name:"DomainName"`
	EndTime       string `position:"Query" name:"EndTime"`
	StartTime     string `position:"Query" name:"StartTime"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
	Version       string `position:"Query" name:"Version"`
	SecurityToken string `position:"Query" name:"SecurityToken"`
}

func (req *DescribeDomainQoSRtRequest) Invoke(client *sdk.Client) (resp *DescribeDomainQoSRtResponse, err error) {
	req.InitWithApiInfo("Cdn", "2014-11-11", "DescribeDomainQoSRt", "", "")
	resp = &DescribeDomainQoSRtResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeDomainQoSRtResponse struct {
	responses.BaseResponse
	DomainName string
	StartTime  string
	EndTime    string
	Ip         string
	Content    DescribeDomainQoSRtDataList
}

type DescribeDomainQoSRtData struct {
	More5s string
	Time   string
	More3s string
}

type DescribeDomainQoSRtDataList []DescribeDomainQoSRtData

func (list *DescribeDomainQoSRtDataList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDomainQoSRtData)
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
