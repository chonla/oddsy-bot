package translator

import (
	"log"
	"os"

	"google.golang.org/api/option"

	"cloud.google.com/go/translate"
	"golang.org/x/net/context"
	"golang.org/x/text/language"
)

// Translator is translator
type Translator struct {
	ctx    context.Context
	client *translate.Client
	logger *log.Logger
	token  string
}

// Configuration is for translator
type Configuration struct {
	GcpToken string
}

// NewTranslator creates a new translator
func NewTranslator(conf *Configuration) *Translator {
	t := &Translator{
		token:  conf.GcpToken,
		logger: log.New(os.Stdout, "translator: ", log.Lshortfile|log.LstdFlags),
	}

	envToken := os.Getenv("GCP_TOKEN")
	if envToken != "" {
		t.SetToken(envToken)
	}

	ctx := context.Background()
	c, e := translate.NewClient(ctx, option.WithAPIKey(t.token))
	if e != nil {
		log.Fatalf("Cannot create client: %v", e)
	}

	t.ctx = ctx
	t.client = c
	return t
}

// SetToken overrides configure token
func (t *Translator) SetToken(s string) {
	t.token = s
	t.logger.Println("GCP token is overwritten by environment variable.")
}

// Translate to translate language to target language
func (t *Translator) Translate(d string) (r string, e error) {
	target, e := language.Parse("th")
	if e != nil {
		log.Fatalf("Cannot parse language: %v", e)
		return
	}

	langs, e := t.client.DetectLanguage(t.ctx, []string{d})
	if e != nil {
		log.Fatalf("Cannot detect language: %v", e)
		return
	}

	translations, e := t.client.Translate(t.ctx, []string{d}, target, nil)
	if e != nil {
		log.Fatalf("Cannot translate: %v", e)
		return
	}

	r = "[" + langs[0][0].Language.String() + "] " + translations[0].Text

	return
}
