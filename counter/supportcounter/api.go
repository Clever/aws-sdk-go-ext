// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package supportcounter

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/support"
)

// AddAttachmentsToSetRequest is a passthrough to the underlying AddAttachmentsToSetRequest.
// It will increment the count of requests made to AddAttachmentsToSet.
func (c *Support) AddAttachmentsToSetRequest(input *support.AddAttachmentsToSetInput) (req *request.Request, output *support.AddAttachmentsToSetOutput) {
	c.inc("AddAttachmentsToSet")
	return c.svc.AddAttachmentsToSetRequest(input)
}

// AddAttachmentsToSet is a passthrough to the underlying AddAttachmentsToSet method.
// It will increment the count of requests made to AddAttachmentsToSet.
func (c *Support) AddAttachmentsToSet(input *support.AddAttachmentsToSetInput) (*support.AddAttachmentsToSetOutput, error) {
	c.inc("AddAttachmentsToSet")
	return c.svc.AddAttachmentsToSet(input)
}

// AddAttachmentsToSetWithContext is a passthrough to the underlying AddAttachmentsToSetWithContext method.
// It will increment the count of requests made to AddAttachmentsToSet.
func (c *Support) AddAttachmentsToSetWithContext(ctx aws.Context, input *support.AddAttachmentsToSetInput, opts ...request.Option) (*support.AddAttachmentsToSetOutput, error) {
	c.inc("AddAttachmentsToSet")
	return c.svc.AddAttachmentsToSetWithContext(ctx, input, opts...)
}

// AddCommunicationToCaseRequest is a passthrough to the underlying AddCommunicationToCaseRequest.
// It will increment the count of requests made to AddCommunicationToCase.
func (c *Support) AddCommunicationToCaseRequest(input *support.AddCommunicationToCaseInput) (req *request.Request, output *support.AddCommunicationToCaseOutput) {
	c.inc("AddCommunicationToCase")
	return c.svc.AddCommunicationToCaseRequest(input)
}

// AddCommunicationToCase is a passthrough to the underlying AddCommunicationToCase method.
// It will increment the count of requests made to AddCommunicationToCase.
func (c *Support) AddCommunicationToCase(input *support.AddCommunicationToCaseInput) (*support.AddCommunicationToCaseOutput, error) {
	c.inc("AddCommunicationToCase")
	return c.svc.AddCommunicationToCase(input)
}

// AddCommunicationToCaseWithContext is a passthrough to the underlying AddCommunicationToCaseWithContext method.
// It will increment the count of requests made to AddCommunicationToCase.
func (c *Support) AddCommunicationToCaseWithContext(ctx aws.Context, input *support.AddCommunicationToCaseInput, opts ...request.Option) (*support.AddCommunicationToCaseOutput, error) {
	c.inc("AddCommunicationToCase")
	return c.svc.AddCommunicationToCaseWithContext(ctx, input, opts...)
}

// CreateCaseRequest is a passthrough to the underlying CreateCaseRequest.
// It will increment the count of requests made to CreateCase.
func (c *Support) CreateCaseRequest(input *support.CreateCaseInput) (req *request.Request, output *support.CreateCaseOutput) {
	c.inc("CreateCase")
	return c.svc.CreateCaseRequest(input)
}

// CreateCase is a passthrough to the underlying CreateCase method.
// It will increment the count of requests made to CreateCase.
func (c *Support) CreateCase(input *support.CreateCaseInput) (*support.CreateCaseOutput, error) {
	c.inc("CreateCase")
	return c.svc.CreateCase(input)
}

// CreateCaseWithContext is a passthrough to the underlying CreateCaseWithContext method.
// It will increment the count of requests made to CreateCase.
func (c *Support) CreateCaseWithContext(ctx aws.Context, input *support.CreateCaseInput, opts ...request.Option) (*support.CreateCaseOutput, error) {
	c.inc("CreateCase")
	return c.svc.CreateCaseWithContext(ctx, input, opts...)
}

// DescribeAttachmentRequest is a passthrough to the underlying DescribeAttachmentRequest.
// It will increment the count of requests made to DescribeAttachment.
func (c *Support) DescribeAttachmentRequest(input *support.DescribeAttachmentInput) (req *request.Request, output *support.DescribeAttachmentOutput) {
	c.inc("DescribeAttachment")
	return c.svc.DescribeAttachmentRequest(input)
}

// DescribeAttachment is a passthrough to the underlying DescribeAttachment method.
// It will increment the count of requests made to DescribeAttachment.
func (c *Support) DescribeAttachment(input *support.DescribeAttachmentInput) (*support.DescribeAttachmentOutput, error) {
	c.inc("DescribeAttachment")
	return c.svc.DescribeAttachment(input)
}

// DescribeAttachmentWithContext is a passthrough to the underlying DescribeAttachmentWithContext method.
// It will increment the count of requests made to DescribeAttachment.
func (c *Support) DescribeAttachmentWithContext(ctx aws.Context, input *support.DescribeAttachmentInput, opts ...request.Option) (*support.DescribeAttachmentOutput, error) {
	c.inc("DescribeAttachment")
	return c.svc.DescribeAttachmentWithContext(ctx, input, opts...)
}

// DescribeCasesRequest is a passthrough to the underlying DescribeCasesRequest.
// It will increment the count of requests made to DescribeCases.
func (c *Support) DescribeCasesRequest(input *support.DescribeCasesInput) (req *request.Request, output *support.DescribeCasesOutput) {
	c.inc("DescribeCases")
	return c.svc.DescribeCasesRequest(input)
}

// DescribeCases is a passthrough to the underlying DescribeCases method.
// It will increment the count of requests made to DescribeCases.
func (c *Support) DescribeCases(input *support.DescribeCasesInput) (*support.DescribeCasesOutput, error) {
	c.inc("DescribeCases")
	return c.svc.DescribeCases(input)
}

// DescribeCasesWithContext is a passthrough to the underlying DescribeCasesWithContext method.
// It will increment the count of requests made to DescribeCases.
func (c *Support) DescribeCasesWithContext(ctx aws.Context, input *support.DescribeCasesInput, opts ...request.Option) (*support.DescribeCasesOutput, error) {
	c.inc("DescribeCases")
	return c.svc.DescribeCasesWithContext(ctx, input, opts...)
}

// DescribeCasesPages is a passthrough to the underlying DescribeCasesPages method.
// It will increment the count of requests made to DescribeCases on each page.
// NOTE: this is slightly inaccurate in the case of errors, since the function will not be called.
// Use DescribeCasesPagesWithContext to avoid this.
func (c *Support) DescribeCasesPages(input *support.DescribeCasesInput, fn func(*support.DescribeCasesOutput, bool) bool) error {
	wrappedFn := func(page *support.DescribeCasesOutput, lastPage bool) bool {
		c.inc("DescribeCases")
		return fn(page, lastPage)
	}
	return c.svc.DescribeCasesPages(input, wrappedFn)
}

// DescribeCasesPagesWithContext is a passthrough to the underlying DescribeCasesPagesWithContext method.
// It will add a request.Option that will increment the count of requests made to DescribeCases when applied to the request.
func (c *Support) DescribeCasesPagesWithContext(ctx aws.Context, input *support.DescribeCasesInput, fn func(*support.DescribeCasesOutput, bool) bool, opts ...request.Option) error {
	opts = append(opts, c.incViaRequestOption("DescribeCases"))
	return c.svc.DescribeCasesPagesWithContext(ctx, input, fn, opts...)
}

// DescribeCommunicationsRequest is a passthrough to the underlying DescribeCommunicationsRequest.
// It will increment the count of requests made to DescribeCommunications.
func (c *Support) DescribeCommunicationsRequest(input *support.DescribeCommunicationsInput) (req *request.Request, output *support.DescribeCommunicationsOutput) {
	c.inc("DescribeCommunications")
	return c.svc.DescribeCommunicationsRequest(input)
}

// DescribeCommunications is a passthrough to the underlying DescribeCommunications method.
// It will increment the count of requests made to DescribeCommunications.
func (c *Support) DescribeCommunications(input *support.DescribeCommunicationsInput) (*support.DescribeCommunicationsOutput, error) {
	c.inc("DescribeCommunications")
	return c.svc.DescribeCommunications(input)
}

// DescribeCommunicationsWithContext is a passthrough to the underlying DescribeCommunicationsWithContext method.
// It will increment the count of requests made to DescribeCommunications.
func (c *Support) DescribeCommunicationsWithContext(ctx aws.Context, input *support.DescribeCommunicationsInput, opts ...request.Option) (*support.DescribeCommunicationsOutput, error) {
	c.inc("DescribeCommunications")
	return c.svc.DescribeCommunicationsWithContext(ctx, input, opts...)
}

// DescribeCommunicationsPages is a passthrough to the underlying DescribeCommunicationsPages method.
// It will increment the count of requests made to DescribeCommunications on each page.
// NOTE: this is slightly inaccurate in the case of errors, since the function will not be called.
// Use DescribeCommunicationsPagesWithContext to avoid this.
func (c *Support) DescribeCommunicationsPages(input *support.DescribeCommunicationsInput, fn func(*support.DescribeCommunicationsOutput, bool) bool) error {
	wrappedFn := func(page *support.DescribeCommunicationsOutput, lastPage bool) bool {
		c.inc("DescribeCommunications")
		return fn(page, lastPage)
	}
	return c.svc.DescribeCommunicationsPages(input, wrappedFn)
}

// DescribeCommunicationsPagesWithContext is a passthrough to the underlying DescribeCommunicationsPagesWithContext method.
// It will add a request.Option that will increment the count of requests made to DescribeCommunications when applied to the request.
func (c *Support) DescribeCommunicationsPagesWithContext(ctx aws.Context, input *support.DescribeCommunicationsInput, fn func(*support.DescribeCommunicationsOutput, bool) bool, opts ...request.Option) error {
	opts = append(opts, c.incViaRequestOption("DescribeCommunications"))
	return c.svc.DescribeCommunicationsPagesWithContext(ctx, input, fn, opts...)
}

// DescribeServicesRequest is a passthrough to the underlying DescribeServicesRequest.
// It will increment the count of requests made to DescribeServices.
func (c *Support) DescribeServicesRequest(input *support.DescribeServicesInput) (req *request.Request, output *support.DescribeServicesOutput) {
	c.inc("DescribeServices")
	return c.svc.DescribeServicesRequest(input)
}

// DescribeServices is a passthrough to the underlying DescribeServices method.
// It will increment the count of requests made to DescribeServices.
func (c *Support) DescribeServices(input *support.DescribeServicesInput) (*support.DescribeServicesOutput, error) {
	c.inc("DescribeServices")
	return c.svc.DescribeServices(input)
}

// DescribeServicesWithContext is a passthrough to the underlying DescribeServicesWithContext method.
// It will increment the count of requests made to DescribeServices.
func (c *Support) DescribeServicesWithContext(ctx aws.Context, input *support.DescribeServicesInput, opts ...request.Option) (*support.DescribeServicesOutput, error) {
	c.inc("DescribeServices")
	return c.svc.DescribeServicesWithContext(ctx, input, opts...)
}

// DescribeSeverityLevelsRequest is a passthrough to the underlying DescribeSeverityLevelsRequest.
// It will increment the count of requests made to DescribeSeverityLevels.
func (c *Support) DescribeSeverityLevelsRequest(input *support.DescribeSeverityLevelsInput) (req *request.Request, output *support.DescribeSeverityLevelsOutput) {
	c.inc("DescribeSeverityLevels")
	return c.svc.DescribeSeverityLevelsRequest(input)
}

// DescribeSeverityLevels is a passthrough to the underlying DescribeSeverityLevels method.
// It will increment the count of requests made to DescribeSeverityLevels.
func (c *Support) DescribeSeverityLevels(input *support.DescribeSeverityLevelsInput) (*support.DescribeSeverityLevelsOutput, error) {
	c.inc("DescribeSeverityLevels")
	return c.svc.DescribeSeverityLevels(input)
}

// DescribeSeverityLevelsWithContext is a passthrough to the underlying DescribeSeverityLevelsWithContext method.
// It will increment the count of requests made to DescribeSeverityLevels.
func (c *Support) DescribeSeverityLevelsWithContext(ctx aws.Context, input *support.DescribeSeverityLevelsInput, opts ...request.Option) (*support.DescribeSeverityLevelsOutput, error) {
	c.inc("DescribeSeverityLevels")
	return c.svc.DescribeSeverityLevelsWithContext(ctx, input, opts...)
}

// DescribeTrustedAdvisorCheckRefreshStatusesRequest is a passthrough to the underlying DescribeTrustedAdvisorCheckRefreshStatusesRequest.
// It will increment the count of requests made to DescribeTrustedAdvisorCheckRefreshStatuses.
func (c *Support) DescribeTrustedAdvisorCheckRefreshStatusesRequest(input *support.DescribeTrustedAdvisorCheckRefreshStatusesInput) (req *request.Request, output *support.DescribeTrustedAdvisorCheckRefreshStatusesOutput) {
	c.inc("DescribeTrustedAdvisorCheckRefreshStatuses")
	return c.svc.DescribeTrustedAdvisorCheckRefreshStatusesRequest(input)
}

// DescribeTrustedAdvisorCheckRefreshStatuses is a passthrough to the underlying DescribeTrustedAdvisorCheckRefreshStatuses method.
// It will increment the count of requests made to DescribeTrustedAdvisorCheckRefreshStatuses.
func (c *Support) DescribeTrustedAdvisorCheckRefreshStatuses(input *support.DescribeTrustedAdvisorCheckRefreshStatusesInput) (*support.DescribeTrustedAdvisorCheckRefreshStatusesOutput, error) {
	c.inc("DescribeTrustedAdvisorCheckRefreshStatuses")
	return c.svc.DescribeTrustedAdvisorCheckRefreshStatuses(input)
}

// DescribeTrustedAdvisorCheckRefreshStatusesWithContext is a passthrough to the underlying DescribeTrustedAdvisorCheckRefreshStatusesWithContext method.
// It will increment the count of requests made to DescribeTrustedAdvisorCheckRefreshStatuses.
func (c *Support) DescribeTrustedAdvisorCheckRefreshStatusesWithContext(ctx aws.Context, input *support.DescribeTrustedAdvisorCheckRefreshStatusesInput, opts ...request.Option) (*support.DescribeTrustedAdvisorCheckRefreshStatusesOutput, error) {
	c.inc("DescribeTrustedAdvisorCheckRefreshStatuses")
	return c.svc.DescribeTrustedAdvisorCheckRefreshStatusesWithContext(ctx, input, opts...)
}

// DescribeTrustedAdvisorCheckResultRequest is a passthrough to the underlying DescribeTrustedAdvisorCheckResultRequest.
// It will increment the count of requests made to DescribeTrustedAdvisorCheckResult.
func (c *Support) DescribeTrustedAdvisorCheckResultRequest(input *support.DescribeTrustedAdvisorCheckResultInput) (req *request.Request, output *support.DescribeTrustedAdvisorCheckResultOutput) {
	c.inc("DescribeTrustedAdvisorCheckResult")
	return c.svc.DescribeTrustedAdvisorCheckResultRequest(input)
}

// DescribeTrustedAdvisorCheckResult is a passthrough to the underlying DescribeTrustedAdvisorCheckResult method.
// It will increment the count of requests made to DescribeTrustedAdvisorCheckResult.
func (c *Support) DescribeTrustedAdvisorCheckResult(input *support.DescribeTrustedAdvisorCheckResultInput) (*support.DescribeTrustedAdvisorCheckResultOutput, error) {
	c.inc("DescribeTrustedAdvisorCheckResult")
	return c.svc.DescribeTrustedAdvisorCheckResult(input)
}

// DescribeTrustedAdvisorCheckResultWithContext is a passthrough to the underlying DescribeTrustedAdvisorCheckResultWithContext method.
// It will increment the count of requests made to DescribeTrustedAdvisorCheckResult.
func (c *Support) DescribeTrustedAdvisorCheckResultWithContext(ctx aws.Context, input *support.DescribeTrustedAdvisorCheckResultInput, opts ...request.Option) (*support.DescribeTrustedAdvisorCheckResultOutput, error) {
	c.inc("DescribeTrustedAdvisorCheckResult")
	return c.svc.DescribeTrustedAdvisorCheckResultWithContext(ctx, input, opts...)
}

// DescribeTrustedAdvisorCheckSummariesRequest is a passthrough to the underlying DescribeTrustedAdvisorCheckSummariesRequest.
// It will increment the count of requests made to DescribeTrustedAdvisorCheckSummaries.
func (c *Support) DescribeTrustedAdvisorCheckSummariesRequest(input *support.DescribeTrustedAdvisorCheckSummariesInput) (req *request.Request, output *support.DescribeTrustedAdvisorCheckSummariesOutput) {
	c.inc("DescribeTrustedAdvisorCheckSummaries")
	return c.svc.DescribeTrustedAdvisorCheckSummariesRequest(input)
}

// DescribeTrustedAdvisorCheckSummaries is a passthrough to the underlying DescribeTrustedAdvisorCheckSummaries method.
// It will increment the count of requests made to DescribeTrustedAdvisorCheckSummaries.
func (c *Support) DescribeTrustedAdvisorCheckSummaries(input *support.DescribeTrustedAdvisorCheckSummariesInput) (*support.DescribeTrustedAdvisorCheckSummariesOutput, error) {
	c.inc("DescribeTrustedAdvisorCheckSummaries")
	return c.svc.DescribeTrustedAdvisorCheckSummaries(input)
}

// DescribeTrustedAdvisorCheckSummariesWithContext is a passthrough to the underlying DescribeTrustedAdvisorCheckSummariesWithContext method.
// It will increment the count of requests made to DescribeTrustedAdvisorCheckSummaries.
func (c *Support) DescribeTrustedAdvisorCheckSummariesWithContext(ctx aws.Context, input *support.DescribeTrustedAdvisorCheckSummariesInput, opts ...request.Option) (*support.DescribeTrustedAdvisorCheckSummariesOutput, error) {
	c.inc("DescribeTrustedAdvisorCheckSummaries")
	return c.svc.DescribeTrustedAdvisorCheckSummariesWithContext(ctx, input, opts...)
}

// DescribeTrustedAdvisorChecksRequest is a passthrough to the underlying DescribeTrustedAdvisorChecksRequest.
// It will increment the count of requests made to DescribeTrustedAdvisorChecks.
func (c *Support) DescribeTrustedAdvisorChecksRequest(input *support.DescribeTrustedAdvisorChecksInput) (req *request.Request, output *support.DescribeTrustedAdvisorChecksOutput) {
	c.inc("DescribeTrustedAdvisorChecks")
	return c.svc.DescribeTrustedAdvisorChecksRequest(input)
}

// DescribeTrustedAdvisorChecks is a passthrough to the underlying DescribeTrustedAdvisorChecks method.
// It will increment the count of requests made to DescribeTrustedAdvisorChecks.
func (c *Support) DescribeTrustedAdvisorChecks(input *support.DescribeTrustedAdvisorChecksInput) (*support.DescribeTrustedAdvisorChecksOutput, error) {
	c.inc("DescribeTrustedAdvisorChecks")
	return c.svc.DescribeTrustedAdvisorChecks(input)
}

// DescribeTrustedAdvisorChecksWithContext is a passthrough to the underlying DescribeTrustedAdvisorChecksWithContext method.
// It will increment the count of requests made to DescribeTrustedAdvisorChecks.
func (c *Support) DescribeTrustedAdvisorChecksWithContext(ctx aws.Context, input *support.DescribeTrustedAdvisorChecksInput, opts ...request.Option) (*support.DescribeTrustedAdvisorChecksOutput, error) {
	c.inc("DescribeTrustedAdvisorChecks")
	return c.svc.DescribeTrustedAdvisorChecksWithContext(ctx, input, opts...)
}

// RefreshTrustedAdvisorCheckRequest is a passthrough to the underlying RefreshTrustedAdvisorCheckRequest.
// It will increment the count of requests made to RefreshTrustedAdvisorCheck.
func (c *Support) RefreshTrustedAdvisorCheckRequest(input *support.RefreshTrustedAdvisorCheckInput) (req *request.Request, output *support.RefreshTrustedAdvisorCheckOutput) {
	c.inc("RefreshTrustedAdvisorCheck")
	return c.svc.RefreshTrustedAdvisorCheckRequest(input)
}

// RefreshTrustedAdvisorCheck is a passthrough to the underlying RefreshTrustedAdvisorCheck method.
// It will increment the count of requests made to RefreshTrustedAdvisorCheck.
func (c *Support) RefreshTrustedAdvisorCheck(input *support.RefreshTrustedAdvisorCheckInput) (*support.RefreshTrustedAdvisorCheckOutput, error) {
	c.inc("RefreshTrustedAdvisorCheck")
	return c.svc.RefreshTrustedAdvisorCheck(input)
}

// RefreshTrustedAdvisorCheckWithContext is a passthrough to the underlying RefreshTrustedAdvisorCheckWithContext method.
// It will increment the count of requests made to RefreshTrustedAdvisorCheck.
func (c *Support) RefreshTrustedAdvisorCheckWithContext(ctx aws.Context, input *support.RefreshTrustedAdvisorCheckInput, opts ...request.Option) (*support.RefreshTrustedAdvisorCheckOutput, error) {
	c.inc("RefreshTrustedAdvisorCheck")
	return c.svc.RefreshTrustedAdvisorCheckWithContext(ctx, input, opts...)
}

// ResolveCaseRequest is a passthrough to the underlying ResolveCaseRequest.
// It will increment the count of requests made to ResolveCase.
func (c *Support) ResolveCaseRequest(input *support.ResolveCaseInput) (req *request.Request, output *support.ResolveCaseOutput) {
	c.inc("ResolveCase")
	return c.svc.ResolveCaseRequest(input)
}

// ResolveCase is a passthrough to the underlying ResolveCase method.
// It will increment the count of requests made to ResolveCase.
func (c *Support) ResolveCase(input *support.ResolveCaseInput) (*support.ResolveCaseOutput, error) {
	c.inc("ResolveCase")
	return c.svc.ResolveCase(input)
}

// ResolveCaseWithContext is a passthrough to the underlying ResolveCaseWithContext method.
// It will increment the count of requests made to ResolveCase.
func (c *Support) ResolveCaseWithContext(ctx aws.Context, input *support.ResolveCaseInput, opts ...request.Option) (*support.ResolveCaseOutput, error) {
	c.inc("ResolveCase")
	return c.svc.ResolveCaseWithContext(ctx, input, opts...)
}
