package example

import (
	"fmt"

	"github.com/Vicomtech/itzuli-api-lib/go/itzuli"
)

func main() {

	// Example of package usage
	itzuli := itzuli.Itzuli{}
	itzuli.Init("api key")

	// Traslation
	translationResponse, err := itzuli.Translate("kaixo zer moduz", "eu", "es", "a")
	if err == nil {
		fmt.Println(translationResponse.TranslatedText)
	} else {
		panic(err)
	}

	// Quota information
	if responseQuota, err := itzuli.GetQuota(); err == nil {
		fmt.Println(responseQuota.ConsumedQuota)
	} else {
		panic(err)
	}

	// Feedback info
	if responseFeedback, err := itzuli.Feedback(translationResponse.FeedbackId, "hau horrela da", 1); err == nil {
		fmt.Println(responseFeedback.Message)
	} else {
		panic(err)
	}

}
