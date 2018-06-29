package main

import (
	"encoding/json"
	"fmt"
	"gokit-grpc/car-client/client"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Startig client")
	router := mux.NewRouter()

	router.HandleFunc("/car/{id}", func(w http.ResponseWriter, req *http.Request) {
		params := mux.Vars(req)

		cli := client.NewCarGRPCClient("localhost:9090")
		id, err := strconv.ParseInt(params["id"], 10, 32)
		if err != nil {
			fmt.Errorf("id must be an intgeger")
		}
		var result = cli.GetCar(int32(id))

		encodeCarResponse(w, carResponse{
			Id:           result.Id,
			Name:         result.Name,
			Manifacturer: result.Manifacturer,
		})
	}).Methods("GET")

	router.HandleFunc("/car", func(w http.ResponseWriter, req *http.Request) {
		request, _ := decodeCreateCarRequest(req)
		params := request.(createCarRequest)

		cli := client.NewCarGRPCClient("localhost:9090")
		var result = cli.CreateCar(params.Name, params.Manifacturer)

		encodeCarResponse(w, carResponse{
			Id:           result.Id,
			Name:         result.Name,
			Manifacturer: result.Manifacturer,
		})
	}).Methods("POST")

	http.ListenAndServe(":3000", router)
}

func decodeCreateCarRequest(r *http.Request) (interface{}, error) {
	var request createCarRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func encodeCarResponse(w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

type createCarRequest struct {
	Name         string `json:"name"`
	Manifacturer string `json:"manifacturer"`
}

type carResponse struct {
	Id           int32  `json:"id"`
	Name         string `json:"name"`
	Manifacturer string `json:"manifacturer"`
}
