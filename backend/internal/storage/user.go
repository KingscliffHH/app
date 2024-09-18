package storage

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/Infinities-ICT-Solutions/project-dashboard/data"
)

type UserStorage interface {
	GetById(hex string) (*data.User, error)
	GetAll() ([]*data.User, error)
	Create(user *data.User) (*data.User, error)
	Update(hex string, user *data.User) (*data.User, error)
	Delete(hex string) error
}

type auth0Storage struct {
	Audience     string
	ClientID     string
	ClientSecret string
	Domain       string
	PStorage     PreferenceStorage
	MemberRoleId string
	ClientRoleId string
}

type Auth0User struct {
	Email        string `json:"email"`
	Name         string `json:"name"`
	Nickname     string `json:"nickname,omitempty"`
	Picture      string `json:"picture,omitempty"`
	UpdatedAt    string `json:"updated_at,omitempty"`
	UserID       string `json:"user_id,omitempty"`
	LastLogin    string `json:"last_login,omitempty"`
	Connection   string `json:"connection,omitempty"`
	Password     string `json:"password,omitempty"`
	UserMetadata struct {
		UserRole     string `json:"user_role"`
		Bio          string `json:"bio"`
		Organisation string `json:"organisation"`
		ClientRole   string `json:"client_role"`
	} `json:"user_metadata"`
}

func NewUserStorage(Domain, ClientID, ClientSecret, Audience string, pst PreferenceStorage) (UserStorage, error) {
	storage := &auth0Storage{
		Audience:     Audience,
		ClientID:     ClientID,
		ClientSecret: ClientSecret,
		Domain:       Domain,
		PStorage:     pst,
	}

	token, err := storage.getToken()
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	// get roles
	url := fmt.Sprintf("https://%s/api/v2/roles", Domain)
	fmt.Println("url", url)
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", token)

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	fmt.Println("body", string(body))
	// fmt.Println(string(body))

	roles := []*struct {
		ID          string `json:"id"`
		Name        string `json:"name"`
		Description string `json:"description"`
	}{}
	err = json.Unmarshal(body, &roles)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	for _, r := range roles {
		if r.Name == "member" {
			storage.MemberRoleId = r.ID
		} else if r.Name == "client" {
			storage.ClientRoleId = r.ID
		}
	}

	if storage.MemberRoleId == "" || storage.ClientRoleId == "" {
		return nil, fmt.Errorf("auth0 roles not found")
	}

	return storage, nil

}

func (p *auth0Storage) getToken() (string, error) {
	url := fmt.Sprintf("https://%s/oauth/token", p.Domain)

	data := map[string]string{
		"client_id":     p.ClientID,
		"client_secret": p.ClientSecret,
		"audience":      p.Audience,
		"grant_type":    "client_credentials",
	}

	payload, err := json.Marshal(data)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))
	if err != nil {
		fmt.Println("req", err)
		return "", err
	}

	req.Header.Add("content-type", "application/json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("do", err)
		return "", err
	}

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("readall", err)
		return "", err
	}

	type auth0Token struct {
		AccessToken string `json:"access_token"`
		TokenType   string `json:"token_type"`
	}
	token := &auth0Token{}
	err = json.Unmarshal(body, token)
	if err != nil {
		fmt.Println("unmarshal", err)
		return "", err
	}

	return fmt.Sprintf("%s %s", token.TokenType, token.AccessToken), nil
}

func (p *auth0Storage) GetById(hex string) (*data.User, error) {
	fmt.Println("get by id", hex)
	token, err := p.getToken()
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	url := fmt.Sprintf("https://%s/api/v2/users/%s", p.Domain, hex)
	fmt.Println("url", url)
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", token)

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	fmt.Println("body", string(body))
	// fmt.Println(string(body))
	user := &Auth0User{}
	err = json.Unmarshal(body, user)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	fmt.Println("user", user)

	// convert auth user to data user
	userData := &data.User{
		ID:           user.UserID,
		Type:         user.UserMetadata.UserRole,
		FullName:     user.Name,
		Email:        user.Email,
		Avatar:       user.Picture,
		Bio:          user.UserMetadata.Bio,
		Organisation: user.UserMetadata.Organisation,
		ClientRole:   user.UserMetadata.ClientRole,
		LastAccess:   user.LastLogin,
	}

	return userData, nil
}

func (p *auth0Storage) GetAll() ([]*data.User, error) {
	token, err := p.getToken()
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	url := fmt.Sprintf("https://%s/api/v2/users", p.Domain)
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", token)

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	// fmt.Println(string(body))
	users := []*Auth0User{}
	err = json.Unmarshal(body, &users)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	// convert auth user to data user
	ret := []*data.User{}
	for _, u := range users {
		ret = append(ret, &data.User{
			ID:           u.UserID,
			Type:         u.UserMetadata.UserRole,
			FullName:     u.Name,
			Email:        u.Email,
			Password:     u.Password,
			Avatar:       u.Picture,
			Bio:          u.UserMetadata.Bio,
			Organisation: u.UserMetadata.Organisation,
			ClientRole:   u.UserMetadata.ClientRole,
			LastAccess:   u.LastLogin,
		})
	}

	return ret, nil
}

func (p *auth0Storage) requestPasswordChange(email, token string) error {
	url := fmt.Sprintf("https://%s/dbconnections/change_password", p.Domain)
	method := "POST"

	payload, err := json.Marshal(map[string]string{
		"email":      email,
		"connection": "Username-Password-Authentication",
	})
	if err != nil {
		fmt.Println("jsonparse", err)
		return err
	}

	fmt.Println("payload", string(payload))

	client := &http.Client{}
	req, err := http.NewRequest(method, url, bytes.NewBuffer(payload))

	if err != nil {
		fmt.Println(err)
		return err
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", token)

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return err
	}

	fmt.Println("StatusCode", res.StatusCode)
	if res.StatusCode == 400 {
		response := &struct {
			StatusCode int    `json:"statusCode"`
			Error      string `json:"error"`
			Message    string `json:"message"`
			ErrorCode  string `json:"errorCode"`
		}{}

		_ = json.Unmarshal(body, response)

		fmt.Println("Error requesting password change", response.Message)
		return fmt.Errorf(response.Message)
	}

	fmt.Println("request password change body", string(body))
	return nil
}

func (p *auth0Storage) assignRole(user_id, role, token string) error {
	url := fmt.Sprintf("https://%s/api/v2/users/%s/roles", p.Domain, user_id)
	method := "POST"

	payload, err := json.Marshal(map[string][]string{
		"roles": []string{role},
	})
	if err != nil {
		fmt.Println("jsonparse", err)
		return err
	}

	fmt.Println("payload", string(payload))

	client := &http.Client{}
	req, err := http.NewRequest(method, url, bytes.NewBuffer(payload))

	if err != nil {
		fmt.Println(err)
		return err
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", token)

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return err
	}

	fmt.Println("StatusCode", res.StatusCode)
	if res.StatusCode == 400 {
		response := &struct {
			StatusCode int    `json:"statusCode"`
			Error      string `json:"error"`
			Message    string `json:"message"`
			ErrorCode  string `json:"errorCode"`
		}{}

		_ = json.Unmarshal(body, response)

		fmt.Println("Error setting user roles", response.Message)
		return fmt.Errorf(response.Message)
	}

	fmt.Println("assing role body", string(body))
	return nil
}

func (p *auth0Storage) Create(user *data.User) (*data.User, error) {
	fmt.Println("create user", user)
	userCreate := &Auth0User{
		Email:   user.Email,
		Name:    user.FullName,
		Picture: user.Avatar,
		UserMetadata: struct {
			UserRole     string `json:"user_role"`
			Bio          string `json:"bio"`
			Organisation string `json:"organisation"`
			ClientRole   string `json:"client_role"`
		}{
			UserRole:     user.Type,
			Bio:          user.Bio,
			Organisation: user.Organisation,
			ClientRole:   user.ClientRole,
		},
		Connection: "Username-Password-Authentication",
		Password:   "@Infinities2025..",
	}

	token, err := p.getToken()
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	url := fmt.Sprintf("https://%s/api/v2/users", p.Domain)
	method := "POST"

	payload, err := json.Marshal(userCreate)
	if err != nil {
		fmt.Println("jsonparse", err)
		return nil, err
	}

	fmt.Println("payload", string(payload))

	client := &http.Client{}
	req, err := http.NewRequest(method, url, bytes.NewBuffer(payload))

	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", token)

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	fmt.Println("StatusCode", res.StatusCode)
	if res.StatusCode != 201 {
		response := &struct {
			StatusCode int    `json:"statusCode"`
			Error      string `json:"error"`
			Message    string `json:"message"`
			ErrorCode  string `json:"errorCode"`
		}{}

		_ = json.Unmarshal(body, response)

		fmt.Println("Error creating user", response.Message)
		return nil, fmt.Errorf(response.Message)
	}

	fmt.Println("create user body", string(body))
	userRes := &Auth0User{}
	err = json.Unmarshal(body, userRes)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	role := ""
	if user.Type == "client" {
		role = p.ClientRoleId
	} else if user.Type == "member" {
		role = p.MemberRoleId
	}

	if role != "" {
		err = p.assignRole(userRes.UserID, role, token)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
	}

	err = p.requestPasswordChange(userRes.Email, token)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	ret := &data.User{
		ID:           userRes.UserID,
		Type:         userRes.UserMetadata.UserRole,
		FullName:     userRes.Name,
		Email:        userRes.Email,
		Password:     userRes.Password,
		Avatar:       userRes.Picture,
		Bio:          userRes.UserMetadata.Bio,
		Organisation: userRes.UserMetadata.Organisation,
		ClientRole:   userRes.UserMetadata.ClientRole,
		LastAccess:   userRes.LastLogin,
	}

	if ret.Type == "client" {
		p.PStorage.AddOrganisation(ret.Organisation)
	}

	return ret, nil
}

func (p *auth0Storage) Update(hex string, user *data.User) (*data.User, error) {
	fmt.Println("update user", user)
	userUpdate := &Auth0User{
		Email:   user.Email,
		Name:    user.FullName,
		Picture: user.Avatar,
		UserMetadata: struct {
			UserRole     string `json:"user_role"`
			Bio          string `json:"bio"`
			Organisation string `json:"organisation"`
			ClientRole   string `json:"client_role"`
		}{
			UserRole:     user.Type,
			Bio:          user.Bio,
			Organisation: user.Organisation,
			ClientRole:   user.ClientRole,
		},
	}

	token, err := p.getToken()
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	url := fmt.Sprintf("https://%s/api/v2/users/%s", p.Domain, hex)
	method := "PATCH"

	payload, err := json.Marshal(userUpdate)
	if err != nil {
		fmt.Println("jsonparse", err)
		return nil, err
	}

	fmt.Println("payload", string(payload))

	client := &http.Client{}
	req, err := http.NewRequest(method, url, bytes.NewBuffer(payload))

	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", token)

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	fmt.Println("StatusCode", res.StatusCode)
	if res.StatusCode == 400 {
		response := &struct {
			StatusCode int    `json:"statusCode"`
			Error      string `json:"error"`
			Message    string `json:"message"`
			ErrorCode  string `json:"errorCode"`
		}{}

		_ = json.Unmarshal(body, response)

		fmt.Println("Error updating user", response.Message)
		return nil, fmt.Errorf(response.Message)
	}

	fmt.Println("update user body", string(body))
	userRes := &Auth0User{}
	err = json.Unmarshal(body, userRes)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	ret := &data.User{
		ID:           userRes.UserID,
		Type:         userRes.UserMetadata.UserRole,
		FullName:     userRes.Name,
		Email:        userRes.Email,
		Password:     userRes.Password,
		Avatar:       userRes.Picture,
		Bio:          userRes.UserMetadata.Bio,
		Organisation: userRes.UserMetadata.Organisation,
		ClientRole:   userRes.UserMetadata.ClientRole,
		LastAccess:   userRes.LastLogin,
	}

	if ret.Type == "client" {
		p.PStorage.AddOrganisation(ret.Organisation)
	}

	return ret, nil
}

func (p *auth0Storage) Delete(hex string) error {
	fmt.Println("delete by id", hex)
	token, err := p.getToken()
	if err != nil {
		fmt.Println(err)
		return err
	}

	url := fmt.Sprintf("https://%s/api/v2/users/%s", p.Domain, hex)
	fmt.Println("url", url)
	method := "DELETE"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return err
	}
	req.Header.Add("Authorization", token)

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return err
	}

	fmt.Println("StatusCode", res.StatusCode)
	if res.StatusCode == 400 {
		response := &struct {
			StatusCode int    `json:"statusCode"`
			Error      string `json:"error"`
			Message    string `json:"message"`
			ErrorCode  string `json:"errorCode"`
		}{}

		_ = json.Unmarshal(body, response)

		fmt.Println("Error deleting user", response.Message)
		return fmt.Errorf(response.Message)
	}

	fmt.Println("body", res)

	return nil
}
