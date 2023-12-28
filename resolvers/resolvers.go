package resolvers

import (
	"github.com/leandro-lugaresi/hub"
	"github.com/muety/telepush/config"
	"github.com/muety/telepush/model"
)

const (
	TextType = "TEXT"
	FileType = "FILE"
)

type MessageResolver interface {
	IsValid(*model.Message) error
	Resolve(string, *model.Message, *model.MessageOptions) error
}

func GetResolver(ttype string) MessageResolver {
	switch ttype {
	case FileType:
		return &FileResolver{}
	}
	return &TextResolver{}
}

func logMessage(m *model.Message) {
	ttype := m.Type
	if ttype == "" {
		ttype = TextType
	}

	config.GetHub().Publish(hub.Message{
		Name: config.EventOnMessageDelivered,
		Fields: map[string]interface{}{
			"origin": m.Origin,
			"type":   ttype,
		},
	})
}
