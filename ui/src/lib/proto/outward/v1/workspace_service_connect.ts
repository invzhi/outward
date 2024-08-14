// @generated by protoc-gen-connect-es v1.4.0 with parameter "target=ts"
// @generated from file outward/v1/workspace_service.proto (package outward.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { CreateWorkspaceRequest, CreateWorkspaceResponse, GetWorkspaceListRequest, GetWorkspaceListResponse } from "./workspace_service_pb.js";
import { MethodKind } from "@bufbuild/protobuf";

/**
 * @generated from service outward.v1.WorkspaceService
 */
export const WorkspaceService = {
  typeName: "outward.v1.WorkspaceService",
  methods: {
    /**
     * @generated from rpc outward.v1.WorkspaceService.CreateWorkspace
     */
    createWorkspace: {
      name: "CreateWorkspace",
      I: CreateWorkspaceRequest,
      O: CreateWorkspaceResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc outward.v1.WorkspaceService.GetWorkspaceList
     */
    getWorkspaceList: {
      name: "GetWorkspaceList",
      I: GetWorkspaceListRequest,
      O: GetWorkspaceListResponse,
      kind: MethodKind.Unary,
    },
  }
} as const;

