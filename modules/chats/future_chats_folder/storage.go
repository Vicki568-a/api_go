package future_chats_folder

import (
	"context"
	"github.com/zhashkevych/todo-app/modules/chats/chat_naw"
)
/*интерфейс*/
type StorageChat interface {
	CreateChat(ctx context.Context, chat chat_naw.Chat) (string, error)
	FindAllChat(ctx context.Context) (u []chat_naw.Chat, err error)
	FindOneChat(ctx context.Context, chat_id string) (chat_naw.Chat, error)
	UpdateChat(ctx context.Context, chat chat_naw.Chat) error
	DeleteChat(ctx context.Context, chat_id string) error


}