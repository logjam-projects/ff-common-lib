package httputils

import (
	"cloud.google.com/go/compute/metadata"
	"fmt"
	"log"
	"net/http"
)

type GetServiceToServiceRequest struct {
	//This is usually just the root domain, for example http://cr-service-drerererw12 no / and no uri
	S2SHost string
	//The actual path you need to call too. /alive for example with /
	URLPath string
	// Authenticated determines whether identity token authentication will be used.
	Authenticated bool
}

func (s *GetServiceToServiceRequest) NewRequest(method string) *http.Response {

	completeUrlToCall := s.S2SHost + s.URLPath

	log.Println("Calling out too: " + completeUrlToCall)

	req, err := http.NewRequest("GET", completeUrlToCall, nil)
	if err != nil {
		fmt.Errorf("cannot build request: %s", err)
		return nil
	}

	if s.Authenticated {
		tokenURL := fmt.Sprintf("/instance/service-accounts/default/identity?audience=%s", s.S2SHost)
		log.Println("TokenURL: " + tokenURL)
		idToken, err := metadata.Get(tokenURL)

		if err != nil {
			fmt.Errorf("metadata.Get: failed to query id_token: %+v", err)
			return nil
		}

		//add JWT
		log.Println("Token is : " + idToken)
		req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", idToken))
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Errorf("cannot callout on the request: %s", err)
		return nil
	}
	return res
}
