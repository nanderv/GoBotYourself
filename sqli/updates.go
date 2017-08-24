package sql

import (
	"log"
	"gopkg.in/telegram-bot-api.v4"
)

type Update struct {
	UpdateId             int
	MessageId            int
	FromId               int
	Date                 int
	ChatId               int64
	ReplyToMessageID     int
	EditDate             int
	ForwardFromChatId    int
	ForwardFromMessageId int
	Text                 string
	PhotoFileId          string
	DocumentFileId       string
	DocumentFileName     string
	Caption              string
	NewChatTitle         string
	PinnedMessageId      int
	AudioFileId          string
	AudioFileName        string
	VideoFileId          string
	VoiceFileId          string
	ContactUserId        int
	LocationLongitude    float32
	LocationLatitude     float32
	StickerFileId        string
}

func StoreUpdate(u tgbotapi.Update) {
	var stmt string
	var message *tgbotapi.Message
	var toStore Update
	switch {
	case u.Message != nil:
		message = u.Message
	case u.EditedMessage != nil:
		message = u.EditedMessage
	default:
		return
	}

	// concentrate important info
	toStore.UpdateId = u.UpdateID
	toStore.MessageId = message.MessageID
	toStore.FromId = message.From.ID
	toStore.Date = message.Date
	toStore.ChatId = message.Chat.ID
	if message.ReplyToMessage != nil {
		toStore.ReplyToMessageID = message.ReplyToMessage.MessageID
	}
	toStore.EditDate = message.EditDate
	if message.ForwardFrom != nil {
		toStore.ForwardFromChatId = message.ForwardFrom.ID
	}
	toStore.ForwardFromMessageId = message.ForwardFromMessageID
	toStore.Text = message.Text
	if message.Photo != nil {
		var id string
		maxSize := 0
		for _, p := range *message.Photo {
			if p.FileSize > maxSize {
				maxSize = p.FileSize
				id = p.FileID
			}
		}
		toStore.PhotoFileId = id
	}
	if message.Document != nil {
		toStore.DocumentFileId = message.Document.FileID
		toStore.DocumentFileName = message.Document.FileName
	}
	toStore.Caption = message.Caption
	toStore.NewChatTitle = message.NewChatTitle
	if message.PinnedMessage != nil {
		toStore.PinnedMessageId = message.PinnedMessage.MessageID
	}

	if message.Audio != nil {
		toStore.AudioFileId = message.Audio.FileID
		toStore.AudioFileName = message.Audio.Title
	}
	if message.Video != nil {
		toStore.VideoFileId = message.Video.FileID
	}
	if message.Voice != nil {
		toStore.VoiceFileId = message.Voice.FileID
	}
	if message.Contact != nil {
		toStore.ContactUserId = message.Contact.UserID
	}
	if message.Location != nil {
		toStore.LocationLongitude = float32(message.Location.Longitude)
		toStore.LocationLatitude = float32(message.Location.Latitude)
	}
	if message.Sticker != nil {
		toStore.StickerFileId = message.Sticker.FileID
	}

	// Insert into db
	stmt = "INSERT INTO Updates(UpdateId, MessageId, FromId, Date, ChatId, ReplyToMessageID, EditDate, ForwardFromChatId, ForwardFromMessageId, Text, PhotoFileId, DocumentFileId, DocumentFileName, Caption, NewChatTitle, PinnedMessageId, AudioFileId, AudioFileName, VideoFileId, VoiceFileId, ContactUserId, LocationLongitude, LocationLatitude, StickerFileId) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?);"
	db := Getdb().db

	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}

	preparedStmt, err := tx.Prepare(stmt)
	defer preparedStmt.Close()

	if err != nil {
		log.Fatal(err)
	}
	_, err = preparedStmt.Exec(toStore.UpdateId, toStore.MessageId, toStore.FromId, toStore.Date, toStore.ChatId, toStore.ReplyToMessageID, toStore.EditDate, toStore.ForwardFromChatId, toStore.ForwardFromMessageId, toStore.Text, toStore.PhotoFileId, toStore.DocumentFileId, toStore.DocumentFileName, toStore.Caption, toStore.NewChatTitle, toStore.PinnedMessageId, toStore.AudioFileId, toStore.AudioFileName, toStore.VideoFileId, toStore.VoiceFileId, toStore.ContactUserId, toStore.LocationLongitude, toStore.LocationLatitude, toStore.StickerFileId)
	if err != nil {
		log.Fatal(err)
	} else {
		tx.Commit()
	}
}
