package face

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/validation"
	"io"
	"net/http"
)

// GroupClient is the an API for face detection, verification, and identification.
type GroupClient struct {
	ManagementClient
}

// NewGroupClient creates an instance of the GroupClient client.
func NewGroupClient(subscriptionKey string, azureRegion AzureRegions) GroupClient {
	return GroupClient{New(subscriptionKey, azureRegion)}
}

// Detect detect human faces in an image and returns face locations, and optionally with faceIds, landmarks, and
// attributes.
//
// imageURL is a JSON document with a URL pointing to the image that is to be analyzed. returnFaceID is a value
// indicating whether the operation should return faceIds of detected faces. returnFaceLandmarks is a value indicating
// whether the operation should return landmarks of the detected faces. returnFaceAttributes is analyze and return the
// one or more specified face attributes in the comma-separated string like "returnFaceAttributes=age,gender".
// Supported face attributes include age, gender, headPose, smile, facialHair, glasses and emotion. Note that each face
// attribute analysis has additional computational and time cost.
func (client GroupClient) Detect(imageURL ImageURL, returnFaceID *bool, returnFaceLandmarks *bool, returnFaceAttributes string) (result ListDetectedFace, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: imageURL,
			Constraints: []validation.Constraint{{Target: "imageURL.URL", Name: validation.Null, Rule: true, Chain: nil}}}}); err != nil {
		return result, validation.NewErrorWithValidationError(err, "face.GroupClient", "Detect")
	}

	req, err := client.DetectPreparer(imageURL, returnFaceID, returnFaceLandmarks, returnFaceAttributes)
	if err != nil {
		err = autorest.NewErrorWithError(err, "face.GroupClient", "Detect", nil, "Failure preparing request")
		return
	}

	resp, err := client.DetectSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "face.GroupClient", "Detect", resp, "Failure sending request")
		return
	}

	result, err = client.DetectResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "face.GroupClient", "Detect", resp, "Failure responding to request")
	}

	return
}

// DetectPreparer prepares the Detect request.
func (client GroupClient) DetectPreparer(imageURL ImageURL, returnFaceID *bool, returnFaceLandmarks *bool, returnFaceAttributes string) (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"AzureRegion": client.AzureRegion,
	}

	queryParameters := map[string]interface{}{}
	if returnFaceID != nil {
		queryParameters["returnFaceId"] = autorest.Encode("query", *returnFaceID)
	}
	if returnFaceLandmarks != nil {
		queryParameters["returnFaceLandmarks"] = autorest.Encode("query", *returnFaceLandmarks)
	}
	if len(returnFaceAttributes) > 0 {
		queryParameters["returnFaceAttributes"] = autorest.Encode("query", returnFaceAttributes)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsPost(),
		autorest.WithCustomBaseURL("https://{AzureRegion}.api.cognitive.microsoft.com/face/v1.0", urlParameters),
		autorest.WithPath("/detect"),
		autorest.WithJSON(imageURL),
		autorest.WithQueryParameters(queryParameters),
		autorest.WithHeader("Ocp-Apim-Subscription-Key", client.SubscriptionKey))
	return preparer.Prepare(&http.Request{})
}

// DetectSender sends the Detect request. The method will close the
// http.Response Body if it receives an error.
func (client GroupClient) DetectSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client,
		req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// DetectResponder handles the response to the Detect request. The method always
// closes the http.Response Body.
func (client GroupClient) DetectResponder(resp *http.Response) (result ListDetectedFace, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Value),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// DetectInStream detect human faces in an image and returns face locations, and optionally with faceIds, landmarks,
// and attributes.
//
// imageParameter is an image stream. imageParameter will be closed upon successful return. Callers should ensure
// closure when receiving an error.returnFaceID is a value indicating whether the operation should return faceIds of
// detected faces. returnFaceLandmarks is a value indicating whether the operation should return landmarks of the
// detected faces. returnFaceAttributes is analyze and return the one or more specified face attributes in the
// comma-separated string like "returnFaceAttributes=age,gender". Supported face attributes include age, gender,
// headPose, smile, facialHair, glasses and emotion. Note that each face attribute analysis has additional
// computational and time cost.
func (client GroupClient) DetectInStream(imageParameter io.ReadCloser, returnFaceID *bool, returnFaceLandmarks *bool, returnFaceAttributes string) (result ListDetectedFace, err error) {
	req, err := client.DetectInStreamPreparer(imageParameter, returnFaceID, returnFaceLandmarks, returnFaceAttributes)
	if err != nil {
		err = autorest.NewErrorWithError(err, "face.GroupClient", "DetectInStream", nil, "Failure preparing request")
		return
	}

	resp, err := client.DetectInStreamSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "face.GroupClient", "DetectInStream", resp, "Failure sending request")
		return
	}

	result, err = client.DetectInStreamResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "face.GroupClient", "DetectInStream", resp, "Failure responding to request")
	}

	return
}

// DetectInStreamPreparer prepares the DetectInStream request.
func (client GroupClient) DetectInStreamPreparer(imageParameter io.ReadCloser, returnFaceID *bool, returnFaceLandmarks *bool, returnFaceAttributes string) (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"AzureRegion": client.AzureRegion,
	}

	queryParameters := map[string]interface{}{}
	if returnFaceID != nil {
		queryParameters["returnFaceId"] = autorest.Encode("query", *returnFaceID)
	}
	if returnFaceLandmarks != nil {
		queryParameters["returnFaceLandmarks"] = autorest.Encode("query", *returnFaceLandmarks)
	}
	if len(returnFaceAttributes) > 0 {
		queryParameters["returnFaceAttributes"] = autorest.Encode("query", returnFaceAttributes)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithCustomBaseURL("https://{AzureRegion}.api.cognitive.microsoft.com/face/v1.0", urlParameters),
		autorest.WithPath("/detect"),
		autorest.WithFile(imageParameter),
		autorest.WithQueryParameters(queryParameters),
		autorest.WithHeader("Ocp-Apim-Subscription-Key", client.SubscriptionKey))
	return preparer.Prepare(&http.Request{})
}

// DetectInStreamSender sends the DetectInStream request. The method will close the
// http.Response Body if it receives an error.
func (client GroupClient) DetectInStreamSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client,
		req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// DetectInStreamResponder handles the response to the DetectInStream request. The method always
// closes the http.Response Body.
func (client GroupClient) DetectInStreamResponder(resp *http.Response) (result ListDetectedFace, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Value),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// FindSimilar given query face's faceId, find the similar-looking faces from a faceId array or a faceListId.
//
func (client GroupClient) FindSimilar(body FindSimilarRequest) (result ListSimilarFaceResult, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: body,
			Constraints: []validation.Constraint{{Target: "body.FaceID", Name: validation.Null, Rule: true,
				Chain: []validation.Constraint{{Target: "body.FaceID", Name: validation.MaxLength, Rule: 64, Chain: nil}}},
				{Target: "body.FaceListID", Name: validation.Null, Rule: false,
					Chain: []validation.Constraint{{Target: "body.FaceListID", Name: validation.MaxLength, Rule: 64, Chain: nil},
						{Target: "body.FaceListID", Name: validation.Pattern, Rule: `^[a-z0-9-_]+$`, Chain: nil},
					}},
				{Target: "body.FaceIds", Name: validation.Null, Rule: false,
					Chain: []validation.Constraint{{Target: "body.FaceIds", Name: validation.MaxItems, Rule: 1000, Chain: nil}}},
				{Target: "body.MaxNumOfCandidatesReturned", Name: validation.Null, Rule: false,
					Chain: []validation.Constraint{{Target: "body.MaxNumOfCandidatesReturned", Name: validation.InclusiveMaximum, Rule: 1000, Chain: nil},
						{Target: "body.MaxNumOfCandidatesReturned", Name: validation.InclusiveMinimum, Rule: 1, Chain: nil},
					}}}}}); err != nil {
		return result, validation.NewErrorWithValidationError(err, "face.GroupClient", "FindSimilar")
	}

	req, err := client.FindSimilarPreparer(body)
	if err != nil {
		err = autorest.NewErrorWithError(err, "face.GroupClient", "FindSimilar", nil, "Failure preparing request")
		return
	}

	resp, err := client.FindSimilarSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "face.GroupClient", "FindSimilar", resp, "Failure sending request")
		return
	}

	result, err = client.FindSimilarResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "face.GroupClient", "FindSimilar", resp, "Failure responding to request")
	}

	return
}

// FindSimilarPreparer prepares the FindSimilar request.
func (client GroupClient) FindSimilarPreparer(body FindSimilarRequest) (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"AzureRegion": client.AzureRegion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsPost(),
		autorest.WithCustomBaseURL("https://{AzureRegion}.api.cognitive.microsoft.com/face/v1.0", urlParameters),
		autorest.WithPath("/findsimilars"),
		autorest.WithJSON(body),
		autorest.WithHeader("Ocp-Apim-Subscription-Key", client.SubscriptionKey))
	return preparer.Prepare(&http.Request{})
}

// FindSimilarSender sends the FindSimilar request. The method will close the
// http.Response Body if it receives an error.
func (client GroupClient) FindSimilarSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client,
		req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// FindSimilarResponder handles the response to the FindSimilar request. The method always
// closes the http.Response Body.
func (client GroupClient) FindSimilarResponder(resp *http.Response) (result ListSimilarFaceResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Value),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Group divide candidate faces into groups based on face similarity.
//
func (client GroupClient) Group(body GroupRequest) (result GroupResponse, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: body,
			Constraints: []validation.Constraint{{Target: "body.FaceIds", Name: validation.Null, Rule: true,
				Chain: []validation.Constraint{{Target: "body.FaceIds", Name: validation.MaxItems, Rule: 1000, Chain: nil}}}}}}); err != nil {
		return result, validation.NewErrorWithValidationError(err, "face.GroupClient", "Group")
	}

	req, err := client.GroupPreparer(body)
	if err != nil {
		err = autorest.NewErrorWithError(err, "face.GroupClient", "Group", nil, "Failure preparing request")
		return
	}

	resp, err := client.GroupSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "face.GroupClient", "Group", resp, "Failure sending request")
		return
	}

	result, err = client.GroupResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "face.GroupClient", "Group", resp, "Failure responding to request")
	}

	return
}

// GroupPreparer prepares the Group request.
func (client GroupClient) GroupPreparer(body GroupRequest) (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"AzureRegion": client.AzureRegion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsPost(),
		autorest.WithCustomBaseURL("https://{AzureRegion}.api.cognitive.microsoft.com/face/v1.0", urlParameters),
		autorest.WithPath("/group"),
		autorest.WithJSON(body),
		autorest.WithHeader("Ocp-Apim-Subscription-Key", client.SubscriptionKey))
	return preparer.Prepare(&http.Request{})
}

// GroupSender sends the Group request. The method will close the
// http.Response Body if it receives an error.
func (client GroupClient) GroupSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client,
		req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GroupResponder handles the response to the Group request. The method always
// closes the http.Response Body.
func (client GroupClient) GroupResponder(resp *http.Response) (result GroupResponse, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Identify identify unknown faces from a person group.
//
func (client GroupClient) Identify(body IdentifyRequest) (result ListIdentifyResultItem, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: body,
			Constraints: []validation.Constraint{{Target: "body.PersonGroupID", Name: validation.Null, Rule: true, Chain: nil},
				{Target: "body.FaceIds", Name: validation.Null, Rule: true,
					Chain: []validation.Constraint{{Target: "body.FaceIds", Name: validation.MaxItems, Rule: 1000, Chain: nil}}},
				{Target: "body.MaxNumOfCandidatesReturned", Name: validation.Null, Rule: false,
					Chain: []validation.Constraint{{Target: "body.MaxNumOfCandidatesReturned", Name: validation.InclusiveMaximum, Rule: 1000, Chain: nil},
						{Target: "body.MaxNumOfCandidatesReturned", Name: validation.InclusiveMinimum, Rule: 1, Chain: nil},
					}},
				{Target: "body.ConfidenceThreshold", Name: validation.Null, Rule: false,
					Chain: []validation.Constraint{{Target: "body.ConfidenceThreshold", Name: validation.InclusiveMaximum, Rule: 1, Chain: nil},
						{Target: "body.ConfidenceThreshold", Name: validation.InclusiveMinimum, Rule: 0, Chain: nil},
					}}}}}); err != nil {
		return result, validation.NewErrorWithValidationError(err, "face.GroupClient", "Identify")
	}

	req, err := client.IdentifyPreparer(body)
	if err != nil {
		err = autorest.NewErrorWithError(err, "face.GroupClient", "Identify", nil, "Failure preparing request")
		return
	}

	resp, err := client.IdentifySender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "face.GroupClient", "Identify", resp, "Failure sending request")
		return
	}

	result, err = client.IdentifyResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "face.GroupClient", "Identify", resp, "Failure responding to request")
	}

	return
}

// IdentifyPreparer prepares the Identify request.
func (client GroupClient) IdentifyPreparer(body IdentifyRequest) (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"AzureRegion": client.AzureRegion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsPost(),
		autorest.WithCustomBaseURL("https://{AzureRegion}.api.cognitive.microsoft.com/face/v1.0", urlParameters),
		autorest.WithPath("/identify"),
		autorest.WithJSON(body),
		autorest.WithHeader("Ocp-Apim-Subscription-Key", client.SubscriptionKey))
	return preparer.Prepare(&http.Request{})
}

// IdentifySender sends the Identify request. The method will close the
// http.Response Body if it receives an error.
func (client GroupClient) IdentifySender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client,
		req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// IdentifyResponder handles the response to the Identify request. The method always
// closes the http.Response Body.
func (client GroupClient) IdentifyResponder(resp *http.Response) (result ListIdentifyResultItem, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Value),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Verify verify whether two faces belong to a same person or whether one face belongs to a person.
//
func (client GroupClient) Verify(body VerifyRequest) (result VerifyResult, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: body,
			Constraints: []validation.Constraint{{Target: "body.FaceID", Name: validation.Null, Rule: true,
				Chain: []validation.Constraint{{Target: "body.FaceID", Name: validation.MaxLength, Rule: 64, Chain: nil}}},
				{Target: "body.PersonID", Name: validation.Null, Rule: true, Chain: nil},
				{Target: "body.PersonGroupID", Name: validation.Null, Rule: true, Chain: nil}}}}); err != nil {
		return result, validation.NewErrorWithValidationError(err, "face.GroupClient", "Verify")
	}

	req, err := client.VerifyPreparer(body)
	if err != nil {
		err = autorest.NewErrorWithError(err, "face.GroupClient", "Verify", nil, "Failure preparing request")
		return
	}

	resp, err := client.VerifySender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "face.GroupClient", "Verify", resp, "Failure sending request")
		return
	}

	result, err = client.VerifyResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "face.GroupClient", "Verify", resp, "Failure responding to request")
	}

	return
}

// VerifyPreparer prepares the Verify request.
func (client GroupClient) VerifyPreparer(body VerifyRequest) (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"AzureRegion": client.AzureRegion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsPost(),
		autorest.WithCustomBaseURL("https://{AzureRegion}.api.cognitive.microsoft.com/face/v1.0", urlParameters),
		autorest.WithPath("/verify"),
		autorest.WithJSON(body),
		autorest.WithHeader("Ocp-Apim-Subscription-Key", client.SubscriptionKey))
	return preparer.Prepare(&http.Request{})
}

// VerifySender sends the Verify request. The method will close the
// http.Response Body if it receives an error.
func (client GroupClient) VerifySender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client,
		req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// VerifyResponder handles the response to the Verify request. The method always
// closes the http.Response Body.
func (client GroupClient) VerifyResponder(resp *http.Response) (result VerifyResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}