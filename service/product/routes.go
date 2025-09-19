package product

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/dekko911/start-with-goLang/service/auth"
	"github.com/dekko911/start-with-goLang/types"
	"github.com/dekko911/start-with-goLang/utils"
	"github.com/gorilla/mux"
)

type Handler struct {
	store     types.ProductStore
	userStore types.UserStore
}

func NewHandler(store types.ProductStore, userStore types.UserStore) *Handler {
	return &Handler{
		store:     store,
		userStore: userStore,
	}
}

func (h *Handler) RegisterRoutes(r *mux.Router) {
	r.HandleFunc("/products", h.handleGetProducts).Methods(http.MethodGet)
	r.HandleFunc("/products/{productID}", h.handleGetProductByID).Methods(http.MethodGet)

	// for admin routes
	r.HandleFunc("/products", auth.WithJWTAuth(h.handleCreateProduct, h.userStore)).Methods(http.MethodPost)
}

func (h *Handler) handleGetProducts(w http.ResponseWriter, r *http.Request) {
	products, err := h.store.GetProducts()
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, map[string]any{
		"status":   http.StatusOK,
		"products": products,
	})
}

func (h *Handler) handleGetProductByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	str, ok := vars["productID"]
	if !ok {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("missing id product"))
		return
	}

	productID, err := strconv.Atoi(str)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid type productID"))
		return
	}

	product, err := h.store.GetProductByID(productID)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, map[string]any{
		"status":  http.StatusOK,
		"product": product,
	})

}

func (h *Handler) handleCreateProduct(w http.ResponseWriter, r *http.Request) {
	// get JSON req payload
	payload := types.CreateProductPayload{}
	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	// validate the payload
	if err := utils.Validate.Struct(payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	// create the product
	err := h.store.CreateProduct(payload)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	utils.WriteJSON(w, http.StatusCreated, map[string]any{
		"status":  http.StatusCreated,
		"message": "Product Has Created!",
		"product": payload,
	})
}
