// @generated by protoc-gen-es v1.10.0 with parameter "target=ts"
// @generated from file outward/v1/user_service.proto (package outward.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3, protoInt64 } from "@bufbuild/protobuf";

/**
 * @generated from enum outward.v1.WorkspaceRole
 */
export enum WorkspaceRole {
  /**
   * @generated from enum value: WORKSPACE_ROLE_UNSPECIFIED = 0;
   */
  UNSPECIFIED = 0,

  /**
   * @generated from enum value: WORKSPACE_ROLE_ADMIN = 1;
   */
  ADMIN = 1,
}
// Retrieve enum metadata with: proto3.getEnumType(WorkspaceRole)
proto3.util.setEnumType(WorkspaceRole, "outward.v1.WorkspaceRole", [
  { no: 0, name: "WORKSPACE_ROLE_UNSPECIFIED" },
  { no: 1, name: "WORKSPACE_ROLE_ADMIN" },
]);

/**
 * @generated from message outward.v1.User
 */
export class User extends Message<User> {
  /**
   * @generated from field: int64 id = 1;
   */
  id = protoInt64.zero;

  /**
   * @generated from field: string email = 2;
   */
  email = "";

  /**
   * @generated from field: string first_name = 3;
   */
  firstName = "";

  /**
   * @generated from field: string last_name = 4;
   */
  lastName = "";

  constructor(data?: PartialMessage<User>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "outward.v1.User";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "id", kind: "scalar", T: 3 /* ScalarType.INT64 */ },
    { no: 2, name: "email", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "first_name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 4, name: "last_name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): User {
    return new User().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): User {
    return new User().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): User {
    return new User().fromJsonString(jsonString, options);
  }

  static equals(a: User | PlainMessage<User> | undefined, b: User | PlainMessage<User> | undefined): boolean {
    return proto3.util.equals(User, a, b);
  }
}

/**
 * @generated from message outward.v1.CreateUserRequest
 */
export class CreateUserRequest extends Message<CreateUserRequest> {
  /**
   * @generated from field: string email = 1;
   */
  email = "";

  /**
   * @generated from field: string first_name = 2;
   */
  firstName = "";

  /**
   * @generated from field: string last_name = 3;
   */
  lastName = "";

  /**
   * @generated from field: string password = 4;
   */
  password = "";

  constructor(data?: PartialMessage<CreateUserRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "outward.v1.CreateUserRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "email", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "first_name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "last_name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 4, name: "password", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): CreateUserRequest {
    return new CreateUserRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): CreateUserRequest {
    return new CreateUserRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): CreateUserRequest {
    return new CreateUserRequest().fromJsonString(jsonString, options);
  }

  static equals(a: CreateUserRequest | PlainMessage<CreateUserRequest> | undefined, b: CreateUserRequest | PlainMessage<CreateUserRequest> | undefined): boolean {
    return proto3.util.equals(CreateUserRequest, a, b);
  }
}

/**
 * @generated from message outward.v1.CreateUserResponse
 */
export class CreateUserResponse extends Message<CreateUserResponse> {
  /**
   * @generated from field: outward.v1.User user = 1;
   */
  user?: User;

  constructor(data?: PartialMessage<CreateUserResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "outward.v1.CreateUserResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "user", kind: "message", T: User },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): CreateUserResponse {
    return new CreateUserResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): CreateUserResponse {
    return new CreateUserResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): CreateUserResponse {
    return new CreateUserResponse().fromJsonString(jsonString, options);
  }

  static equals(a: CreateUserResponse | PlainMessage<CreateUserResponse> | undefined, b: CreateUserResponse | PlainMessage<CreateUserResponse> | undefined): boolean {
    return proto3.util.equals(CreateUserResponse, a, b);
  }
}

/**
 * @generated from message outward.v1.GetUserListRequest
 */
export class GetUserListRequest extends Message<GetUserListRequest> {
  /**
   * @generated from field: int32 page_size = 1;
   */
  pageSize = 0;

  /**
   * @generated from field: string page_token = 2;
   */
  pageToken = "";

  /**
   * @generated from field: int64 workspace_id = 3;
   */
  workspaceId = protoInt64.zero;

  constructor(data?: PartialMessage<GetUserListRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "outward.v1.GetUserListRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "page_size", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 2, name: "page_token", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "workspace_id", kind: "scalar", T: 3 /* ScalarType.INT64 */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GetUserListRequest {
    return new GetUserListRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GetUserListRequest {
    return new GetUserListRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GetUserListRequest {
    return new GetUserListRequest().fromJsonString(jsonString, options);
  }

  static equals(a: GetUserListRequest | PlainMessage<GetUserListRequest> | undefined, b: GetUserListRequest | PlainMessage<GetUserListRequest> | undefined): boolean {
    return proto3.util.equals(GetUserListRequest, a, b);
  }
}

/**
 * @generated from message outward.v1.GetUserListResponse
 */
export class GetUserListResponse extends Message<GetUserListResponse> {
  /**
   * @generated from field: string next_page_token = 1;
   */
  nextPageToken = "";

  /**
   * @generated from field: repeated outward.v1.User users = 2;
   */
  users: User[] = [];

  constructor(data?: PartialMessage<GetUserListResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "outward.v1.GetUserListResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "next_page_token", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "users", kind: "message", T: User, repeated: true },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GetUserListResponse {
    return new GetUserListResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GetUserListResponse {
    return new GetUserListResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GetUserListResponse {
    return new GetUserListResponse().fromJsonString(jsonString, options);
  }

  static equals(a: GetUserListResponse | PlainMessage<GetUserListResponse> | undefined, b: GetUserListResponse | PlainMessage<GetUserListResponse> | undefined): boolean {
    return proto3.util.equals(GetUserListResponse, a, b);
  }
}

/**
 * @generated from message outward.v1.CreateWorkspaceMemberRequest
 */
export class CreateWorkspaceMemberRequest extends Message<CreateWorkspaceMemberRequest> {
  /**
   * @generated from field: int64 user_id = 1;
   */
  userId = protoInt64.zero;

  /**
   * @generated from field: int64 workspace_id = 2;
   */
  workspaceId = protoInt64.zero;

  /**
   * @generated from field: outward.v1.WorkspaceRole role = 3;
   */
  role = WorkspaceRole.UNSPECIFIED;

  constructor(data?: PartialMessage<CreateWorkspaceMemberRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "outward.v1.CreateWorkspaceMemberRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "user_id", kind: "scalar", T: 3 /* ScalarType.INT64 */ },
    { no: 2, name: "workspace_id", kind: "scalar", T: 3 /* ScalarType.INT64 */ },
    { no: 3, name: "role", kind: "enum", T: proto3.getEnumType(WorkspaceRole) },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): CreateWorkspaceMemberRequest {
    return new CreateWorkspaceMemberRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): CreateWorkspaceMemberRequest {
    return new CreateWorkspaceMemberRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): CreateWorkspaceMemberRequest {
    return new CreateWorkspaceMemberRequest().fromJsonString(jsonString, options);
  }

  static equals(a: CreateWorkspaceMemberRequest | PlainMessage<CreateWorkspaceMemberRequest> | undefined, b: CreateWorkspaceMemberRequest | PlainMessage<CreateWorkspaceMemberRequest> | undefined): boolean {
    return proto3.util.equals(CreateWorkspaceMemberRequest, a, b);
  }
}

/**
 * @generated from message outward.v1.CreateWorkspaceMemberResponse
 */
export class CreateWorkspaceMemberResponse extends Message<CreateWorkspaceMemberResponse> {
  constructor(data?: PartialMessage<CreateWorkspaceMemberResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "outward.v1.CreateWorkspaceMemberResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): CreateWorkspaceMemberResponse {
    return new CreateWorkspaceMemberResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): CreateWorkspaceMemberResponse {
    return new CreateWorkspaceMemberResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): CreateWorkspaceMemberResponse {
    return new CreateWorkspaceMemberResponse().fromJsonString(jsonString, options);
  }

  static equals(a: CreateWorkspaceMemberResponse | PlainMessage<CreateWorkspaceMemberResponse> | undefined, b: CreateWorkspaceMemberResponse | PlainMessage<CreateWorkspaceMemberResponse> | undefined): boolean {
    return proto3.util.equals(CreateWorkspaceMemberResponse, a, b);
  }
}

/**
 * @generated from message outward.v1.LoginRequest
 */
export class LoginRequest extends Message<LoginRequest> {
  /**
   * @generated from field: string email = 1;
   */
  email = "";

  /**
   * @generated from field: string password = 2;
   */
  password = "";

  constructor(data?: PartialMessage<LoginRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "outward.v1.LoginRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "email", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "password", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): LoginRequest {
    return new LoginRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): LoginRequest {
    return new LoginRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): LoginRequest {
    return new LoginRequest().fromJsonString(jsonString, options);
  }

  static equals(a: LoginRequest | PlainMessage<LoginRequest> | undefined, b: LoginRequest | PlainMessage<LoginRequest> | undefined): boolean {
    return proto3.util.equals(LoginRequest, a, b);
  }
}

/**
 * @generated from message outward.v1.LoginResponse
 */
export class LoginResponse extends Message<LoginResponse> {
  /**
   * @generated from field: string token = 1;
   */
  token = "";

  constructor(data?: PartialMessage<LoginResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "outward.v1.LoginResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "token", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): LoginResponse {
    return new LoginResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): LoginResponse {
    return new LoginResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): LoginResponse {
    return new LoginResponse().fromJsonString(jsonString, options);
  }

  static equals(a: LoginResponse | PlainMessage<LoginResponse> | undefined, b: LoginResponse | PlainMessage<LoginResponse> | undefined): boolean {
    return proto3.util.equals(LoginResponse, a, b);
  }
}

/**
 * @generated from message outward.v1.LogoutRequest
 */
export class LogoutRequest extends Message<LogoutRequest> {
  constructor(data?: PartialMessage<LogoutRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "outward.v1.LogoutRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): LogoutRequest {
    return new LogoutRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): LogoutRequest {
    return new LogoutRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): LogoutRequest {
    return new LogoutRequest().fromJsonString(jsonString, options);
  }

  static equals(a: LogoutRequest | PlainMessage<LogoutRequest> | undefined, b: LogoutRequest | PlainMessage<LogoutRequest> | undefined): boolean {
    return proto3.util.equals(LogoutRequest, a, b);
  }
}

/**
 * @generated from message outward.v1.LogoutResponse
 */
export class LogoutResponse extends Message<LogoutResponse> {
  constructor(data?: PartialMessage<LogoutResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "outward.v1.LogoutResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): LogoutResponse {
    return new LogoutResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): LogoutResponse {
    return new LogoutResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): LogoutResponse {
    return new LogoutResponse().fromJsonString(jsonString, options);
  }

  static equals(a: LogoutResponse | PlainMessage<LogoutResponse> | undefined, b: LogoutResponse | PlainMessage<LogoutResponse> | undefined): boolean {
    return proto3.util.equals(LogoutResponse, a, b);
  }
}

