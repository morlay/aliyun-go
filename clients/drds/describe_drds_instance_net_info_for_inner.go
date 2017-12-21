package drds

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeDrdsInstanceNetInfoForInnerRequest struct {
	requests.RpcRequest
	DrdsInstanceId string `position:"Query" name:"DrdsInstanceId"`
}

func (req *DescribeDrdsInstanceNetInfoForInnerRequest) Invoke(client *sdk.Client) (resp *DescribeDrdsInstanceNetInfoForInnerResponse, err error) {
	req.InitWithApiInfo("Drds", "2017-10-16", "DescribeDrdsInstanceNetInfoForInner", "", "")
	resp = &DescribeDrdsInstanceNetInfoForInnerResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeDrdsInstanceNetInfoForInnerResponse struct {
	responses.BaseResponse
	RequestId      string
	Success        bool
	DrdsInstanceId string
	NetworkType    string
	NetInfos       DescribeDrdsInstanceNetInfoForInnerNetInfoList
}

type DescribeDrdsInstanceNetInfoForInnerNetInfo struct {
	IP       string
	Port     string
	Type     string
	IsForVpc bool
}

type DescribeDrdsInstanceNetInfoForInnerNetInfoList []DescribeDrdsInstanceNetInfoForInnerNetInfo

func (list *DescribeDrdsInstanceNetInfoForInnerNetInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDrdsInstanceNetInfoForInnerNetInfo)
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
