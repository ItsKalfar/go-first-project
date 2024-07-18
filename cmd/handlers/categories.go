package handlers

import (
	"database/sql"
	"encoding/json"
	"firstproject/cmd/types"
	"firstproject/cmd/utils"
	"fmt"
	"net/http"
)

type Handler struct {
	db *sql.DB
}

func CategoryHandler(db *sql.DB) *Handler {
	return &Handler{db: db}
}

func (h *Handler) CategoryRouter(router *http.ServeMux){
	router.HandleFunc("POST /getCategories", h.GetCategories)
}

func (h *Handler) GetCategories(w http.ResponseWriter, r *http.Request){
	var reqData types.GetRequest

	if err := json.NewDecoder(r.Body).Decode(&reqData); err != nil {
		utils.SendResponse(w, http.StatusBadRequest, false, "Failed to decode request body: "+err.Error(), nil)
		return
	}

	defer r.Body.Close() 

	var searchQuery string
    if reqData.Search != "" {
        searchQuery = fmt.Sprintf("WHERE uc.name LIKE '%%%s%%'", reqData.Search)
    }

	countQuery := fmt.Sprintf(`SELECT uc.id,uc.name,uc.added_by,uc.icon FROM categories as uc %s`, searchQuery)

	count, err := h.db.Query(countQuery)
	if err != nil {
        utils.SendResponse(w, http.StatusBadRequest, false, err.Error(), nil)
        return
    }

	defer count.Close()

	var orderQuery string
    if reqData.OrderBy.Key != "" {
        orderQuery = fmt.Sprintf("ORDER BY %s %s", reqData.OrderBy.Key, reqData.OrderBy.Order)
    } else {
        orderQuery = "ORDER BY uc.id, uc.is_approved ASC"
    }

	limit := reqData.PageSize
    offset := (reqData.PageIndex - 1) * reqData.PageSize

	query := fmt.Sprintf(`
    SELECT uc.id, uc.name, uc.added_by, uc.icon as icon_id,
           ci.icon_color as frg_color, ci.icon_bgcolor as bg_color,
           ci.icon_url as url, uc.active, uc.description, uc.is_approved
    FROM categories as uc
    INNER JOIN category_icons as ci on ci.category_icon_id = uc.icon
    %s %s
    LIMIT %d OFFSET %d
	`, searchQuery, orderQuery, limit, offset)

    rows, err := h.db.Query(query)
    if err != nil {
		utils.SendResponse(w, http.StatusBadRequest, false, err.Error(), nil)
        return
    }
    defer rows.Close()

	var categories []types.Category

	for rows.Next() {
		var category types.Category
		err := rows.Scan(
			&category.ID, &category.Name, &category.AddedBy, &category.IconID,
			&category.FrgColor, &category.BgColor, &category.URL,
			&category.Active, &category.Description, &category.IsApproved,
		)
		if err != nil {
			utils.SendResponse(w, http.StatusBadRequest, false, err.Error(), nil)
			return
		}
		
		categories = append(categories, category)
	}
	
	if err = rows.Err(); err != nil {
		utils.SendResponse(w, http.StatusBadRequest, false, err.Error(), nil)
        return
	}

	result := map[string]interface{}{
        "data":  categories,
        "count": count,
    }

	utils.SendResponse(w, http.StatusOK, true, "Categories fetched successfully", result)
}