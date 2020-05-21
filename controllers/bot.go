package controllers

import (
	"encoding/gob"
	"fmt"
	qrcodeTerminal "github.com/Baozisoftware/qrcode-terminal-go"
	"os"
	"strings"
	"time"
	"wa-chattbot/lib"
	"wa-chattbot/models"
)

type waHandler struct {
	wac       *lib.Conn
	startTime uint64
}

func (wh *waHandler) HandleError(err error) {
	fmt.Fprintf(os.Stderr, "error caught in handler: %v\n", err)
}

func (wh *waHandler) HandleTextMessage(message lib.TextMessage) {


	if !strings.Contains(strings.ToLower(message.Text), "covid: today") || message.Info.Timestamp < wh.startTime {
		return
	}

	lampung := models.Response{}
	test := lampung.GetData()

	msg := lib.TextMessage{
		Info: lib.MessageInfo{
			RemoteJid: message.Info.RemoteJid,
		},
		Text: test,
	}

	if _, err := wh.wac.Send(msg); err != nil {
		fmt.Fprintf(os.Stderr, "error sending message: %v\n", err)
	}

	fmt.Printf("echoed message '%v' to user %v\n", test, message.Info.RemoteJid)
}

func login(wac *lib.Conn) error {
	session, err := readSession()
	if err == nil {
		session, err = wac.RestoreWithSession(session)
		if err != nil {
			return fmt.Errorf("restoring session failed: %v", err)
		}
	} else {
		qr := make(chan string)

		go func() {
			terminal := qrcodeTerminal.New()
			terminal.Get(<-qr).Print()
		}()

		session, err = wac.Login(qr)
		if err != nil {
			return fmt.Errorf("error during login: %v", err)
		}
	}

	if err = writeSession(session); err != nil {
		return fmt.Errorf("error saving session: %v", err)
	}

	return nil
}

func readSession() (lib.Session, error) {
	session := lib.Session{}

	file, err := os.Open(os.TempDir() + "/whatsappSession.gob")
	if err != nil {
		return session, err
	}
	defer file.Close()

	decoder := gob.NewDecoder(file)
	if err = decoder.Decode(&session); err != nil {
		return session, err
	}

	return session, nil
}

func writeSession(session lib.Session) error {
	file, err := os.Create(os.TempDir() + "/whatsappSession.gob")
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := gob.NewEncoder(file)
	if err = encoder.Encode(session); err != nil {
		return err
	}

	return nil
}

func Run() {
	wac, err := lib.NewConn(5 * time.Second)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error creating connection: %v\n", err)
		return
	}

	wac.AddHandler(&waHandler{wac, uint64(time.Now().Unix())})

	if err = login(wac); err != nil {
		fmt.Fprintf(os.Stderr, "error logging in: %v\n", err)
		return
	}

	<-time.After(60 * time.Minute)
}