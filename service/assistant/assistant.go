package assistant

type Assistant interface {
	CreateAssistant(name string, vectorStores *[]string) (string, error)
	ModifyAssistant(assistantID string, vectorStores *[]string) error
	UploadFile(path string, vectorStore string) (string, error)
	DeleteFile(fileID string) error

	CreateVectorStore(fileIDList *[]string) (string, error)
	DeleteVectorStore(vectorStoreID string) error
}
