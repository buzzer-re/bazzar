package abuse

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"errors"

	"github.com/aandersonl/bazzar/pkg/utils"
)

//GetLatestSamples updated in MalwareBazzar Database
func GetLatestSamples() Response {
	resp, err := http.PostForm(MALWARE_BAZZAR_API_URL, latestSamplesFormData)
	utils.PanicIfError(err)

	body, err := ioutil.ReadAll(resp.Body)
	utils.PanicIfError(err)

	response := Response{}

	err = json.Unmarshal(body, &response)
	utils.PanicIfError(err)

	return response
}

func GetSampleByTag(tag string) Response {
	sampleByTag.Set("tag", tag)

	resp, err := http.PostForm(MALWARE_BAZZAR_API_URL, sampleByTag)
	utils.PanicIfError(err)

	body, err := ioutil.ReadAll(resp.Body)
	utils.PanicIfError(err)

	response := Response{}
	err = json.Unmarshal(body, &response)

	utils.PanicIfError(err)

	return response
}

func QuerySampleInfo(sampleHash string) (string, SampleQuery) {
	getSampleInfoForm.Set("hash", sampleHash)
	resp, err := http.PostForm(MALWARE_BAZZAR_API_URL, getSampleInfoForm)
	utils.PanicIfError(err)
	body, err := ioutil.ReadAll(resp.Body)
	utils.PanicIfError(err)

	sampleQuery := SampleQuery{}

	err = json.Unmarshal(body, &sampleQuery)
	utils.PanicIfError(err)

	
	return string(body), sampleQuery
}

func GetSample(sampleHash string) ([]byte, error) {
	getSample.Set("sha256_hash", sampleHash)
	resp, err := http.PostForm(MALWARE_BAZZAR_API_URL, getSample)
	utils.PanicIfError(err)

	body, err := ioutil.ReadAll(resp.Body)
	utils.PanicIfError(err)

	queryResponse := QueryStatus{}
	err = json.Unmarshal(body, &queryResponse)

	if err == nil {
		return nil, errors.New(queryResponse.QueryStatus)
	}

	return body, nil
}

func QuerySignature(signature string, limit int) Response {
	querySampleSignature.Set("signature", signature)
	querySampleSignature.Set("limit", fmt.Sprint(limit))

	resp, err := http.PostForm(MALWARE_BAZZAR_API_URL, querySampleSignature)
	utils.PanicIfError(err)

	body, err := ioutil.ReadAll(resp.Body)

	utils.PanicIfError(err)

	response := Response{}
	err = json.Unmarshal(body, &response)
	utils.PanicIfError(err)

	return response
}

func QueryClamavSignature(signature string) Response {
	return Response{}
}
