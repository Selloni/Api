package server

import (
	tr "github.com/Conight/go-googletrans"
	"io"
	"log"
	session "steam/api/proto"
)

type TrServer struct {
	session.UnimplementedTransliterationServer // встраивание
}

//func NewServer() *TrServer {
//	return &TrServer{}
//}

func (s TrServer) EnRu(inStream session.Transliteration_EnRuServer) error {
	t := tr.New()
	for {
		inWord, err := inStream.Recv()
		log.Println("server-get", inWord.Word)
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		ressult, err := t.Translate(inWord.Word, "auto", "en") // перевод слова на английски
		out := &session.Word{
			Word: ressult.Text,
		}
		inStream.Send(out)
	}
}
