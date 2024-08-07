package mail

import (
	"testing"

	"github.com/agolosnichenko/golang-simplebank/simplebank/util"
	"github.com/stretchr/testify/require"
)

func TestSendEmailWithGmail(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	config, err := util.LoadConfig("..")
	require.NoError(t, err)

	sender := NewGmailSender(config.EmailSenderName, config.EmailSenderAddress, config.EmailSenderPassword)

	subject := "A test email"
	content := `
		<h1>Hello world</h1>
		<p>This is a test email from <b>SimpleBank</b>.</p>
	`
	to := []string{"a.golosnichenko@icloud.com"}
	attachFiles := []string{"../../README.md"}

	err = sender.SendEmail(subject, content, to, nil, nil, attachFiles)
	require.NoError(t, err)
}
