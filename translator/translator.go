package translator

import (
	"log"

	"cloud.google.com/go/translate"
	"golang.org/x/net/context"
	"golang.org/x/text/language"
)

// Translator is translator
type Translator struct {
	ctx    context.Context
	client *translate.Client
}

// NewTranslator creates a new translator
func NewTranslator() *Translator {
	ctx := context.Background()
	c, e := translate.NewClient(ctx)
	if e != nil {
		log.Fatalf("Cannot create client: %v", e)
	}
	return &Translator{
		ctx:    ctx,
		client: c,
	}
}

// Translate to translate language to target language
func (t *Translator) Translate(d string) (r string, e error) {
	target, e := language.Parse("th")
	if e != nil {
		log.Fatalf("Cannot parse language: %v", e)
		return
	}

	translations, e := t.client.Translate(t.ctx, []string{d}, target, nil)
	if e != nil {
		log.Fatalf("Cannot translate: %v", e)
		return
	}

	r = translations[0].Text

	return
}
