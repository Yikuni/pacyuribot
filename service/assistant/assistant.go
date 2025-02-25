package assistant

type AssistantService interface {
	CreateAssistant(name string, vectorStores []string) (string, error)
	ModifyAssistant(assistantID string, vectorStores []string) error
	UploadFile(path string, name string) (string, error)
	DeleteFile(fileID string) error

	CreateVectorStore(name string, fileIDList []string) (string, error)
	DeleteVectorStore(vectorStoreID string) error
}
