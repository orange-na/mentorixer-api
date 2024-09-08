package gimini

import (
	"context"
	"fmt"
	"os"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

func GenerateContent(prompt string) (string, error) {
	ctx := context.Background()
	client, err := genai.NewClient(ctx, option.WithAPIKey(os.Getenv("GEMINI_API_KEY")))
	if err != nil {
    	return "", err
	}
	defer client.Close()

	model := client.GenerativeModel("gemini-1.5-flash")
	resp, err := model.GenerateContent(ctx, genai.Text(prompt))
	if err != nil {
		return "", err
	}

	return extractGeneratedContent(resp), nil
}

func extractGeneratedContent(resp *genai.GenerateContentResponse) string {
    var generatedContent string
    for _, cand := range resp.Candidates {
        if cand.Content != nil {
            for _, part := range cand.Content.Parts {
                generatedContent += fmt.Sprintf("%v", part)
            }
        }
    }
    return generatedContent
}
