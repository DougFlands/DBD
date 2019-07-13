package fetch

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

var url = "https://used-api.jd.com/auction/list?pageSize=50&category1=&status=&orderDirection=1&orderType=1&groupId=1000005&callback=__jp4&pageNo="

func Fetchurl(page string) ([]byte, error) {
	resp, err := http.Get(url + page)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil,
			fmt.Errorf("wrong status code %d", resp.StatusCode)
	}

	return ioutil.ReadAll(resp.Body)
}
