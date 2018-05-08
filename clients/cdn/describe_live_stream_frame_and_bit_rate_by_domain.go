package cdn

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DescribeLiveStreamFrameAndBitRateByDomainRequest struct {
	requests.RpcRequest
	AppName       string `position:"Query" name:"AppName"`
	SecurityToken string `position:"Query" name:"SecurityToken"`
	DomainName    string `position:"Query" name:"DomainName"`
	PageSize      int64  `position:"Query" name:"PageSize"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
	PageNumber    int64  `position:"Query" name:"PageNumber"`
}

func (req *DescribeLiveStreamFrameAndBitRateByDomainRequest) Invoke(client *sdk.Client) (resp *DescribeLiveStreamFrameAndBitRateByDomainResponse, err error) {
	req.InitWithApiInfo("Cdn", "2014-11-11", "DescribeLiveStreamFrameAndBitRateByDomain", "", "")
	resp = &DescribeLiveStreamFrameAndBitRateByDomainResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeLiveStreamFrameAndBitRateByDomainResponse struct {
	responses.BaseResponse
	RequestId                common.String
	Count                    common.Long
	PageNumber               common.Long
	PageSize                 common.Long
	FrameRateAndBitRateInfos DescribeLiveStreamFrameAndBitRateByDomainFrameRateAndBitRateInfoList
}

type DescribeLiveStreamFrameAndBitRateByDomainFrameRateAndBitRateInfo struct {
	StreamUrl      common.String
	VideoFrameRate common.Float
	AudioFrameRate common.Float
	BitRate        common.Float
	Time           common.String
}

type DescribeLiveStreamFrameAndBitRateByDomainFrameRateAndBitRateInfoList []DescribeLiveStreamFrameAndBitRateByDomainFrameRateAndBitRateInfo

func (list *DescribeLiveStreamFrameAndBitRateByDomainFrameRateAndBitRateInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeLiveStreamFrameAndBitRateByDomainFrameRateAndBitRateInfo)
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
