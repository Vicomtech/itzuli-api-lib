/*
 * Itzuli API libraries
 * Coded by Aritz Olea <aolea@vicomtech.org>
 * (c) 2020 Vicomtech
 */

package itzuli

import(
  "errors"
  "encoding/json"
  "io/ioutil"
  "net/http"
  "strings"
)

// Constant paths
const URL_API = "https://api.itzuli.vicomtech.org/"
const PATH_TRANSLATE = "translation/get"
const PATH_FEEDBACK = "translation/feedback"
const PATH_QUOTA = "quota/get"


// Makes a translation request to API
// If success, a TranslationResponse object will be resturned
// If any error, an error will be returned
func Translate(text, fromlang, tolang, auth string) (TranslationResponse, error) {

  // Basic error check
  if len(fromlang) <= 0 || len(tolang) <= 0 || len(text) <= 0 || len(auth) <= 0 {
    return TranslationResponse{}, errors.New("Invalid parameters")
  }

  // Make new request object
  reqObj := TranslationRequest{}
  reqObj.FromLang = fromlang
  reqObj.ToLang = tolang
  reqObj.Text = text

  reqObjBytes, err := json.Marshal(reqObj)
	if err != nil {
    return TranslationResponse{}, errors.New("Error creating request bytes")
	}

  // Create a new request using http
  req, err := http.NewRequest("POST", URL_API+PATH_TRANSLATE, strings.NewReader(string(reqObjBytes)))
  if err != nil {
    return TranslationResponse{}, err
  }

  req.Header.Set("Content-Type", "application/json")
  req.Header.Add("Authorization", "Bearer "+auth)

  // Make request
  client := &http.Client{}
  resp, err := client.Do(req)
  if err != nil {
     return TranslationResponse{}, err
  }

  // Check status code
  if resp.StatusCode == 401 {
		return TranslationResponse{}, errors.New("Invalid API key or expired")
	}

  // Check status code
  if resp.StatusCode != 200 {
		return TranslationResponse{}, errors.New("Invalid status code")
	}

  // Read body
  body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return TranslationResponse{}, err
	}

  // Parse object and return
  responseObj := TranslationResponse{}
  err = json.Unmarshal(body, &responseObj)
  if err != nil {
    return TranslationResponse{}, err
  }

  // Return object
  return responseObj, nil
}


// Checks user quota
// If valid, QuotaResponse object will be returned
// if not, an error will be thrown
func GetQuota(auth string) (QuotaResponse, error) {

  // Basic check
  if len(auth) <= 0 {
    return QuotaResponse{}, errors.New("No auth string provided")
  }

  // Create a new request using http
  req, err := http.NewRequest("GET", URL_API+PATH_QUOTA, nil)
  if err != nil {
    return QuotaResponse{}, err
  }

  req.Header.Set("Content-Type", "application/json")
  req.Header.Add("Authorization", "Bearer "+auth)

  // Make request
  client := &http.Client{}
  resp, err := client.Do(req)
  if err != nil {
     return QuotaResponse{}, err
  }

  // Check status code 401
  if resp.StatusCode == 401 {
		return QuotaResponse{}, errors.New("Invalid API key or expired")
	}

  // Check status code not 200
  if resp.StatusCode != 200 {
		return QuotaResponse{}, errors.New("Invalid status code")
	}

  // Read body
  body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return QuotaResponse{}, err
	}

  // Parse object and return
  responseObj := QuotaResponse{}
  err = json.Unmarshal(body, &responseObj)
  if err != nil {
    return QuotaResponse{}, err
  }

  // Return object
  return responseObj, nil
}


// Sends feedback information using translated job identificator
// If valid, a message will be return
// if not, an error will be thrown
func Feedback(id, correction string, evaluation int, auth string) (FeedbackResponse, error) {
  // Basic error check
  if len(id) <= 0 || len(correction) <= 0 || len(auth) <= 0 {
    return FeedbackResponse{}, errors.New("Invalid parameters")
  }

  // Make new request object
  reqObj := FeedbackRequest{}
  reqObj.Id = id
  reqObj.Evaluation = evaluation
  reqObj.Correction = correction

  reqObjBytes, err := json.Marshal(reqObj)
	if err != nil {
    return FeedbackResponse{}, errors.New("Error creating request bytes")
	}

  // Create a new request using http
  req, err := http.NewRequest("POST", URL_API+PATH_FEEDBACK, strings.NewReader(string(reqObjBytes)))
  if err != nil {
    return FeedbackResponse{}, err
  }

  req.Header.Set("Content-Type", "application/json")
  req.Header.Add("Authorization", "Bearer "+auth)

  // Make request
  client := &http.Client{}
  resp, err := client.Do(req)
  if err != nil {
     return FeedbackResponse{}, err
  }

  // Check status code
  if resp.StatusCode == 401 {
		return FeedbackResponse{}, errors.New("Invalid API key or expired")
	}

  // Check status code
  if resp.StatusCode != 200 {
		return FeedbackResponse{}, errors.New("Invalid status code")
	}

  // Read body
  body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return FeedbackResponse{}, err
	}

  // Parse object and return
  responseObj := FeedbackResponse{}
  err = json.Unmarshal(body, &responseObj)
  if err != nil {
    return FeedbackResponse{}, err
  }

  // Return object
  return responseObj, nil
}
