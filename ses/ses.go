package ses

import (
	"bytes"
	"encoding/xml"
	"github.com/therahulprasad/goAws/awsAuth"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

type SuccessXmlResponse struct {
	XMLName   xml.Name `xml:"SendEmailResponse"`
	MessageId string   `xml:"SendEmailResult>MessageId"`
	RequestId string   `xml:"ResponseMetadata>RequestId"`
}

type ErrorXmlResponse struct {
	XMLName   xml.Name `xml:"ErrorResponse"`
	Type      string   `xml:"Error>Type"`
	Code      string   `xml:"Error>Code"`
	Message   string   `xml:"Error>Message"`
	RequestId string   `xml:"RequestId"`
}

func SendSingleMail(to, subject, body, from, toName, fromName, replyTo, replyToName string) (int, SuccessXmlResponse, ErrorXmlResponse, error) {
	successResponse := SuccessXmlResponse{}
	errorResponse := ErrorXmlResponse{}
	data := setData(to, subject, body, from, toName, fromName, replyTo, replyToName)

	dateHeader := time.Now().UTC().Format(http.TimeFormat)
	signature := awsAuth.AuthorizationString(dateHeader)

	r, _ := http.NewRequest("POST", Apiurl, bytes.NewBufferString(data.Encode()))
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	r.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))
	r.Header.Set("Date", dateHeader)
	r.Header.Add("X-Amzn-Authorization", signature)

	client := &http.Client{}
	resp, err := client.Do(r)

	if err != nil {
		return -1, successResponse, errorResponse, err
	} else {
		contents, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return resp.StatusCode, successResponse, errorResponse, err
		}

		err = xml.Unmarshal(contents, &successResponse)
		if err != nil {
			err = xml.Unmarshal(contents, &errorResponse)
			if err != nil {
				return resp.StatusCode, successResponse, errorResponse, err
			}
			return resp.StatusCode, successResponse, errorResponse, nil
		}

		return resp.StatusCode, successResponse, errorResponse, nil
	}

	// Code flow wont reach here
	return resp.StatusCode, successResponse, errorResponse, nil
}

func setData(to, subject, body, from, toName, fromName, replyTo, replyToName string) url.Values {
	// TODO: Check if email address is valid
	data := url.Values{}
	data.Set("AWSAccessKeyId", awsAuth.AccessKey)
	data.Add("Action", "SendEmail")
	if DryRun == false {
		if toName == "" {
			data.Add("Destination.ToAddresses.member.1", to)
		} else {
			data.Add("Destination.ToAddresses.member.1", toName+" <"+to+">")
		}
	} else {
		data.Add("Destination.ToAddresses.member.1", DryRunTo)
	}

	if replyTo != "" {
		if replyToName == "" {
			data.Add("ReplyToAddresses.member.1", replyTo)
		} else {
			data.Add("ReplyToAddresses.member.1", replyToName+" <"+replyTo+">")
		}
	}

	data.Add("Message.Body.Html.Data", body)
	data.Add("Message.Body.Html.Charset", "UTF-8")
	data.Add("Message.Subject.Data", subject)
	if fromName == "" {
		data.Add("Source", from)
	} else {
		data.Add("Source", fromName+" <"+from+">")
	}

	return data
}

/*
	1. Send using SES along with configurations
	2. Fetch data from database pass it to Sender
	3. Push data into SQS
	4. Pull data from SQS and pass it to Sender
*/
