package vpc

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DeleteRouteEntryRequest struct {
	ResourceOwnerId      int64                            `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string                           `position:"Query" name:"ResourceOwnerAccount"`
	DestinationCidrBlock string                           `position:"Query" name:"DestinationCidrBlock"`
	OwnerAccount         string                           `position:"Query" name:"OwnerAccount"`
	NextHopId            string                           `position:"Query" name:"NextHopId"`
	OwnerId              int64                            `position:"Query" name:"OwnerId"`
	NextHopLists         *DeleteRouteEntryNextHopListList `position:"Query" type:"Repeated" name:"NextHopList"`
	RouteTableId         string                           `position:"Query" name:"RouteTableId"`
}

func (r DeleteRouteEntryRequest) Invoke(client *sdk.Client) (response *DeleteRouteEntryResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DeleteRouteEntryRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Vpc", "2016-04-28", "DeleteRouteEntry", "vpc", "")

	resp := struct {
		*responses.BaseResponse
		DeleteRouteEntryResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DeleteRouteEntryResponse

	err = client.DoAction(&req, &resp)
	return
}

type DeleteRouteEntryNextHopList struct {
	NextHopType string `name:"NextHopType"`
	NextHopId   string `name:"NextHopId"`
}

type DeleteRouteEntryResponse struct {
	RequestId string
}

type DeleteRouteEntryNextHopListList []DeleteRouteEntryNextHopList

func (list *DeleteRouteEntryNextHopListList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DeleteRouteEntryNextHopList)
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
