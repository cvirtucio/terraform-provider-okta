// DO NOT EDIT LOCAL SDK - USE v3 okta-sdk-golang FOR API CALLS THAT DO NOT EXIST IN LOCAL SDK
package sdk

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"time"
)

type OrgSettingResource resource

type OrgSetting struct {
	Links                 interface{} `json:"_links,omitempty"`
	Address1              string      `json:"address1,omitempty"`
	Address2              string      `json:"address2,omitempty"`
	City                  string      `json:"city,omitempty"`
	CompanyName           string      `json:"companyName,omitempty"`
	Country               string      `json:"country,omitempty"`
	Created               *time.Time  `json:"created,omitempty"`
	EndUserSupportHelpURL string      `json:"endUserSupportHelpURL,omitempty"`
	ExpiresAt             *time.Time  `json:"expiresAt,omitempty"`
	Id                    string      `json:"id,omitempty"`
	LastUpdated           *time.Time  `json:"lastUpdated,omitempty"`
	PhoneNumber           string      `json:"phoneNumber,omitempty"`
	PostalCode            string      `json:"postalCode,omitempty"`
	State                 string      `json:"state,omitempty"`
	Status                string      `json:"status,omitempty"`
	Subdomain             string      `json:"subdomain,omitempty"`
	SupportPhoneNumber    string      `json:"supportPhoneNumber,omitempty"`
	Website               string      `json:"website,omitempty"`
}

// Get settings of your organization.
func (m *OrgSettingResource) GetOrgSettings(ctx context.Context) (*OrgSetting, *Response, error) {
	url := "/api/v1/org"

	rq := m.client.CloneRequestExecutor()

	req, err := rq.WithAccept("application/json").WithContentType("application/json").NewRequest("GET", url, nil)
	if err != nil {
		return nil, nil, err
	}

	var orgSetting *OrgSetting

	resp, err := rq.Do(ctx, req, &orgSetting)
	if err != nil {
		return nil, resp, err
	}

	return orgSetting, resp, nil
}

// Update settings of your organization.
func (m *OrgSettingResource) UpdateOrgSetting(ctx context.Context, body OrgSetting) (*OrgSetting, *Response, error) {
	url := "/api/v1/org"

	rq := m.client.CloneRequestExecutor()

	req, err := rq.WithAccept("application/json").WithContentType("application/json").NewRequest("PUT", url, body)
	if err != nil {
		return nil, nil, err
	}

	var orgSetting *OrgSetting

	resp, err := rq.Do(ctx, req, &orgSetting)
	if err != nil {
		return nil, resp, err
	}

	return orgSetting, resp, nil
}

// Partial update settings of your organization.
func (m *OrgSettingResource) PartialUpdateOrgSetting(ctx context.Context, body OrgSetting) (*OrgSetting, *Response, error) {
	url := "/api/v1/org"

	rq := m.client.CloneRequestExecutor()

	req, err := rq.WithAccept("application/json").WithContentType("application/json").NewRequest("POST", url, body)
	if err != nil {
		return nil, nil, err
	}

	var orgSetting *OrgSetting

	resp, err := rq.Do(ctx, req, &orgSetting)
	if err != nil {
		return nil, resp, err
	}

	return orgSetting, resp, nil
}

// Gets Contact Types of your organization.
func (m *OrgSettingResource) GetOrgContactTypes(ctx context.Context) ([]*OrgContactTypeObj, *Response, error) {
	url := "/api/v1/org/contacts"

	rq := m.client.CloneRequestExecutor()

	req, err := rq.WithAccept("application/json").WithContentType("application/json").NewRequest("GET", url, nil)
	if err != nil {
		return nil, nil, err
	}

	var orgContactTypeObj []*OrgContactTypeObj

	resp, err := rq.Do(ctx, req, &orgContactTypeObj)
	if err != nil {
		return nil, resp, err
	}

	return orgContactTypeObj, resp, nil
}

// Retrieves the URL of the User associated with the specified Contact Type.
func (m *OrgSettingResource) GetOrgContactUser(ctx context.Context, contactType string) (*OrgContactUser, *Response, error) {
	url := fmt.Sprintf("/api/v1/org/contacts/%v", contactType)

	rq := m.client.CloneRequestExecutor()

	req, err := rq.WithAccept("application/json").WithContentType("application/json").NewRequest("GET", url, nil)
	if err != nil {
		return nil, nil, err
	}

	var orgContactUser *OrgContactUser

	resp, err := rq.Do(ctx, req, &orgContactUser)
	if err != nil {
		return nil, resp, err
	}

	return orgContactUser, resp, nil
}

// Updates the User associated with the specified Contact Type.
func (m *OrgSettingResource) UpdateOrgContactUser(ctx context.Context, contactType string, body UserIdString) (*OrgContactUser, *Response, error) {
	url := fmt.Sprintf("/api/v1/org/contacts/%v", contactType)

	rq := m.client.CloneRequestExecutor()

	req, err := rq.WithAccept("application/json").WithContentType("application/json").NewRequest("PUT", url, body)
	if err != nil {
		return nil, nil, err
	}

	var orgContactUser *OrgContactUser

	resp, err := rq.Do(ctx, req, &orgContactUser)
	if err != nil {
		return nil, resp, err
	}

	return orgContactUser, resp, nil
}

// Updates the logo for your organization.
func (m *OrgSettingResource) UpdateOrgLogo(ctx context.Context, file string) (*Response, error) {
	url := "/api/v1/org/logo"

	rq := m.client.CloneRequestExecutor()

	fo, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer fo.Close()
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	fw, err := writer.CreateFormFile("file", file)
	if err != nil {
		return nil, err
	}
	_, err = io.Copy(fw, fo)
	if err != nil {
		return nil, err
	}
	_ = writer.Close()

	req, err := rq.WithAccept("application/json").WithContentType(writer.FormDataContentType()).NewRequest("POST", url, body)
	if err != nil {
		return nil, err
	}

	resp, err := m.client.requestExecutor.Do(ctx, req, nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// Gets preferences of your organization.
func (m *OrgSettingResource) GetOrgPreferences(ctx context.Context) (*OrgPreferences, *Response, error) {
	url := "/api/v1/org/preferences"

	rq := m.client.CloneRequestExecutor()

	req, err := rq.WithAccept("application/json").WithContentType("application/json").NewRequest("GET", url, nil)
	if err != nil {
		return nil, nil, err
	}

	var orgPreferences *OrgPreferences

	resp, err := rq.Do(ctx, req, &orgPreferences)
	if err != nil {
		return nil, resp, err
	}

	return orgPreferences, resp, nil
}

// Hide the Okta UI footer for all end users of your organization.
func (m *OrgSettingResource) HideOktaUIFooter(ctx context.Context) (*OrgPreferences, *Response, error) {
	url := "/api/v1/org/preferences/hideEndUserFooter"

	rq := m.client.CloneRequestExecutor()

	req, err := rq.WithAccept("application/json").WithContentType("application/json").NewRequest("POST", url, nil)
	if err != nil {
		return nil, nil, err
	}

	var orgPreferences *OrgPreferences

	resp, err := rq.Do(ctx, req, &orgPreferences)
	if err != nil {
		return nil, resp, err
	}

	return orgPreferences, resp, nil
}

// Makes the Okta UI footer visible for all end users of your organization.
func (m *OrgSettingResource) ShowOktaUIFooter(ctx context.Context) (*OrgPreferences, *Response, error) {
	url := "/api/v1/org/preferences/showEndUserFooter"

	rq := m.client.CloneRequestExecutor()

	req, err := rq.WithAccept("application/json").WithContentType("application/json").NewRequest("POST", url, nil)
	if err != nil {
		return nil, nil, err
	}

	var orgPreferences *OrgPreferences

	resp, err := rq.Do(ctx, req, &orgPreferences)
	if err != nil {
		return nil, resp, err
	}

	return orgPreferences, resp, nil
}

// Gets Okta Communication Settings of your organization.
func (m *OrgSettingResource) GetOktaCommunicationSettings(ctx context.Context) (*OrgOktaCommunicationSetting, *Response, error) {
	url := "/api/v1/org/privacy/oktaCommunication"

	rq := m.client.CloneRequestExecutor()

	req, err := rq.WithAccept("application/json").WithContentType("application/json").NewRequest("GET", url, nil)
	if err != nil {
		return nil, nil, err
	}

	var orgOktaCommunicationSetting *OrgOktaCommunicationSetting

	resp, err := rq.Do(ctx, req, &orgOktaCommunicationSetting)
	if err != nil {
		return nil, resp, err
	}

	return orgOktaCommunicationSetting, resp, nil
}

// Opts in all users of this org to Okta Communication emails.
func (m *OrgSettingResource) OptInUsersToOktaCommunicationEmails(ctx context.Context) (*OrgOktaCommunicationSetting, *Response, error) {
	url := "/api/v1/org/privacy/oktaCommunication/optIn"

	rq := m.client.CloneRequestExecutor()

	req, err := rq.WithAccept("application/json").WithContentType("application/json").NewRequest("POST", url, nil)
	if err != nil {
		return nil, nil, err
	}

	var orgOktaCommunicationSetting *OrgOktaCommunicationSetting

	resp, err := rq.Do(ctx, req, &orgOktaCommunicationSetting)
	if err != nil {
		return nil, resp, err
	}

	return orgOktaCommunicationSetting, resp, nil
}

// Opts out all users of this org from Okta Communication emails.
func (m *OrgSettingResource) OptOutUsersFromOktaCommunicationEmails(ctx context.Context) (*OrgOktaCommunicationSetting, *Response, error) {
	url := "/api/v1/org/privacy/oktaCommunication/optOut"

	rq := m.client.CloneRequestExecutor()

	req, err := rq.WithAccept("application/json").WithContentType("application/json").NewRequest("POST", url, nil)
	if err != nil {
		return nil, nil, err
	}

	var orgOktaCommunicationSetting *OrgOktaCommunicationSetting

	resp, err := rq.Do(ctx, req, &orgOktaCommunicationSetting)
	if err != nil {
		return nil, resp, err
	}

	return orgOktaCommunicationSetting, resp, nil
}

// Gets Okta Support Settings of your organization.
func (m *OrgSettingResource) GetOrgOktaSupportSettings(ctx context.Context) (*OrgOktaSupportSettingsObj, *Response, error) {
	url := "/api/v1/org/privacy/oktaSupport"

	rq := m.client.CloneRequestExecutor()

	req, err := rq.WithAccept("application/json").WithContentType("application/json").NewRequest("GET", url, nil)
	if err != nil {
		return nil, nil, err
	}

	var orgOktaSupportSettingsObj *OrgOktaSupportSettingsObj

	resp, err := rq.Do(ctx, req, &orgOktaSupportSettingsObj)
	if err != nil {
		return nil, resp, err
	}

	return orgOktaSupportSettingsObj, resp, nil
}

// Extends the length of time that Okta Support can access your org by 24 hours. This means that 24 hours are added to the remaining access time.
func (m *OrgSettingResource) ExtendOktaSupport(ctx context.Context) (*OrgOktaSupportSettingsObj, *Response, error) {
	url := "/api/v1/org/privacy/oktaSupport/extend"

	rq := m.client.CloneRequestExecutor()

	req, err := rq.WithAccept("application/json").WithContentType("application/json").NewRequest("POST", url, nil)
	if err != nil {
		return nil, nil, err
	}

	var orgOktaSupportSettingsObj *OrgOktaSupportSettingsObj

	resp, err := rq.Do(ctx, req, &orgOktaSupportSettingsObj)
	if err != nil {
		return nil, resp, err
	}

	return orgOktaSupportSettingsObj, resp, nil
}

// Enables you to temporarily allow Okta Support to access your org as an administrator for eight hours.
func (m *OrgSettingResource) GrantOktaSupport(ctx context.Context) (*OrgOktaSupportSettingsObj, *Response, error) {
	url := "/api/v1/org/privacy/oktaSupport/grant"

	rq := m.client.CloneRequestExecutor()

	req, err := rq.WithAccept("application/json").WithContentType("application/json").NewRequest("POST", url, nil)
	if err != nil {
		return nil, nil, err
	}

	var orgOktaSupportSettingsObj *OrgOktaSupportSettingsObj

	resp, err := rq.Do(ctx, req, &orgOktaSupportSettingsObj)
	if err != nil {
		return nil, resp, err
	}

	return orgOktaSupportSettingsObj, resp, nil
}

// Revokes Okta Support access to your organization.
func (m *OrgSettingResource) RevokeOktaSupport(ctx context.Context) (*OrgOktaSupportSettingsObj, *Response, error) {
	url := "/api/v1/org/privacy/oktaSupport/revoke"

	rq := m.client.CloneRequestExecutor()

	req, err := rq.WithAccept("application/json").WithContentType("application/json").NewRequest("POST", url, nil)
	if err != nil {
		return nil, nil, err
	}

	var orgOktaSupportSettingsObj *OrgOktaSupportSettingsObj

	resp, err := rq.Do(ctx, req, &orgOktaSupportSettingsObj)
	if err != nil {
		return nil, resp, err
	}

	return orgOktaSupportSettingsObj, resp, nil
}
