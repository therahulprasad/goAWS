<h1>goSES</h1>

Golang library for Amazon Web Services - Simple Email Service (AWS - SES)

<h3>Installation</h3>
<pre>
go get github.com/therahulprasad/goAwsSES
</pre>

<h3>Usage</h3>
<pre>
/**
  * @Argument:
  *     STRING: to, subject, body, from, toName, fromName, replyTo, replyToName
  * @Returns
  *     http://docs.aws.amazon.com/ses/latest/DeveloperGuide/query-interface-responses.html
  *     INT StatusCode => Code return by SES 
  *     SuccessXmlResponse by SES (Struct defined in ses.go, it contains MessageId and RequestId)
  *     ErrorXmlResponse by SES (Struct defined in ses.go, it contains Type, Code, Message, RequestId)
  *     err => error while Unmarshaling, error during http request
  */
StatusCode, SuccessXmlResponse, ErrorXmlResponse, err := SendSingleMail(to, 
                                                                        subject, 
                                                                        body, 
                                                                        from, 
                                                                        toName, 
                                                                        fromName, 
                                                                        replyTo, 
                                                                        replyToName);

// If you dont worry about error code 
_, _, _, err := SendSingleMail(to, subject, body, from, toName, fromName, replyTo, replyToName);
if err != nil {
 // Email reached SES
}

// If you care about error
_, _, ErrorXmlResponse, err := SendSingleMail(to, subject, body, from, toName, fromName, replyTo, replyToName);
if err != nil && ErrorXmlResponse.Code == "" {
 // Email reached SES and There was no Error on SES
} else if err != nil {
 // Email reached SES by There was an error on ses
 fmt.Println(ErrorXmlResponse.Message)
}

</pre>

<h4>Note:</h4>
Volunteers needed, alot can be enhanced in this library.
<ol>
  <li>Need someone to write test cases</li>
  <li>Need help to add alot of features</li>
</ol>
