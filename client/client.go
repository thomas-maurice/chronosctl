package client

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/spf13/viper"
	"github.com/thomas-maurice/chronosctl/types"
	"io/ioutil"
	"net/http"
	"net/http/httputil"
	"strings"
)

type Client struct {
	Username string
	Password string
	URL      string
	Debug    bool
}

func NewClient(debug bool) *Client {
	var client Client
	client.Username = viper.GetString("username")
	client.Password = viper.GetString("password")
	client.URL = viper.GetString("url")
	client.Debug = debug
	return &client
}

func DumpRequest(request *http.Request) {
	dump, err := httputil.DumpRequestOut(request, true)
	if err != nil {
		fmt.Println("Could not dump request: %s", err)
		return
	}
	fmt.Printf("%s\n", bytes.NewBuffer(dump).String())
}

func DumpResponse(request *http.Response) {
	dump, err := httputil.DumpResponse(request, true)
	if err != nil {
		fmt.Println("Could not dump response: %s", err)
		return
	}
	fmt.Printf("%s\n", bytes.NewBuffer(dump).String())
}

func (self *Client) Get(resource string, resultObj interface{}, expectedCodes ...int) (*http.Response, error) {
	req, err := http.NewRequest("GET", self.URL+resource, nil)
	client := &http.Client{}

	if self.Username != "" && self.Password != "" {
		req.SetBasicAuth(self.Username, self.Password)
	}

	if self.Debug {
		DumpRequest(req)
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if self.Debug {
		DumpResponse(resp)
	}

	result, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resultObj != nil {
		err = json.Unmarshal(result, &resultObj)
		if err != nil {
			return nil, err
		}
	}

	for _, code := range expectedCodes {
		if code == resp.StatusCode {
			return resp, nil
		}
	}

	return resp, errors.New(fmt.Sprintf("Unexpected return code, got %s", resp.Status))
}

func (self *Client) Post(resource string, object interface{}, resultObj interface{}, errorObj interface{}, expectedCodes ...int) (*http.Response, error) {
	payload, err := json.Marshal(&object)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", self.URL+resource, bytes.NewBuffer(payload))
	client := &http.Client{}

	if self.Username != "" && self.Password != "" {
		req.SetBasicAuth(self.Username, self.Password)
	}

	req.Header.Add("Content-Type", "application/json")

	if self.Debug {
		DumpRequest(req)
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if self.Debug {
		DumpResponse(resp)
	}

	result, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	if resultObj != nil {
		err = json.Unmarshal(result, &resultObj)
		if err != nil {
			return nil, err
		}
	}

	if errorObj != nil {
		err = json.Unmarshal(result, &errorObj)
		if err != nil {
			return nil, err
		}
	}

	for _, code := range expectedCodes {
		if code == resp.StatusCode {
			return resp, nil
		}
	}

	return resp, errors.New(fmt.Sprintf("Unexpected return code, got %s", resp.Status))
}

func (self *Client) Put(resource string, object interface{}, resultObj interface{}, errorObj interface{}, expectedCodes ...int) (*http.Response, error) {
	payload, err := json.Marshal(&object)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("PUT", self.URL+resource, bytes.NewBuffer(payload))
	client := &http.Client{}

	if self.Username != "" && self.Password != "" {
		req.SetBasicAuth(self.Username, self.Password)
	}

	req.Header.Add("Content-Type", "application/json")

	if self.Debug {
		DumpRequest(req)
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if self.Debug {
		DumpResponse(resp)
	}

	result, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	if resultObj != nil {
		err = json.Unmarshal(result, &resultObj)
		if err != nil {
			return nil, err
		}
	}

	if errorObj != nil {
		err = json.Unmarshal(result, &errorObj)
		if err != nil {
			return nil, err
		}
	}

	for _, code := range expectedCodes {
		if code == resp.StatusCode {
			return resp, nil
		}
	}

	return resp, errors.New(fmt.Sprintf("Unexpected return code, got %s", resp.Status))
}

func (self *Client) Delete(resource string, errorObj interface{}, expectedCodes ...int) (*http.Response, error) {
	req, err := http.NewRequest("DELETE", self.URL+resource, nil)
	client := &http.Client{}

	if self.Username != "" && self.Password != "" {
		req.SetBasicAuth(self.Username, self.Password)
	}

	if self.Debug {
		DumpRequest(req)
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if err != nil {
		return nil, err
	}

	if self.Debug {
		DumpResponse(resp)
	}

	result, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	if errorObj != nil {
		err = json.Unmarshal(result, &errorObj)
		if err != nil {
			return nil, err
		}
	}

	for _, code := range expectedCodes {
		if code == resp.StatusCode {
			return resp, nil
		}
	}

	return resp, errors.New(fmt.Sprintf("Unexpected return code, got %s", resp.Status))
}

// This function exists because the /graph/csv endpoint does not return json, so
// the standard Get function cannot do much to retrieve data
func (self *Client) GetJobsStatus() ([]types.ChronosJobStatus, error) {
	var resultArray []types.ChronosJobStatus
	req, err := http.NewRequest("GET", self.URL+"/scheduler/graph/csv", nil)
	client := &http.Client{}

	if self.Username != "" && self.Password != "" {
		req.SetBasicAuth(self.Username, self.Password)
	}

	if self.Debug {
		DumpRequest(req)
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if self.Debug {
		DumpResponse(resp)
	}

	result, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, errors.New(fmt.Sprintf("Unexpected return code, got %s", resp.Status))
	}

	for _, line := range strings.Split(string(result), "\n") {
		fields := strings.Split(line, ",")
		if len(fields) != 4 {
			continue
		}
		resultArray = append(resultArray, types.ChronosJobStatus{
			Name:        fields[1],
			LastOutcome: fields[2],
			Status:      fields[3],
		})
	}

	return resultArray, nil
}
