//SPDX-License-Identifier: Apache-2.0
//Copyright © 2019-2020 Intel Corporation

package af_test

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/otcshare/epcforedge/ngc/pkg/af"
)

func TestAf(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "AF Suite")
}

type KeyType string

var _ = Describe("AF", func() {

	var (
		ctx         context.Context
		srvCancel   context.CancelFunc
		afIsRunning bool
	)

	ctx, srvCancel = context.WithCancel(context.Background())

	Describe("Cnca client request methods to AF : ", func() {

		Context("Subscription GET ALL", func() {

			By("Starting AF server")
			var err error
			ctx, srvCancel = context.WithCancel(context.Background())
			_ = srvCancel
			afRunFail := make(chan bool)
			go func() {

				err = af.Run(ctx, "./testdata/testconfigs/af.json")

				Expect(err).ShouldNot(HaveOccurred())
				if err != nil {
					fmt.Printf("Run() exited with error: %#v", err)
					afIsRunning = false
					afRunFail <- true
				}
			}()
			_ = afIsRunning
		})
	})

	Describe("Cnca client request methods to AF : ", func() {

		Context("Subscription POST", func() {
			Specify("Sending POST 001 request", func() {
				By("Reading json file")
				reqBody, err := ioutil.ReadFile(
					"./testdata/100_AF_NB_SUB_POST001.json")
				Expect(err).ShouldNot(HaveOccurred())

				By("Preparing request")
				reqBodyBytes := bytes.NewReader(reqBody)
				req, err := http.NewRequest(http.MethodPost,
					"http://localhost:8080/af/v1/subscriptions",
					reqBodyBytes)
				Expect(err).ShouldNot(HaveOccurred())

				By("Sending request")
				resp := httptest.NewRecorder()
				ctx := context.WithValue(req.Context(),
					KeyType("af-ctx"), af.AfCtx)
				af.AfRouter.ServeHTTP(resp, req.WithContext(ctx))

				Expect(resp.Code).To(Equal(http.StatusCreated))

			})
			Specify("Sending POST 002 Request", func() {
				By("Reading json file")
				reqBody, err := ioutil.ReadFile(
					"./testdata/100_AF_NB_SUB_POST002.json")
				Expect(err).ShouldNot(HaveOccurred())

				By("Preparing request")
				reqBodyBytes := bytes.NewReader(reqBody)
				req, err := http.NewRequest(http.MethodPost,
					"http://localhost:8080/af/v1/subscriptions",
					reqBodyBytes)
				Expect(err).ShouldNot(HaveOccurred())

				By("Sending request")
				resp := httptest.NewRecorder()
				ctx := context.WithValue(req.Context(),
					KeyType("af-ctx"), af.AfCtx)
				af.AfRouter.ServeHTTP(resp, req.WithContext(ctx))

				Expect(resp.Code).To(Equal(http.StatusCreated))

			})

			Specify("Sending POST 003 request", func() {
				By("Reading json file")
				reqBody, err := ioutil.ReadFile(
					"./testdata/100_AF_NB_SUB_POST003.json")
				Expect(err).ShouldNot(HaveOccurred())

				By("Preparing request")
				reqBodyBytes := bytes.NewReader(reqBody)
				req, err := http.NewRequest(http.MethodPost,
					"http://localhost:8080/af/v1/subscriptions",
					reqBodyBytes)
				Expect(err).ShouldNot(HaveOccurred())

				By("Sending request")
				resp := httptest.NewRecorder()
				ctx := context.WithValue(req.Context(),
					KeyType("af-ctx"), af.AfCtx)
				af.AfRouter.ServeHTTP(resp, req.WithContext(ctx))

				Expect(resp.Code).To(Equal(http.StatusCreated))

			})
			Specify("Sending POST 004 request", func() {
				By("Reading json file")
				reqBody, err := ioutil.ReadFile(
					"./testdata/100_AF_NB_SUB_POST004.json")
				Expect(err).ShouldNot(HaveOccurred())

				By("Preparing request with subID")
				reqBodyBytes := bytes.NewReader(reqBody)
				req, err := http.NewRequest(http.MethodPost,
					"http://localhost:8080/af/v1/subscriptions/1000",
					reqBodyBytes)
				Expect(err).ShouldNot(HaveOccurred())

				By("Sending request")
				resp := httptest.NewRecorder()
				ctx := context.WithValue(req.Context(),
					KeyType("af-ctx"), af.AfCtx)
				af.AfRouter.ServeHTTP(resp, req.WithContext(ctx))

				Expect(resp.Code).To(Equal(http.StatusMethodNotAllowed))

			})

		})

		Context("Subscription GET ALL", func() {
			Specify("Read all subscriptions", func() {

				req, err := http.NewRequest(http.MethodGet,
					"http://localhost:8080/af/v1/subscriptions",
					nil)
				Expect(err).ShouldNot(HaveOccurred())

				resp := httptest.NewRecorder()
				ctx := context.WithValue(req.Context(),
					KeyType("af-ctx"), af.AfCtx)
				af.AfRouter.ServeHTTP(resp, req.WithContext(ctx))

				Expect(resp.Code).To(Equal(http.StatusOK))
			})

			Specify("Read all subscriptions", func() {
				By("sending wrong url")
				req, err := http.NewRequest(http.MethodGet,
					"http://localhost:8080/af/v2/subscriptions",
					nil)
				Expect(err).ShouldNot(HaveOccurred())

				resp := httptest.NewRecorder()
				ctx := context.WithValue(req.Context(),
					KeyType("af-ctx"), af.AfCtx)
				af.AfRouter.ServeHTTP(resp, req.WithContext(ctx))

				Expect(resp.Code).To(Equal(http.StatusNotFound))

			})
		})

		Context("Subscription ID GET", func() {
			Specify("", func() {
				req, err := http.NewRequest(http.MethodGet,
					"http://localhost:8080/af/v1/subscriptions/11112",
					nil)
				Expect(err).ShouldNot(HaveOccurred())

				resp := httptest.NewRecorder()
				ctx := context.WithValue(req.Context(),
					KeyType("af-ctx"), af.AfCtx)
				af.AfRouter.ServeHTTP(resp, req.WithContext(ctx))

				Expect(resp.Code).To(Equal(http.StatusOK))

			})

		})

		Context("Subscription ID PUT", func() {
			Specify("", func() {
				By("Reading json file")
				reqBody, err := ioutil.ReadFile(
					"./testdata/300_AF_NB_SUB_SUBID_PUT001.json")
				Expect(err).ShouldNot(HaveOccurred())

				By("Preparing request")
				reqBodyBytes := bytes.NewReader(reqBody)
				req, err := http.NewRequest(http.MethodPut,
					"http://localhost:8080/af/v1/subscriptions/11113",
					reqBodyBytes)
				Expect(err).ShouldNot(HaveOccurred())

				By("Sending request")
				resp := httptest.NewRecorder()
				ctx := context.WithValue(req.Context(),
					KeyType("af-ctx"), af.AfCtx)
				af.AfRouter.ServeHTTP(resp, req.WithContext(ctx))

				Expect(resp.Code).To(Equal(http.StatusOK))

			})

		})

		Context("Subscription ID PATCH", func() {
			Specify("", func() {
				By("Reading json file")
				reqBody, err := ioutil.ReadFile(
					"./testdata/400_AF_NB_SUB_SUBID_PATCH001.json")
				Expect(err).ShouldNot(HaveOccurred())

				By("Preparing request")
				reqBodyBytes := bytes.NewReader(reqBody)
				req, err := http.NewRequest(http.MethodPatch,
					"http://localhost:8080/af/v1/subscriptions/11112",
					reqBodyBytes)
				Expect(err).ShouldNot(HaveOccurred())

				By("Sending request")
				resp := httptest.NewRecorder()
				ctx := context.WithValue(
					req.Context(), KeyType("af-ctx"), af.AfCtx)
				af.AfRouter.ServeHTTP(resp, req.WithContext(ctx))

				Expect(resp.Code).To(Equal(http.StatusOK))

			})

		})

		Context("Subscription ID DELETE", func() {
			Specify("DELETE Subcription 01", func() {

				req, err := http.NewRequest(http.MethodDelete,
					"http://localhost:8080/af/v1/subscriptions/11111",
					nil)
				Expect(err).ShouldNot(HaveOccurred())

				resp := httptest.NewRecorder()
				ctx := context.WithValue(req.Context(),
					KeyType("af-ctx"), af.AfCtx)
				af.AfRouter.ServeHTTP(resp, req.WithContext(ctx))
				Expect(resp.Code).To(Equal(http.StatusNoContent))
			})

			Specify("DELETE Subcription 02", func() {
				req, err := http.NewRequest(http.MethodDelete,
					"http://localhost:8080/af/v1/subscriptions/11112",
					nil)
				Expect(err).ShouldNot(HaveOccurred())

				resp := httptest.NewRecorder()
				ctx := context.WithValue(req.Context(),
					KeyType("af-ctx"), af.AfCtx)
				af.AfRouter.ServeHTTP(resp, req.WithContext(ctx))
				Expect(resp.Code).To(Equal(http.StatusNoContent))
			})
			Specify("DELETE Subcription 03", func() {
				req, err := http.NewRequest(http.MethodDelete,
					"http://localhost:8080/af/v1/subscriptions/11113",
					nil)
				Expect(err).ShouldNot(HaveOccurred())

				resp := httptest.NewRecorder()
				ctx := context.WithValue(req.Context(),
					KeyType("af-ctx"), af.AfCtx)
				af.AfRouter.ServeHTTP(resp, req.WithContext(ctx))
				Expect(resp.Code).To(Equal(http.StatusNoContent))
			})
			Specify("DELETE Subcription 04", func() {
				req, err := http.NewRequest(http.MethodDelete,
					"http://localhost:8080/af/v1/subscriptions/11114",
					nil)
				Expect(err).ShouldNot(HaveOccurred())

				resp := httptest.NewRecorder()
				ctx := context.WithValue(req.Context(),
					KeyType("af-ctx"), af.AfCtx)
				af.AfRouter.ServeHTTP(resp, req.WithContext(ctx))
				Expect(resp.Code).To(Equal(http.StatusNotFound))
			})
		})

		PContext("PFD  GET ALL - NO PFDS ", func() {
			PSpecify("Read all PFD transactions", func() {

				req, err := http.NewRequest(http.MethodGet,
					"http://localhost:8080/af/v1/pfd/transactions",
					nil)
				Expect(err).ShouldNot(HaveOccurred())

				resp := httptest.NewRecorder()
				ctx := context.WithValue(req.Context(),
					KeyType("af-ctx"), af.AfCtx)
				af.AfRouter.ServeHTTP(resp, req.WithContext(ctx))

				Expect(resp.Code).To(Equal(http.StatusOK))
			})
		})

		PContext("PFD Transaction POST", func() {
			PSpecify("Sending PFD POST 001 request", func() {
				By("Reading json file")
				reqBody, err := ioutil.ReadFile(
					"./testdata/pfd/AF_NB_PFD_POST001.json")
				Expect(err).ShouldNot(HaveOccurred())

				By("Preparing request")
				reqBodyBytes := bytes.NewReader(reqBody)
				req, err := http.NewRequest(http.MethodPost,
					"http://localhost:8080/af/v1/pfd/transactions",
					reqBodyBytes)
				Expect(err).ShouldNot(HaveOccurred())

				By("Sending request")
				resp := httptest.NewRecorder()
				ctx := context.WithValue(req.Context(),
					KeyType("af-ctx"), af.AfCtx)
				af.AfRouter.ServeHTTP(resp, req.WithContext(ctx))

				Expect(resp.Code).To(Equal(http.StatusCreated))

			})

		})

		PContext("PFD Transaction POST", func() {
			PSpecify("Sending PFD POST 002 request", func() {
				By("Reading json file")
				reqBody, err := ioutil.ReadFile(
					"./testdata/pfd/AF_NB_PFD_POST002.json")
				Expect(err).ShouldNot(HaveOccurred())

				By("Preparing request")
				reqBodyBytes := bytes.NewReader(reqBody)
				req, err := http.NewRequest(http.MethodPost,
					"http://localhost:8080/af/v1/pfd/transactions",
					reqBodyBytes)
				Expect(err).ShouldNot(HaveOccurred())

				By("Sending request")
				resp := httptest.NewRecorder()
				ctx := context.WithValue(req.Context(),
					KeyType("af-ctx"), af.AfCtx)
				af.AfRouter.ServeHTTP(resp, req.WithContext(ctx))

				Expect(resp.Code).To(Equal(http.StatusCreated))

			})

		})

		PContext("PFD Transaction INVALID POST", func() {
			PSpecify("Sending PFD POST request", func() {
				By("Reading json file")
				reqBody, err := ioutil.ReadFile(
					"./testdata/pfd/AF_NB_PFD_POST001.json")
				Expect(err).ShouldNot(HaveOccurred())

				By("Preparing request")
				reqBodyBytes := bytes.NewReader(reqBody)
				req, err := http.NewRequest(http.MethodPost,
					"http://localhost:8080/af/v1/pfd/transactions",
					reqBodyBytes)
				Expect(err).ShouldNot(HaveOccurred())

				By("Sending request")
				resp := httptest.NewRecorder()
				ctx := context.WithValue(req.Context(),
					KeyType("af-ctx"), af.AfCtx)
				af.AfRouter.ServeHTTP(resp, req.WithContext(ctx))

				Expect(resp.Code).To(Equal(http.StatusInternalServerError))

			})

		})

		PContext("PFD Transaction INVALID POST", func() {
			PSpecify("Decode error", func() {
				By("Reading json file")
				reqBody, err := ioutil.ReadFile(
					"./testdata/pfd/AF_NB_PFD_POST003.json")
				Expect(err).ShouldNot(HaveOccurred())

				By("Preparing request")
				reqBodyBytes := bytes.NewReader(reqBody)
				req, err := http.NewRequest(http.MethodPost,
					"http://localhost:8080/af/v1/pfd/transactions",
					reqBodyBytes)
				Expect(err).ShouldNot(HaveOccurred())

				By("Sending request")
				resp := httptest.NewRecorder()
				ctx := context.WithValue(req.Context(),
					KeyType("af-ctx"), af.AfCtx)
				af.AfRouter.ServeHTTP(resp, req.WithContext(ctx))

				Expect(resp.Code).To(Equal(http.StatusInternalServerError))

			})

		})

		PContext("PFD  GET ALL", func() {
			PSpecify("Read all PFD transactions", func() {

				req, err := http.NewRequest(http.MethodGet,
					"http://localhost:8080/af/v1/pfd/transactions",
					nil)
				Expect(err).ShouldNot(HaveOccurred())

				resp := httptest.NewRecorder()
				ctx := context.WithValue(req.Context(),
					KeyType("af-ctx"), af.AfCtx)
				af.AfRouter.ServeHTTP(resp, req.WithContext(ctx))

				Expect(resp.Code).To(Equal(http.StatusOK))
			})

			PSpecify("Read all PFD transactions", func() {
				By("sending wrong url")
				req, err := http.NewRequest(http.MethodGet,
					"http://localhost:8080/af/v2/pfd/transactions",
					nil)
				Expect(err).ShouldNot(HaveOccurred())

				resp := httptest.NewRecorder()
				ctx := context.WithValue(req.Context(),
					KeyType("af-ctx"), af.AfCtx)
				af.AfRouter.ServeHTTP(resp, req.WithContext(ctx))

				Expect(resp.Code).To(Equal(http.StatusNotFound))

			})

		})

		PContext("PFD transaction ID GET", func() {
			PSpecify("", func() {
				req, err := http.NewRequest(http.MethodGet,
					"http://localhost:8080/af/v1/pfd/transactions/10000",
					nil)
				Expect(err).ShouldNot(HaveOccurred())

				resp := httptest.NewRecorder()
				ctx := context.WithValue(req.Context(),
					KeyType("af-ctx"), af.AfCtx)
				af.AfRouter.ServeHTTP(resp, req.WithContext(ctx))

				Expect(resp.Code).To(Equal(http.StatusOK))

			})

			PSpecify("INVALID GET PFD TRANS", func() {
				req, err := http.NewRequest(http.MethodGet,
					"http://localhost:8080/af/v1/pfd/transactions/11000",
					nil)
				Expect(err).ShouldNot(HaveOccurred())

				resp := httptest.NewRecorder()
				ctx := context.WithValue(req.Context(),
					KeyType("af-ctx"), af.AfCtx)
				af.AfRouter.ServeHTTP(resp, req.WithContext(ctx))

				Expect(resp.Code).To(Equal(http.StatusNotFound))

			})

		})

		PContext("PFD Transcation ID PUT", func() {
			PSpecify("", func() {
				By("Reading json file")
				reqBody, err := ioutil.ReadFile(
					"./testdata/pfd/AF_NB_PFD_PUT001.json")
				Expect(err).ShouldNot(HaveOccurred())

				By("Preparing request")
				reqBodyBytes := bytes.NewReader(reqBody)
				req, err := http.NewRequest(http.MethodPut,
					"http://localhost:8080/af/v1/pfd/transactions/10000",
					reqBodyBytes)
				Expect(err).ShouldNot(HaveOccurred())

				By("Sending request")
				resp := httptest.NewRecorder()
				ctx := context.WithValue(req.Context(),
					KeyType("af-ctx"), af.AfCtx)
				af.AfRouter.ServeHTTP(resp, req.WithContext(ctx))

				Expect(resp.Code).To(Equal(http.StatusOK))

			})

			PSpecify("INVALID PUT", func() {
				By("Reading json file")
				reqBody, err := ioutil.ReadFile(
					"./testdata/pfd/AF_NB_PFD_PUT002.json")
				Expect(err).ShouldNot(HaveOccurred())

				By("Preparing request")
				reqBodyBytes := bytes.NewReader(reqBody)
				req, err := http.NewRequest(http.MethodPut,
					"http://localhost:8080/af/v1/pfd/transactions/10000",
					reqBodyBytes)
				Expect(err).ShouldNot(HaveOccurred())

				By("Sending request")
				resp := httptest.NewRecorder()
				ctx := context.WithValue(req.Context(),
					KeyType("af-ctx"), af.AfCtx)
				af.AfRouter.ServeHTTP(resp, req.WithContext(ctx))

				Expect(resp.Code).To(Equal(http.StatusBadRequest))

			})

			PSpecify("INVALID PUT - Decode", func() {
				By("Reading json file")
				reqBody, err := ioutil.ReadFile(
					"./testdata/pfd/AF_NB_PFD_POST003.json")
				Expect(err).ShouldNot(HaveOccurred())

				By("Preparing request")
				reqBodyBytes := bytes.NewReader(reqBody)
				req, err := http.NewRequest(http.MethodPut,
					"http://localhost:8080/af/v1/pfd/transactions/10000",
					reqBodyBytes)
				Expect(err).ShouldNot(HaveOccurred())

				By("Sending request")
				resp := httptest.NewRecorder()
				ctx := context.WithValue(req.Context(),
					KeyType("af-ctx"), af.AfCtx)
				af.AfRouter.ServeHTTP(resp, req.WithContext(ctx))

				Expect(resp.Code).To(Equal(http.StatusInternalServerError))

			})

		})

		PContext("PFD Transcation DELETE", func() {
			PSpecify("DELETE PFD Transaction 02", func() {

				req, err := http.NewRequest(http.MethodDelete,
					"http://localhost:8080/af/v1/pfd/transactions/10001",
					nil)
				Expect(err).ShouldNot(HaveOccurred())

				resp := httptest.NewRecorder()
				ctx := context.WithValue(req.Context(),
					KeyType("af-ctx"), af.AfCtx)
				af.AfRouter.ServeHTTP(resp, req.WithContext(ctx))
				Expect(resp.Code).To(Equal(http.StatusNoContent))
			})
			PSpecify("INVALID DELETE PFD Transaction 10", func() {

				req, err := http.NewRequest(http.MethodDelete,
					"http://localhost:8080/af/v1/pfd/transactions/11000",
					nil)
				Expect(err).ShouldNot(HaveOccurred())

				resp := httptest.NewRecorder()
				ctx := context.WithValue(req.Context(),
					KeyType("af-ctx"), af.AfCtx)
				af.AfRouter.ServeHTTP(resp, req.WithContext(ctx))
				Expect(resp.Code).To(Equal(http.StatusNotFound))
			})

		})

		PContext("PFD transaction Application GET", func() {
			PSpecify("", func() {
				req, err := http.NewRequest(http.MethodGet,
					"http://localhost:8080/af/v1/pfd/transactions/10000/"+
						"applications/app1", nil)
				Expect(err).ShouldNot(HaveOccurred())

				resp := httptest.NewRecorder()
				ctx := context.WithValue(req.Context(),
					KeyType("af-ctx"), af.AfCtx)
				af.AfRouter.ServeHTTP(resp, req.WithContext(ctx))

				Expect(resp.Code).To(Equal(http.StatusOK))

			})

			PSpecify("INVALID GET PFD TRANS 10000 and app10", func() {
				req, err := http.NewRequest(http.MethodGet,
					"http://localhost:8080/af/v1/pfd/transactions/10000/"+
						"applications/app10", nil)
				Expect(err).ShouldNot(HaveOccurred())

				resp := httptest.NewRecorder()
				ctx := context.WithValue(req.Context(),
					KeyType("af-ctx"), af.AfCtx)
				af.AfRouter.ServeHTTP(resp, req.WithContext(ctx))

				Expect(resp.Code).To(Equal(http.StatusNotFound))

			})

		})

		PContext("PFD transaction Application PUT", func() {
			PSpecify("", func() {

				By("Reading json file")
				reqBody, err := ioutil.ReadFile(
					"./testdata/pfd/AF_NB_PFD_APP_PUT_01.json")
				Expect(err).ShouldNot(HaveOccurred())

				By("Preparing request")
				reqBodyBytes := bytes.NewReader(reqBody)
				req, err := http.NewRequest(http.MethodPut,
					"http://localhost:8080/af/v1/pfd/transactions/10000/"+
						"applications/app1", reqBodyBytes)
				Expect(err).ShouldNot(HaveOccurred())

				By("Sending request")
				resp := httptest.NewRecorder()
				ctx := context.WithValue(req.Context(),
					KeyType("af-ctx"), af.AfCtx)
				af.AfRouter.ServeHTTP(resp, req.WithContext(ctx))

				Expect(resp.Code).To(Equal(http.StatusOK))

			})

			PSpecify("INVALID PUT FOR PFD TRANS 10000 and app1", func() {

				By("Reading json file")
				reqBody, err := ioutil.ReadFile(
					"./testdata/pfd/AF_NB_PFD_APP_PUT_02.json")
				Expect(err).ShouldNot(HaveOccurred())

				By("Preparing request")
				reqBodyBytes := bytes.NewReader(reqBody)
				req, err := http.NewRequest(http.MethodPut,
					"http://localhost:8080/af/v1/pfd/transactions/10000/"+
						"applications/app1", reqBodyBytes)
				Expect(err).ShouldNot(HaveOccurred())

				By("Sending request")
				resp := httptest.NewRecorder()
				ctx := context.WithValue(req.Context(),
					KeyType("af-ctx"), af.AfCtx)
				af.AfRouter.ServeHTTP(resp, req.WithContext(ctx))

				Expect(resp.Code).To(Equal(http.StatusBadRequest))

			})

			PSpecify("INVALID PUT FOR PFD TRANS/ APP - Decode error", func() {

				By("Reading json file")
				reqBody, err := ioutil.ReadFile(
					"./testdata/pfd/AF_NB_PFD_APP_PUT_03.json")
				Expect(err).ShouldNot(HaveOccurred())

				By("Preparing request")
				reqBodyBytes := bytes.NewReader(reqBody)
				req, err := http.NewRequest(http.MethodPut,
					"http://localhost:8080/af/v1/pfd/transactions/10000/"+
						"applications/app1", reqBodyBytes)
				Expect(err).ShouldNot(HaveOccurred())

				By("Sending request")
				resp := httptest.NewRecorder()
				ctx := context.WithValue(req.Context(),
					KeyType("af-ctx"), af.AfCtx)
				af.AfRouter.ServeHTTP(resp, req.WithContext(ctx))

				Expect(resp.Code).To(Equal(http.StatusInternalServerError))

			})

		})

		PContext("PFD transaction Application PATCH", func() {
			PSpecify("", func() {

				By("Reading json file")
				reqBody, err := ioutil.ReadFile(
					"./testdata/pfd/AF_NB_PFD_APP_PATCH_01.json")
				Expect(err).ShouldNot(HaveOccurred())

				By("Preparing request")
				reqBodyBytes := bytes.NewReader(reqBody)
				req, err := http.NewRequest(http.MethodPatch,
					"http://localhost:8080/af/v1/pfd/transactions/10000/"+
						"applications/app1", reqBodyBytes)
				Expect(err).ShouldNot(HaveOccurred())

				By("Sending request")
				resp := httptest.NewRecorder()
				ctx := context.WithValue(req.Context(),
					KeyType("af-ctx"), af.AfCtx)
				af.AfRouter.ServeHTTP(resp, req.WithContext(ctx))

				Expect(resp.Code).To(Equal(http.StatusOK))

			})

			PSpecify("INVALID PATCH PFD TRANS 10000 and app1", func() {

				By("Reading json file")
				reqBody, err := ioutil.ReadFile(
					"./testdata/pfd/AF_NB_PFD_APP_PATCH_02.json")
				Expect(err).ShouldNot(HaveOccurred())

				By("Preparing request")
				reqBodyBytes := bytes.NewReader(reqBody)
				req, err := http.NewRequest(http.MethodPatch,
					"http://localhost:8080/af/v1/pfd/transactions/10000/"+
						"applications/app1", reqBodyBytes)
				Expect(err).ShouldNot(HaveOccurred())

				By("Sending request")
				resp := httptest.NewRecorder()
				ctx := context.WithValue(req.Context(),
					KeyType("af-ctx"), af.AfCtx)
				af.AfRouter.ServeHTTP(resp, req.WithContext(ctx))

				Expect(resp.Code).To(Equal(http.StatusBadRequest))

			})

			PSpecify("INVALID PATCH FOR PFD TRANS/ APP - Decode error", func() {

				By("Reading json file")
				reqBody, err := ioutil.ReadFile(
					"./testdata/pfd/AF_NB_PFD_APP_PUT_03.json")
				Expect(err).ShouldNot(HaveOccurred())

				By("Preparing request")
				reqBodyBytes := bytes.NewReader(reqBody)
				req, err := http.NewRequest(http.MethodPatch,
					"http://localhost:8080/af/v1/pfd/transactions/10000/"+
						"applications/app1", reqBodyBytes)
				Expect(err).ShouldNot(HaveOccurred())

				By("Sending request")
				resp := httptest.NewRecorder()
				ctx := context.WithValue(req.Context(),
					KeyType("af-ctx"), af.AfCtx)
				af.AfRouter.ServeHTTP(resp, req.WithContext(ctx))

				Expect(resp.Code).To(Equal(http.StatusInternalServerError))

			})

		})

		PContext("PFD transaction Application DELETE", func() {
			PSpecify("", func() {
				req, err := http.NewRequest(http.MethodDelete,
					"http://localhost:8080/af/v1/pfd/transactions/10000/"+
						"applications/app1", nil)
				Expect(err).ShouldNot(HaveOccurred())

				resp := httptest.NewRecorder()
				ctx := context.WithValue(req.Context(),
					KeyType("af-ctx"), af.AfCtx)
				af.AfRouter.ServeHTTP(resp, req.WithContext(ctx))

				Expect(resp.Code).To(Equal(http.StatusNoContent))

			})

			PSpecify("INVALID DELETE TRANSACTION APPLICATION", func() {
				req, err := http.NewRequest(http.MethodDelete,
					"http://localhost:8080/af/v1/pfd/transactions/10000/"+
						"applications/app10", nil)
				Expect(err).ShouldNot(HaveOccurred())

				resp := httptest.NewRecorder()
				ctx := context.WithValue(req.Context(),
					KeyType("af-ctx"), af.AfCtx)
				af.AfRouter.ServeHTTP(resp, req.WithContext(ctx))

				Expect(resp.Code).To(Equal(http.StatusNotFound))

			})

		})

	})

	Describe("Stop the AF Server", func() {
		It("Disconnect AF Server", func() {
			srvCancel()
		})
	})
})
