package cart

import (
	"fmt"
	"net/http"

	"github.com/dekko911/start-with-goLang/service/auth"
	"github.com/dekko911/start-with-goLang/types"
	"github.com/dekko911/start-with-goLang/utils"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

type Handler struct {
	store      types.ProductStore
	userStore  types.UserStore
	orderStore types.OrderStore
}

func NewHandler(store types.ProductStore, userStore types.UserStore, orderStore types.OrderStore) *Handler {
	return &Handler{store: store, userStore: userStore, orderStore: orderStore}
}

func (h *Handler) RegisterRoutes(r *mux.Router) {
	r.HandleFunc("/cart/checkout", auth.WithJWTAuth(h.handleCheckout, h.userStore)).Methods(http.MethodPost)
}

func (h *Handler) handleCheckout(w http.ResponseWriter, r *http.Request) {
	userID := auth.GetUserIDFromContext(r.Context())

	var payload types.CartCheckoutPayload
	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if err := utils.Validate.Struct(payload); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload: %v", errors))
		return
	}

	productIDs, err := getCartItemsIDs(payload.Items)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	// GET products
	products, err := h.store.GetProductsByID(productIDs)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	orderID, totalPrice, err := h.createOrder(products, payload.Items, userID)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, map[string]any{
		"status":      http.StatusOK,
		"total_price": totalPrice,
		"order_id":    orderID,
	})

}
