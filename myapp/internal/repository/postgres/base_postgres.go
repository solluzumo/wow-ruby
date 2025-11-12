package postgres

import (
	"context"
	"fmt"
	"reflect"
	"strings"
	"wow-ruby/internal/dto"
	"wow-ruby/internal/models"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type BaseRepository[T models.Tabler] struct {
	DB *sqlx.DB
}

func NewBaseRepository[T models.Tabler](db *sqlx.DB) *BaseRepository[T] {
	return &BaseRepository[T]{DB: db}
}

func (br *BaseRepository[T]) GetById(ctx context.Context, idString string) (*T, error) {

	var model T

	tableName := model.TableName()

	parsedId, err := uuid.Parse(idString)

	if err != nil {
		return nil, err
	}

	query := fmt.Sprintf("SELECT * FROM %s WHERE id = $1", tableName)

	err = br.DB.Get(&model, query, parsedId)

	if err != nil {
		return nil, err
	}

	return &model, nil
}

func (br *BaseRepository[T]) GetList(ctx context.Context, req *dto.ListRequest) (*dto.ListResponse[T], error) {
	var total int
	var model T
	var models []T

	tableName := model.TableName()
	baseQuery := fmt.Sprintf("FROM %s WHERE 1=1", tableName)

	args := []interface{}{}

	filterQuery := "SELECT * " + baseQuery

	filterQuery, args = br.applyFilters(filterQuery, args, req.Filters)

	countQuery := "SELECT COUNT(*) " + baseQuery
	countQuery, countArgs := br.applyFilters(countQuery, []interface{}{}, req.Filters)

	if req.Search != "" {
		searchParam := "%" + strings.ToLower(req.Search) + "%"
		filterQuery += " AND lower(name) LIKE $" + fmt.Sprintf("%d", len(args)+1)
		args = append(args, searchParam)

		countQuery += " AND lower(name) LIKE $" + fmt.Sprintf("%d", len(countArgs)+1)
		countArgs = append(countArgs, searchParam)
	}

	if err := br.DB.GetContext(ctx, &total, countQuery, countArgs...); err != nil {
		return nil, err
	}

	if req.SortBy != "" {
		safeSortBy := br.validateSortField(req.SortBy)
		safeSortOrder := br.validateSortOrder(req.SortOrder)
		filterQuery += " ORDER BY " + safeSortBy + " " + safeSortOrder
	} else {
		filterQuery += " ORDER BY name DESC"
	}

	// Пагинация
	filterQuery += " LIMIT $" + fmt.Sprintf("%d", len(args)+1) + " OFFSET $" + fmt.Sprintf("%d", len(args)+2)
	args = append(args, req.PageSize, (req.Page-1)*req.PageSize)

	// Выборка
	if err := br.DB.SelectContext(ctx, &models, filterQuery, args...); err != nil {
		return nil, err
	}

	// Ответ
	return &dto.ListResponse[T]{
		Data:       models,
		Total:      total,
		Page:       req.Page,
		PageSize:   req.PageSize,
		TotalPages: (total + req.PageSize - 1) / req.PageSize,
	}, nil
}

func (br *BaseRepository[T]) applyFilters(
	baseQuery string,
	args []interface{},
	filters map[string]interface{},
) (string, []interface{}) {

	if len(filters) == 0 {
		return baseQuery, args
	}

	query := baseQuery
	namedParams := make(map[string]interface{})
	filterParts := make([]string, 0, len(filters))
	filterSliceParts := make(map[string]interface{})
	i := 0

	for field, value := range filters {

		columnName := field

		if value == nil {
			continue
		}

		val := reflect.ValueOf(value)

		if val.Kind() == reflect.Slice || val.Kind() == reflect.Array {
			val := reflect.ValueOf(value)
			partsArray := make([]string, 0, val.Len())
			for j := i; j < val.Len()+i; j++ {
				paramName := fmt.Sprintf("filter_%d", j)
				partsArray = append(partsArray, fmt.Sprintf("%s = :%s", columnName, paramName))
				namedParams[paramName] = val.Index(j - i).Interface()
			}
			filterSliceParts[columnName] = partsArray
			partsArray = nil
			i += val.Len()
			continue
		}

		paramName := fmt.Sprintf("filter_%d", i)

		filterParts = append(filterParts, fmt.Sprintf("%s = :%s", columnName, paramName))

		namedParams[paramName] = value

		i++
	}

	if (len(filterParts) == 0) && (len(filterSliceParts) == 0) {
		return baseQuery, args
	}

	for field := range filterSliceParts {

		var array []string
		if strSlice, ok := filterSliceParts[field].([]string); ok {
			array = strSlice
			filterParts = append(filterParts, "("+strings.Join(array, " OR ")+")")
		} else {
			fmt.Println(ok)
		}

	}

	query += " AND " + strings.Join(filterParts, " AND ")
	fmt.Println(query)
	finalQuery, finalArgs, err := sqlx.Named(query, namedParams)
	if err != nil {
		return baseQuery, args
	}

	fmt.Println(finalQuery)

	finalQuery = sqlx.Rebind(sqlx.DOLLAR, finalQuery)

	args = append(args, finalArgs...)

	return finalQuery, args
}

func (br *BaseRepository[T]) validateSortField(field string) string {
	allowedFields := []string{"name", "price", "required_level", "rarity", "item_type",
		"max_stack", "reward_money", "required_character_level", "quest_level", "is_repeatable", "difficulty",
		"health", "mana", "level", "tameable", "faction", "reaction", "location", "respawn_time"}

	for _, el := range allowedFields {
		if field == el {
			return field
		}
	}
	return "name"
}

func (br *BaseRepository[T]) validateSortOrder(order string) string {
	order = strings.ToUpper(order)
	if order == "ASC" || order == "DESC" {
		return order
	}
	return "DESC" // значение по умолчанию
}
