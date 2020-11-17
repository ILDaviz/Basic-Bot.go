package util

import (
	"time"

	"github.com/andersfylling/disgord"
)

type messageEmbed struct {
	*disgord.Embed
}

func NewEmbed() *messageEmbed {
	return &messageEmbed{&disgord.Embed{}}
}

func (e *messageEmbed) SetTitle(title string) *messageEmbed {
	e.Title = title
	return e
}

func (e *messageEmbed) SetDescription(descrption string) *messageEmbed {
	e.Description = descrption
	return e
}

func (e *messageEmbed) SetUrl(url string) *messageEmbed {
	e.URL = url
	return e
}

func (e *messageEmbed) SetColor(color int) *messageEmbed {
	e.Color = color
	return e
}

func (e *messageEmbed) SetTimestamp() *messageEmbed {
	e.Timestamp = disgord.Time{Time: time.Now()}
	return e
}

func (e *messageEmbed) SetFooter(args ...string) *messageEmbed {
	var (
		text    string
		iconURL string
	)

	switch {
	case len(args) > 1:
		iconURL = args[1]
		fallthrough
	case len(args) > 0:
		text = args[0]
	case len(args) == 0:
		return e
	}

	e.Footer = &disgord.EmbedFooter{
		Text:    text,
		IconURL: iconURL,
	}

	return e
}

func (e *messageEmbed) SetImage(url string) *messageEmbed {
	e.Image = &disgord.EmbedImage{
		URL: url,
	}
	return e
}

func (e *messageEmbed) SetThumbnail(url string) *messageEmbed {
	e.Thumbnail = &disgord.EmbedThumbnail{
		URL: url,
	}
	return e
}

func (e *messageEmbed) SetAuthor(args ...string) *messageEmbed {
	var (
		name    string
		iconURL string
	)

	switch {
	case len(args) > 1:
		iconURL = args[1]
		fallthrough
	case len(args) > 0:
		name = args[0]
	case len(args) == 0:
		return e
	}

	e.Author = &disgord.EmbedAuthor{
		Name:    name,
		IconURL: iconURL,
	}

	return e
}

func (e *messageEmbed) AddField(name string, value string, inline bool) *messageEmbed {
	if len(name) > 1024 {
		name = name[:1024]
	}

	if len(value) > 1024 {
		value = value[:1024]
	}

	e.Fields = append(e.Fields, &disgord.EmbedField{
		Name:   name,
		Value:  value,
		Inline: inline,
	})

	return e
}

func (e *messageEmbed) ToMessage() *disgord.CreateMessageParams {
	return &disgord.CreateMessageParams{Embed: e.Embed}
}
