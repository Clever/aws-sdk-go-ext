// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package route53domainscounter

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/route53domains"
)

// CheckDomainAvailabilityRequest is a passthrough to the underlying CheckDomainAvailabilityRequest.
// It will increment the count of requests made to CheckDomainAvailability.
func (c *Route53Domains) CheckDomainAvailabilityRequest(input *route53domains.CheckDomainAvailabilityInput) (req *request.Request, output *route53domains.CheckDomainAvailabilityOutput) {
	c.inc("CheckDomainAvailability")
	return c.svc.CheckDomainAvailabilityRequest(input)
}

// CheckDomainAvailability is a passthrough to the underlying CheckDomainAvailability method.
// It will increment the count of requests made to CheckDomainAvailability.
func (c *Route53Domains) CheckDomainAvailability(input *route53domains.CheckDomainAvailabilityInput) (*route53domains.CheckDomainAvailabilityOutput, error) {
	c.inc("CheckDomainAvailability")
	return c.svc.CheckDomainAvailability(input)
}

// CheckDomainAvailabilityWithContext is a passthrough to the underlying CheckDomainAvailabilityWithContext method.
// It will increment the count of requests made to CheckDomainAvailability.
func (c *Route53Domains) CheckDomainAvailabilityWithContext(ctx aws.Context, input *route53domains.CheckDomainAvailabilityInput, opts ...request.Option) (*route53domains.CheckDomainAvailabilityOutput, error) {
	c.inc("CheckDomainAvailability")
	return c.svc.CheckDomainAvailabilityWithContext(ctx, input, opts...)
}

// CheckDomainTransferabilityRequest is a passthrough to the underlying CheckDomainTransferabilityRequest.
// It will increment the count of requests made to CheckDomainTransferability.
func (c *Route53Domains) CheckDomainTransferabilityRequest(input *route53domains.CheckDomainTransferabilityInput) (req *request.Request, output *route53domains.CheckDomainTransferabilityOutput) {
	c.inc("CheckDomainTransferability")
	return c.svc.CheckDomainTransferabilityRequest(input)
}

// CheckDomainTransferability is a passthrough to the underlying CheckDomainTransferability method.
// It will increment the count of requests made to CheckDomainTransferability.
func (c *Route53Domains) CheckDomainTransferability(input *route53domains.CheckDomainTransferabilityInput) (*route53domains.CheckDomainTransferabilityOutput, error) {
	c.inc("CheckDomainTransferability")
	return c.svc.CheckDomainTransferability(input)
}

// CheckDomainTransferabilityWithContext is a passthrough to the underlying CheckDomainTransferabilityWithContext method.
// It will increment the count of requests made to CheckDomainTransferability.
func (c *Route53Domains) CheckDomainTransferabilityWithContext(ctx aws.Context, input *route53domains.CheckDomainTransferabilityInput, opts ...request.Option) (*route53domains.CheckDomainTransferabilityOutput, error) {
	c.inc("CheckDomainTransferability")
	return c.svc.CheckDomainTransferabilityWithContext(ctx, input, opts...)
}

// DeleteTagsForDomainRequest is a passthrough to the underlying DeleteTagsForDomainRequest.
// It will increment the count of requests made to DeleteTagsForDomain.
func (c *Route53Domains) DeleteTagsForDomainRequest(input *route53domains.DeleteTagsForDomainInput) (req *request.Request, output *route53domains.DeleteTagsForDomainOutput) {
	c.inc("DeleteTagsForDomain")
	return c.svc.DeleteTagsForDomainRequest(input)
}

// DeleteTagsForDomain is a passthrough to the underlying DeleteTagsForDomain method.
// It will increment the count of requests made to DeleteTagsForDomain.
func (c *Route53Domains) DeleteTagsForDomain(input *route53domains.DeleteTagsForDomainInput) (*route53domains.DeleteTagsForDomainOutput, error) {
	c.inc("DeleteTagsForDomain")
	return c.svc.DeleteTagsForDomain(input)
}

// DeleteTagsForDomainWithContext is a passthrough to the underlying DeleteTagsForDomainWithContext method.
// It will increment the count of requests made to DeleteTagsForDomain.
func (c *Route53Domains) DeleteTagsForDomainWithContext(ctx aws.Context, input *route53domains.DeleteTagsForDomainInput, opts ...request.Option) (*route53domains.DeleteTagsForDomainOutput, error) {
	c.inc("DeleteTagsForDomain")
	return c.svc.DeleteTagsForDomainWithContext(ctx, input, opts...)
}

// DisableDomainAutoRenewRequest is a passthrough to the underlying DisableDomainAutoRenewRequest.
// It will increment the count of requests made to DisableDomainAutoRenew.
func (c *Route53Domains) DisableDomainAutoRenewRequest(input *route53domains.DisableDomainAutoRenewInput) (req *request.Request, output *route53domains.DisableDomainAutoRenewOutput) {
	c.inc("DisableDomainAutoRenew")
	return c.svc.DisableDomainAutoRenewRequest(input)
}

// DisableDomainAutoRenew is a passthrough to the underlying DisableDomainAutoRenew method.
// It will increment the count of requests made to DisableDomainAutoRenew.
func (c *Route53Domains) DisableDomainAutoRenew(input *route53domains.DisableDomainAutoRenewInput) (*route53domains.DisableDomainAutoRenewOutput, error) {
	c.inc("DisableDomainAutoRenew")
	return c.svc.DisableDomainAutoRenew(input)
}

// DisableDomainAutoRenewWithContext is a passthrough to the underlying DisableDomainAutoRenewWithContext method.
// It will increment the count of requests made to DisableDomainAutoRenew.
func (c *Route53Domains) DisableDomainAutoRenewWithContext(ctx aws.Context, input *route53domains.DisableDomainAutoRenewInput, opts ...request.Option) (*route53domains.DisableDomainAutoRenewOutput, error) {
	c.inc("DisableDomainAutoRenew")
	return c.svc.DisableDomainAutoRenewWithContext(ctx, input, opts...)
}

// DisableDomainTransferLockRequest is a passthrough to the underlying DisableDomainTransferLockRequest.
// It will increment the count of requests made to DisableDomainTransferLock.
func (c *Route53Domains) DisableDomainTransferLockRequest(input *route53domains.DisableDomainTransferLockInput) (req *request.Request, output *route53domains.DisableDomainTransferLockOutput) {
	c.inc("DisableDomainTransferLock")
	return c.svc.DisableDomainTransferLockRequest(input)
}

// DisableDomainTransferLock is a passthrough to the underlying DisableDomainTransferLock method.
// It will increment the count of requests made to DisableDomainTransferLock.
func (c *Route53Domains) DisableDomainTransferLock(input *route53domains.DisableDomainTransferLockInput) (*route53domains.DisableDomainTransferLockOutput, error) {
	c.inc("DisableDomainTransferLock")
	return c.svc.DisableDomainTransferLock(input)
}

// DisableDomainTransferLockWithContext is a passthrough to the underlying DisableDomainTransferLockWithContext method.
// It will increment the count of requests made to DisableDomainTransferLock.
func (c *Route53Domains) DisableDomainTransferLockWithContext(ctx aws.Context, input *route53domains.DisableDomainTransferLockInput, opts ...request.Option) (*route53domains.DisableDomainTransferLockOutput, error) {
	c.inc("DisableDomainTransferLock")
	return c.svc.DisableDomainTransferLockWithContext(ctx, input, opts...)
}

// EnableDomainAutoRenewRequest is a passthrough to the underlying EnableDomainAutoRenewRequest.
// It will increment the count of requests made to EnableDomainAutoRenew.
func (c *Route53Domains) EnableDomainAutoRenewRequest(input *route53domains.EnableDomainAutoRenewInput) (req *request.Request, output *route53domains.EnableDomainAutoRenewOutput) {
	c.inc("EnableDomainAutoRenew")
	return c.svc.EnableDomainAutoRenewRequest(input)
}

// EnableDomainAutoRenew is a passthrough to the underlying EnableDomainAutoRenew method.
// It will increment the count of requests made to EnableDomainAutoRenew.
func (c *Route53Domains) EnableDomainAutoRenew(input *route53domains.EnableDomainAutoRenewInput) (*route53domains.EnableDomainAutoRenewOutput, error) {
	c.inc("EnableDomainAutoRenew")
	return c.svc.EnableDomainAutoRenew(input)
}

// EnableDomainAutoRenewWithContext is a passthrough to the underlying EnableDomainAutoRenewWithContext method.
// It will increment the count of requests made to EnableDomainAutoRenew.
func (c *Route53Domains) EnableDomainAutoRenewWithContext(ctx aws.Context, input *route53domains.EnableDomainAutoRenewInput, opts ...request.Option) (*route53domains.EnableDomainAutoRenewOutput, error) {
	c.inc("EnableDomainAutoRenew")
	return c.svc.EnableDomainAutoRenewWithContext(ctx, input, opts...)
}

// EnableDomainTransferLockRequest is a passthrough to the underlying EnableDomainTransferLockRequest.
// It will increment the count of requests made to EnableDomainTransferLock.
func (c *Route53Domains) EnableDomainTransferLockRequest(input *route53domains.EnableDomainTransferLockInput) (req *request.Request, output *route53domains.EnableDomainTransferLockOutput) {
	c.inc("EnableDomainTransferLock")
	return c.svc.EnableDomainTransferLockRequest(input)
}

// EnableDomainTransferLock is a passthrough to the underlying EnableDomainTransferLock method.
// It will increment the count of requests made to EnableDomainTransferLock.
func (c *Route53Domains) EnableDomainTransferLock(input *route53domains.EnableDomainTransferLockInput) (*route53domains.EnableDomainTransferLockOutput, error) {
	c.inc("EnableDomainTransferLock")
	return c.svc.EnableDomainTransferLock(input)
}

// EnableDomainTransferLockWithContext is a passthrough to the underlying EnableDomainTransferLockWithContext method.
// It will increment the count of requests made to EnableDomainTransferLock.
func (c *Route53Domains) EnableDomainTransferLockWithContext(ctx aws.Context, input *route53domains.EnableDomainTransferLockInput, opts ...request.Option) (*route53domains.EnableDomainTransferLockOutput, error) {
	c.inc("EnableDomainTransferLock")
	return c.svc.EnableDomainTransferLockWithContext(ctx, input, opts...)
}

// GetContactReachabilityStatusRequest is a passthrough to the underlying GetContactReachabilityStatusRequest.
// It will increment the count of requests made to GetContactReachabilityStatus.
func (c *Route53Domains) GetContactReachabilityStatusRequest(input *route53domains.GetContactReachabilityStatusInput) (req *request.Request, output *route53domains.GetContactReachabilityStatusOutput) {
	c.inc("GetContactReachabilityStatus")
	return c.svc.GetContactReachabilityStatusRequest(input)
}

// GetContactReachabilityStatus is a passthrough to the underlying GetContactReachabilityStatus method.
// It will increment the count of requests made to GetContactReachabilityStatus.
func (c *Route53Domains) GetContactReachabilityStatus(input *route53domains.GetContactReachabilityStatusInput) (*route53domains.GetContactReachabilityStatusOutput, error) {
	c.inc("GetContactReachabilityStatus")
	return c.svc.GetContactReachabilityStatus(input)
}

// GetContactReachabilityStatusWithContext is a passthrough to the underlying GetContactReachabilityStatusWithContext method.
// It will increment the count of requests made to GetContactReachabilityStatus.
func (c *Route53Domains) GetContactReachabilityStatusWithContext(ctx aws.Context, input *route53domains.GetContactReachabilityStatusInput, opts ...request.Option) (*route53domains.GetContactReachabilityStatusOutput, error) {
	c.inc("GetContactReachabilityStatus")
	return c.svc.GetContactReachabilityStatusWithContext(ctx, input, opts...)
}

// GetDomainDetailRequest is a passthrough to the underlying GetDomainDetailRequest.
// It will increment the count of requests made to GetDomainDetail.
func (c *Route53Domains) GetDomainDetailRequest(input *route53domains.GetDomainDetailInput) (req *request.Request, output *route53domains.GetDomainDetailOutput) {
	c.inc("GetDomainDetail")
	return c.svc.GetDomainDetailRequest(input)
}

// GetDomainDetail is a passthrough to the underlying GetDomainDetail method.
// It will increment the count of requests made to GetDomainDetail.
func (c *Route53Domains) GetDomainDetail(input *route53domains.GetDomainDetailInput) (*route53domains.GetDomainDetailOutput, error) {
	c.inc("GetDomainDetail")
	return c.svc.GetDomainDetail(input)
}

// GetDomainDetailWithContext is a passthrough to the underlying GetDomainDetailWithContext method.
// It will increment the count of requests made to GetDomainDetail.
func (c *Route53Domains) GetDomainDetailWithContext(ctx aws.Context, input *route53domains.GetDomainDetailInput, opts ...request.Option) (*route53domains.GetDomainDetailOutput, error) {
	c.inc("GetDomainDetail")
	return c.svc.GetDomainDetailWithContext(ctx, input, opts...)
}

// GetDomainSuggestionsRequest is a passthrough to the underlying GetDomainSuggestionsRequest.
// It will increment the count of requests made to GetDomainSuggestions.
func (c *Route53Domains) GetDomainSuggestionsRequest(input *route53domains.GetDomainSuggestionsInput) (req *request.Request, output *route53domains.GetDomainSuggestionsOutput) {
	c.inc("GetDomainSuggestions")
	return c.svc.GetDomainSuggestionsRequest(input)
}

// GetDomainSuggestions is a passthrough to the underlying GetDomainSuggestions method.
// It will increment the count of requests made to GetDomainSuggestions.
func (c *Route53Domains) GetDomainSuggestions(input *route53domains.GetDomainSuggestionsInput) (*route53domains.GetDomainSuggestionsOutput, error) {
	c.inc("GetDomainSuggestions")
	return c.svc.GetDomainSuggestions(input)
}

// GetDomainSuggestionsWithContext is a passthrough to the underlying GetDomainSuggestionsWithContext method.
// It will increment the count of requests made to GetDomainSuggestions.
func (c *Route53Domains) GetDomainSuggestionsWithContext(ctx aws.Context, input *route53domains.GetDomainSuggestionsInput, opts ...request.Option) (*route53domains.GetDomainSuggestionsOutput, error) {
	c.inc("GetDomainSuggestions")
	return c.svc.GetDomainSuggestionsWithContext(ctx, input, opts...)
}

// GetOperationDetailRequest is a passthrough to the underlying GetOperationDetailRequest.
// It will increment the count of requests made to GetOperationDetail.
func (c *Route53Domains) GetOperationDetailRequest(input *route53domains.GetOperationDetailInput) (req *request.Request, output *route53domains.GetOperationDetailOutput) {
	c.inc("GetOperationDetail")
	return c.svc.GetOperationDetailRequest(input)
}

// GetOperationDetail is a passthrough to the underlying GetOperationDetail method.
// It will increment the count of requests made to GetOperationDetail.
func (c *Route53Domains) GetOperationDetail(input *route53domains.GetOperationDetailInput) (*route53domains.GetOperationDetailOutput, error) {
	c.inc("GetOperationDetail")
	return c.svc.GetOperationDetail(input)
}

// GetOperationDetailWithContext is a passthrough to the underlying GetOperationDetailWithContext method.
// It will increment the count of requests made to GetOperationDetail.
func (c *Route53Domains) GetOperationDetailWithContext(ctx aws.Context, input *route53domains.GetOperationDetailInput, opts ...request.Option) (*route53domains.GetOperationDetailOutput, error) {
	c.inc("GetOperationDetail")
	return c.svc.GetOperationDetailWithContext(ctx, input, opts...)
}

// ListDomainsRequest is a passthrough to the underlying ListDomainsRequest.
// It will increment the count of requests made to ListDomains.
func (c *Route53Domains) ListDomainsRequest(input *route53domains.ListDomainsInput) (req *request.Request, output *route53domains.ListDomainsOutput) {
	c.inc("ListDomains")
	return c.svc.ListDomainsRequest(input)
}

// ListDomains is a passthrough to the underlying ListDomains method.
// It will increment the count of requests made to ListDomains.
func (c *Route53Domains) ListDomains(input *route53domains.ListDomainsInput) (*route53domains.ListDomainsOutput, error) {
	c.inc("ListDomains")
	return c.svc.ListDomains(input)
}

// ListDomainsWithContext is a passthrough to the underlying ListDomainsWithContext method.
// It will increment the count of requests made to ListDomains.
func (c *Route53Domains) ListDomainsWithContext(ctx aws.Context, input *route53domains.ListDomainsInput, opts ...request.Option) (*route53domains.ListDomainsOutput, error) {
	c.inc("ListDomains")
	return c.svc.ListDomainsWithContext(ctx, input, opts...)
}

// ListDomainsPages is a passthrough to the underlying ListDomainsPages method.
// It will increment the count of requests made to ListDomains on each page.
// NOTE: this is slightly inaccurate in the case of errors, since the function will not be called.
// Use ListDomainsPagesWithContext to avoid this.
func (c *Route53Domains) ListDomainsPages(input *route53domains.ListDomainsInput, fn func(*route53domains.ListDomainsOutput, bool) bool) error {
	wrappedFn := func(page *route53domains.ListDomainsOutput, lastPage bool) bool {
		c.inc("ListDomains")
		return fn(page, lastPage)
	}
	return c.svc.ListDomainsPages(input, wrappedFn)
}

// ListDomainsPagesWithContext is a passthrough to the underlying ListDomainsPagesWithContext method.
// It will add a request.Option that will increment the count of requests made to ListDomains when applied to the request.
func (c *Route53Domains) ListDomainsPagesWithContext(ctx aws.Context, input *route53domains.ListDomainsInput, fn func(*route53domains.ListDomainsOutput, bool) bool, opts ...request.Option) error {
	opts = append(opts, c.incViaRequestOption("ListDomains"))
	return c.svc.ListDomainsPagesWithContext(ctx, input, fn, opts...)
}

// ListOperationsRequest is a passthrough to the underlying ListOperationsRequest.
// It will increment the count of requests made to ListOperations.
func (c *Route53Domains) ListOperationsRequest(input *route53domains.ListOperationsInput) (req *request.Request, output *route53domains.ListOperationsOutput) {
	c.inc("ListOperations")
	return c.svc.ListOperationsRequest(input)
}

// ListOperations is a passthrough to the underlying ListOperations method.
// It will increment the count of requests made to ListOperations.
func (c *Route53Domains) ListOperations(input *route53domains.ListOperationsInput) (*route53domains.ListOperationsOutput, error) {
	c.inc("ListOperations")
	return c.svc.ListOperations(input)
}

// ListOperationsWithContext is a passthrough to the underlying ListOperationsWithContext method.
// It will increment the count of requests made to ListOperations.
func (c *Route53Domains) ListOperationsWithContext(ctx aws.Context, input *route53domains.ListOperationsInput, opts ...request.Option) (*route53domains.ListOperationsOutput, error) {
	c.inc("ListOperations")
	return c.svc.ListOperationsWithContext(ctx, input, opts...)
}

// ListOperationsPages is a passthrough to the underlying ListOperationsPages method.
// It will increment the count of requests made to ListOperations on each page.
// NOTE: this is slightly inaccurate in the case of errors, since the function will not be called.
// Use ListOperationsPagesWithContext to avoid this.
func (c *Route53Domains) ListOperationsPages(input *route53domains.ListOperationsInput, fn func(*route53domains.ListOperationsOutput, bool) bool) error {
	wrappedFn := func(page *route53domains.ListOperationsOutput, lastPage bool) bool {
		c.inc("ListOperations")
		return fn(page, lastPage)
	}
	return c.svc.ListOperationsPages(input, wrappedFn)
}

// ListOperationsPagesWithContext is a passthrough to the underlying ListOperationsPagesWithContext method.
// It will add a request.Option that will increment the count of requests made to ListOperations when applied to the request.
func (c *Route53Domains) ListOperationsPagesWithContext(ctx aws.Context, input *route53domains.ListOperationsInput, fn func(*route53domains.ListOperationsOutput, bool) bool, opts ...request.Option) error {
	opts = append(opts, c.incViaRequestOption("ListOperations"))
	return c.svc.ListOperationsPagesWithContext(ctx, input, fn, opts...)
}

// ListTagsForDomainRequest is a passthrough to the underlying ListTagsForDomainRequest.
// It will increment the count of requests made to ListTagsForDomain.
func (c *Route53Domains) ListTagsForDomainRequest(input *route53domains.ListTagsForDomainInput) (req *request.Request, output *route53domains.ListTagsForDomainOutput) {
	c.inc("ListTagsForDomain")
	return c.svc.ListTagsForDomainRequest(input)
}

// ListTagsForDomain is a passthrough to the underlying ListTagsForDomain method.
// It will increment the count of requests made to ListTagsForDomain.
func (c *Route53Domains) ListTagsForDomain(input *route53domains.ListTagsForDomainInput) (*route53domains.ListTagsForDomainOutput, error) {
	c.inc("ListTagsForDomain")
	return c.svc.ListTagsForDomain(input)
}

// ListTagsForDomainWithContext is a passthrough to the underlying ListTagsForDomainWithContext method.
// It will increment the count of requests made to ListTagsForDomain.
func (c *Route53Domains) ListTagsForDomainWithContext(ctx aws.Context, input *route53domains.ListTagsForDomainInput, opts ...request.Option) (*route53domains.ListTagsForDomainOutput, error) {
	c.inc("ListTagsForDomain")
	return c.svc.ListTagsForDomainWithContext(ctx, input, opts...)
}

// RegisterDomainRequest is a passthrough to the underlying RegisterDomainRequest.
// It will increment the count of requests made to RegisterDomain.
func (c *Route53Domains) RegisterDomainRequest(input *route53domains.RegisterDomainInput) (req *request.Request, output *route53domains.RegisterDomainOutput) {
	c.inc("RegisterDomain")
	return c.svc.RegisterDomainRequest(input)
}

// RegisterDomain is a passthrough to the underlying RegisterDomain method.
// It will increment the count of requests made to RegisterDomain.
func (c *Route53Domains) RegisterDomain(input *route53domains.RegisterDomainInput) (*route53domains.RegisterDomainOutput, error) {
	c.inc("RegisterDomain")
	return c.svc.RegisterDomain(input)
}

// RegisterDomainWithContext is a passthrough to the underlying RegisterDomainWithContext method.
// It will increment the count of requests made to RegisterDomain.
func (c *Route53Domains) RegisterDomainWithContext(ctx aws.Context, input *route53domains.RegisterDomainInput, opts ...request.Option) (*route53domains.RegisterDomainOutput, error) {
	c.inc("RegisterDomain")
	return c.svc.RegisterDomainWithContext(ctx, input, opts...)
}

// RenewDomainRequest is a passthrough to the underlying RenewDomainRequest.
// It will increment the count of requests made to RenewDomain.
func (c *Route53Domains) RenewDomainRequest(input *route53domains.RenewDomainInput) (req *request.Request, output *route53domains.RenewDomainOutput) {
	c.inc("RenewDomain")
	return c.svc.RenewDomainRequest(input)
}

// RenewDomain is a passthrough to the underlying RenewDomain method.
// It will increment the count of requests made to RenewDomain.
func (c *Route53Domains) RenewDomain(input *route53domains.RenewDomainInput) (*route53domains.RenewDomainOutput, error) {
	c.inc("RenewDomain")
	return c.svc.RenewDomain(input)
}

// RenewDomainWithContext is a passthrough to the underlying RenewDomainWithContext method.
// It will increment the count of requests made to RenewDomain.
func (c *Route53Domains) RenewDomainWithContext(ctx aws.Context, input *route53domains.RenewDomainInput, opts ...request.Option) (*route53domains.RenewDomainOutput, error) {
	c.inc("RenewDomain")
	return c.svc.RenewDomainWithContext(ctx, input, opts...)
}

// ResendContactReachabilityEmailRequest is a passthrough to the underlying ResendContactReachabilityEmailRequest.
// It will increment the count of requests made to ResendContactReachabilityEmail.
func (c *Route53Domains) ResendContactReachabilityEmailRequest(input *route53domains.ResendContactReachabilityEmailInput) (req *request.Request, output *route53domains.ResendContactReachabilityEmailOutput) {
	c.inc("ResendContactReachabilityEmail")
	return c.svc.ResendContactReachabilityEmailRequest(input)
}

// ResendContactReachabilityEmail is a passthrough to the underlying ResendContactReachabilityEmail method.
// It will increment the count of requests made to ResendContactReachabilityEmail.
func (c *Route53Domains) ResendContactReachabilityEmail(input *route53domains.ResendContactReachabilityEmailInput) (*route53domains.ResendContactReachabilityEmailOutput, error) {
	c.inc("ResendContactReachabilityEmail")
	return c.svc.ResendContactReachabilityEmail(input)
}

// ResendContactReachabilityEmailWithContext is a passthrough to the underlying ResendContactReachabilityEmailWithContext method.
// It will increment the count of requests made to ResendContactReachabilityEmail.
func (c *Route53Domains) ResendContactReachabilityEmailWithContext(ctx aws.Context, input *route53domains.ResendContactReachabilityEmailInput, opts ...request.Option) (*route53domains.ResendContactReachabilityEmailOutput, error) {
	c.inc("ResendContactReachabilityEmail")
	return c.svc.ResendContactReachabilityEmailWithContext(ctx, input, opts...)
}

// RetrieveDomainAuthCodeRequest is a passthrough to the underlying RetrieveDomainAuthCodeRequest.
// It will increment the count of requests made to RetrieveDomainAuthCode.
func (c *Route53Domains) RetrieveDomainAuthCodeRequest(input *route53domains.RetrieveDomainAuthCodeInput) (req *request.Request, output *route53domains.RetrieveDomainAuthCodeOutput) {
	c.inc("RetrieveDomainAuthCode")
	return c.svc.RetrieveDomainAuthCodeRequest(input)
}

// RetrieveDomainAuthCode is a passthrough to the underlying RetrieveDomainAuthCode method.
// It will increment the count of requests made to RetrieveDomainAuthCode.
func (c *Route53Domains) RetrieveDomainAuthCode(input *route53domains.RetrieveDomainAuthCodeInput) (*route53domains.RetrieveDomainAuthCodeOutput, error) {
	c.inc("RetrieveDomainAuthCode")
	return c.svc.RetrieveDomainAuthCode(input)
}

// RetrieveDomainAuthCodeWithContext is a passthrough to the underlying RetrieveDomainAuthCodeWithContext method.
// It will increment the count of requests made to RetrieveDomainAuthCode.
func (c *Route53Domains) RetrieveDomainAuthCodeWithContext(ctx aws.Context, input *route53domains.RetrieveDomainAuthCodeInput, opts ...request.Option) (*route53domains.RetrieveDomainAuthCodeOutput, error) {
	c.inc("RetrieveDomainAuthCode")
	return c.svc.RetrieveDomainAuthCodeWithContext(ctx, input, opts...)
}

// TransferDomainRequest is a passthrough to the underlying TransferDomainRequest.
// It will increment the count of requests made to TransferDomain.
func (c *Route53Domains) TransferDomainRequest(input *route53domains.TransferDomainInput) (req *request.Request, output *route53domains.TransferDomainOutput) {
	c.inc("TransferDomain")
	return c.svc.TransferDomainRequest(input)
}

// TransferDomain is a passthrough to the underlying TransferDomain method.
// It will increment the count of requests made to TransferDomain.
func (c *Route53Domains) TransferDomain(input *route53domains.TransferDomainInput) (*route53domains.TransferDomainOutput, error) {
	c.inc("TransferDomain")
	return c.svc.TransferDomain(input)
}

// TransferDomainWithContext is a passthrough to the underlying TransferDomainWithContext method.
// It will increment the count of requests made to TransferDomain.
func (c *Route53Domains) TransferDomainWithContext(ctx aws.Context, input *route53domains.TransferDomainInput, opts ...request.Option) (*route53domains.TransferDomainOutput, error) {
	c.inc("TransferDomain")
	return c.svc.TransferDomainWithContext(ctx, input, opts...)
}

// UpdateDomainContactRequest is a passthrough to the underlying UpdateDomainContactRequest.
// It will increment the count of requests made to UpdateDomainContact.
func (c *Route53Domains) UpdateDomainContactRequest(input *route53domains.UpdateDomainContactInput) (req *request.Request, output *route53domains.UpdateDomainContactOutput) {
	c.inc("UpdateDomainContact")
	return c.svc.UpdateDomainContactRequest(input)
}

// UpdateDomainContact is a passthrough to the underlying UpdateDomainContact method.
// It will increment the count of requests made to UpdateDomainContact.
func (c *Route53Domains) UpdateDomainContact(input *route53domains.UpdateDomainContactInput) (*route53domains.UpdateDomainContactOutput, error) {
	c.inc("UpdateDomainContact")
	return c.svc.UpdateDomainContact(input)
}

// UpdateDomainContactWithContext is a passthrough to the underlying UpdateDomainContactWithContext method.
// It will increment the count of requests made to UpdateDomainContact.
func (c *Route53Domains) UpdateDomainContactWithContext(ctx aws.Context, input *route53domains.UpdateDomainContactInput, opts ...request.Option) (*route53domains.UpdateDomainContactOutput, error) {
	c.inc("UpdateDomainContact")
	return c.svc.UpdateDomainContactWithContext(ctx, input, opts...)
}

// UpdateDomainContactPrivacyRequest is a passthrough to the underlying UpdateDomainContactPrivacyRequest.
// It will increment the count of requests made to UpdateDomainContactPrivacy.
func (c *Route53Domains) UpdateDomainContactPrivacyRequest(input *route53domains.UpdateDomainContactPrivacyInput) (req *request.Request, output *route53domains.UpdateDomainContactPrivacyOutput) {
	c.inc("UpdateDomainContactPrivacy")
	return c.svc.UpdateDomainContactPrivacyRequest(input)
}

// UpdateDomainContactPrivacy is a passthrough to the underlying UpdateDomainContactPrivacy method.
// It will increment the count of requests made to UpdateDomainContactPrivacy.
func (c *Route53Domains) UpdateDomainContactPrivacy(input *route53domains.UpdateDomainContactPrivacyInput) (*route53domains.UpdateDomainContactPrivacyOutput, error) {
	c.inc("UpdateDomainContactPrivacy")
	return c.svc.UpdateDomainContactPrivacy(input)
}

// UpdateDomainContactPrivacyWithContext is a passthrough to the underlying UpdateDomainContactPrivacyWithContext method.
// It will increment the count of requests made to UpdateDomainContactPrivacy.
func (c *Route53Domains) UpdateDomainContactPrivacyWithContext(ctx aws.Context, input *route53domains.UpdateDomainContactPrivacyInput, opts ...request.Option) (*route53domains.UpdateDomainContactPrivacyOutput, error) {
	c.inc("UpdateDomainContactPrivacy")
	return c.svc.UpdateDomainContactPrivacyWithContext(ctx, input, opts...)
}

// UpdateDomainNameserversRequest is a passthrough to the underlying UpdateDomainNameserversRequest.
// It will increment the count of requests made to UpdateDomainNameservers.
func (c *Route53Domains) UpdateDomainNameserversRequest(input *route53domains.UpdateDomainNameserversInput) (req *request.Request, output *route53domains.UpdateDomainNameserversOutput) {
	c.inc("UpdateDomainNameservers")
	return c.svc.UpdateDomainNameserversRequest(input)
}

// UpdateDomainNameservers is a passthrough to the underlying UpdateDomainNameservers method.
// It will increment the count of requests made to UpdateDomainNameservers.
func (c *Route53Domains) UpdateDomainNameservers(input *route53domains.UpdateDomainNameserversInput) (*route53domains.UpdateDomainNameserversOutput, error) {
	c.inc("UpdateDomainNameservers")
	return c.svc.UpdateDomainNameservers(input)
}

// UpdateDomainNameserversWithContext is a passthrough to the underlying UpdateDomainNameserversWithContext method.
// It will increment the count of requests made to UpdateDomainNameservers.
func (c *Route53Domains) UpdateDomainNameserversWithContext(ctx aws.Context, input *route53domains.UpdateDomainNameserversInput, opts ...request.Option) (*route53domains.UpdateDomainNameserversOutput, error) {
	c.inc("UpdateDomainNameservers")
	return c.svc.UpdateDomainNameserversWithContext(ctx, input, opts...)
}

// UpdateTagsForDomainRequest is a passthrough to the underlying UpdateTagsForDomainRequest.
// It will increment the count of requests made to UpdateTagsForDomain.
func (c *Route53Domains) UpdateTagsForDomainRequest(input *route53domains.UpdateTagsForDomainInput) (req *request.Request, output *route53domains.UpdateTagsForDomainOutput) {
	c.inc("UpdateTagsForDomain")
	return c.svc.UpdateTagsForDomainRequest(input)
}

// UpdateTagsForDomain is a passthrough to the underlying UpdateTagsForDomain method.
// It will increment the count of requests made to UpdateTagsForDomain.
func (c *Route53Domains) UpdateTagsForDomain(input *route53domains.UpdateTagsForDomainInput) (*route53domains.UpdateTagsForDomainOutput, error) {
	c.inc("UpdateTagsForDomain")
	return c.svc.UpdateTagsForDomain(input)
}

// UpdateTagsForDomainWithContext is a passthrough to the underlying UpdateTagsForDomainWithContext method.
// It will increment the count of requests made to UpdateTagsForDomain.
func (c *Route53Domains) UpdateTagsForDomainWithContext(ctx aws.Context, input *route53domains.UpdateTagsForDomainInput, opts ...request.Option) (*route53domains.UpdateTagsForDomainOutput, error) {
	c.inc("UpdateTagsForDomain")
	return c.svc.UpdateTagsForDomainWithContext(ctx, input, opts...)
}

// ViewBillingRequest is a passthrough to the underlying ViewBillingRequest.
// It will increment the count of requests made to ViewBilling.
func (c *Route53Domains) ViewBillingRequest(input *route53domains.ViewBillingInput) (req *request.Request, output *route53domains.ViewBillingOutput) {
	c.inc("ViewBilling")
	return c.svc.ViewBillingRequest(input)
}

// ViewBilling is a passthrough to the underlying ViewBilling method.
// It will increment the count of requests made to ViewBilling.
func (c *Route53Domains) ViewBilling(input *route53domains.ViewBillingInput) (*route53domains.ViewBillingOutput, error) {
	c.inc("ViewBilling")
	return c.svc.ViewBilling(input)
}

// ViewBillingWithContext is a passthrough to the underlying ViewBillingWithContext method.
// It will increment the count of requests made to ViewBilling.
func (c *Route53Domains) ViewBillingWithContext(ctx aws.Context, input *route53domains.ViewBillingInput, opts ...request.Option) (*route53domains.ViewBillingOutput, error) {
	c.inc("ViewBilling")
	return c.svc.ViewBillingWithContext(ctx, input, opts...)
}
