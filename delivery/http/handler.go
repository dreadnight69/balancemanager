package http

import (
	"encoding/json"
	"net/http"
	"strconv"
)

func (s *server) getBalance(w http.ResponseWriter, r *http.Request) {
	userID := r.URL.Query().Get("user_id")
	if userID == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	userIDint, err := strconv.ParseInt(userID, 10, 64)
	if userID == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	result, err := s.BalanceManager.GetUserInfo(r.Context(), userIDint)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(err.Error()))
		return
	}
	resp := userInfoToHttpResponse(result)
	if err := json.NewEncoder(w).Encode(&resp); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (s *server) makeDeposit(w http.ResponseWriter, r *http.Request) {
	var req UpdateBalanceRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	result, err := s.BalanceManager.MakeDeposit(r.Context(), req.UserID, req.Amount, req.Description)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(err.Error()))
		return
	}
	resp := userInfoToHttpResponse(result)
	if err := json.NewEncoder(w).Encode(&resp); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (s *server) withdrawFunds(w http.ResponseWriter, r *http.Request) {
	var req UpdateBalanceRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	result, err := s.BalanceManager.WithdrawFunds(r.Context(), req.UserID, req.Amount, req.Description)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(err.Error()))
		return
	}
	resp := userInfoToHttpResponse(result)
	if err := json.NewEncoder(w).Encode(&resp); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (s *server) sendFunds(w http.ResponseWriter, r *http.Request) {
	var req SendFundsRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	result, err := s.BalanceManager.SendFunds(r.Context(), req.SenderID, req.RecipientID, req.Amount, req.Description)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(err.Error()))
		return
	}
	resp := sendFundsResponseToHttp(result)
	if err := json.NewEncoder(w).Encode(&resp); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (s *server) getTransactions(w http.ResponseWriter, r *http.Request) {
	userID := r.URL.Query().Get("user_id")
	if userID == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	limit := r.URL.Query().Get("limit")
	if limit == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	cursor := r.URL.Query().Get("cursor")
	result, nextCursor, err := s.BalanceManager.GetTransactions(r.Context(), userID, limit, cursor)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(err.Error()))
		return
	}
	resp := transactionsToHttpResponse(result, nextCursor)
	if err := json.NewEncoder(w).Encode(&resp); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (s *server) createUser(w http.ResponseWriter, r *http.Request) {
	result, err := s.BalanceManager.CreateUser(r.Context())
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(err.Error()))
		return
	}
	resp := userInfoToHttpResponse(result)
	if err := json.NewEncoder(w).Encode(&resp); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
