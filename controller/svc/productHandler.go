package svc

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

type ProductInfo struct {
	ProductKey   string `json:"ProductKey"`
	ProductName  string `json:"ProductName"`
	Description  string `json:"Description"`
	DeviceCount  int32  `json:"DeviceCount"`
	AccessPoints string `json:"AccessPoints"`
	CreateAt     string `json:"CreateAt"`
	UpdateAt     string `json:"UpdateAt"`

	ProductSecret string `json:"ProductSecret,omitempty"`
	DeleteFlag    int8
}

func handleCreateProduct(w http.ResponseWriter, r *http.Request) {
	productName := r.URL.Query().Get("ProductName")
	productDesc := r.URL.Query().Get("Description")
	// create product
	productInfo := &ProductInfo{
		ProductKey:    uuid.New().String(),
		ProductSecret: uuid.New().String(),
		ProductName:   productName,
		Description:   productDesc,
		DeviceCount:   0,
		AccessPoints:  "",
		CreateAt:      "",
		UpdateAt:      "",
		DeleteFlag:    0,
	}
	logrus.Infof("create product name : %s, description : %s", productInfo.ProductName, productInfo.Description)
	w.Header().Set("Content-Type", "application/json")
	_, err := Ctx.Dbsvc.CreateProduct(*productInfo)
	if err != nil {
		errInfo := &ErrInfo{
			Code:    "500",
			Message: err.Error(),
		}
		b, _ := json.Marshal(errInfo)
		w.Write(b)
		w.WriteHeader(http.StatusOK)
	} else {
		b, _ := json.Marshal(productInfo)
		w.Write(b)
		w.WriteHeader(http.StatusOK)
	}
}

func handleQueryProduct(w http.ResponseWriter, r *http.Request) {
	productKey := r.URL.Query().Get("ProductKey")
	logrus.Infof("query product info, product key : %s", productKey)
	w.Header().Set("Content-Type", "application/json")
	productInfo, err := Ctx.Dbsvc.QueryProductInfo(productKey)
	if err != nil {
		errInfo := &ErrInfo{
			Code:    "500",
			Message: err.Error(),
		}
		b, _ := json.Marshal(errInfo)
		w.Write(b)
		w.WriteHeader(http.StatusOK)
	} else {
		productInfo.ProductKey = productKey
		b, _ := json.Marshal(productInfo)
		w.Write(b)
		w.WriteHeader(http.StatusOK)
	}
}

func handleQueryProductList(w http.ResponseWriter, r *http.Request) {
}

func handleUpdateProduct(w http.ResponseWriter, r *http.Request) {
}

func handleDeleteProduct(w http.ResponseWriter, r *http.Request) {
}

func handleQueryProductQuota(w http.ResponseWriter, r *http.Request) {
}

func handleModifyProductQuota(w http.ResponseWriter, r *http.Request) {
}

func handleCheckProductName(w http.ResponseWriter, r *http.Request) {
}
