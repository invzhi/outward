// @generated by protoc-gen-es v2.0.0 with parameter "target=ts"
// @generated from file outward/v1/workspace_service.proto (package outward.v1, syntax proto3)
/* eslint-disable */

import type { GenEnum, GenFile, GenMessage, GenService } from "@bufbuild/protobuf/codegenv1";
import { enumDesc, fileDesc, messageDesc, serviceDesc } from "@bufbuild/protobuf/codegenv1";
import { file_google_api_annotations } from "../../google/api/annotations_pb";
import { file_buf_validate_validate } from "../../buf/validate/validate_pb";
import type { Message } from "@bufbuild/protobuf";

/**
 * Describes the file outward/v1/workspace_service.proto.
 */
export const file_outward_v1_workspace_service: GenFile = /*@__PURE__*/
  fileDesc("CiJvdXR3YXJkL3YxL3dvcmtzcGFjZV9zZXJ2aWNlLnByb3RvEgpvdXR3YXJkLnYxIlIKCVdvcmtzcGFjZRIKCgJpZBgBIAEoAxIMCgRuYW1lGAIgASgJEisKBnJlZ2lvbhgDIAEoDjIbLm91dHdhcmQudjEuV29ya3NwYWNlUmVnaW9uIlwKFkNyZWF0ZVdvcmtzcGFjZVJlcXVlc3QSFQoEbmFtZRgBIAEoCUIHukgEcgIQARIrCgZyZWdpb24YAiABKA4yGy5vdXR3YXJkLnYxLldvcmtzcGFjZVJlZ2lvbiJDChdDcmVhdGVXb3Jrc3BhY2VSZXNwb25zZRIoCgl3b3Jrc3BhY2UYASABKAsyFS5vdXR3YXJkLnYxLldvcmtzcGFjZSIZChdHZXRXb3Jrc3BhY2VMaXN0UmVxdWVzdCJFChhHZXRXb3Jrc3BhY2VMaXN0UmVzcG9uc2USKQoKd29ya3NwYWNlcxgBIAMoCzIVLm91dHdhcmQudjEuV29ya3NwYWNlKlMKD1dvcmtzcGFjZVJlZ2lvbhIgChxXT1JLU1BBQ0VfUkVHSU9OX1VOU1BFQ0lGSUVEEAASHgoaV09SS1NQQUNFX1JFR0lPTl9TSU5HQVBPUkUQATKAAgoQV29ya3NwYWNlU2VydmljZRJ1Cg9DcmVhdGVXb3Jrc3BhY2USIi5vdXR3YXJkLnYxLkNyZWF0ZVdvcmtzcGFjZVJlcXVlc3QaIy5vdXR3YXJkLnYxLkNyZWF0ZVdvcmtzcGFjZVJlc3BvbnNlIhmC0+STAhM6ASoiDi92MS93b3Jrc3BhY2VzEnUKEEdldFdvcmtzcGFjZUxpc3QSIy5vdXR3YXJkLnYxLkdldFdvcmtzcGFjZUxpc3RSZXF1ZXN0GiQub3V0d2FyZC52MS5HZXRXb3Jrc3BhY2VMaXN0UmVzcG9uc2UiFoLT5JMCEBIOL3YxL3dvcmtzcGFjZXNCfAoOY29tLm91dHdhcmQudjFCFVdvcmtzcGFjZVNlcnZpY2VQcm90b1ABWgpvdXR3YXJkL3YxogIDT1hYqgIKT3V0d2FyZC5WMcoCCk91dHdhcmRcVjHiAhZPdXR3YXJkXFYxXEdQQk1ldGFkYXRh6gILT3V0d2FyZDo6VjFiBnByb3RvMw", [file_google_api_annotations, file_buf_validate_validate]);

/**
 * @generated from message outward.v1.Workspace
 */
export type Workspace = Message<"outward.v1.Workspace"> & {
  /**
   * @generated from field: int64 id = 1;
   */
  id: bigint;

  /**
   * @generated from field: string name = 2;
   */
  name: string;

  /**
   * @generated from field: outward.v1.WorkspaceRegion region = 3;
   */
  region: WorkspaceRegion;
};

/**
 * Describes the message outward.v1.Workspace.
 * Use `create(WorkspaceSchema)` to create a new message.
 */
export const WorkspaceSchema: GenMessage<Workspace> = /*@__PURE__*/
  messageDesc(file_outward_v1_workspace_service, 0);

/**
 * @generated from message outward.v1.CreateWorkspaceRequest
 */
export type CreateWorkspaceRequest = Message<"outward.v1.CreateWorkspaceRequest"> & {
  /**
   * @generated from field: string name = 1;
   */
  name: string;

  /**
   * @generated from field: outward.v1.WorkspaceRegion region = 2;
   */
  region: WorkspaceRegion;
};

/**
 * Describes the message outward.v1.CreateWorkspaceRequest.
 * Use `create(CreateWorkspaceRequestSchema)` to create a new message.
 */
export const CreateWorkspaceRequestSchema: GenMessage<CreateWorkspaceRequest> = /*@__PURE__*/
  messageDesc(file_outward_v1_workspace_service, 1);

/**
 * @generated from message outward.v1.CreateWorkspaceResponse
 */
export type CreateWorkspaceResponse = Message<"outward.v1.CreateWorkspaceResponse"> & {
  /**
   * @generated from field: outward.v1.Workspace workspace = 1;
   */
  workspace?: Workspace;
};

/**
 * Describes the message outward.v1.CreateWorkspaceResponse.
 * Use `create(CreateWorkspaceResponseSchema)` to create a new message.
 */
export const CreateWorkspaceResponseSchema: GenMessage<CreateWorkspaceResponse> = /*@__PURE__*/
  messageDesc(file_outward_v1_workspace_service, 2);

/**
 * @generated from message outward.v1.GetWorkspaceListRequest
 */
export type GetWorkspaceListRequest = Message<"outward.v1.GetWorkspaceListRequest"> & {
};

/**
 * Describes the message outward.v1.GetWorkspaceListRequest.
 * Use `create(GetWorkspaceListRequestSchema)` to create a new message.
 */
export const GetWorkspaceListRequestSchema: GenMessage<GetWorkspaceListRequest> = /*@__PURE__*/
  messageDesc(file_outward_v1_workspace_service, 3);

/**
 * @generated from message outward.v1.GetWorkspaceListResponse
 */
export type GetWorkspaceListResponse = Message<"outward.v1.GetWorkspaceListResponse"> & {
  /**
   * @generated from field: repeated outward.v1.Workspace workspaces = 1;
   */
  workspaces: Workspace[];
};

/**
 * Describes the message outward.v1.GetWorkspaceListResponse.
 * Use `create(GetWorkspaceListResponseSchema)` to create a new message.
 */
export const GetWorkspaceListResponseSchema: GenMessage<GetWorkspaceListResponse> = /*@__PURE__*/
  messageDesc(file_outward_v1_workspace_service, 4);

/**
 * @generated from enum outward.v1.WorkspaceRegion
 */
export enum WorkspaceRegion {
  /**
   * @generated from enum value: WORKSPACE_REGION_UNSPECIFIED = 0;
   */
  UNSPECIFIED = 0,

  /**
   * @generated from enum value: WORKSPACE_REGION_SINGAPORE = 1;
   */
  SINGAPORE = 1,
}

/**
 * Describes the enum outward.v1.WorkspaceRegion.
 */
export const WorkspaceRegionSchema: GenEnum<WorkspaceRegion> = /*@__PURE__*/
  enumDesc(file_outward_v1_workspace_service, 0);

/**
 * @generated from service outward.v1.WorkspaceService
 */
export const WorkspaceService: GenService<{
  /**
   * @generated from rpc outward.v1.WorkspaceService.CreateWorkspace
   */
  createWorkspace: {
    methodKind: "unary";
    input: typeof CreateWorkspaceRequestSchema;
    output: typeof CreateWorkspaceResponseSchema;
  },
  /**
   * @generated from rpc outward.v1.WorkspaceService.GetWorkspaceList
   */
  getWorkspaceList: {
    methodKind: "unary";
    input: typeof GetWorkspaceListRequestSchema;
    output: typeof GetWorkspaceListResponseSchema;
  },
}> = /*@__PURE__*/
  serviceDesc(file_outward_v1_workspace_service, 0);

