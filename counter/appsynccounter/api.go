// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package appsynccounter

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/appsync"
)

// CreateApiKeyRequest is a passthrough to the underlying CreateApiKeyRequest.
// It will increment the count of requests made to CreateApiKey.
func (c *AppSync) CreateApiKeyRequest(input *appsync.CreateApiKeyInput) (req *request.Request, output *appsync.CreateApiKeyOutput) {
	c.inc("CreateApiKey")
	return c.svc.CreateApiKeyRequest(input)
}

// CreateApiKey is a passthrough to the underlying CreateApiKey method.
// It will increment the count of requests made to CreateApiKey.
func (c *AppSync) CreateApiKey(input *appsync.CreateApiKeyInput) (*appsync.CreateApiKeyOutput, error) {
	c.inc("CreateApiKey")
	return c.svc.CreateApiKey(input)
}

// CreateApiKeyWithContext is a passthrough to the underlying CreateApiKeyWithContext method.
// It will increment the count of requests made to CreateApiKey.
func (c *AppSync) CreateApiKeyWithContext(ctx aws.Context, input *appsync.CreateApiKeyInput, opts ...request.Option) (*appsync.CreateApiKeyOutput, error) {
	c.inc("CreateApiKey")
	return c.svc.CreateApiKeyWithContext(ctx, input, opts...)
}

// CreateDataSourceRequest is a passthrough to the underlying CreateDataSourceRequest.
// It will increment the count of requests made to CreateDataSource.
func (c *AppSync) CreateDataSourceRequest(input *appsync.CreateDataSourceInput) (req *request.Request, output *appsync.CreateDataSourceOutput) {
	c.inc("CreateDataSource")
	return c.svc.CreateDataSourceRequest(input)
}

// CreateDataSource is a passthrough to the underlying CreateDataSource method.
// It will increment the count of requests made to CreateDataSource.
func (c *AppSync) CreateDataSource(input *appsync.CreateDataSourceInput) (*appsync.CreateDataSourceOutput, error) {
	c.inc("CreateDataSource")
	return c.svc.CreateDataSource(input)
}

// CreateDataSourceWithContext is a passthrough to the underlying CreateDataSourceWithContext method.
// It will increment the count of requests made to CreateDataSource.
func (c *AppSync) CreateDataSourceWithContext(ctx aws.Context, input *appsync.CreateDataSourceInput, opts ...request.Option) (*appsync.CreateDataSourceOutput, error) {
	c.inc("CreateDataSource")
	return c.svc.CreateDataSourceWithContext(ctx, input, opts...)
}

// CreateGraphqlApiRequest is a passthrough to the underlying CreateGraphqlApiRequest.
// It will increment the count of requests made to CreateGraphqlApi.
func (c *AppSync) CreateGraphqlApiRequest(input *appsync.CreateGraphqlApiInput) (req *request.Request, output *appsync.CreateGraphqlApiOutput) {
	c.inc("CreateGraphqlApi")
	return c.svc.CreateGraphqlApiRequest(input)
}

// CreateGraphqlApi is a passthrough to the underlying CreateGraphqlApi method.
// It will increment the count of requests made to CreateGraphqlApi.
func (c *AppSync) CreateGraphqlApi(input *appsync.CreateGraphqlApiInput) (*appsync.CreateGraphqlApiOutput, error) {
	c.inc("CreateGraphqlApi")
	return c.svc.CreateGraphqlApi(input)
}

// CreateGraphqlApiWithContext is a passthrough to the underlying CreateGraphqlApiWithContext method.
// It will increment the count of requests made to CreateGraphqlApi.
func (c *AppSync) CreateGraphqlApiWithContext(ctx aws.Context, input *appsync.CreateGraphqlApiInput, opts ...request.Option) (*appsync.CreateGraphqlApiOutput, error) {
	c.inc("CreateGraphqlApi")
	return c.svc.CreateGraphqlApiWithContext(ctx, input, opts...)
}

// CreateResolverRequest is a passthrough to the underlying CreateResolverRequest.
// It will increment the count of requests made to CreateResolver.
func (c *AppSync) CreateResolverRequest(input *appsync.CreateResolverInput) (req *request.Request, output *appsync.CreateResolverOutput) {
	c.inc("CreateResolver")
	return c.svc.CreateResolverRequest(input)
}

// CreateResolver is a passthrough to the underlying CreateResolver method.
// It will increment the count of requests made to CreateResolver.
func (c *AppSync) CreateResolver(input *appsync.CreateResolverInput) (*appsync.CreateResolverOutput, error) {
	c.inc("CreateResolver")
	return c.svc.CreateResolver(input)
}

// CreateResolverWithContext is a passthrough to the underlying CreateResolverWithContext method.
// It will increment the count of requests made to CreateResolver.
func (c *AppSync) CreateResolverWithContext(ctx aws.Context, input *appsync.CreateResolverInput, opts ...request.Option) (*appsync.CreateResolverOutput, error) {
	c.inc("CreateResolver")
	return c.svc.CreateResolverWithContext(ctx, input, opts...)
}

// CreateTypeRequest is a passthrough to the underlying CreateTypeRequest.
// It will increment the count of requests made to CreateType.
func (c *AppSync) CreateTypeRequest(input *appsync.CreateTypeInput) (req *request.Request, output *appsync.CreateTypeOutput) {
	c.inc("CreateType")
	return c.svc.CreateTypeRequest(input)
}

// CreateType is a passthrough to the underlying CreateType method.
// It will increment the count of requests made to CreateType.
func (c *AppSync) CreateType(input *appsync.CreateTypeInput) (*appsync.CreateTypeOutput, error) {
	c.inc("CreateType")
	return c.svc.CreateType(input)
}

// CreateTypeWithContext is a passthrough to the underlying CreateTypeWithContext method.
// It will increment the count of requests made to CreateType.
func (c *AppSync) CreateTypeWithContext(ctx aws.Context, input *appsync.CreateTypeInput, opts ...request.Option) (*appsync.CreateTypeOutput, error) {
	c.inc("CreateType")
	return c.svc.CreateTypeWithContext(ctx, input, opts...)
}

// DeleteApiKeyRequest is a passthrough to the underlying DeleteApiKeyRequest.
// It will increment the count of requests made to DeleteApiKey.
func (c *AppSync) DeleteApiKeyRequest(input *appsync.DeleteApiKeyInput) (req *request.Request, output *appsync.DeleteApiKeyOutput) {
	c.inc("DeleteApiKey")
	return c.svc.DeleteApiKeyRequest(input)
}

// DeleteApiKey is a passthrough to the underlying DeleteApiKey method.
// It will increment the count of requests made to DeleteApiKey.
func (c *AppSync) DeleteApiKey(input *appsync.DeleteApiKeyInput) (*appsync.DeleteApiKeyOutput, error) {
	c.inc("DeleteApiKey")
	return c.svc.DeleteApiKey(input)
}

// DeleteApiKeyWithContext is a passthrough to the underlying DeleteApiKeyWithContext method.
// It will increment the count of requests made to DeleteApiKey.
func (c *AppSync) DeleteApiKeyWithContext(ctx aws.Context, input *appsync.DeleteApiKeyInput, opts ...request.Option) (*appsync.DeleteApiKeyOutput, error) {
	c.inc("DeleteApiKey")
	return c.svc.DeleteApiKeyWithContext(ctx, input, opts...)
}

// DeleteDataSourceRequest is a passthrough to the underlying DeleteDataSourceRequest.
// It will increment the count of requests made to DeleteDataSource.
func (c *AppSync) DeleteDataSourceRequest(input *appsync.DeleteDataSourceInput) (req *request.Request, output *appsync.DeleteDataSourceOutput) {
	c.inc("DeleteDataSource")
	return c.svc.DeleteDataSourceRequest(input)
}

// DeleteDataSource is a passthrough to the underlying DeleteDataSource method.
// It will increment the count of requests made to DeleteDataSource.
func (c *AppSync) DeleteDataSource(input *appsync.DeleteDataSourceInput) (*appsync.DeleteDataSourceOutput, error) {
	c.inc("DeleteDataSource")
	return c.svc.DeleteDataSource(input)
}

// DeleteDataSourceWithContext is a passthrough to the underlying DeleteDataSourceWithContext method.
// It will increment the count of requests made to DeleteDataSource.
func (c *AppSync) DeleteDataSourceWithContext(ctx aws.Context, input *appsync.DeleteDataSourceInput, opts ...request.Option) (*appsync.DeleteDataSourceOutput, error) {
	c.inc("DeleteDataSource")
	return c.svc.DeleteDataSourceWithContext(ctx, input, opts...)
}

// DeleteGraphqlApiRequest is a passthrough to the underlying DeleteGraphqlApiRequest.
// It will increment the count of requests made to DeleteGraphqlApi.
func (c *AppSync) DeleteGraphqlApiRequest(input *appsync.DeleteGraphqlApiInput) (req *request.Request, output *appsync.DeleteGraphqlApiOutput) {
	c.inc("DeleteGraphqlApi")
	return c.svc.DeleteGraphqlApiRequest(input)
}

// DeleteGraphqlApi is a passthrough to the underlying DeleteGraphqlApi method.
// It will increment the count of requests made to DeleteGraphqlApi.
func (c *AppSync) DeleteGraphqlApi(input *appsync.DeleteGraphqlApiInput) (*appsync.DeleteGraphqlApiOutput, error) {
	c.inc("DeleteGraphqlApi")
	return c.svc.DeleteGraphqlApi(input)
}

// DeleteGraphqlApiWithContext is a passthrough to the underlying DeleteGraphqlApiWithContext method.
// It will increment the count of requests made to DeleteGraphqlApi.
func (c *AppSync) DeleteGraphqlApiWithContext(ctx aws.Context, input *appsync.DeleteGraphqlApiInput, opts ...request.Option) (*appsync.DeleteGraphqlApiOutput, error) {
	c.inc("DeleteGraphqlApi")
	return c.svc.DeleteGraphqlApiWithContext(ctx, input, opts...)
}

// DeleteResolverRequest is a passthrough to the underlying DeleteResolverRequest.
// It will increment the count of requests made to DeleteResolver.
func (c *AppSync) DeleteResolverRequest(input *appsync.DeleteResolverInput) (req *request.Request, output *appsync.DeleteResolverOutput) {
	c.inc("DeleteResolver")
	return c.svc.DeleteResolverRequest(input)
}

// DeleteResolver is a passthrough to the underlying DeleteResolver method.
// It will increment the count of requests made to DeleteResolver.
func (c *AppSync) DeleteResolver(input *appsync.DeleteResolverInput) (*appsync.DeleteResolverOutput, error) {
	c.inc("DeleteResolver")
	return c.svc.DeleteResolver(input)
}

// DeleteResolverWithContext is a passthrough to the underlying DeleteResolverWithContext method.
// It will increment the count of requests made to DeleteResolver.
func (c *AppSync) DeleteResolverWithContext(ctx aws.Context, input *appsync.DeleteResolverInput, opts ...request.Option) (*appsync.DeleteResolverOutput, error) {
	c.inc("DeleteResolver")
	return c.svc.DeleteResolverWithContext(ctx, input, opts...)
}

// DeleteTypeRequest is a passthrough to the underlying DeleteTypeRequest.
// It will increment the count of requests made to DeleteType.
func (c *AppSync) DeleteTypeRequest(input *appsync.DeleteTypeInput) (req *request.Request, output *appsync.DeleteTypeOutput) {
	c.inc("DeleteType")
	return c.svc.DeleteTypeRequest(input)
}

// DeleteType is a passthrough to the underlying DeleteType method.
// It will increment the count of requests made to DeleteType.
func (c *AppSync) DeleteType(input *appsync.DeleteTypeInput) (*appsync.DeleteTypeOutput, error) {
	c.inc("DeleteType")
	return c.svc.DeleteType(input)
}

// DeleteTypeWithContext is a passthrough to the underlying DeleteTypeWithContext method.
// It will increment the count of requests made to DeleteType.
func (c *AppSync) DeleteTypeWithContext(ctx aws.Context, input *appsync.DeleteTypeInput, opts ...request.Option) (*appsync.DeleteTypeOutput, error) {
	c.inc("DeleteType")
	return c.svc.DeleteTypeWithContext(ctx, input, opts...)
}

// GetDataSourceRequest is a passthrough to the underlying GetDataSourceRequest.
// It will increment the count of requests made to GetDataSource.
func (c *AppSync) GetDataSourceRequest(input *appsync.GetDataSourceInput) (req *request.Request, output *appsync.GetDataSourceOutput) {
	c.inc("GetDataSource")
	return c.svc.GetDataSourceRequest(input)
}

// GetDataSource is a passthrough to the underlying GetDataSource method.
// It will increment the count of requests made to GetDataSource.
func (c *AppSync) GetDataSource(input *appsync.GetDataSourceInput) (*appsync.GetDataSourceOutput, error) {
	c.inc("GetDataSource")
	return c.svc.GetDataSource(input)
}

// GetDataSourceWithContext is a passthrough to the underlying GetDataSourceWithContext method.
// It will increment the count of requests made to GetDataSource.
func (c *AppSync) GetDataSourceWithContext(ctx aws.Context, input *appsync.GetDataSourceInput, opts ...request.Option) (*appsync.GetDataSourceOutput, error) {
	c.inc("GetDataSource")
	return c.svc.GetDataSourceWithContext(ctx, input, opts...)
}

// GetGraphqlApiRequest is a passthrough to the underlying GetGraphqlApiRequest.
// It will increment the count of requests made to GetGraphqlApi.
func (c *AppSync) GetGraphqlApiRequest(input *appsync.GetGraphqlApiInput) (req *request.Request, output *appsync.GetGraphqlApiOutput) {
	c.inc("GetGraphqlApi")
	return c.svc.GetGraphqlApiRequest(input)
}

// GetGraphqlApi is a passthrough to the underlying GetGraphqlApi method.
// It will increment the count of requests made to GetGraphqlApi.
func (c *AppSync) GetGraphqlApi(input *appsync.GetGraphqlApiInput) (*appsync.GetGraphqlApiOutput, error) {
	c.inc("GetGraphqlApi")
	return c.svc.GetGraphqlApi(input)
}

// GetGraphqlApiWithContext is a passthrough to the underlying GetGraphqlApiWithContext method.
// It will increment the count of requests made to GetGraphqlApi.
func (c *AppSync) GetGraphqlApiWithContext(ctx aws.Context, input *appsync.GetGraphqlApiInput, opts ...request.Option) (*appsync.GetGraphqlApiOutput, error) {
	c.inc("GetGraphqlApi")
	return c.svc.GetGraphqlApiWithContext(ctx, input, opts...)
}

// GetIntrospectionSchemaRequest is a passthrough to the underlying GetIntrospectionSchemaRequest.
// It will increment the count of requests made to GetIntrospectionSchema.
func (c *AppSync) GetIntrospectionSchemaRequest(input *appsync.GetIntrospectionSchemaInput) (req *request.Request, output *appsync.GetIntrospectionSchemaOutput) {
	c.inc("GetIntrospectionSchema")
	return c.svc.GetIntrospectionSchemaRequest(input)
}

// GetIntrospectionSchema is a passthrough to the underlying GetIntrospectionSchema method.
// It will increment the count of requests made to GetIntrospectionSchema.
func (c *AppSync) GetIntrospectionSchema(input *appsync.GetIntrospectionSchemaInput) (*appsync.GetIntrospectionSchemaOutput, error) {
	c.inc("GetIntrospectionSchema")
	return c.svc.GetIntrospectionSchema(input)
}

// GetIntrospectionSchemaWithContext is a passthrough to the underlying GetIntrospectionSchemaWithContext method.
// It will increment the count of requests made to GetIntrospectionSchema.
func (c *AppSync) GetIntrospectionSchemaWithContext(ctx aws.Context, input *appsync.GetIntrospectionSchemaInput, opts ...request.Option) (*appsync.GetIntrospectionSchemaOutput, error) {
	c.inc("GetIntrospectionSchema")
	return c.svc.GetIntrospectionSchemaWithContext(ctx, input, opts...)
}

// GetResolverRequest is a passthrough to the underlying GetResolverRequest.
// It will increment the count of requests made to GetResolver.
func (c *AppSync) GetResolverRequest(input *appsync.GetResolverInput) (req *request.Request, output *appsync.GetResolverOutput) {
	c.inc("GetResolver")
	return c.svc.GetResolverRequest(input)
}

// GetResolver is a passthrough to the underlying GetResolver method.
// It will increment the count of requests made to GetResolver.
func (c *AppSync) GetResolver(input *appsync.GetResolverInput) (*appsync.GetResolverOutput, error) {
	c.inc("GetResolver")
	return c.svc.GetResolver(input)
}

// GetResolverWithContext is a passthrough to the underlying GetResolverWithContext method.
// It will increment the count of requests made to GetResolver.
func (c *AppSync) GetResolverWithContext(ctx aws.Context, input *appsync.GetResolverInput, opts ...request.Option) (*appsync.GetResolverOutput, error) {
	c.inc("GetResolver")
	return c.svc.GetResolverWithContext(ctx, input, opts...)
}

// GetSchemaCreationStatusRequest is a passthrough to the underlying GetSchemaCreationStatusRequest.
// It will increment the count of requests made to GetSchemaCreationStatus.
func (c *AppSync) GetSchemaCreationStatusRequest(input *appsync.GetSchemaCreationStatusInput) (req *request.Request, output *appsync.GetSchemaCreationStatusOutput) {
	c.inc("GetSchemaCreationStatus")
	return c.svc.GetSchemaCreationStatusRequest(input)
}

// GetSchemaCreationStatus is a passthrough to the underlying GetSchemaCreationStatus method.
// It will increment the count of requests made to GetSchemaCreationStatus.
func (c *AppSync) GetSchemaCreationStatus(input *appsync.GetSchemaCreationStatusInput) (*appsync.GetSchemaCreationStatusOutput, error) {
	c.inc("GetSchemaCreationStatus")
	return c.svc.GetSchemaCreationStatus(input)
}

// GetSchemaCreationStatusWithContext is a passthrough to the underlying GetSchemaCreationStatusWithContext method.
// It will increment the count of requests made to GetSchemaCreationStatus.
func (c *AppSync) GetSchemaCreationStatusWithContext(ctx aws.Context, input *appsync.GetSchemaCreationStatusInput, opts ...request.Option) (*appsync.GetSchemaCreationStatusOutput, error) {
	c.inc("GetSchemaCreationStatus")
	return c.svc.GetSchemaCreationStatusWithContext(ctx, input, opts...)
}

// GetTypeRequest is a passthrough to the underlying GetTypeRequest.
// It will increment the count of requests made to GetType.
func (c *AppSync) GetTypeRequest(input *appsync.GetTypeInput) (req *request.Request, output *appsync.GetTypeOutput) {
	c.inc("GetType")
	return c.svc.GetTypeRequest(input)
}

// GetType is a passthrough to the underlying GetType method.
// It will increment the count of requests made to GetType.
func (c *AppSync) GetType(input *appsync.GetTypeInput) (*appsync.GetTypeOutput, error) {
	c.inc("GetType")
	return c.svc.GetType(input)
}

// GetTypeWithContext is a passthrough to the underlying GetTypeWithContext method.
// It will increment the count of requests made to GetType.
func (c *AppSync) GetTypeWithContext(ctx aws.Context, input *appsync.GetTypeInput, opts ...request.Option) (*appsync.GetTypeOutput, error) {
	c.inc("GetType")
	return c.svc.GetTypeWithContext(ctx, input, opts...)
}

// ListApiKeysRequest is a passthrough to the underlying ListApiKeysRequest.
// It will increment the count of requests made to ListApiKeys.
func (c *AppSync) ListApiKeysRequest(input *appsync.ListApiKeysInput) (req *request.Request, output *appsync.ListApiKeysOutput) {
	c.inc("ListApiKeys")
	return c.svc.ListApiKeysRequest(input)
}

// ListApiKeys is a passthrough to the underlying ListApiKeys method.
// It will increment the count of requests made to ListApiKeys.
func (c *AppSync) ListApiKeys(input *appsync.ListApiKeysInput) (*appsync.ListApiKeysOutput, error) {
	c.inc("ListApiKeys")
	return c.svc.ListApiKeys(input)
}

// ListApiKeysWithContext is a passthrough to the underlying ListApiKeysWithContext method.
// It will increment the count of requests made to ListApiKeys.
func (c *AppSync) ListApiKeysWithContext(ctx aws.Context, input *appsync.ListApiKeysInput, opts ...request.Option) (*appsync.ListApiKeysOutput, error) {
	c.inc("ListApiKeys")
	return c.svc.ListApiKeysWithContext(ctx, input, opts...)
}

// ListDataSourcesRequest is a passthrough to the underlying ListDataSourcesRequest.
// It will increment the count of requests made to ListDataSources.
func (c *AppSync) ListDataSourcesRequest(input *appsync.ListDataSourcesInput) (req *request.Request, output *appsync.ListDataSourcesOutput) {
	c.inc("ListDataSources")
	return c.svc.ListDataSourcesRequest(input)
}

// ListDataSources is a passthrough to the underlying ListDataSources method.
// It will increment the count of requests made to ListDataSources.
func (c *AppSync) ListDataSources(input *appsync.ListDataSourcesInput) (*appsync.ListDataSourcesOutput, error) {
	c.inc("ListDataSources")
	return c.svc.ListDataSources(input)
}

// ListDataSourcesWithContext is a passthrough to the underlying ListDataSourcesWithContext method.
// It will increment the count of requests made to ListDataSources.
func (c *AppSync) ListDataSourcesWithContext(ctx aws.Context, input *appsync.ListDataSourcesInput, opts ...request.Option) (*appsync.ListDataSourcesOutput, error) {
	c.inc("ListDataSources")
	return c.svc.ListDataSourcesWithContext(ctx, input, opts...)
}

// ListGraphqlApisRequest is a passthrough to the underlying ListGraphqlApisRequest.
// It will increment the count of requests made to ListGraphqlApis.
func (c *AppSync) ListGraphqlApisRequest(input *appsync.ListGraphqlApisInput) (req *request.Request, output *appsync.ListGraphqlApisOutput) {
	c.inc("ListGraphqlApis")
	return c.svc.ListGraphqlApisRequest(input)
}

// ListGraphqlApis is a passthrough to the underlying ListGraphqlApis method.
// It will increment the count of requests made to ListGraphqlApis.
func (c *AppSync) ListGraphqlApis(input *appsync.ListGraphqlApisInput) (*appsync.ListGraphqlApisOutput, error) {
	c.inc("ListGraphqlApis")
	return c.svc.ListGraphqlApis(input)
}

// ListGraphqlApisWithContext is a passthrough to the underlying ListGraphqlApisWithContext method.
// It will increment the count of requests made to ListGraphqlApis.
func (c *AppSync) ListGraphqlApisWithContext(ctx aws.Context, input *appsync.ListGraphqlApisInput, opts ...request.Option) (*appsync.ListGraphqlApisOutput, error) {
	c.inc("ListGraphqlApis")
	return c.svc.ListGraphqlApisWithContext(ctx, input, opts...)
}

// ListResolversRequest is a passthrough to the underlying ListResolversRequest.
// It will increment the count of requests made to ListResolvers.
func (c *AppSync) ListResolversRequest(input *appsync.ListResolversInput) (req *request.Request, output *appsync.ListResolversOutput) {
	c.inc("ListResolvers")
	return c.svc.ListResolversRequest(input)
}

// ListResolvers is a passthrough to the underlying ListResolvers method.
// It will increment the count of requests made to ListResolvers.
func (c *AppSync) ListResolvers(input *appsync.ListResolversInput) (*appsync.ListResolversOutput, error) {
	c.inc("ListResolvers")
	return c.svc.ListResolvers(input)
}

// ListResolversWithContext is a passthrough to the underlying ListResolversWithContext method.
// It will increment the count of requests made to ListResolvers.
func (c *AppSync) ListResolversWithContext(ctx aws.Context, input *appsync.ListResolversInput, opts ...request.Option) (*appsync.ListResolversOutput, error) {
	c.inc("ListResolvers")
	return c.svc.ListResolversWithContext(ctx, input, opts...)
}

// ListTypesRequest is a passthrough to the underlying ListTypesRequest.
// It will increment the count of requests made to ListTypes.
func (c *AppSync) ListTypesRequest(input *appsync.ListTypesInput) (req *request.Request, output *appsync.ListTypesOutput) {
	c.inc("ListTypes")
	return c.svc.ListTypesRequest(input)
}

// ListTypes is a passthrough to the underlying ListTypes method.
// It will increment the count of requests made to ListTypes.
func (c *AppSync) ListTypes(input *appsync.ListTypesInput) (*appsync.ListTypesOutput, error) {
	c.inc("ListTypes")
	return c.svc.ListTypes(input)
}

// ListTypesWithContext is a passthrough to the underlying ListTypesWithContext method.
// It will increment the count of requests made to ListTypes.
func (c *AppSync) ListTypesWithContext(ctx aws.Context, input *appsync.ListTypesInput, opts ...request.Option) (*appsync.ListTypesOutput, error) {
	c.inc("ListTypes")
	return c.svc.ListTypesWithContext(ctx, input, opts...)
}

// StartSchemaCreationRequest is a passthrough to the underlying StartSchemaCreationRequest.
// It will increment the count of requests made to StartSchemaCreation.
func (c *AppSync) StartSchemaCreationRequest(input *appsync.StartSchemaCreationInput) (req *request.Request, output *appsync.StartSchemaCreationOutput) {
	c.inc("StartSchemaCreation")
	return c.svc.StartSchemaCreationRequest(input)
}

// StartSchemaCreation is a passthrough to the underlying StartSchemaCreation method.
// It will increment the count of requests made to StartSchemaCreation.
func (c *AppSync) StartSchemaCreation(input *appsync.StartSchemaCreationInput) (*appsync.StartSchemaCreationOutput, error) {
	c.inc("StartSchemaCreation")
	return c.svc.StartSchemaCreation(input)
}

// StartSchemaCreationWithContext is a passthrough to the underlying StartSchemaCreationWithContext method.
// It will increment the count of requests made to StartSchemaCreation.
func (c *AppSync) StartSchemaCreationWithContext(ctx aws.Context, input *appsync.StartSchemaCreationInput, opts ...request.Option) (*appsync.StartSchemaCreationOutput, error) {
	c.inc("StartSchemaCreation")
	return c.svc.StartSchemaCreationWithContext(ctx, input, opts...)
}

// UpdateApiKeyRequest is a passthrough to the underlying UpdateApiKeyRequest.
// It will increment the count of requests made to UpdateApiKey.
func (c *AppSync) UpdateApiKeyRequest(input *appsync.UpdateApiKeyInput) (req *request.Request, output *appsync.UpdateApiKeyOutput) {
	c.inc("UpdateApiKey")
	return c.svc.UpdateApiKeyRequest(input)
}

// UpdateApiKey is a passthrough to the underlying UpdateApiKey method.
// It will increment the count of requests made to UpdateApiKey.
func (c *AppSync) UpdateApiKey(input *appsync.UpdateApiKeyInput) (*appsync.UpdateApiKeyOutput, error) {
	c.inc("UpdateApiKey")
	return c.svc.UpdateApiKey(input)
}

// UpdateApiKeyWithContext is a passthrough to the underlying UpdateApiKeyWithContext method.
// It will increment the count of requests made to UpdateApiKey.
func (c *AppSync) UpdateApiKeyWithContext(ctx aws.Context, input *appsync.UpdateApiKeyInput, opts ...request.Option) (*appsync.UpdateApiKeyOutput, error) {
	c.inc("UpdateApiKey")
	return c.svc.UpdateApiKeyWithContext(ctx, input, opts...)
}

// UpdateDataSourceRequest is a passthrough to the underlying UpdateDataSourceRequest.
// It will increment the count of requests made to UpdateDataSource.
func (c *AppSync) UpdateDataSourceRequest(input *appsync.UpdateDataSourceInput) (req *request.Request, output *appsync.UpdateDataSourceOutput) {
	c.inc("UpdateDataSource")
	return c.svc.UpdateDataSourceRequest(input)
}

// UpdateDataSource is a passthrough to the underlying UpdateDataSource method.
// It will increment the count of requests made to UpdateDataSource.
func (c *AppSync) UpdateDataSource(input *appsync.UpdateDataSourceInput) (*appsync.UpdateDataSourceOutput, error) {
	c.inc("UpdateDataSource")
	return c.svc.UpdateDataSource(input)
}

// UpdateDataSourceWithContext is a passthrough to the underlying UpdateDataSourceWithContext method.
// It will increment the count of requests made to UpdateDataSource.
func (c *AppSync) UpdateDataSourceWithContext(ctx aws.Context, input *appsync.UpdateDataSourceInput, opts ...request.Option) (*appsync.UpdateDataSourceOutput, error) {
	c.inc("UpdateDataSource")
	return c.svc.UpdateDataSourceWithContext(ctx, input, opts...)
}

// UpdateGraphqlApiRequest is a passthrough to the underlying UpdateGraphqlApiRequest.
// It will increment the count of requests made to UpdateGraphqlApi.
func (c *AppSync) UpdateGraphqlApiRequest(input *appsync.UpdateGraphqlApiInput) (req *request.Request, output *appsync.UpdateGraphqlApiOutput) {
	c.inc("UpdateGraphqlApi")
	return c.svc.UpdateGraphqlApiRequest(input)
}

// UpdateGraphqlApi is a passthrough to the underlying UpdateGraphqlApi method.
// It will increment the count of requests made to UpdateGraphqlApi.
func (c *AppSync) UpdateGraphqlApi(input *appsync.UpdateGraphqlApiInput) (*appsync.UpdateGraphqlApiOutput, error) {
	c.inc("UpdateGraphqlApi")
	return c.svc.UpdateGraphqlApi(input)
}

// UpdateGraphqlApiWithContext is a passthrough to the underlying UpdateGraphqlApiWithContext method.
// It will increment the count of requests made to UpdateGraphqlApi.
func (c *AppSync) UpdateGraphqlApiWithContext(ctx aws.Context, input *appsync.UpdateGraphqlApiInput, opts ...request.Option) (*appsync.UpdateGraphqlApiOutput, error) {
	c.inc("UpdateGraphqlApi")
	return c.svc.UpdateGraphqlApiWithContext(ctx, input, opts...)
}

// UpdateResolverRequest is a passthrough to the underlying UpdateResolverRequest.
// It will increment the count of requests made to UpdateResolver.
func (c *AppSync) UpdateResolverRequest(input *appsync.UpdateResolverInput) (req *request.Request, output *appsync.UpdateResolverOutput) {
	c.inc("UpdateResolver")
	return c.svc.UpdateResolverRequest(input)
}

// UpdateResolver is a passthrough to the underlying UpdateResolver method.
// It will increment the count of requests made to UpdateResolver.
func (c *AppSync) UpdateResolver(input *appsync.UpdateResolverInput) (*appsync.UpdateResolverOutput, error) {
	c.inc("UpdateResolver")
	return c.svc.UpdateResolver(input)
}

// UpdateResolverWithContext is a passthrough to the underlying UpdateResolverWithContext method.
// It will increment the count of requests made to UpdateResolver.
func (c *AppSync) UpdateResolverWithContext(ctx aws.Context, input *appsync.UpdateResolverInput, opts ...request.Option) (*appsync.UpdateResolverOutput, error) {
	c.inc("UpdateResolver")
	return c.svc.UpdateResolverWithContext(ctx, input, opts...)
}

// UpdateTypeRequest is a passthrough to the underlying UpdateTypeRequest.
// It will increment the count of requests made to UpdateType.
func (c *AppSync) UpdateTypeRequest(input *appsync.UpdateTypeInput) (req *request.Request, output *appsync.UpdateTypeOutput) {
	c.inc("UpdateType")
	return c.svc.UpdateTypeRequest(input)
}

// UpdateType is a passthrough to the underlying UpdateType method.
// It will increment the count of requests made to UpdateType.
func (c *AppSync) UpdateType(input *appsync.UpdateTypeInput) (*appsync.UpdateTypeOutput, error) {
	c.inc("UpdateType")
	return c.svc.UpdateType(input)
}

// UpdateTypeWithContext is a passthrough to the underlying UpdateTypeWithContext method.
// It will increment the count of requests made to UpdateType.
func (c *AppSync) UpdateTypeWithContext(ctx aws.Context, input *appsync.UpdateTypeInput, opts ...request.Option) (*appsync.UpdateTypeOutput, error) {
	c.inc("UpdateType")
	return c.svc.UpdateTypeWithContext(ctx, input, opts...)
}
