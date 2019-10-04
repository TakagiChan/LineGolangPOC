package main

import (
	"github.com/apache/thrift/lib/go/thrift"
	"github.com/okanon/linethrift.golang/LineThrift"
)

type MessageInfo struct {
	To string
	Text string
}

func authLoginAndSendMessage(auth string){

	var mes MessageInfo
	mes.To = "gid"
	mes.Text = "test"

	message := LineThrift.Message{
		To: mes.To,
		Text: mes.Text,
	}

	transport, _ := thrift.NewTHttpClient("https://gd2.line.naver.jp/S4")
	httpTransport := transport.(*thrift.THttpClient)
	httpTransport.SetHeader("X-Line-Application", "DESKTOPWIN\t12.1.0\tTEST\t1.12.0")
	httpTransport.SetHeader("X-Line-Access","test")
	var iprot,oprot thrift.TProtocol
	client := LineThrift.NewTalkServiceClientProtocol(transport,iprot,oprot)
	client.SendMessage(client.SeqId,&message)
}

func main(){
	authLoginAndSendMessage("auth")
}