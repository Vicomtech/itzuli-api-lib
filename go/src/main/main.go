package main

import(
  "itzuli"
)


func main() {

  // Example of package usage

  // Traslation
  if response, err := itzuli.Translate("kaixo zer moduz", "eu", "es", "apikey"); err != nil {
    fmt.Println(response.TranslatedText)
  }

  // Quota information
  if responseQuota, err := itzuli.GetQuota("apikey"); err != nil {
    fmt.Println(responseQuota.ConsumedQuota)
  }

  // Feedback info
  if responseFeedback, err := itzuli.Feedback(translationResponse.FeedbackId, "hau horrela da", 1, "apikey"); err != nil {
    fmt.Println(responseFeedback.Message)
  }


}
