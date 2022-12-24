package config

/** GLOBAL */

const GlobalErrAuthFailed = "auth:err:auth-failed"
const GlobalErrUnknown = "global:err:failed-unknown"
const GlobalErrNotAllowed = "global:err:not-allowed"
const GlobalErrInvalidFormat = "global:err:invalid-format"

/** AUTH */

const AuthErrToken = "auth:err:err-token"
const AuthErrGetProfile = "auth:err:get-profile"
const AuthErrRefreshToken = "auth:err:get-refresh-token"
const AuthErruserNotActive = "auth:err:user-not-active"
const AuthErrInvalid = "auth:err:invalid-auth"

/** CLIENT */

const ClientErrDuplicate = "client:err:duplicate"
const ClientErrAlreadyExists = "client:err:read-exists"
const ClientErrUpdate = "client:err:update:failed"
const ClientErrGet = "client:err:read-exists"

/** MENU */

const MenuErrNotFound = "menu:err:menu-not-found"
const MenuErrCreate = "menu:err:create"
const MenuErrGet = "menu:err:load"

/** ROLE */

const RoleErrNotExists = "role:err:read-exists"
const RoleErrAlreadyExists = "role:err:read-available"

/** ROLE ASSIGNMENT */

const RoleAssignErrUnknown = "role-assign:err:failed-unknown"
const RoleAssignErrRemovePermit = "role-assign:err:failed-remove-permit"
const RoleAssignErrAlreadyExists = "role-assign:err:exists"
const RoleAssignErrFailed = "role-assign:err:failed"
const RoleAssignErrLoad = "role-assign:err:load"

/** USER */

const UserErrRegister = "user:err:register-failed"
const UserErrActivationNotFound = "user:err:activation-not-found"
const UserErrNotFound = "user:err:user-not-found"
const UserErrAlreadyActive = "user:err:activate-already-active"
const UserErrActivationExpired = "user:err:activation-expired"
const UserErrPassMatch = "user:err:pass-match"
const UserErrCodeAlreadyUsed = "user:err:pass-code-used"
const UserErrNotActive = "user:err:user-not-active"
const UserErrInvited = "user:err:user-invited"
const UserErrCreateActivation = "user:err:create-invitation"
const UserErrActivationFailed = "user:err:activation-failed"

/** ROUTER */

const RouteClientUpdate = "route:client:update"
const RouteRoleAdd = "route:role:add"
const RouteRoleList = "route:role:list"
const RouteRoleUpdate = "route:role:update"
const RouteRoleDelete = "route:role:delete"
const RouteClientRoleAdd = "route:client:role:add"
const RouteClientRoleList = "route:client:role:list"
const RouteClientRoleUpdate = "route:client:role:update"
const RouteClientRoleDelete = "route:client:role:delete"
const RouteClientRoleAssignmentAdd = "route:client:role:assignment:add"
const RouteCLientRoleAssignmentDelete = "route:client:role:assignment:remove"
const RouteClientRoleAssignment = "route:client:role:assignment:assign"
const RouteClientRoleViewAssignment = "route:client:role:assignment:load"
const RouteUserCreate = "route:user:create"
const RouteClientUserInvitation = "route:client:user:invitation"
const RouteMenuCreate = "route:menu:create"
const RouteMenuGetAll = "route:menu:get:all"
const RouteMenuGetAllSa = "route:menu:get:all:sa"
const RouteMenuGetAllCl = "route:menu:get:all:cl"
