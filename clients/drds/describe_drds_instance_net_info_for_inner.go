package drds

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeDrdsInstanceNetInfoForInnerRequest struct {
	DrdsInstanceId string `position:"Query" name:"DrdsInstanceId"`
}

func (r DescribeDrdsInstanceNetInfoForInnerRequest) Invoke(client *sdk.Client) (response *DescribeDrdsInstanceNetInfoForInnerResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DescribeDrdsInstanceNetInfoForInnerRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Drds", "2017-10-16", "DescribeDrdsInstanceNetInfoForInner", "", "")

	resp := struct {
		*responses.BaseResponse
		DescribeDrdsInstanceNetInfoForInnerResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DescribeDrdsInstanceNetInfoForInnerResponse

	err = client.DoAction(&req, &resp)
	return
}

type DescribeDrdsInstanceNetInfoForInnerResponse struct {
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
