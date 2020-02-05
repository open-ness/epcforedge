/* SPDX-License-Identifier: Apache-2.0
* Copyright (c) 2019 Intel Corporation
 */
package ngcnef_test

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	ngcnef "github.com/otcshare/epcforedge/ngc/pkg/nef"
)

const basePFDAPIURL = "http://localhost:8091/3gpp-pfd-management/" +
	"v1/AF_01/transactions"

func CreatePFDReqForNEF(ctx context.Context, method string, pfdTrans string,
	appID string, body []byte) (*httptest.ResponseRecorder, *http.Request) {
	var req *http.Request
	if len(pfdTrans) > 0 {
		if len(appID) > 0 {
			//TBD for application level GET, PUT, PATCH, DELETE
		} else {
			if body != nil {
				//PUT
				req, _ = http.NewRequest(method, basePFDAPIURL+"/"+pfdTrans,
					bytes.NewBuffer(body))
			} else {
				//GET PFD
				req, _ = http.NewRequest(method, basePFDAPIURL+"/"+pfdTrans,
					nil)
			}
			//TBD for DELETE
		}

	} else {
		if body != nil {
			//POST
			req, _ = http.NewRequest(method, basePFDAPIURL,
				bytes.NewBuffer(body))
		} else {
			//GET ALL
			req, _ = http.NewRequest(method, basePFDAPIURL, nil)
		}
	}

	rr := httptest.NewRecorder()
	return rr, req
}

var _ = Describe("Test NEF Server PFD NB API's ", func() {
	var ctx context.Context
	var cancel func()

	ctx, cancel = context.WithCancel(context.Background())
	defer cancel()

	Describe("Start the NEF Server: To be done to start NEF PFD API testing",
		func() {
			It("Will init NefServer",
				func() {
					ctx, cancel = context.WithCancel(context.Background())
					defer cancel()
					go func() {
						err := ngcnef.Run(ctx, NefTestCfgBasepath+"valid.json")
						Expect(err).To(BeNil())
					}()
					time.Sleep(2 * time.Second)
				})
		})

	Describe("REQ to NEF GET ALL", func() {

		It("Send valid GET all to NEF -No Data as no PFD exists",
			func() {
				rr, req := CreatePFDReqForNEF(ctx, "GET", "", "", nil)
				ngcnef.NefAppG.NefRouter.ServeHTTP(rr, req.WithContext(ctx))
				Expect(rr.Code).Should(Equal(http.StatusOK))

				//Validate PFD
				//Read Body from response
				resp := rr.Result()
				b, _ := ioutil.ReadAll(resp.Body)

				//Convert the body(json data) into PFD Management Struct data
				var pfdBody []ngcnef.PfdManagement
				err := json.Unmarshal(b, &pfdBody)
				Expect(err).Should(BeNil())
				fmt.Print("Body Received: ")
				fmt.Println(pfdBody)
				resp.Body.Close()
				Expect(len(pfdBody)).Should(Equal(0))
			})
	})

	Describe("End the NEF Server: To be done to end NEF PFD API testing",
		func() {
			It("Will stop NefServer", func() {
				cancel()
				time.Sleep(2 * time.Second)
			})
		})

})
