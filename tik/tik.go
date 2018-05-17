package tik

import (
	"log"
	"os"

	"google.golang.org/api/option"

	"cloud.google.com/go/firestore"

	firebase "firebase.google.com/go"
	"golang.org/x/net/context"
)

// Tik is tik
type Tik struct {
	ctx    context.Context
	client *firestore.Client
	logger *log.Logger
	token  string
}

// Configuration is for translator
type Configuration struct {
	GcpToken          string
	FirebaseProjectID string
}

// NewTik creates a new tik
func NewTik(conf *Configuration) *Tik {
	t := &Tik{
		token:  conf.GcpToken,
		logger: log.New(os.Stdout, "tik: ", log.Lshortfile|log.LstdFlags),
	}

	envToken := os.Getenv("GCP_TOKEN")
	if envToken != "" {
		t.SetToken(envToken)
	}

	ctx := context.Background()
	fconf := &firebase.Config{
		ProjectID:   conf.FirebaseProjectID,
		DatabaseURL: "https://" + conf.FirebaseProjectID + ".firebaseio.com",
	}

	a, e := firebase.NewApp(ctx, fconf, option.WithAPIKey(t.token))
	if e != nil {
		log.Fatalf("Cannot create app: %v", e)
	}

	c, e := a.Firestore(ctx)
	if e != nil {
		log.Fatalf("Cannot create client: %v", e)
	}

	t.ctx = ctx
	t.client = c
	return t
}

// SetToken overrides configure token
func (t *Tik) SetToken(s string) {
	t.token = s
	t.logger.Println("GCP token is overwritten by environment variable.")
}

// Release to release client
func (t *Tik) Release() {
	t.client.Close()
}
