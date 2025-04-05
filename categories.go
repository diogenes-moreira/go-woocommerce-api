package woocommerce

import "net/http"

type CategoriesService service

type Category struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Slug        string `json:"slug"`
	Parent      int    `json:"parent"`
	Description string `json:"description"`
	Display     string `json:"display"`
	Image       Image  `json:"image"`
	MenuOrder   int    `json:"menu_order"`
	Count       int    `json:"count"`
	Links       Links  `json:"_links"`
}

func (service *CategoriesService) Get(couponID string) (*Category, *http.Response, error) {
	url := "/products/categories/" + couponID
	req, _ := service.client.NewRequest("GET", url, nil, nil)

	category := new(Category)
	response, err := service.client.Do(req, category)

	if err != nil {
		return nil, response, err
	}

	return category, response, nil
}

type ListCategoriesParams struct {
	Context string `url:"context,omitempty"`
	Page    int    `url:"page,omitempty"`
	PerPage int    `url:"per_page,omitempty"`
	Search  string `url:"search,omitempty"`
	Exclude *[]int `url:"exclude,omitempty"`
	Include *[]int `url:"include,omitempty"`
	Offset  int    `url:"offset,omitempty"`
	Order   string `url:"order,omitempty"`
	OrderBy string `url:"orderby,omitempty"`

	HideEmpty bool   `url:"hide_empty,omitempty"`
	Parent    int    `url:"parent,omitempty"`
	Products  int    `url:"products,omitempty"`
	Slug      string `url:"slug,omitempty"`
}

func (service *CategoriesService) List(opts *ListCategoriesParams) (*[]Category, *http.Response, error) {
	url := "/products/categories"
	req, _ := service.client.NewRequest("GET", url, opts, nil)

	categories := new([]Category)
	response, err := service.client.Do(req, categories)

	if err != nil {
		return nil, response, err
	}

	return categories, response, nil
}

func (service *CategoriesService) Update(categoryID string, category *Category) (*Category, *http.Response, error) {
	url := "/products/categories/" + categoryID
	req, _ := service.client.NewRequest("PUT", url, nil, category)

	updatedCategory := new(Category)
	response, err := service.client.Do(req, updatedCategory)

	if err != nil {
		return nil, response, err
	}

	return updatedCategory, response, nil
}

type DeleteCategoryParams struct {
	Force string `json:"force,omitempty"`
}

func (service *CategoriesService) Delete(categoryID string, param *DeleteCategoryParams) (*http.Response, error) {
	url := "/products/categories/" + categoryID
	req, _ := service.client.NewRequest("DELETE", url, param, nil)

	category := new(Category)
	response, err := service.client.Do(req, category)

	if err != nil {
		return response, err
	}

	return response, nil
}

func (service *CategoriesService) Create(category *Category) (*Category, *http.Response, error) {
	url := "/products/categories"
	req, _ := service.client.NewRequest("POST", url, nil, category)

	createdCategory := new(Category)
	response, err := service.client.Do(req, createdCategory)

	if err != nil {
		return nil, response, err
	}

	return createdCategory, response, nil

}

type BatchCategoryUpdate struct {
	Create *[]Category `json:"create,omitempty"`
	Update *[]Category `json:"update,omitempty"`
	Delete *[]int      `json:"delete,omitempty"`
}

type BatchCategoryUpdateResponse struct {
	Create *[]Category `json:"create,omitempty"`
	Update *[]Category `json:"update,omitempty"`
	Delete *[]Category `json:"delete,omitempty"`
}

func (service *CategoriesService) Batch(opts *BatchCategoryUpdate) (*BatchCategoryUpdateResponse, *http.Response, error) {
	url := "/products/categories/batch"
	req, _ := service.client.NewRequest("POST", url, nil, opts)

	batchCategory := new(BatchCategoryUpdateResponse)
	response, err := service.client.Do(req, batchCategory)

	if err != nil {
		return nil, response, err
	}

	return batchCategory, response, nil

}
