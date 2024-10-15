package v2

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"ppo/domain"
	"ppo/internal/app"
	"ppo/pkg/errs"
	"ppo/web"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

// LoginHandler godoc
//
//	@Summary		Аутентификация
//	@ID				login
//	@Produce json
//	@Description	Метод для получения bearer-токена для auth
//	@Param data body web.LoginReq true "Login data"
//	@Tags			users
//	@Success		200	{object} web.SuccessResponseStruct
//	@Failure		400	{object} web.ErrorResponseStruct
//	@Router			/login [post]
func LoginHandler(app *app.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		prompt := "LoginHandler"
		start := time.Now()

		wrappedWriter := &web.StatusResponseWriter{ResponseWriter: w, StatusCodeOuter: http.StatusOK}

		defer func() {
			web.ObserveRequest(time.Since(start), wrappedWriter.StatusCode(), r.Method, prompt)
		}()

		var req web.LoginReq

		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			app.Logger.Infof("%s: %v", prompt, err)
			web.ErrorResponse(wrappedWriter, fmt.Errorf("%s: %w", prompt, err).Error(), http.StatusBadRequest)
			return
		}

		ua := &domain.UserAuth{Username: req.Login, Password: req.Password}
		token, err := app.AuthSvc.Login(r.Context(), ua)
		if err != nil {
			app.Logger.Infof("%s: %v", prompt, err)
			web.ErrorResponse(wrappedWriter, fmt.Errorf("%s: %w", prompt, err).Error(), http.StatusUnauthorized)
			return
		}

		cookie := http.Cookie{
			Name:    "access_token",
			Value:   token,
			Path:    "/",
			Secure:  true,
			Expires: time.Now().Add(3600 * 24 * time.Second),
		}

		http.SetCookie(w, &cookie)
		web.SuccessResponse(wrappedWriter, http.StatusOK, map[string]string{"token": token})
	}
}

// RegisterHandler godoc
//
//	@Summary		Регистрация
//	@ID				signup
//	@Produce json
//	@Description	Метод для регистрации
//	@Param data body web.RegisterReq true "Signup data"
//	@Tags			users
//	@Success		200	{object} web.SuccessResponseStruct
//	@Failure		400	{object} web.ErrorResponseStruct
//	@Router			/signup [post]
func RegisterHandler(app *app.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		prompt := "RegisterHandler"
		start := time.Now()

		wrappedWriter := &web.StatusResponseWriter{ResponseWriter: w, StatusCodeOuter: http.StatusOK}

		defer func() {
			web.ObserveRequest(time.Since(start), wrappedWriter.StatusCode(), r.Method, prompt)
		}()

		var req web.RegisterReq

		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			app.Logger.Infof("%s: %v", prompt, err)
			web.ErrorResponse(wrappedWriter, fmt.Errorf("%s: %w", prompt, err).Error(), http.StatusBadRequest)
			return
		}

		ua := &domain.UserAuth{Username: req.Login, Password: req.Password}
		err = app.AuthSvc.Register(r.Context(), ua)
		if err != nil {
			app.Logger.Infof("%s: %v", prompt, err)
			web.ErrorResponse(wrappedWriter, fmt.Errorf("%s: %w", prompt, err).Error(), http.StatusBadRequest)
			return
		}

		web.SuccessResponse(wrappedWriter, http.StatusOK, nil)
	}
}

//	@Param name query string false "FIO entry"

// ListEntrepreneurs godoc
//
//	@Summary		Список предпринимателей
//	@ID				listEntrepreneurs
//	@Produce json
//	@Description	Метод для получения пагинированного списка предпринимателей
//	@Tags			entrepreneurs
//	@Param page query integer true "Page number"
//	@Success		200	{object} EntrepreneursResponse
//	@Failure		400	{object} web.ErrorResponseStruct
//	@Failure 		500 {object} web.ErrorResponseStruct
//	@Router			/entrepreneurs [get]
func ListEntrepreneurs(app *app.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		prompt := "ListEntrepreneursHandler"
		start := time.Now()

		wrappedWriter := &web.StatusResponseWriter{ResponseWriter: w, StatusCodeOuter: http.StatusOK}

		defer func() {
			web.ObserveRequest(time.Since(start), wrappedWriter.StatusCode(), r.Method, prompt)
		}()

		page := r.URL.Query().Get("page")
		if page == "" {
			app.Logger.Infof("%s: пустой номер страницы", prompt)
			web.ErrorResponse(wrappedWriter, fmt.Errorf("%s: пустой номер страницы", prompt).Error(), http.StatusBadRequest)
			return
		}

		pageInt, err := strconv.Atoi(page)
		if err != nil {
			app.Logger.Infof("%s: преобразование номера страницы к int: %v", prompt, err)
			web.ErrorResponse(wrappedWriter, fmt.Errorf("%s: преобразование номера страницы к int: %w", prompt, err).Error(), http.StatusBadRequest)
			return
		}

		users, numPages, err := app.UserSvc.GetAll(r.Context(), pageInt)
		if err != nil {
			app.Logger.Infof("%s: %v", prompt, err)
			web.ErrorResponse(wrappedWriter, fmt.Errorf("%s: %w", prompt, err).Error(), http.StatusInternalServerError)
			return
		}

		usersTransport := make([]web.User, len(users))
		for i, user := range users {
			usersTransport[i] = web.ToUserTransport(user)
		}

		web.SuccessResponse(wrappedWriter, http.StatusOK, EntrepreneursResponse{Pages: numPages, Entrepreneurs: usersTransport})
	}
}

// UpdateEntrepreneur godoc
//
//	@Summary		Обновление информации о предпринимателе
//	@ID				updateEntrepreneur
//	@Produce json
//	@Tags			entrepreneurs
//	@Param id path string true "Entrepreneur`s id"
//	@Param data body web.User true "Entrepreneur data"
//	@Success		204	{object} web.SuccessResponseStruct
//	@Failure		400	{object} web.ErrorResponseStruct
//	@Failure		401
//	@Failure		403
//	@Failure 		500 {object} web.ErrorResponseStruct
//	@Security		BearerAuth
//	@Router			/entrepreneurs/{id} [patch]
func UpdateEntrepreneur(app *app.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		prompt := "UpdateEntrepreneurHandler"
		start := time.Now()

		wrappedWriter := &web.StatusResponseWriter{ResponseWriter: w, StatusCodeOuter: http.StatusOK}

		defer func() {
			web.ObserveRequest(time.Since(start), wrappedWriter.StatusCode(), r.Method, prompt)
		}()

		id := chi.URLParam(r, "id")
		if id == "" {
			app.Logger.Infof("%s: пустой id", prompt)
			web.ErrorResponse(wrappedWriter, fmt.Errorf("%s: пустой id", prompt).Error(), http.StatusBadRequest)
			return
		}

		idUuid, err := uuid.Parse(id)
		if err != nil {
			app.Logger.Infof("%s: преобразование id к uuid: %v", prompt, err)
			web.ErrorResponse(wrappedWriter, fmt.Errorf("%s: преобразование id к uuid: %w", prompt, err).Error(), http.StatusBadRequest)
			return
		}

		var req web.User

		err = json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			app.Logger.Infof("%s: %v", prompt, err)
			web.ErrorResponse(wrappedWriter, err.Error(), http.StatusBadRequest)
			return
		}

		req.ID = idUuid
		userModel := web.ToUserModel(&req)

		err = app.UserSvc.Update(r.Context(), &userModel)
		if err != nil {
			app.Logger.Infof("%s: %v", prompt, err)
			web.ErrorResponse(wrappedWriter, fmt.Errorf("%s: %w", prompt, err).Error(), http.StatusBadRequest)
			return
		}

		web.SuccessResponse(wrappedWriter, http.StatusNoContent, nil)
	}
}

// DeleteEntrepreneur godoc
//
//	@Summary		Удаление информации о предпринимателе
//	@ID				deleteEntrepreneur
//	@Produce json
//	@Tags			entrepreneurs
//	@Param id path string true "Entrepreneur`s id"
//	@Success		204	{object} web.SuccessResponseStruct
//	@Failure		400	{object} web.ErrorResponseStruct
//	@Failure		401
//	@Failure		403
//	@Failure 		500 {object} web.ErrorResponseStruct
//	@Security		BearerAuth
//	@Router			/entrepreneurs/{id} [delete]
func DeleteEntrepreneur(app *app.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		prompt := "DeleteEntrepreneurHandler"
		start := time.Now()

		wrappedWriter := &web.StatusResponseWriter{ResponseWriter: w, StatusCodeOuter: http.StatusOK}

		defer func() {
			web.ObserveRequest(time.Since(start), wrappedWriter.StatusCode(), r.Method, prompt)
		}()

		id := chi.URLParam(r, "id")
		if id == "" {
			app.Logger.Infof("%s: пустой id", prompt)
			web.ErrorResponse(wrappedWriter, fmt.Errorf("%s: пустой id", prompt).Error(), http.StatusBadRequest)
			return
		}

		idUuid, err := uuid.Parse(id)
		if err != nil {
			app.Logger.Infof("%s: преобразование id к uuid: %v", prompt, err)
			web.ErrorResponse(wrappedWriter, fmt.Errorf("%s: преобразование id к uuid: %w", prompt, err).Error(), http.StatusBadRequest)
			return
		}

		err = app.UserSvc.DeleteById(r.Context(), idUuid)
		if err != nil {
			app.Logger.Infof("%s: %v", prompt, err)
			web.ErrorResponse(wrappedWriter, fmt.Errorf("%s: %w", prompt, err).Error(), http.StatusInternalServerError)
			return
		}

		web.SuccessResponse(wrappedWriter, http.StatusNoContent, nil)
	}
}

// GetEntrepreneur godoc
//
//	@Summary		Получение информации о предпринимателе
//	@ID				getEntrepreneur
//	@Produce json
//	@Tags			entrepreneurs
//	@Param id path string true "Entrepreneur`s id"
//	@Success		200	{object} EntrepreneurResponse
//	@Failure		400	{object} web.ErrorResponseStruct
//	@Failure		401
//	@Failure		403
//	@Failure		404
//	@Failure 		500 {object} web.ErrorResponseStruct
//	@Security		BearerAuth
//	@Router			/entrepreneurs/{id} [get]
func GetEntrepreneur(app *app.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		prompt := "GetEntrepreneurHandler"
		start := time.Now()

		wrappedWriter := &web.StatusResponseWriter{ResponseWriter: w, StatusCodeOuter: http.StatusOK}

		defer func() {
			web.ObserveRequest(time.Since(start), wrappedWriter.StatusCode(), r.Method, prompt)
		}()

		id := chi.URLParam(r, "id")
		if id == "" {
			app.Logger.Infof("%s: пустой id", prompt)
			web.ErrorResponse(wrappedWriter, fmt.Errorf("%s: пустой id", prompt).Error(), http.StatusBadRequest)
			return
		}

		idUuid, err := uuid.Parse(id)
		if err != nil {
			app.Logger.Infof("%s: преобразование id к uuid: %v", prompt, err)
			web.ErrorResponse(wrappedWriter, fmt.Errorf("%s: преобразование id к uuid: %w", prompt, err).Error(), http.StatusBadRequest)
			return
		}

		user, err := app.UserSvc.GetById(r.Context(), idUuid)
		if err != nil {
			if errors.Is(err, errs.ErrNotFound) {
				web.ErrorResponse(wrappedWriter, fmt.Errorf("%s: %w", prompt, err).Error(), http.StatusNotFound)
			} else {
				web.ErrorResponse(wrappedWriter, fmt.Errorf("%s: %w", prompt, err).Error(), http.StatusInternalServerError)
			}
			app.Logger.Infof("%s: %v", prompt, err)
			return
		}

		web.SuccessResponse(wrappedWriter, http.StatusOK, EntrepreneurResponse{Entrepreneur: web.ToUserTransport(user)})
	}
}

// CreateContact godoc
//
//	@Summary		Добавление средства связи
//	@ID				createContact
//	@Produce json
//	@Tags			contacts
//	@Param data body web.Contact true "Contact data"
//	@Success		204	{object} web.SuccessResponseStruct
//	@Failure		400	{object} web.ErrorResponseStruct
//	@Failure		401
//	@Failure 		500 {object} web.ErrorResponseStruct
//	@Security		BearerAuth
//	@Router			/contacts [post]
func CreateContact(app *app.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		prompt := "CreateContactHandler"
		start := time.Now()

		wrappedWriter := &web.StatusResponseWriter{ResponseWriter: w, StatusCodeOuter: http.StatusOK}

		defer func() {
			web.ObserveRequest(time.Since(start), wrappedWriter.StatusCode(), r.Method, prompt)
		}()

		idStr, err := web.GetStringClaimFromJWT(r.Context(), "sub")
		if err != nil {
			app.Logger.Infof("%s: получение записей из JWT: %v", prompt, err)
			web.ErrorResponse(wrappedWriter, fmt.Errorf("%s: получение записей из JWT: %w", prompt, err).Error(), http.StatusBadRequest)
			return
		}

		idUuid, err := uuid.Parse(idStr)
		if err != nil {
			app.Logger.Infof("%s: преобразование id к uuid: %v", prompt, err)
			web.ErrorResponse(wrappedWriter, fmt.Errorf("%s: преобразование id к uuid: %w", prompt, err).Error(), http.StatusInternalServerError)
			return
		}

		var req web.Contact
		err = json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			app.Logger.Infof("%s: %v", prompt, err)
			web.ErrorResponse(wrappedWriter, err.Error(), http.StatusBadRequest)
			return
		}

		contact := web.ToContactModel(&req)
		contact.OwnerID = idUuid

		err = app.ConSvc.Create(r.Context(), &contact)
		if err != nil {
			app.Logger.Infof("%s: %v", prompt, err)
			web.ErrorResponse(wrappedWriter, fmt.Errorf("%s: %w", prompt, err).Error(), http.StatusBadRequest)
			return
		}

		web.SuccessResponse(wrappedWriter, http.StatusNoContent, nil)
	}
}

// DeleteContact godoc
//
//	@Summary		Удаление средства связи
//	@ID				deleteContact
//	@Produce json
//	@Tags			contacts
//	@Param id path string true "Contact id"
//	@Success		204	{object} web.SuccessResponseStruct
//	@Failure		400	{object} web.ErrorResponseStruct
//	@Failure		401
//	@Failure 		500 {object} web.ErrorResponseStruct
//	@Security		BearerAuth
//	@Router			/contacts/{id} [delete]
func DeleteContact(app *app.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		prompt := "DeleteContactHandler"
		start := time.Now()

		wrappedWriter := &web.StatusResponseWriter{ResponseWriter: w, StatusCodeOuter: http.StatusOK}

		defer func() {
			web.ObserveRequest(time.Since(start), wrappedWriter.StatusCode(), r.Method, prompt)
		}()

		ownerIdStr, err := web.GetStringClaimFromJWT(r.Context(), "sub")
		if err != nil {
			app.Logger.Infof("%s: получение записей из JWT: %v", prompt, err)
			web.ErrorResponse(wrappedWriter, fmt.Errorf("получение записей из JWT: %w", err).Error(), http.StatusBadRequest)
			return
		}

		ownerIdUuid, err := uuid.Parse(ownerIdStr)
		if err != nil {
			app.Logger.Infof("%s: преобразование строки к uuid: %v", prompt, err)
			web.ErrorResponse(wrappedWriter, fmt.Errorf("преобразование строки к uuid: %w", err).Error(), http.StatusInternalServerError)
			return
		}

		id := chi.URLParam(r, "id")
		if id == "" {
			app.Logger.Infof("%s: пустой id", prompt)
			web.ErrorResponse(wrappedWriter, fmt.Errorf("пустой id").Error(), http.StatusBadRequest)
			return
		}

		idUuid, err := uuid.Parse(id)
		if err != nil {
			app.Logger.Infof("%s: преобразование строки к uuid: %v", prompt, err)
			web.ErrorResponse(wrappedWriter, fmt.Errorf("преобразование строки к uuid: %w", err).Error(), http.StatusBadRequest)
			return
		}

		err = app.ConSvc.DeleteById(r.Context(), idUuid, ownerIdUuid)
		if err != nil {
			app.Logger.Infof("%s: удаление средства связи по id: %v", prompt, err)
			web.ErrorResponse(wrappedWriter, fmt.Errorf("удаление средства связи по id: %w", err).Error(), http.StatusInternalServerError)
			return
		}

		web.SuccessResponse(wrappedWriter, http.StatusNoContent, nil)
	}
}

// UpdateContact godoc
//
//	@Summary		Обновление средства связи
//	@ID				updateContact
//	@Produce json
//	@Tags			contacts
//	@Param id path string true "Contact id"
//	@Param data body web.Contact true "Contact data"
//	@Success		204	{object} web.SuccessResponseStruct
//	@Failure		400	{object} web.ErrorResponseStruct
//	@Failure		401
//	@Failure 		500 {object} web.ErrorResponseStruct
//	@Security		BearerAuth
//	@Router			/contacts/{id} [patch]
func UpdateContact(app *app.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		prompt := "UpdateContactHandler"
		start := time.Now()

		wrappedWriter := &web.StatusResponseWriter{ResponseWriter: w, StatusCodeOuter: http.StatusOK}

		defer func() {
			web.ObserveRequest(time.Since(start), wrappedWriter.StatusCode(), r.Method, prompt)
		}()

		ownerIdStr, err := web.GetStringClaimFromJWT(r.Context(), "sub")
		if err != nil {
			app.Logger.Infof("%s: получение записей из JWT: %v", prompt, err)
			web.ErrorResponse(wrappedWriter, fmt.Errorf("получение записей из JWT: %w", err).Error(), http.StatusBadRequest)
			return
		}

		ownerIdUuid, err := uuid.Parse(ownerIdStr)
		if err != nil {
			app.Logger.Infof("%s: преобразование строки к uuid: %v", prompt, err)
			web.ErrorResponse(wrappedWriter, fmt.Errorf("преобразование строки к uuid: %w", err).Error(), http.StatusInternalServerError)
			return
		}

		id := chi.URLParam(r, "id")
		if id == "" {
			app.Logger.Infof("%s: пустой id", prompt)
			web.ErrorResponse(wrappedWriter, fmt.Errorf("пустой id").Error(), http.StatusBadRequest)
			return
		}

		idUuid, err := uuid.Parse(id)
		if err != nil {
			app.Logger.Infof("%s: преобразование id к uuid: %v", prompt, err)
			web.ErrorResponse(wrappedWriter, fmt.Errorf("преобразование id к uuid: %w", err).Error(), http.StatusBadRequest)
			return
		}

		var req web.Contact

		err = json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			app.Logger.Infof("%s: %v", err)
			web.ErrorResponse(wrappedWriter, err.Error(), http.StatusBadRequest)
			return
		}
		req.ID = idUuid
		model := web.ToContactModel(&req)

		err = app.ConSvc.Update(r.Context(), &model, ownerIdUuid)
		if err != nil {
			app.Logger.Infof("%s: обновление информации о средстве связи: %v", prompt, err)
			web.ErrorResponse(wrappedWriter, fmt.Errorf("обновление информации о средстве связи: %w", err).Error(), http.StatusInternalServerError)
			return
		}

		web.SuccessResponse(wrappedWriter, http.StatusNoContent, nil)
	}
}

// GetContact godoc
//
//	@Summary		Получение информации о средстве связи
//	@ID				getContact
//	@Produce json
//	@Tags			contacts
//	@Param id path string true "Contact`s id"
//	@Success		200	{object} ContactResponse
//	@Failure		400	{object} web.ErrorResponseStruct
//	@Failure		401
//	@Failure		404
//	@Failure 		500 {object} web.ErrorResponseStruct
//	@Security		BearerAuth
//	@Router			/contacts/{id} [get]
func GetContact(app *app.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		prompt := "GetContactHandler"
		start := time.Now()

		wrappedWriter := &web.StatusResponseWriter{ResponseWriter: w, StatusCodeOuter: http.StatusOK}

		defer func() {
			web.ObserveRequest(time.Since(start), wrappedWriter.StatusCode(), r.Method, prompt)
		}()

		id := chi.URLParam(r, "id")
		if id == "" {
			app.Logger.Infof("%s: пустой id", prompt)
			web.ErrorResponse(wrappedWriter, fmt.Errorf("пустой id").Error(), http.StatusBadRequest)
			return
		}

		idUuid, err := uuid.Parse(id)
		if err != nil {
			app.Logger.Infof("%s: преобразование id к uuid: %v", prompt, err)
			web.ErrorResponse(wrappedWriter, fmt.Errorf("преобразование id к uuid: %w", err).Error(), http.StatusBadRequest)
			return
		}

		contact, err := app.ConSvc.GetById(r.Context(), idUuid)
		if err != nil {
			if errors.Is(err, errs.ErrNotFound) {
				web.ErrorResponse(wrappedWriter, fmt.Errorf("получение средства связи по id: %w", err).Error(), http.StatusNotFound)
			} else {
				web.ErrorResponse(wrappedWriter, fmt.Errorf("получение средства связи по id: %w", err).Error(), http.StatusInternalServerError)
			}
			app.Logger.Infof("%s: получение средства связи по id: %v", prompt, err)
			return
		}

		web.SuccessResponse(wrappedWriter, http.StatusOK, ContactResponse{Contact: web.ToContactTransport(contact)})
	}
}

// ListEntrepreneurContacts godoc
//
//	@Summary		Список средств связи предпринимателя
//	@ID				listEntrepreneurContacts
//	@Produce json
//	@Description	Метод для получения списка средств связи с предпринимателем
//	@Tags			entrepreneurs
//	@Param id path string true "Entrepreneur ID"
//	@Success		200	{object} ContactsResponse
//	@Failure		400	{object} web.ErrorResponseStruct
//	@Failure		404
//	@Failure 		500 {object} web.ErrorResponseStruct
//	@Router			/entrepreneurs/{id}/contacts [get]
func ListEntrepreneurContacts(app *app.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		prompt := "ListEntrepreneursContactsHandler"
		start := time.Now()

		wrappedWriter := &web.StatusResponseWriter{ResponseWriter: w, StatusCodeOuter: http.StatusOK}

		defer func() {
			web.ObserveRequest(time.Since(start), wrappedWriter.StatusCode(), r.Method, prompt)
		}()

		entId := chi.URLParam(r, "id")
		if entId == "" {
			app.Logger.Infof("%s: пустой id предпринимателя", prompt)
			web.ErrorResponse(wrappedWriter, fmt.Errorf("пустой id предпринимателя").Error(), http.StatusBadRequest)
			return
		}

		entUuid, err := uuid.Parse(entId)
		if err != nil {
			app.Logger.Infof("%s: преобразование id к uuid: %v", prompt, err)
			web.ErrorResponse(wrappedWriter, fmt.Errorf("преобразование id к uuid: %w", err).Error(), http.StatusInternalServerError)
			return
		}

		contacts, err := app.ConSvc.GetByOwnerId(r.Context(), entUuid)
		if err != nil {
			if errors.Is(err, errs.ErrNotFound) {
				web.ErrorResponse(wrappedWriter, fmt.Errorf("получение списка контактов: %w", err).Error(), http.StatusNotFound)
			} else {
				web.ErrorResponse(wrappedWriter, fmt.Errorf("получение списка контактов: %w", err).Error(), http.StatusInternalServerError)
			}
			app.Logger.Infof("%s: получение списка контактов: %v", prompt, err)
			return
		}

		contactsTransport := make([]web.Contact, len(contacts))
		for i, contact := range contacts {
			contactsTransport[i] = web.ToContactTransport(contact)
		}

		web.SuccessResponse(wrappedWriter, http.StatusOK, ContactsResponse{EntrepreneurId: entUuid, Contacts: contactsTransport})
	}
}

// CreateActivityField godoc
//
//	@Summary		Добавление сферы деятельности
//	@ID				createActField
//	@Produce json
//	@Tags			activityFields
//	@Param data body web.ActivityField true "Activity Field data"
//	@Success		204	{object} web.SuccessResponseStruct
//	@Failure		400	{object} web.ErrorResponseStruct
//	@Failure		401
//	@Failure		403
//	@Failure 		500 {object} web.ErrorResponseStruct
//	@Security		BearerAuth
//	@Router			/activity_fields [post]
func CreateActivityField(app *app.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		prompt := "CreateActivityFieldHandler"
		start := time.Now()

		wrappedWriter := &web.StatusResponseWriter{ResponseWriter: w, StatusCodeOuter: http.StatusOK}

		defer func() {
			web.ObserveRequest(time.Since(start), wrappedWriter.StatusCode(), r.Method, prompt)
		}()

		var req web.ActivityField
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			app.Logger.Infof("%s: %v", prompt, err)
			web.ErrorResponse(wrappedWriter, err.Error(), http.StatusBadRequest)
			return
		}

		actField := web.ToActFieldModel(&req)

		err = app.ActFieldSvc.Create(r.Context(), &actField)
		if err != nil {
			app.Logger.Infof("%s: создание сферы деятельности: %v", prompt, err)
			web.ErrorResponse(wrappedWriter, fmt.Errorf("создание сферы деятельности: %w", err).Error(), http.StatusBadRequest)
			return
		}

		web.SuccessResponse(wrappedWriter, http.StatusNoContent, nil)
	}
}

// DeleteActivityField godoc
//
//	@Summary		Удаление информации о сфере деятельности
//	@ID				deleteActField
//	@Produce json
//	@Tags			activityFields
//	@Param id path string true "Activity Field`s id"
//	@Success		204	{object} web.SuccessResponseStruct
//	@Failure		400	{object} web.ErrorResponseStruct
//	@Failure		401
//	@Failure		403
//	@Failure 		500 {object} web.ErrorResponseStruct
//	@Security		BearerAuth
//	@Router			/activity_fields/{id} [delete]
func DeleteActivityField(app *app.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		prompt := "DeleteActivityFieldHandler"
		start := time.Now()

		wrappedWriter := &web.StatusResponseWriter{ResponseWriter: w, StatusCodeOuter: http.StatusOK}

		defer func() {
			web.ObserveRequest(time.Since(start), wrappedWriter.StatusCode(), r.Method, prompt)
		}()

		id := chi.URLParam(r, "id")
		if id == "" {
			app.Logger.Infof("%s: пустой id", prompt)
			web.ErrorResponse(wrappedWriter, fmt.Errorf("пустой id").Error(), http.StatusBadRequest)
			return
		}

		idUuid, err := uuid.Parse(id)
		if err != nil {
			app.Logger.Infof("%s: преобразование id к uuid: %v", prompt, err)
			web.ErrorResponse(wrappedWriter, fmt.Errorf("преобразование id к uuid: %w", err).Error(), http.StatusBadRequest)
			return
		}

		err = app.ActFieldSvc.DeleteById(r.Context(), idUuid)
		if err != nil {
			app.Logger.Infof("%s: удаление сферы деятельности по id: %v", prompt, err)
			web.ErrorResponse(wrappedWriter, fmt.Errorf("удаление сферы деятельности по id: %w", err).Error(), http.StatusInternalServerError)
			return
		}

		web.SuccessResponse(wrappedWriter, http.StatusNoContent, nil)
	}
}

// UpdateActivityField godoc
//
//	@Summary		Обновление информации о сфере деятельности
//	@ID				updateActivityField
//	@Produce json
//	@Tags			activityFields
//	@Param id path string true "Activity Field id"
//	@Param data body web.ActivityField true "Activity Field data"
//	@Success		204	{object} web.SuccessResponseStruct
//	@Failure		400	{object} web.ErrorResponseStruct
//	@Failure		401
//	@Failure		403
//	@Failure 		500 {object} web.ErrorResponseStruct
//	@Security		BearerAuth
//	@Router			/activity_fields/{id} [patch]
func UpdateActivityField(app *app.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		prompt := "UpdateActivityFieldHandler"
		start := time.Now()

		wrappedWriter := &web.StatusResponseWriter{ResponseWriter: w, StatusCodeOuter: http.StatusOK}

		defer func() {
			web.ObserveRequest(time.Since(start), wrappedWriter.StatusCode(), r.Method, prompt)
		}()

		id := chi.URLParam(r, "id")
		if id == "" {
			app.Logger.Infof("%s: пустой id", prompt)
			web.ErrorResponse(wrappedWriter, fmt.Errorf("пустой id").Error(), http.StatusBadRequest)
			return
		}

		idUuid, err := uuid.Parse(id)
		if err != nil {
			app.Logger.Infof("%s: преобразование id к uuid: %v", prompt, err)
			web.ErrorResponse(wrappedWriter, fmt.Errorf("преобразование id к uuid: %w", err).Error(), http.StatusBadRequest)
			return
		}

		var req web.ActivityField

		err = json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			app.Logger.Infof("%s: %v", prompt, err)
			web.ErrorResponse(wrappedWriter, err.Error(), http.StatusBadRequest)
			return
		}

		req.ID = idUuid

		model := web.ToActFieldModel(&req)

		err = app.ActFieldSvc.Update(r.Context(), &model)
		if err != nil {
			app.Logger.Infof("%s: обновление информации о сфере деятельности: %v", prompt, err)
			web.ErrorResponse(wrappedWriter, fmt.Errorf("обновление информации о сфере деятельности: %w", err).Error(), http.StatusInternalServerError)
			return
		}

		web.SuccessResponse(wrappedWriter, http.StatusNoContent, nil)
	}
}

// GetActivityField godoc
//
//	@Summary		Получение информации о сфере деятельности
//	@ID				getActField
//	@Produce json
//	@Tags			activityFields
//	@Param id path string true "Activity field`s id"
//	@Success		200	{object} ActFieldResponse
//	@Failure		400	{object} web.ErrorResponseStruct
//	@Failure		401
//	@Failure		404
//	@Failure 		500 {object} web.ErrorResponseStruct
//	@Security		BearerAuth
//	@Router			/activity_fields/{id} [get]
func GetActivityField(app *app.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		prompt := "GetActivityFieldHandler"
		start := time.Now()

		wrappedWriter := &web.StatusResponseWriter{ResponseWriter: w, StatusCodeOuter: http.StatusOK}

		defer func() {
			web.ObserveRequest(time.Since(start), wrappedWriter.StatusCode(), r.Method, prompt)
		}()

		id := chi.URLParam(r, "id")
		if id == "" {
			app.Logger.Infof("%s: пустой id", prompt)
			web.ErrorResponse(wrappedWriter, fmt.Errorf("пустой id").Error(), http.StatusBadRequest)
			return
		}

		idUuid, err := uuid.Parse(id)
		if err != nil {
			app.Logger.Infof("%s: преобразование id к uuid: %v", prompt, err)
			web.ErrorResponse(wrappedWriter, fmt.Errorf("преобразование id к uuid: %w", err).Error(), http.StatusBadRequest)
			return
		}

		actField, err := app.ActFieldSvc.GetById(r.Context(), idUuid)
		if err != nil {
			if errors.Is(err, errs.ErrNotFound) {
				web.ErrorResponse(wrappedWriter, fmt.Errorf("получение сферы деятельности по id: %w", err).Error(), http.StatusNotFound)
			} else {
				web.ErrorResponse(wrappedWriter, fmt.Errorf("получение сферы деятельности по id: %w", err).Error(), http.StatusInternalServerError)
			}
			app.Logger.Infof("%s: получение сферы деятельности по id: %v", prompt, err)
			return
		}

		web.SuccessResponse(wrappedWriter, http.StatusOK, ActFieldResponse{ActField: web.ToActFieldTransport(actField)})
	}
}

// ListActivityFields godoc
//
//	@Summary		Получение списка сфер деятельности
//	@ID				listActivityFields
//	@Produce json
//	@Tags			activityFields
//	@Param page query int true "Page number"
//	@Success		200	{object} ActFieldsResponse
//	@Failure		400	{object} web.ErrorResponseStruct
//	@Failure		401
//	@Failure 		500 {object} web.ErrorResponseStruct
//	@Security		BearerAuth
//	@Router			/activity_fields [get]
func ListActivityFields(app *app.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		prompt := "ListActivityFieldsHandler"
		start := time.Now()

		wrappedWriter := &web.StatusResponseWriter{ResponseWriter: w, StatusCodeOuter: http.StatusOK}

		defer func() {
			web.ObserveRequest(time.Since(start), wrappedWriter.StatusCode(), r.Method, prompt)
		}()

		var paginated bool
		var pageInt int
		var err error

		page := r.URL.Query().Get("page")
		if page != "" {
			paginated = true

			pageInt, err = strconv.Atoi(page)
			if err != nil {
				app.Logger.Infof("%s: преобразование страницы к int: %v", prompt, err)
				web.ErrorResponse(wrappedWriter, fmt.Errorf("преобразование страницы к int: %w", err).Error(), http.StatusBadRequest)
				return
			}
		}

		actFields, numPages, err := app.ActFieldSvc.GetAll(r.Context(), pageInt, paginated)
		if err != nil {
			app.Logger.Infof("%s: получение списка сфер деятельности: %v", prompt, err)
			web.ErrorResponse(wrappedWriter, fmt.Errorf("получение списка сфер деятельности: %w", err).Error(), http.StatusInternalServerError)
			return
		}

		actFieldsTransport := make([]web.ActivityField, len(actFields))
		for i, actField := range actFields {
			actFieldsTransport[i] = web.ToActFieldTransport(actField)
		}

		web.SuccessResponse(wrappedWriter, http.StatusOK, ActFieldsResponse{Pages: numPages, ActFields: actFieldsTransport})
	}
}

// CreateCompany godoc
//
//	@Summary		Добавление компании
//	@ID				createCompany
//	@Produce json
//	@Tags			companies
//	@Param data body web.Company true "Company data"
//	@Success		204	{object} web.SuccessResponseStruct
//	@Failure		400	{object} web.ErrorResponseStruct
//	@Failure		401
//	@Failure 		500 {object} web.ErrorResponseStruct
//	@Security		BearerAuth
//	@Router			/companies [post]
func CreateCompany(app *app.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		prompt := "CreateCompanyHandler"
		start := time.Now()

		wrappedWriter := &web.StatusResponseWriter{ResponseWriter: w, StatusCodeOuter: http.StatusOK}

		defer func() {
			web.ObserveRequest(time.Since(start), wrappedWriter.StatusCode(), r.Method, prompt)
		}()

		idStr, err := web.GetStringClaimFromJWT(r.Context(), "sub")
		if err != nil {
			app.Logger.Infof("%s: получение записей из JWT: %v", prompt, err)
			web.ErrorResponse(wrappedWriter, fmt.Errorf("получение записей из JWT: %w", err).Error(), http.StatusBadRequest)
			return
		}

		idUuid, err := uuid.Parse(idStr)
		if err != nil {
			app.Logger.Infof("%s: преобразование строки к uuid: %v", prompt, err)
			web.ErrorResponse(wrappedWriter, fmt.Errorf("преобразование строки к uuid: %w", err).Error(), http.StatusInternalServerError)
			return
		}

		var req web.Company
		err = json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			app.Logger.Infof("%s: %v", prompt, err)
			web.ErrorResponse(wrappedWriter, err.Error(), http.StatusBadRequest)
			return
		}

		company := web.ToCompanyModel(&req)
		company.OwnerID = idUuid

		err = app.CompSvc.Create(r.Context(), &company)
		if err != nil {
			app.Logger.Infof("%s: создание компании: %v", prompt, err)
			web.ErrorResponse(wrappedWriter, fmt.Errorf("создание компании: %w", err).Error(), http.StatusBadRequest)
			return
		}

		web.SuccessResponse(wrappedWriter, http.StatusNoContent, nil)
	}
}

// DeleteCompany godoc
//
//	@Summary		Удаление информации о компании
//	@ID				deleteCompany
//	@Produce json
//	@Tags			companies
//	@Param id path string true "Company`s id"
//	@Success		204	{object} web.SuccessResponseStruct
//	@Failure		400	{object} web.ErrorResponseStruct
//	@Failure		401
//	@Failure 		500 {object} web.ErrorResponseStruct
//	@Security		BearerAuth
//	@Router			/companies/{id} [delete]
func DeleteCompany(app *app.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		prompt := "DeleteCompanyHandler"
		start := time.Now()

		wrappedWriter := &web.StatusResponseWriter{ResponseWriter: w, StatusCodeOuter: http.StatusOK}

		defer func() {
			web.ObserveRequest(time.Since(start), wrappedWriter.StatusCode(), r.Method, prompt)
		}()

		ownerIdStr, err := web.GetStringClaimFromJWT(r.Context(), "sub")
		if err != nil {
			app.Logger.Infof("%s: получение записей из JWT: %v", prompt, err)
			web.ErrorResponse(wrappedWriter, fmt.Errorf("получение записей из JWT: %w", err).Error(), http.StatusBadRequest)
			return
		}

		ownerIdUuid, err := uuid.Parse(ownerIdStr)
		if err != nil {
			app.Logger.Infof("%s: преобразование строки к uuid: %v", prompt, err)
			web.ErrorResponse(wrappedWriter, fmt.Errorf("преобразование строки к uuid: %w", err).Error(), http.StatusInternalServerError)
			return
		}

		id := chi.URLParam(r, "id")
		if id == "" {
			app.Logger.Infof("%s: пустой id", prompt)
			web.ErrorResponse(wrappedWriter, fmt.Errorf("пустой id").Error(), http.StatusBadRequest)
			return
		}

		idUuid, err := uuid.Parse(id)
		if err != nil {
			app.Logger.Infof("%s: преобразование id к uuid: %v", prompt, err)
			web.ErrorResponse(wrappedWriter, fmt.Errorf("преобразование id к uuid: %w", err).Error(), http.StatusBadRequest)
			return
		}

		err = app.CompSvc.DeleteById(r.Context(), idUuid, ownerIdUuid)
		if err != nil {
			app.Logger.Infof("%s: удаление компании по id: %v", prompt, err)
			web.ErrorResponse(wrappedWriter, fmt.Errorf("удаление компании по id: %w", err).Error(), http.StatusInternalServerError)
			return
		}

		web.SuccessResponse(wrappedWriter, http.StatusNoContent, nil)
	}
}

// UpdateCompany godoc
//
//	@Summary		Обновление информации о компании
//	@ID				updateCompany
//	@Produce json
//	@Tags			companies
//	@Param id path string true "Company id"
//	@Param data body web.Company true "Company data"
//	@Success		204	{object} web.SuccessResponseStruct
//	@Failure		400	{object} web.ErrorResponseStruct
//	@Failure		401
//	@Failure 		500 {object} web.ErrorResponseStruct
//	@Security		BearerAuth
//	@Router			/companies/{id} [patch]
func UpdateCompany(app *app.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		prompt := "UpdateCompanyHandler"
		start := time.Now()

		wrappedWriter := &web.StatusResponseWriter{ResponseWriter: w, StatusCodeOuter: http.StatusOK}

		defer func() {
			web.ObserveRequest(time.Since(start), wrappedWriter.StatusCode(), r.Method, prompt)
		}()

		ownerIdStr, err := web.GetStringClaimFromJWT(r.Context(), "sub")
		if err != nil {
			app.Logger.Infof("%s: получение записей из JWT: %v", prompt, err)
			web.ErrorResponse(wrappedWriter, fmt.Errorf("получение записей из JWT: %w", err).Error(), http.StatusBadRequest)
			return
		}

		ownerIdUuid, err := uuid.Parse(ownerIdStr)
		if err != nil {
			app.Logger.Infof("%s: преобразование строки к uuid: %v", prompt, err)
			web.ErrorResponse(wrappedWriter, fmt.Errorf("преобразование строки к uuid: %w", err).Error(), http.StatusInternalServerError)
			return
		}

		id := chi.URLParam(r, "id")
		if id == "" {
			app.Logger.Infof("%s: пустой id", prompt)
			web.ErrorResponse(wrappedWriter, fmt.Errorf("пустой id").Error(), http.StatusBadRequest)
			return
		}

		idUuid, err := uuid.Parse(id)
		if err != nil {
			app.Logger.Infof("%s: преобразование id к uuid: %v", prompt, err)
			web.ErrorResponse(wrappedWriter, fmt.Errorf("преобразование id к uuid: %w", err).Error(), http.StatusBadRequest)
			return
		}

		var req web.Company

		err = json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			app.Logger.Infof("%s: %v", prompt, err)
			web.ErrorResponse(wrappedWriter, err.Error(), http.StatusBadRequest)
			return
		}
		req.ID = idUuid
		model := web.ToCompanyModel(&req)

		err = app.CompSvc.Update(r.Context(), &model, ownerIdUuid)
		if err != nil {
			app.Logger.Infof("%s: обновление информации о компании: %v", prompt, err)
			web.ErrorResponse(wrappedWriter, fmt.Errorf("обновление информации о компании: %w", err).Error(), http.StatusInternalServerError)
			return
		}

		web.SuccessResponse(wrappedWriter, http.StatusNoContent, nil)
	}
}

// GetCompany godoc
//
//	@Summary		Получение информации о компании
//	@ID				getCompany
//	@Produce json
//	@Tags			companies
//	@Param id path string true "Company`s id"
//	@Success		200	{object} CompanyResponse
//	@Failure		400	{object} web.ErrorResponseStruct
//	@Failure		401
//	@Failure		404
//	@Failure 		500 {object} web.ErrorResponseStruct
//	@Security		BearerAuth
//	@Router			/companies/{id} [get]
func GetCompany(app *app.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		prompt := "GetCompanyHandler"
		start := time.Now()

		wrappedWriter := &web.StatusResponseWriter{ResponseWriter: w, StatusCodeOuter: http.StatusOK}

		defer func() {
			web.ObserveRequest(time.Since(start), wrappedWriter.StatusCode(), r.Method, prompt)
		}()

		id := chi.URLParam(r, "id")
		if id == "" {
			app.Logger.Infof("%s: пустой id", prompt)
			web.ErrorResponse(wrappedWriter, fmt.Errorf("пустой id").Error(), http.StatusBadRequest)
			return
		}

		idUuid, err := uuid.Parse(id)
		if err != nil {
			app.Logger.Infof("%s: преобразование id к uuid: %v", prompt, err)
			web.ErrorResponse(wrappedWriter, fmt.Errorf("преобразование id к uuid: %w", err).Error(), http.StatusBadRequest)
			return
		}

		company, err := app.CompSvc.GetById(r.Context(), idUuid)
		if err != nil {
			if errors.Is(err, errs.ErrNotFound) {
				web.ErrorResponse(wrappedWriter, fmt.Errorf("получение компании по id: %w", err).Error(), http.StatusNotFound)
			} else {
				web.ErrorResponse(wrappedWriter, fmt.Errorf("получение компании по id: %w", err).Error(), http.StatusInternalServerError)
			}
			app.Logger.Infof("%s: получение компании по id: %v", prompt, err)
			return
		}

		web.SuccessResponse(wrappedWriter, http.StatusOK, CompanyResponse{Company: web.ToCompanyTransport(company)})
	}
}

// ListEntrepreneurCompanies godoc
//
//	@Summary		Получение информации о компаниях предпринимателя
//	@ID				getEntrepreneurCompanies
//	@Produce json
//	@Tags			companies
//	@Param page query int true "Page number"
//	@Param id path string true "Entrepreneur ID"
//	@Success		200	{object} CompaniesResponse
//	@Failure		400	{object} web.ErrorResponseStruct
//	@Failure		401
//	@Failure 		500 {object} web.ErrorResponseStruct
//	@Security		BearerAuth
//	@Router			/entrepreneurs/{id}/companies [get]
func ListEntrepreneurCompanies(app *app.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		prompt := "ListEntrepreneurCompaniesHandler"
		start := time.Now()

		wrappedWriter := &web.StatusResponseWriter{ResponseWriter: w, StatusCodeOuter: http.StatusOK}

		defer func() {
			web.ObserveRequest(time.Since(start), wrappedWriter.StatusCode(), r.Method, prompt)
		}()

		page := r.URL.Query().Get("page")
		if page == "" {
			app.Logger.Infof("%s: пустой номер страницы", prompt)
			web.ErrorResponse(wrappedWriter, fmt.Errorf("пустой номер страницы").Error(), http.StatusBadRequest)
			return
		}

		pageInt, err := strconv.Atoi(page)
		if err != nil {
			app.Logger.Infof("%s: преобразование к int: %v", prompt, err)
			web.ErrorResponse(wrappedWriter, fmt.Errorf("преобразование к int: %w", err).Error(), http.StatusBadRequest)
			return
		}

		//entId := r.URL.Query().Get("id")
		entId := chi.URLParam(r, "id")
		if page == "" {
			app.Logger.Infof("%s: пустой id предпринимателя", prompt)
			web.ErrorResponse(wrappedWriter, fmt.Errorf("пустой id предпринимателя").Error(), http.StatusBadRequest)
			return
		}

		entUuid, err := uuid.Parse(entId)
		if err != nil {
			app.Logger.Infof("%s: преобразование id предпринимателя к uuid: %v", prompt, err)
			web.ErrorResponse(wrappedWriter, fmt.Errorf("преобразование id предпринимателя к uuid: %w", err).Error(), http.StatusInternalServerError)
			return
		}

		companies, numPages, err := app.CompSvc.GetByOwnerId(r.Context(), entUuid, pageInt, true)
		if err != nil {
			app.Logger.Infof("%s: получение списка компаний: %v", prompt, err)
			web.ErrorResponse(wrappedWriter, fmt.Errorf("получение списка компаний: %w", err).Error(), http.StatusInternalServerError)
			return
		}

		companiesTransport := make([]web.Company, len(companies))
		for i, company := range companies {
			companiesTransport[i] = web.ToCompanyTransport(company)
		}

		web.SuccessResponse(wrappedWriter, http.StatusOK, CompaniesResponse{Pages: numPages, EntrepreneurId: entUuid, Companies: companiesTransport})
	}
}

// CreateReport godoc
//
//	@Summary		Добавление финансового отчета
//	@ID				createFinReport
//	@Produce json
//	@Tags			financialsReports
//	@Param data body web.FinancialReport true "Financial Report data"
//	@Param id path string true "Company ID"
//	@Success		204	{object} web.SuccessResponseStruct
//	@Failure		400	{object} web.ErrorResponseStruct
//	@Failure		401
//	@Failure 		500 {object} web.ErrorResponseStruct
//	@Security		BearerAuth
//	@Router			/companies/{id}/financials [post]
func CreateReport(app *app.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		prompt := "CreateReportHandler"
		start := time.Now()

		wrappedWriter := &web.StatusResponseWriter{ResponseWriter: w, StatusCodeOuter: http.StatusOK}

		defer func() {
			web.ObserveRequest(time.Since(start), wrappedWriter.StatusCode(), r.Method, prompt)
		}()

		userIdStr, err := web.GetStringClaimFromJWT(r.Context(), "sub")
		if err != nil {
			app.Logger.Infof("%s: получение списка компаний: %v", prompt, err)
			web.ErrorResponse(wrappedWriter, fmt.Errorf("получение записей из JWT: %w", err).Error(), http.StatusBadRequest)
			return
		}

		userIdUuid, err := uuid.Parse(userIdStr)
		if err != nil {
			app.Logger.Infof("%s: преобразование строки к uuid: %v", prompt, err)
			web.ErrorResponse(wrappedWriter, fmt.Errorf("преобразование строки к uuid: %w", err).Error(), http.StatusInternalServerError)
			return
		}

		compIdStr := chi.URLParam(r, "id")
		if compIdStr == "" {
			app.Logger.Infof("%s: пустой id компании", prompt)
			web.ErrorResponse(wrappedWriter, fmt.Errorf("пустой id компании").Error(), http.StatusBadRequest)
			return
		}

		compIdUuid, err := uuid.Parse(compIdStr)
		if err != nil {
			app.Logger.Infof("%s: преобразование строки к uuid: %v", prompt, err)
			web.ErrorResponse(wrappedWriter, fmt.Errorf("преобразование строки к uuid: %w", err).Error(), http.StatusInternalServerError)
			return
		}

		company, err := app.CompSvc.GetById(r.Context(), compIdUuid)
		if err != nil {
			app.Logger.Infof("%s: создание финансового отчета: %v", prompt, err)
			web.ErrorResponse(wrappedWriter, fmt.Errorf("создание финансового отчета: %w", err).Error(), http.StatusInternalServerError)
			return
		}

		if company.OwnerID != userIdUuid {
			app.Logger.Infof("%s: только владелец компании может добавлять финансовые отчеты", prompt)
			web.ErrorResponse(wrappedWriter, fmt.Errorf("только владелец компании может добавлять финансовые отчеты").Error(), http.StatusInternalServerError)
			return
		}

		var req web.FinancialReport
		err = json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			app.Logger.Infof("%s: %м", prompt, err)
			web.ErrorResponse(wrappedWriter, err.Error(), http.StatusBadRequest)
			return
		}

		report := web.ToFinReportModel(&req)
		report.CompanyID = compIdUuid

		err = app.FinSvc.Create(r.Context(), &report)
		if err != nil {
			app.Logger.Infof("%s: создание финансового отчета: %v", prompt, err)
			web.ErrorResponse(wrappedWriter, fmt.Errorf("создание финансового отчета: %w", err).Error(), http.StatusInternalServerError)
			return
		}

		web.SuccessResponse(wrappedWriter, http.StatusNoContent, nil)
	}
}

// DeleteFinReport godoc
//
//	@Summary		Удаление информации об отчете
//	@ID				deleteFinReport
//	@Produce json
//	@Tags			financialsReports
//	@Param id path string true "Financial Report`s id"
//	@Success		204	{object} web.SuccessResponseStruct
//	@Failure		400	{object} web.ErrorResponseStruct
//	@Failure		401
//	@Failure 		500 {object} web.ErrorResponseStruct
//	@Security		BearerAuth
//	@Router			/financials/{id} [delete]
func DeleteFinReport(app *app.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		prompt := "DeleteFinReportHandler"
		start := time.Now()

		wrappedWriter := &web.StatusResponseWriter{ResponseWriter: w, StatusCodeOuter: http.StatusOK}

		defer func() {
			web.ObserveRequest(time.Since(start), wrappedWriter.StatusCode(), r.Method, prompt)
		}()

		userIdStr, err := web.GetStringClaimFromJWT(r.Context(), "sub")
		if err != nil {
			app.Logger.Infof("%s: получение записей из JWT: %v", prompt, err)
			web.ErrorResponse(wrappedWriter, fmt.Errorf("получение записей из JWT: %w", err).Error(), http.StatusBadRequest)
			return
		}

		userIdUuid, err := uuid.Parse(userIdStr)
		if err != nil {
			app.Logger.Infof("%s: преобразование строки к uuid: %v", prompt, err)
			web.ErrorResponse(wrappedWriter, fmt.Errorf("преобразование строки к uuid: %w", err).Error(), http.StatusInternalServerError)
			return
		}

		reportIdStr := chi.URLParam(r, "id")
		if reportIdStr == "" {
			app.Logger.Infof("%s: пустой id отчета", prompt)
			web.ErrorResponse(wrappedWriter, fmt.Errorf("пустой id отчета").Error(), http.StatusBadRequest)
			return
		}

		reportIdUuid, err := uuid.Parse(reportIdStr)
		if err != nil {
			app.Logger.Infof("%s: преобразование строки к uuid: %v", prompt, err)
			web.ErrorResponse(wrappedWriter, fmt.Errorf("преобразование строки к uuid: %w", err).Error(), http.StatusInternalServerError)
			return
		}

		err = app.FinSvc.DeleteById(r.Context(), reportIdUuid, userIdUuid)
		if err != nil {
			app.Logger.Infof("%s: только владелец компании может удалять финансовые отчеты", prompt)
			web.ErrorResponse(wrappedWriter, fmt.Errorf("удаление финансового отчета по id: %w", err).Error(), http.StatusInternalServerError)
			return
		}

		web.SuccessResponse(wrappedWriter, http.StatusNoContent, nil)
	}
}

// UpdateFinReport godoc
//
//	@Summary		Обновление отчета
//	@ID				updateFinReport
//	@Produce json
//	@Tags			financialsReports
//	@Param id path string true "Financial Report id"
//	@Param data body web.FinancialReport true "Financial Report data"
//	@Success		204	{object} web.SuccessResponseStruct
//	@Failure		400	{object} web.ErrorResponseStruct
//	@Failure		401
//	@Failure 		500 {object} web.ErrorResponseStruct
//	@Security		BearerAuth
//	@Router			/financials/{id} [patch]
func UpdateFinReport(app *app.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		prompt := "UpdateFinReportHandler"
		start := time.Now()

		wrappedWriter := &web.StatusResponseWriter{ResponseWriter: w, StatusCodeOuter: http.StatusOK}

		defer func() {
			web.ObserveRequest(time.Since(start), wrappedWriter.StatusCode(), r.Method, prompt)
		}()

		userIdStr, err := web.GetStringClaimFromJWT(r.Context(), "sub")
		if err != nil {
			app.Logger.Infof("%s: получение записей из JWT: %v", prompt, err)
			web.ErrorResponse(wrappedWriter, fmt.Errorf("получение записей из JWT: %w", err).Error(), http.StatusBadRequest)
			return
		}

		userIdUuid, err := uuid.Parse(userIdStr)
		if err != nil {
			app.Logger.Infof("%s: преобразование строки к uuid: %v", prompt, err)
			web.ErrorResponse(wrappedWriter, fmt.Errorf("преобразование строки к uuid: %w", err).Error(), http.StatusInternalServerError)
			return
		}

		reportIdStr := chi.URLParam(r, "id")
		if reportIdStr == "" {
			app.Logger.Infof("%s: пустой id отчета", prompt)
			web.ErrorResponse(wrappedWriter, fmt.Errorf("пустой id отчета").Error(), http.StatusBadRequest)
			return
		}

		reportIdUuid, err := uuid.Parse(reportIdStr)
		if err != nil {
			app.Logger.Infof("%s: преобразование строки к uuid: %v", prompt, err)
			web.ErrorResponse(wrappedWriter, fmt.Errorf("преобразование строки к uuid: %w", err).Error(), http.StatusInternalServerError)
			return
		}

		var req web.FinancialReport

		err = json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			app.Logger.Infof("%s: %v", prompt, err)
			web.ErrorResponse(wrappedWriter, err.Error(), http.StatusBadRequest)
			return
		}
		req.ID = reportIdUuid
		model := web.ToFinReportModel(&req)

		err = app.FinSvc.Update(r.Context(), &model, userIdUuid)
		if err != nil {
			app.Logger.Infof("%s: обновление информации о финансовом отчете: %v", prompt, err)
			web.ErrorResponse(wrappedWriter, fmt.Errorf("обновление информации о финансовом отчете: %w", err).Error(), http.StatusInternalServerError)
			return
		}

		web.SuccessResponse(wrappedWriter, http.StatusNoContent, nil)
	}
}

//// GetFinReport godoc
////
////	@Summary		Получение информации об отчете
////	@ID				getFinReport
////	@Produce json
////	@Tags			financialsReports
////	@Param id path string true "Financial Report`s id"
////	@Success		200	{object} FinReportResponse
////	@Failure		400	{object} web.ErrorResponseStruct
////	@Failure		401	{object} ErrorResponse
////	@Failure 		500 {object} web.ErrorResponseStruct
////	@Router			/financials [get]
//func GetFinReport(app *app.App) http.HandlerFunc {
//	return func(w http.ResponseWriter, r *http.Request) {
//		prompt := "GetFinReportHandler"
//		start := time.Now()
//
//		wrappedWriter := &statusResponseWriter{ResponseWriter: w, StatusCodeOuter: http.StatusOK}
//
//		defer func() {
//			observeRequest(time.Since(start), wrappedWriter.StatusCode(), r.Method, prompt)
//		}()
//
//		id := chi.URLParam(r, "id")
//		if id == "" {
//			app.Logger.Infof("%s: пустой id", prompt)
//			errorResponse(wrappedWriter, fmt.Errorf("пустой id").Error(), http.StatusBadRequest)
//			return
//		}
//
//		idUuid, err := uuid.Parse(id)
//		if err != nil {
//			app.Logger.Infof("%s: преобразование id к uuid: %v", prompt, err)
//			errorResponse(wrappedWriter, fmt.Errorf("преобразование id к uuid: %w", err).Error(), http.StatusBadRequest)
//			return
//		}
//
//		report, err := app.FinSvc.GetById(r.Context(), idUuid)
//		if err != nil {
//			app.Logger.Infof("%s: получение финансового отчета по id: %v", prompt, err)
//			errorResponse(wrappedWriter, fmt.Errorf("получение финансового отчета по id: %w", err).Error(), http.StatusInternalServerError)
//			return
//		}
//
//		successResponse(wrappedWriter, http.StatusOK, FinReportResponse{FinReport: toFinReportTransport(report)})
//	}
//}

// ListCompanyReports godoc
//
//	@Summary		Получение финансовых отчетов компании
//	@ID				getCompaniesReports
//	@Produce json
//	@Tags			companies
//	@Param id path string true "Company ID"
//	@Param start-year query int true "Start year"
//	@Param start-quarter query int true "Start quarter"
//	@Param end-year query int true "End year"
//	@Param end-quarter query int true "End quarter"
//	@Success		200	{object} FinReportByPeriodResponse
//	@Failure		400	{object} web.ErrorResponseStruct
//	@Failure		401
//	@Failure 		500 {object} web.ErrorResponseStruct
//	@Security		BearerAuth
//	@Router			/companies/{id}/financials [get]
func ListCompanyReports(app *app.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		prompt := "ListCompanyReportsHandler"
		start := time.Now()

		wrappedWriter := &web.StatusResponseWriter{ResponseWriter: w, StatusCodeOuter: http.StatusOK}

		defer func() {
			web.ObserveRequest(time.Since(start), wrappedWriter.StatusCode(), r.Method, prompt)
		}()

		period, err := web.ParsePeriodFromURL(r)
		if err != nil {
			app.Logger.Infof("%s: парсинг периода из URL: %v", prompt, err)
			web.ErrorResponse(wrappedWriter, fmt.Errorf("парсинг периода из URL: %w", err).Error(), http.StatusBadRequest)
			return
		}

		compIdUuid, err := web.ParseUUIDFromURL(r, "id", "company")
		if err != nil {
			app.Logger.Infof("%s: парсинг id компании из URL: %v", prompt, err)
			web.ErrorResponse(wrappedWriter, fmt.Errorf("парсинг id компании из URL: %w", err).Error(), http.StatusBadRequest)
			return
		}

		reports, err := app.FinSvc.GetByCompany(r.Context(), compIdUuid, period)
		if err != nil {
			app.Logger.Infof("%s: получение отчетов компании: %v", prompt, err)
			web.ErrorResponse(wrappedWriter, fmt.Errorf("получение отчетов компании: %w", err).Error(), http.StatusInternalServerError)
			return
		}

		reportsTransport := make([]web.FinancialReport, len(reports.Reports))
		for i, rep := range reports.Reports {
			reportsTransport[i] = web.ToFinReportTransport(&rep)
		}

		web.SuccessResponse(wrappedWriter, http.StatusOK,
			FinReportByPeriodResponse{
				CompanyId: compIdUuid,
				Period:    web.ToPeriodTransport(period),
				Revenue:   reports.Revenue(),
				Costs:     reports.Costs(),
				Profit:    reports.Profit(),
				Reports:   reportsTransport,
			},
		)
	}
}

// CalculateRating godoc
//
//	@Summary		Получение финансовых отчетов предпринимателя
//	@ID				calculateRating
//	@Produce json
//	@Tags			entrepreneurs
//	@Param id path string true "Entrepreneur ID"
//	@Success		200	{object} RatingResponse
//	@Failure		400	{object} web.ErrorResponseStruct
//	@Failure		401
//	@Failure 		500 {object} web.ErrorResponseStruct
//	@Security		BearerAuth
//	@Router			/entrepreneurs/{id}/rating [get]
func CalculateRating(app *app.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		prompt := "CalculateRatingHandler"
		start := time.Now()

		wrappedWriter := &web.StatusResponseWriter{ResponseWriter: w, StatusCodeOuter: http.StatusOK}

		defer func() {
			web.ObserveRequest(time.Since(start), wrappedWriter.StatusCode(), r.Method, prompt)
		}()

		id := chi.URLParam(r, "id")
		if id == "" {
			app.Logger.Infof("%s: пустой id", prompt)
			web.ErrorResponse(wrappedWriter, fmt.Errorf("пустой id").Error(), http.StatusBadRequest)
			return
		}

		idUuid, err := uuid.Parse(id)
		if err != nil {
			app.Logger.Infof("%s: преобразование id к uuid: %v", prompt, err)
			web.ErrorResponse(wrappedWriter, fmt.Errorf("преобразование id к uuid: %w", err).Error(), http.StatusBadRequest)
			return
		}

		rating, err := app.Interactor.CalculateUserRating(r.Context(), idUuid)
		if err != nil {
			app.Logger.Infof("%s: вычисление рейтинга предпринимателя: %v", prompt, err)
			web.ErrorResponse(wrappedWriter, fmt.Errorf("вычисление рейтинга предпринимателя: %w", err).Error(), http.StatusInternalServerError)
			return
		}

		web.SuccessResponse(wrappedWriter, http.StatusOK, RatingResponse{Rating: rating})
	}
}

// GetEntrepreneurFinancials godoc
//
//	@Summary		Получение информации об отчете
//	@ID				getEntrepreneurFinReport
//	@Produce json
//	@Tags			financialsReports
//	@Param entrepreneur-id query string true "Entrepreneur`s id"
//	@Success		200	{object} EntrepreneurReportResponse
//	@Failure		400	{object} web.ErrorResponseStruct
//	@Failure		401
//	@Failure 		500 {object} web.ErrorResponseStruct
//	@Security		BearerAuth
//	@Router			/financials [get]
func GetEntrepreneurFinancials(app *app.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		prompt := "GetEntrepreneurFinancials"
		start := time.Now()

		wrappedWriter := &web.StatusResponseWriter{ResponseWriter: w, StatusCodeOuter: http.StatusOK}

		defer func() {
			web.ObserveRequest(time.Since(start), wrappedWriter.StatusCode(), r.Method, prompt)
		}()

		id := r.URL.Query().Get("entrepreneur-id")
		if id == "" {
			app.Logger.Infof("%s: пустой id предпринимателя", prompt)
			web.ErrorResponse(wrappedWriter, fmt.Errorf("пустой id предпринимателя").Error(), http.StatusBadRequest)
			return
		}

		idUuid, err := uuid.Parse(id)
		if err != nil {
			app.Logger.Infof("%s: преобразование id к uuid: %v", prompt, err)
			web.ErrorResponse(wrappedWriter, fmt.Errorf("преобразование id к uuid: %w", err).Error(), http.StatusBadRequest)
			return
		}

		prevYear := time.Now().AddDate(-1, 0, 0).Year()
		period := &domain.Period{
			StartYear:    prevYear,
			EndYear:      prevYear,
			StartQuarter: 1,
			EndQuarter:   4,
		}

		rep, err := app.Interactor.GetUserFinancialReport(r.Context(), idUuid, period)
		if err != nil {
			app.Logger.Infof("%s: получение финансового отчета предпринимателя: %v", prompt, err)
			web.ErrorResponse(wrappedWriter, fmt.Errorf("получение финансового отчета предпринимателя: %w", err).Error(), http.StatusInternalServerError)
			return
		}

		web.SuccessResponse(wrappedWriter, http.StatusOK, EntrepreneurReportResponse{
			Revenue: rep.Revenue(),
			Costs:   rep.Costs(),
			Profit:  rep.Profit(),
			Taxes:   rep.Taxes,
			TaxLoad: rep.TaxLoad,
		})
	}
}
