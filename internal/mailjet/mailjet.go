package mailjet

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type SendRequest struct {
	Messages []Messages
}

type Messages struct {
	From     MailAccount
	To       []MailAccount
	Subject  string
	TextPart string
	HtmlPart string
}

type MailAccount struct {
	Email string
	Name  string
}

type Mailjet struct {
	publicKey         string
	privateKey        string
	defaultSender     MailAccount
	defaultRecipients []MailAccount
	baseUri           string
}

func NewMailjet(publicKey string, privateKey string) Mailjet {
	return Mailjet{
		publicKey:  publicKey,
		privateKey: privateKey,
		defaultSender: MailAccount{
			Email: "hb-contact@willowispsoftware.com",
			Name:  "Herbal Bones Site",
		},
		defaultRecipients: []MailAccount{
			{
				Email: "sami@herbalbones.com",
				Name:  "Sami",
			},
			{
				Email: "will@wrhensel.com",
				Name:  "Will",
			},
		},
		baseUri: "https://api.mailjet.com/v3.1",
	}
}

func (mj *Mailjet) makePostRequest(route string, body []byte) ([]byte, error) {
	client := &http.Client{}

	reqBodyReader := bytes.NewReader(body)

	req, err := http.NewRequest("POST", route, reqBodyReader)
	if err != nil {
		return []byte{}, err
	}

	authData := fmt.Sprintf("%s:%s", mj.publicKey, mj.privateKey)
	authString := base64.StdEncoding.EncodeToString([]byte(authData))

	req.Header.Add("Authorization", "Basic "+authString)
	req.Header.Add("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return []byte{}, err
	}
	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return []byte{}, err
	}

	if resp.StatusCode != http.StatusOK {
		return []byte{}, fmt.Errorf("request failed\n\n%s", string(bodyBytes))
	}

	return bodyBytes, nil
}

func (mj *Mailjet) SendContactMessage(name string, email string, message string) error {

	body := SendRequest{
		Messages: []Messages{
			{
				To:       mj.defaultRecipients,
				From:     mj.defaultSender,
				Subject:  fmt.Sprintf("Herbal Bones Site Contact Form - New message from %s", name),
				TextPart: fmt.Sprintf("You received the following message from\n\tName:%s\n\tEmail:%s\n\nMessage:\n%s", name, email, message),
				HtmlPart: fmt.Sprintf("<p>You received the following message from<br><span style='margin-left: 3rem;'>Name: <b>%s</b></span><br><span style='margin-left: 3rem;'>Email: <b>%s</b><span/><br><br>%s</p>", name, email, message),
			},
		},
	}
	bodyJson, err := json.Marshal(body)
	if err != nil {
		return err
	}

	route, err := url.JoinPath(mj.baseUri, "send")
	if err != nil {
		return err
	}

	mj.makePostRequest(route, bodyJson)

	return nil
}
