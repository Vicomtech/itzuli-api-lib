/*
 * Itzuli API libraries
 * Coded by Aritz Olea <aolea@vicomtech.org>
 * (c) 2020 Vicomtech
 */

package itzuli

// Translation request
type TranslationRequest struct {
	FromLang string   `json:"sourcelanguage"`
	ToLang   string   `json:"targetlanguage"`
	Text     string   `json:"text"`
	Tags     []string `json:"tags,omitempty"`
}

// Response to translation request
type TranslationResponse struct {
	ConsumedQuota     string `json:"consumedquota"`
	DisclaimerBasque  string `json:"disclaimereu"`
	DisclaimerEnglish string `json:"disclaimeren"`
	FeedbackId        string `json:"feedbackid"`
	FromLang          string `json:"source_lang"`
	ToLang            string `json:"target_lang"`
	OriginalText      string `json:"source_text"`
	TranslatedText    string `json:"translated_text"`
}

// Quota response
type QuotaResponse struct {
	ConsumedQuota string `json:"consumedquota"`
}

// Feedback request
type FeedbackRequest struct {
	Id         string `json:"id"`
	Evaluation int    `json:"evaluation"`
	Correction string `json:"correction"`
}

// Feedback response
type FeedbackResponse struct {
	DisclaimerBasque  string `json:"disclaimer_eu"`
	DisclaimerEnglish string `json:"disclaimer_en"`
	Message           string `json:"message"`
}
