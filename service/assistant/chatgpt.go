package assistant

import (
	"context"
	"github.com/sashabaranov/go-openai"
	"pacyuribot/global"
	"pacyuribot/utils"
)

type ChatGPTService struct {
}

// CreateAssistant
// @Description 创建assistant
// @Param name   string 助手名字
// @Param vectorStores *[]string	创建时带上的vector store
// @Return string	助手ID
// @Return error	err
func (c *ChatGPTService) CreateAssistant(name string, vectorStores []string) (string, error) {
	cli := utils.GetChatGPTClient()
	assistant, err := cli.CreateAssistant(context.Background(), openai.AssistantRequest{
		Model:        openai.GPT4o,
		Name:         &name,
		Description:  nil,
		Instructions: &global.Config.Chatgpt.Instruction,
		Tools: []openai.AssistantTool{{
			Type:     openai.AssistantToolTypeFileSearch,
			Function: nil,
		}},
		FileIDs:  nil,
		Metadata: nil,
		ToolResources: &openai.AssistantToolResource{
			FileSearch: &openai.AssistantToolFileSearch{
				VectorStoreIDs: vectorStores,
			},
			CodeInterpreter: nil,
		},
		ResponseFormat: nil,
		Temperature:    nil,
		TopP:           nil,
	})
	if err != nil {
		return "", err
	}
	return assistant.ID, nil
}

func (c *ChatGPTService) ModifyAssistant(assistantID string, vectorStores []string) error {
	cli := utils.GetChatGPTClient()
	_, err := cli.ModifyAssistant(context.Background(), assistantID,
		openai.AssistantRequest{
			Model:        openai.GPT4o,
			Name:         nil,
			Description:  nil,
			Instructions: nil,
			Tools: []openai.AssistantTool{{
				Type:     openai.AssistantToolTypeFileSearch,
				Function: nil,
			}},
			FileIDs:  nil,
			Metadata: nil,
			ToolResources: &openai.AssistantToolResource{
				FileSearch: &openai.AssistantToolFileSearch{
					VectorStoreIDs: vectorStores,
				},
				CodeInterpreter: nil,
			},
			ResponseFormat: nil,
			Temperature:    nil,
			TopP:           nil,
		})
	return err
}

func (c *ChatGPTService) UploadFile(path string, name string) (string, error) {
	cli := utils.GetChatGPTClient()
	file, err := cli.CreateFile(context.Background(), openai.FileRequest{
		FileName: name,
		FilePath: path,
		Purpose:  "assistants",
	})
	if err != nil {
		return "", err
	}
	return file.ID, err
}

func (c *ChatGPTService) DeleteFile(fileID string) error {
	cli := utils.GetChatGPTClient()
	err := cli.DeleteFile(context.Background(), fileID)
	return err
}

func (c *ChatGPTService) CreateVectorStore(name string, fileIDList []string) (string, error) {
	cli := utils.GetChatGPTClient()
	store, err := cli.CreateVectorStore(context.Background(), openai.VectorStoreRequest{
		Name:    name,
		FileIDs: fileIDList,
		ExpiresAfter: &openai.VectorStoreExpires{
			Anchor: "last_active_at",
			Days:   3,
		},
		Metadata: nil,
	})
	if err != nil {
		return "", err
	}
	return store.ID, nil
}

func (c *ChatGPTService) DeleteVectorStore(vectorStoreID string) error {
	cli := utils.GetChatGPTClient()
	_, err := cli.DeleteVectorStore(context.Background(), vectorStoreID)
	return err
}
