package OtxApiClient

import (
	"encoding/json"
	"net/http"
	"path"
)

type indicatorsService struct {
	client *Client
}

type IndicatorType interface {
	GetIndicatorName() string
}

func (i *indicatorsService) GetIndicators(indicatorType IndicatorType, subject, section string) error {
	baseUrl := path.Join("/indicators", indicatorType.GetIndicatorName(), subject, section)
	resp, err := i.client.doHttpRequest(http.MethodGet, baseUrl, nil)
	if err != nil {
		return err
	}
	return json.NewDecoder(resp.Body).Decode(indicatorType)
}
