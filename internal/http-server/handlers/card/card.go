package card

import (
	"credit_card_validator/internal/algo"
	"credit_card_validator/internal/lib/api/response"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type CardResponse struct {
	Verified bool `json:"verified"`
	response.Response
}

type CardRequest struct {
	Card string `json:"card"`
}

func New() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		d := json.NewDecoder(r.Body)
		body := &CardRequest{}
		err := d.Decode(body)
		if err != nil {
			err, _ := json.Marshal(response.Error("bad request"))
			http.Error(w, string(err), http.StatusBadRequest)
			log.Fatal(err)
			return
		}
		verified := algo.CheckCard(body.toIntArr())
		resp := CardResponse{
			Verified: verified,
			Response: response.OK(),
		}
		result, _ := json.Marshal(resp)
		fmt.Fprintf(w, string(result))
	}
}

func (cr CardRequest) toIntArr() [16]uint8 {
	var res [16]uint8
	for i, numS := range strings.Split(cr.Card, "") {
		num, err := strconv.Atoi(numS)
		if err == nil {
			fmt.Println(i, num, numS)
			res[i] = uint8(num)
		} else {
			log.Fatal(err)
		}
	}
	return res
}
