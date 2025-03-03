package assistant

import "github.com/sashabaranov/go-openai"

type ChatMessageCallback func(stream openai.AssistantStreamEvent)
type ChatFinishCallback func()
type AssistantService interface {
	CreateAssistant(name string, vectorStores []string) (string, error)
	ModifyAssistant(assistantID string, vectorStores []string) error
	UploadFile(path string, name string) (string, error)
	DeleteFile(fileID string) error

	CreateVectorStore(name string, fileIDList []string) (string, error)
	DeleteVectorStore(vectorStoreID string) error

	Chat(assistantID string, messages []openai.ThreadMessage, onMessage ChatMessageCallback, onFinish ChatFinishCallback)
}
