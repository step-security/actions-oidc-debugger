package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/github/actions-oidc-debugger/actionsoidc"
)

func validateSubscription() {
	repo := os.Getenv("GITHUB_REPOSITORY")
	apiURL := fmt.Sprintf("https://agent.api.stepsecurity.io/v1/github/%s/actions/subscription", repo)

	client := http.Client{
		Timeout: 3 * time.Second,
	}

	resp, err := client.Get(apiURL)
	if err != nil {
		fmt.Println("Timeout or API not reachable. Continuing to next step.")
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusForbidden {
		fmt.Fprintln(os.Stderr, "Subscription is not valid. Reach out to support@stepsecurity.io")
		os.Exit(1)
	}
}

func main() {
	validateSubscription()
	audience := flag.String("audience", "https://github.com/", "the audience for the requested jwt")
	flag.Parse()

	if *audience == "https://github.com/" {
		actionsoidc.QuitOnErr(fmt.Errorf("-audience cli argument must be specified"))
	}

	c := actionsoidc.DefaultOIDCClient(*audience)
	jwt, err := c.GetJWT()
	actionsoidc.QuitOnErr(err)

	jwt.Parse()
	fmt.Print(jwt.PrettyPrintClaims())
}
