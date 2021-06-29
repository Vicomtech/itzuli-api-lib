# Itzuli API libraries

In this repository, you can find example ready-to-use libraries. For more information, please visit https://itzuli.vicomtech.org/api/


## Go

### Functions in package itzuli

    Translate(text, fromlang, tolang, auth string) (TranslationResponse, error)
    GetQuota(auth string) (QuotaResponse, error)
    Feedback(id, correction string, evaluation int, auth string) (FeedbackResponse, error)

### Example
Get the package
> go get -u github.com/Vicomtech/itzuli-api-lib/go/src/itzuli

Import:

    import("github.com/Vicomtech/itzuli-api-lib/go/src/itzuli")

Send a translation:

    if response, err := itzuli.Translate("kaixo! zer moduz?", "eu", "es", "key"); err != nil {
	    fmt.Println(response.TranslatedText)
	}

  Get used quota:


    if responseQuota, err := itzuli.GetQuota("apikey"); err != nil {
	    fmt.Println(responseQuota.ConsumedQuota)
	}

Feedback:

    if responseFeedback, err := itzuli.Feedback(translationResponse.FeedbackId, "hau horrela da", 1, "apikey"); err != nil {
	    fmt.Println(responseFeedback.Message)
	}


## Python 3
### Functions in "Itzuli" package

    getTranslation(text, fromlang, tolang)
    getQuota()
    sendFeedback(id, correction, evaluation)

An exception will be thrown if any error detected.

### Example

Install the package

> pip install https://github.com/Vicomtech/itzuli-api-lib/blob/master/python/dist/itzuli-pkg-vicomtech-1.0.0.tar.gz

Import:

    import itzuli

Initialize with your api key:

    itzuliapi = itzuli.Itzuli("apikey")

Make a translation:

    itzuliapi.getTranslation("kaixo! zer moduz?", "eu", "es")

Get quota:

    itzuliapi.getQuota()

Send feedback:

    itzuliapi.sendFeedback("id from translation, "corrected text", 1)
