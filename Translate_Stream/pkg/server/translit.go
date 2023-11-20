package server

import (
	"fmt"
	tr "github.com/Conight/go-googletrans"
	"io"
	session "steam/api/proto"
)

type TrServer struct {
	session.UnimplementedTransliterationServer // встраивание
}

func (s TrServer) EnRu(inStream session.Transliteration_EnRuServer) error {
	t := tr.New()
	for {
		inWord, err := inStream.Recv()
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
		fmt.Println("server", out)
		inStream.Send(out)
	}
}
