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
