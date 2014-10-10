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
  *     SuccessXmlResponse by SES
  *     ErrorXmlResponse by SES
  *     err => error while Unmarshaling, error during http request
  */
StatusCode, SuccessXmlResponse, ErrorXmlResponse, err = SendSingleMail(to, subject, body, from, toName, fromName, replyTo, replyToName);
</pre>
