package main
  
import (
        "fmt"
        "net/http"

	"github.com/coreos/go-oidc/v3/oidc"
	"golang.org/x/net/context"
	"golang.org/x/oauth2"
)

var (
	clientID     = "o9xW343UeS6T9E7BvBZ5i9hKUqp1zwJR"
	clientSecret = "x28M5R39q37Tv05p40wqX60Sa1952egbFm4h1790N3q3ct8i"
)

func main() {
	ctx := context.Background()

	provider, err := oidc.NewProvider(ctx, "https://yourdomain.oneauth.cn/oauth/v1")
	if err != nil {
		fmt.Println(err)
		return
	}

	config := oauth2.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		Endpoint:     provider.Endpoint(),
		RedirectURL:  "http://127.0.0.1:5556/oneauth/callback",
		Scopes:       []string{oidc.ScopeOpenID, "profile", "email"},
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
                fmt.Println("Ready goto oneauth...")
		http.Redirect(w, r, config.AuthCodeURL(""), http.StatusFound)
        })

        http.HandleFunc("/oneauth/callback", func(w http.ResponseWriter, r *http.Request) {
                fmt.Println("Oneauth done...")
                w.Write([]byte("hahahahhahha"))
        })


        fmt.Printf("listening on http://%s/", "0.0.0.0:5556")
        fmt.Println(http.ListenAndServe(":5556", nil))
}
