package cdn

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeLiveStreamFrameAndBitRateByDomainRequest struct {
	AppName       string `position:"Query" name:"AppName"`
	SecurityToken string `position:"Query" name:"SecurityToken"`
	DomainName    string `position:"Query" name:"DomainName"`
	PageSize      int64  `position:"Query" name:"PageSize"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
	PageNumber    int64  `position:"Query" name:"PageNumber"`
}

func (r DescribeLiveStreamFrameAndBitRateByDomainRequest) Invoke(client *sdk.Client) (response *DescribeLiveStreamFrameAndBitRateByDomainResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DescribeLiveStreamFrameAndBitRateByDomainRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Cdn", "2014-11-11", "DescribeLiveStreamFrameAndBitRateByDomain", "", "")

	resp := struct {
		*responses.BaseResponse
		DescribeLiveStreamFrameAndBitRateByDomainResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DescribeLiveStreamFrameAndBitRateByDomainResponse

	err = client.DoAction(&req, &resp)
	return
}

type DescribeLiveStreamFrameAndBitRateByDomainResponse struct {
	RequestId                string
	Count                    int64
	PageNumber               int64
	PageSize                 int64
	FrameRateAndBitRateInfos DescribeLiveStreamFrameAndBitRateByDomainFrameRateAndBitRateInfoList
}

type DescribeLiveStreamFrameAndBitRateByDomainFrameRateAndBitRateInfo struct {
	StreamUrl      string
	VideoFrameRate float32
	AudioFrameRate float32
	BitRate        float32
	Time           string
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
