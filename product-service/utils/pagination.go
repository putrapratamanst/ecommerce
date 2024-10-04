package utils

import "math"

type Pagination struct {
    TotalRecords int64 `json:"totalRecords"`
    TotalPages   int   `json:"totalPages"`
    CurrentPage  int   `json:"currentPage"`
    Limit        int   `json:"limit"`
}

func Paginate(totalRecords int64, currentPage, limit int) Pagination {
    totalPages := int(math.Ceil(float64(totalRecords) / float64(limit)))

    return Pagination{
        TotalRecords: totalRecords,
        TotalPages:   totalPages,
        CurrentPage:  currentPage,
        Limit:        limit,
    }
}
