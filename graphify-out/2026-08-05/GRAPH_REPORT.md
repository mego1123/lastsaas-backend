# Graph Report - lastsaas  (2026-08-05)

## Corpus Check
- 239 files · ~183,445 words
- Verdict: corpus is large enough that graph structure adds value.

## Summary
- 2591 nodes · 6503 edges · 153 communities (134 shown, 19 thin omitted)
- Extraction: 78% EXTRACTED · 22% INFERRED · 0% AMBIGUOUS · INFERRED: 1411 edges (avg confidence: 0.8)
- Token cost: 0 input · 0 output

## Graph Freshness
- Built from commit: `c692923e`
- Run `git rev-parse HEAD` and compare to check if the graph is stale.
- Run `graphify update .` after code changes (no API cost).

## Community Hubs (Navigation)
- setupTestServer
- Event
- getErrorMessage
- useAuth
- types/index.ts
- Service
- AuthHandler
- oauth_test.go
- App.tsx
- Service
- MongoDB
- Client
- Store
- config_test.go
- HealthPage.tsx
- middleware/tenant_test.go
- PMPage.tsx
- setupIsolationEnv
- AdminHandler
- JWTService
- Validate
- NewTOTPService
- openapi.go
- RateLimiter
- syslog_test.go
- LoadingSpinner.tsx
- respondWithError
- Write
- respondWithJSON
- New
- ui/index.ts
- compilerOptions
- PMHandler
- ResendService
- GetUserFromContext
- cmd_mcp.go
- MustConnectTestDB
- models/health.go
- dependencies
- Service
- devDependencies
- compilerOptions
- clr
- connectDB
- AllSchemas
- Logger
- BillingHandler
- LogHandler
- middleware/auth_test.go
- ParseUserAgent
- Tenant
- testutil.go
- main
- APIPage.tsx
- Features
- process.go
- Emitter
- handlers/auth.go
- BootstrapHandler
- MemberRole
- RequireRole
- manifest.json
- LastSaaS
- BundlesHandler
- GetTenantFromContext
- init
- models_test.go
- cmdUsersGet
- Service
- PlanPage.tsx
- ConfigPage.tsx
- keywords
- cmdLogs
- WebhookEventType
- MCP Examples
- .TrackAuthenticated
- Context
- tokens.go
- Load
- server
- v1.2 — March 1, 2026
- MessageHandler
- BodySizeLimit
- server.json
- UsageHandler
- Service
- Plan
- scripts
- ErrorBoundary
- models/branding.go
- User
- api_key
- Setting Up Stripe Billing
- MCP Server Setup
- cmdDoctor
- admin.go
- counter_test.go
- emitter_test.go
- CheckAndMigrate
- LastSaaS Development Rules
- server_url
- Quick Start
- Seed
- Invitation
- SSOConnection
- package.json
- cmdDBStats
- NewEventDefinitionsHandler
- TestPasswordValidation
- Announcement
- CreditBundle
- EventDefinition
- TestRoleHasPermission
- Message
- SystemConfig
- UsageEvent
- WebAuthnCredential
- Seed
- React + TypeScript + Vite
- glama.json
- platforms
- NewAnnouncementsHandler
- APIVersion
- Recovery
- tsconfig.json
- author
- repository
- Deployment
- eslint
- @eslint/js
- eslint-plugin-react-hooks
- globals
- msw
- @playwright/test
- @testing-library/jest-dom
- @testing-library/user-event
- @types/dompurify
- @vitejs/plugin-react
- vitest
- @vitest/coverage-v8
- setup.sh
- lastsaas

## God Nodes (most connected - your core abstractions)
1. `setupTestServer()` - 152 edges
2. `respondWithError()` - 146 edges
3. `respondWithJSON()` - 144 edges
4. `CreateTestUser()` - 81 edges
5. `createAdminEnv()` - 76 edges
6. `MarkSystemInitialized()` - 64 edges
7. `AuthHandler` - 60 edges
8. `GetUserFromContext()` - 58 edges
9. `CreateTestTenant()` - 56 edges
10. `getErrorMessage()` - 56 edges

## Surprising Connections (you probably didn't know these)
- `main()` --calls--> `parseGlobalFlags()`  [INFERRED]
  backend/cmd/lastsaas/main.go → backend/cmd/lastsaas/output.go
- `main()` --calls--> `cmdDB()`  [INFERRED]
  backend/cmd/lastsaas/main.go → backend/cmd/lastsaas/cmd_db.go
- `cmdDBStats()` --calls--> `connectDB()`  [INFERRED]
  backend/cmd/lastsaas/cmd_db.go → backend/cmd/lastsaas/main.go
- `cmdDBStats()` --calls--> `bold()`  [INFERRED]
  backend/cmd/lastsaas/cmd_db.go → backend/cmd/lastsaas/output.go
- `cmdDBStats()` --calls--> `formatBytes()`  [INFERRED]
  backend/cmd/lastsaas/cmd_db.go → backend/cmd/lastsaas/output.go

## Import Cycles
- None detected.

## Communities (153 total, 19 thin omitted)

### Community 0 - "setupTestServer"
Cohesion: 0.05
Nodes (151): T, TestIntegration_AdminCancelRootInvitation(), TestIntegration_AdminChangeRootMemberRole(), TestIntegration_AdminDashboard(), TestIntegration_AdminGetTenant(), TestIntegration_AdminGetTenantNotFound(), TestIntegration_AdminGetUser(), TestIntegration_AdminGetUserNotFound() (+143 more)

### Community 1 - "Event"
Cohesion: 0.05
Nodes (57): extractInstanceFromEvent(), Context, MongoDB, ObjectID, Request, ResponseWriter, Service, indexOf() (+49 more)

### Community 2 - "getErrorMessage"
Cohesion: 0.07
Nodes (54): adminApi, messagesApi, setTenantHeader(), AdminLayout(), AdminRoute(), ConfirmModal(), ConfirmModalProps, TableSkeleton() (+46 more)

### Community 3 - "useAuth"
Cohesion: 0.06
Nodes (49): dompurify, dompurify, announcementsApi, bundlesApi, setAuthToken(), BrandingThemeInjector(), generatePalette(), isValidHex() (+41 more)

### Community 4 - "types/index.ts"
Cohesion: 0.05
Nodes (55): api, brandingAdminApi, brandingApi, refreshSubscribers, usageApi, BrandingPage(), Tab, ALL_SEVERITIES (+47 more)

### Community 5 - "Service"
Cohesion: 0.08
Nodes (29): ObjectID, Time, buildFunnelSteps(), Context, M, MongoDB, Mutex, ObjectID (+21 more)

### Community 6 - "AuthHandler"
Cohesion: 0.14
Nodes (14): GoogleOAuthService, Context, Duration, MongoDB, ObjectID, Request, ResponseWriter, Service (+6 more)

### Community 7 - "oauth_test.go"
Cohesion: 0.08
Nodes (39): GoogleUserInfo, MicrosoftOAuthService, MicrosoftUserInfo, mockTransport, NewGitHubOAuthService(), Context, Token, NewGoogleOAuthService() (+31 more)

### Community 8 - "App.tsx"
Cohesion: 0.05
Nodes (38): bootstrapApi, tenantApi, AdminAboutPage, AdminAnnouncementsPage, AdminAPIPage, AdminBrandingPage, AdminConfigPage, AdminDashboardPage (+30 more)

### Community 9 - "Service"
Cohesion: 0.07
Nodes (28): Context, Service, IntegrationCheck, MongoDB, MongoMetrics, RWMutex, SystemMetric, WaitGroup (+20 more)

### Community 10 - "MongoDB"
Cohesion: 0.08
Nodes (5): Collection, Context, Database, MongoDB, NewMongoDB()

### Community 11 - "Client"
Cohesion: 0.09
Nodes (22): GitHubEmail, GitHubOAuthService, GitHubUserInfo, Context, Token, Context, IntegrationCheck, LogSeverity (+14 more)

### Community 12 - "Store"
Cohesion: 0.08
Nodes (25): MongoDB, Request, ResponseWriter, NewConfigHandler(), Context, MongoDB, ObjectID, Request (+17 more)

### Community 13 - "config_test.go"
Cohesion: 0.11
Nodes (42): expandEnvVars(), Load(), LoadEnvFile(), findConfigDir(), T, hasYAMLFiles(), setupTestEnv(), TestEnvVarExpansion() (+34 more)

### Community 14 - "HealthPage.tsx"
Cohesion: 0.09
Nodes (34): ChartCard(), ChartCardProps, avg(), CurrentStatusPanel(), CurrentStatusPanelProps, formatBytes(), formatMs(), formatPercent() (+26 more)

### Community 15 - "middleware/tenant_test.go"
Cohesion: 0.10
Nodes (33): GetAPIKeyFromContext(), GetImpersonatedBy(), APIKey, Context, Handler, MongoDB, Request, ResponseWriter (+25 more)

### Community 16 - "PMPage.tsx"
Cohesion: 0.07
Nodes (33): pmApi, Card(), CardProps, paddingClasses, Modal(), ModalProps, EventDefinitionModal(), FormData (+25 more)

### Community 17 - "setupIsolationEnv"
Cohesion: 0.11
Nodes (30): PasswordService, T, User, setupIsolationEnv(), TestIntegration_Admin_CanInviteUsers_NotAdmins(), TestIntegration_NonRootAdmin_CannotAccessDashboard(), TestIntegration_NonRootAdmin_CannotListAPIKeys(), TestIntegration_NonRootAdmin_CannotListLogs() (+22 more)

### Community 18 - "AdminHandler"
Cohesion: 0.18
Nodes (13): decodeJSON(), Context, ObjectID, Request, ResponseWriter, Service, escapeRegexInput(), sanitizeCSVField() (+5 more)

### Community 19 - "JWTService"
Cohesion: 0.11
Nodes (24): AccessTokenClaims, JWTService, RefreshTokenClaims, Duration, T, newTestJWTService(), TestAccessTokenCantValidateAsRefresh(), TestDefaultTTLValues() (+16 more)

### Community 20 - "Validate"
Cohesion: 0.16
Nodes (34): formatFieldError(), T, User, TestValidate_APIKeyInvalidAuthority(), TestValidate_ConfigVarInvalidType(), TestValidate_CreditBundleZeroCredits(), TestValidate_ErrorFormatting(), TestValidate_InvitationInvalidStatus() (+26 more)

### Community 21 - "NewTOTPService"
Cohesion: 0.13
Nodes (24): TOTPService, NewTOTPService(), NewTOTPServiceWithEncryption(), T, TestGenerateRecoveryCodes_Count(), TestGenerateRecoveryCodes_Format(), TestGenerateRecoveryCodes_HashesMatchPlain(), TestGenerateRecoveryCodes_Uniqueness() (+16 more)

### Community 22 - "openapi.go"
Cohesion: 0.11
Nodes (32): apiReference(), authBadge(), authLabel(), DocsHTML(), DocsMarkdown(), Request, ResponseWriter, stripHTML() (+24 more)

### Community 23 - "RateLimiter"
Cohesion: 0.12
Nodes (21): Collection, Database, Duration, HandlerFunc, Request, RWMutex, Time, NewDistributedRateLimiter() (+13 more)

### Community 24 - "syslog_test.go"
Cohesion: 0.18
Nodes (28): detectInjection(), New(), sanitize(), T, TestDetectInjectionCleanString(), TestDetectInjectionEmbed(), TestDetectInjectionIframe(), TestDetectInjectionJavascript() (+20 more)

### Community 25 - "LoadingSpinner.tsx"
Cohesion: 0.12
Nodes (15): authApi, billingApi, plansApi, LoadingSpinner(), LoadingSpinnerProps, BuyCreditsPage(), formatPrice(), InvoiceModal() (+7 more)

### Community 26 - "respondWithError"
Cohesion: 0.17
Nodes (10): Request, ResponseWriter, Request, ResponseWriter, ResponseWriter, respondWithError(), Request, ResponseWriter (+2 more)

### Community 27 - "Write"
Cohesion: 0.15
Nodes (22): Code, Response, Request, ResponseWriter, BadRequest(), Conflict(), Forbidden(), Request (+14 more)

### Community 28 - "respondWithJSON"
Cohesion: 0.20
Nodes (9): defaultBrandingConfig(), MongoDB, Request, ResponseWriter, NewBrandingHandler(), isValidEmail(), respondWithJSON(), BrandingConfig (+1 more)

### Community 29 - "New"
Cohesion: 0.23
Nodes (26): New(), HandlerFunc, Server, T, setupMockStripe(), TestCancelSubscriptionAtPeriodEnd(), TestCancelSubscriptionAtPeriodEndNoItems(), TestCancelSubscriptionImmediately() (+18 more)

### Community 30 - "ui/index.ts"
Cohesion: 0.07
Nodes (19): AlertProps, AlertVariant, variantClasses, Badge(), BadgeProps, BadgeVariant, variantClasses, Button() (+11 more)

### Community 31 - "compilerOptions"
Cohesion: 0.07
Nodes (26): compilerOptions, allowImportingTsExtensions, erasableSyntaxOnly, jsx, lib, module, moduleDetection, moduleResolution (+18 more)

### Community 32 - "PMHandler"
Cohesion: 0.18
Nodes (13): Context, ObjectID, Request, ResponseWriter, MongoDB, Request, ResponseWriter, Service (+5 more)

### Community 33 - "ResendService"
Cohesion: 0.15
Nodes (10): Request, ResponseWriter, Service, Time, NewHealthHandler(), parseTimeRange(), NewResendService(), emailRequest (+2 more)

### Community 34 - "GetUserFromContext"
Cohesion: 0.21
Nodes (13): Context, MongoDB, Plan, Request, ResponseWriter, Service, NewPlansHandler(), validatePlanRequest() (+5 more)

### Community 35 - "cmd_mcp.go"
Cohesion: 0.38
Nodes (22): buildQuery(), cmdMCP(), Context, newMCPClient(), prettyJSON(), registerAboutTools(), registerAnnouncementTools(), registerConfigTools() (+14 more)

### Community 36 - "MustConnectTestDB"
Cohesion: 0.26
Nodes (23): TestRequireTenantAlreadyInContext(), TestRequireTenantIntegration(), TestRequireTenantInvalidID(), TestRequireTenantMissingHeader(), TestRequireTenantNotAMember(), NewTenantMiddleware(), CleanupCollections(), CreateTestWebhook() (+15 more)

### Community 37 - "models/health.go"
Cohesion: 0.10
Nodes (23): MongoMetrics, ObjectID, Time, CPUMetrics, DiskMetrics, GoRuntimeMetrics, HTTPMetrics, IntegrationCountMetrics (+15 more)

### Community 38 - "dependencies"
Cohesion: 0.09
Nodes (23): axios, dependencies, axios, @hookform/resolvers, lucide-react, react, react-dom, react-hook-form (+15 more)

### Community 39 - "Service"
Cohesion: 0.15
Nodes (9): Context, MongoDB, ObjectID, Time, CheckoutSession, CheckoutLineItem, CheckoutRequest, Service (+1 more)

### Community 40 - "devDependencies"
Cohesion: 0.09
Nodes (23): eslint-plugin-react-refresh, devDependencies, eslint-plugin-react-refresh, jsdom, tailwindcss, @tailwindcss/vite, @testing-library/react, @types/node (+15 more)

### Community 41 - "compilerOptions"
Cohesion: 0.09
Nodes (22): compilerOptions, allowImportingTsExtensions, erasableSyntaxOnly, lib, module, moduleDetection, moduleResolution, noEmit (+14 more)

### Community 42 - "clr"
Cohesion: 0.21
Nodes (18): cmdFinancial(), cmdFinancialMetrics(), cmdFinancialSummary(), cmdFinancialTransactions(), cmdHealth(), printLogEntry(), cmdStats(), bold() (+10 more)

### Community 43 - "connectDB"
Cohesion: 0.21
Nodes (21): cmdChangePassword(), cmdConfig(), cmdConfigGet(), cmdConfigList(), cmdConfigReset(), cmdConfigSet(), cmdSendMessage(), cmdSetup() (+13 more)

### Community 44 - "AllSchemas"
Cohesion: 0.23
Nodes (20): AllSchemas(), announcementsSchema(), apiKeysSchema(), configVarsSchema(), creditBundlesSchema(), customPagesSchema(), eventDefinitionsSchema(), financialTransactionsSchema() (+12 more)

### Community 45 - "Logger"
Cohesion: 0.26
Nodes (7): Context, LogCategory, LogSeverity, MongoDB, ObjectID, SystemLog, Logger

### Community 46 - "BillingHandler"
Cohesion: 0.29
Nodes (5): Context, Request, ResponseWriter, Service, BillingHandler

### Community 47 - "LogHandler"
Cohesion: 0.16
Nodes (14): getFirst(), M, MongoDB, Request, ResponseWriter, SystemLog, NewLogHandler(), ObjectID (+6 more)

### Community 48 - "middleware/auth_test.go"
Cohesion: 0.28
Nodes (18): NewJWTService(), NewAuthMiddleware(), T, setupAuthMiddleware(), TestGetClientIPFlyClientIP(), TestGetClientIPFlyClientIPInvalid(), TestRateLimiterCleanupExpired(), TestRequireAuthAdminAPIKey() (+10 more)

### Community 49 - "ParseUserAgent"
Cohesion: 0.26
Nodes (17): parseBrowser(), parseOS(), ParseUserAgent(), T, TestParseUserAgentAndroid(), TestParseUserAgentBrowserOnly(), TestParseUserAgentChrome(), TestParseUserAgentChromeOS() (+9 more)

### Community 50 - "Tenant"
Cohesion: 0.20
Nodes (16): cmdTenants(), cmdTenantsGet(), cmdTenantsList(), countMembersPerTenant(), Context, MongoDB, ObjectID, resolvePlanNames() (+8 more)

### Community 51 - "testutil.go"
Cohesion: 0.20
Nodes (16): APIKeyAuthority, M, TestMain(), ConnectTestDB(), CreateTestAPIKey(), CreateTestInvitation(), findAndSetConfigDir(), APIKey (+8 more)

### Community 52 - "main"
Cohesion: 0.28
Nodes (15): main(), MongoDB, NewWebhooksHandler(), Context, NewDataDogChecker(), NewGitHubOAuthChecker(), NewGoogleOAuthChecker(), NewMicrosoftOAuthChecker() (+7 more)

### Community 53 - "APIPage.tsx"
Cohesion: 0.17
Nodes (12): APIKeysSection(), APIPage(), CreateKeyModal(), formatDate(), timeAgo(), WebhookDetailModal(), WebhookFormModal(), WebhooksSection() (+4 more)

### Community 54 - "Features"
Cohesion: 0.12
Nodes (16): Admin Interface, API Keys, Authentication & Identity, Billing & Credits (Stripe), Built-in API Documentation, CLI Administration, Features, MCP Server (AI Admin Access) (+8 more)

### Community 55 - "process.go"
Cohesion: 0.35
Nodes (14): capitalizeStr(), cmdRestart(), cmdStart(), cmdStop(), ensurePIDDir(), findProjectRoot(), isDirAt(), mustFindProjectRoot() (+6 more)

### Community 56 - "Emitter"
Cohesion: 0.19
Nodes (11): MongoDB, NewAdminHandler(), MongoDB, NewAPIKeysHandler(), MongoDB, NewBillingHandler(), MongoDB, NewTenantHandler() (+3 more)

### Community 57 - "handlers/auth.go"
Cohesion: 0.14
Nodes (13): User, AcceptInvitationRequest, AuthResponse, ChangePasswordRequest, ForgotPasswordRequest, LoginRequest, MFARequiredResponse, RefreshRequest (+5 more)

### Community 58 - "BootstrapHandler"
Cohesion: 0.23
Nodes (8): Handler, MongoDB, Request, ResponseWriter, RWMutex, NewBootstrapHandler(), BootstrapHandler, bootstrapStatusResponse

### Community 59 - "MemberRole"
Cohesion: 0.19
Nodes (11): Time, ObjectID, Time, RoleHasPermission(), ChangeRoleRequest, InviteMemberRequest, MemberResponse, MembershipInfo (+3 more)

### Community 60 - "RequireRole"
Cohesion: 0.22
Nodes (11): T, TestContextHelpers(), TestRequireRole(), TestRequireRoleMissingContext(), TestRequireRootTenant(), TestSecurityHeaders(), Handler, RequireRole() (+3 more)

### Community 61 - "manifest.json"
Cohesion: 0.14
Nodes (13): description, display_name, documentation, homepage, license, long_description, manifest_version, name (+5 more)

### Community 62 - "LastSaaS"
Cohesion: 0.14
Nodes (13): API Documentation, Configuration, Environment Variables, Fork It and Keep Building with AI, How It Compares, LastSaaS, License, Optional (+5 more)

### Community 63 - "BundlesHandler"
Cohesion: 0.32
Nodes (7): MongoDB, Request, ResponseWriter, NewBundlesHandler(), validateBundleRequest(), bundleRequest, BundlesHandler

### Community 64 - "GetTenantFromContext"
Cohesion: 0.40
Nodes (5): Request, ResponseWriter, Service, GetTenantFromContext(), TenantHandler

### Community 65 - "init"
Cohesion: 0.23
Nodes (11): ObjectID, Time, ValidAPIKeyAuthority(), ObjectID, Time, ValidConfigVarType(), init(), APIKey (+3 more)

### Community 66 - "models_test.go"
Cohesion: 0.28
Nodes (12): T, TestAllWebhookEventTypesNoDuplicates(), TestAuthMethodConstants(), TestBillingStatusConstants(), TestUserHasAuthMethod(), TestUserHasAuthMethodEmpty(), TestUserIsLockedFuture(), TestUserIsLockedNil() (+4 more)

### Community 67 - "cmdUsersGet"
Cohesion: 0.29
Nodes (11): cmdUsers(), cmdUsersGet(), cmdUsersList(), cmdUsersRevokeSessions(), cmdUsersSetActive(), Context, MongoDB, ObjectID (+3 more)

### Community 68 - "Service"
Cohesion: 0.29
Nodes (4): Context, MongoDB, New(), Service

### Community 69 - "PlanPage.tsx"
Cohesion: 0.29
Nodes (10): telemetryApi, getSessionId(), useTelemetry(), annualPrice(), annualTotal(), currencySymbols, formatPrice(), getCurrencySymbol() (+2 more)

### Community 70 - "ConfigPage.tsx"
Cohesion: 0.21
Nodes (9): ConfigPage(), CreateConfigModal(), EditConfigModal(), parseEnumOptions(), serializeEnumOptions(), typeLabels, ConfigVar, ConfigVarType (+1 more)

### Community 71 - "keywords"
Cohesion: 0.17
Nodes (12): keywords, admin, ai-native, billing, dashboard, go, health, monitoring (+4 more)

### Community 72 - "cmdLogs"
Cohesion: 0.35
Nodes (10): buildLogFilter(), cmdLogs(), Context, M, MongoDB, SystemLog, Time, logsFollow() (+2 more)

### Community 73 - "WebhookEventType"
Cohesion: 0.33
Nodes (9): validateWebhookRequest(), validateWebhookURL(), ObjectID, Time, ValidWebhookEventType(), webhookRequest, Webhook, WebhookDelivery (+1 more)

### Community 74 - "MCP Examples"
Cohesion: 0.18
Nodes (11): Example 10: Health Metrics Deep Dive, Example 1: Check Business Overview, Example 2: Investigate Revenue Trend, Example 3: Find a Specific Tenant, Example 4: Search for Critical Errors, Example 5: Monitor System Health, Example 6: Audit API Key Usage, Example 7: Review Subscription Plans (+3 more)

### Community 75 - ".TrackAuthenticated"
Cohesion: 0.40
Nodes (6): Request, ResponseWriter, Service, NewTelemetryHandler(), sanitizeProperties(), TelemetryHandler

### Community 76 - "Context"
Cohesion: 0.38
Nodes (5): Context, Service, SystemMetric, Time, SystemNode

### Community 77 - "tokens.go"
Cohesion: 0.42
Nodes (9): ObjectID, Time, AuthCode, AuthCodeTokenData, OAuthState, RefreshToken, RevokedToken, TokenType (+1 more)

### Community 78 - "Load"
Cohesion: 0.42
Nodes (8): Load(), T, TestBuildVersionTakesPrecedence(), TestLoadFallbackToUnknown(), TestLoadFromBuildVersion(), TestLoadFromVersionFile(), TestLoadTrimsWhitespace(), TestLoadWalksUpDirectories()

### Community 79 - "server"
Cohesion: 0.20
Nodes (10): LASTSAAS_API_KEY, LASTSAAS_URL, args, command, env, server, entry_point, mcp_config (+2 more)

### Community 80 - "v1.2 — March 1, 2026"
Cohesion: 0.20
Nodes (9): CI/CD & Testing (New), Infrastructure & Quality, Initial Public Release, LastSaaS Version Notes, MCP Server Improvements, Product Analytics & Telemetry (New), Security Hardening, v1.0 — February 25, 2026 (+1 more)

### Community 81 - "MessageHandler"
Cohesion: 0.39
Nodes (5): MongoDB, Request, ResponseWriter, NewMessageHandler(), MessageHandler

### Community 82 - "BodySizeLimit"
Cohesion: 0.36
Nodes (7): BodySizeLimit(), Handler, T, TestBodySizeLimitAllowsSmallBody(), TestBodySizeLimitBlocksOversizedBody(), TestBodySizeLimitNilBody(), TestMaxBodySizeConstant()

### Community 83 - "server.json"
Cohesion: 0.22
Nodes (8): description, name, packages, repository, source, url, $schema, version

### Community 84 - "UsageHandler"
Cohesion: 0.39
Nodes (5): MongoDB, Request, ResponseWriter, NewUsageHandler(), UsageHandler

### Community 86 - "Plan"
Cohesion: 0.39
Nodes (7): ObjectID, Time, CreditResetPolicy, EntitlementType, EntitlementValue, Plan, PricingModel

### Community 87 - "scripts"
Cohesion: 0.25
Nodes (8): scripts, build, dev, lint, preview, test, test:coverage, test:watch

### Community 88 - "ErrorBoundary"
Cohesion: 0.25
Nodes (3): ErrorBoundary, Props, State

### Community 89 - "models/branding.go"
Cohesion: 0.52
Nodes (6): ObjectID, Time, BrandingAsset, BrandingConfig, CustomPage, NavItem

### Community 90 - "User"
Cohesion: 0.38
Nodes (4): ObjectID, Time, AuthMethod, User

### Community 91 - "api_key"
Cohesion: 0.29
Nodes (7): description, required, sensitive, title, type, user_config, api_key

### Community 92 - "Setting Up Stripe Billing"
Cohesion: 0.29
Nodes (7): 1. Create a Stripe account, 2. Get your API keys, 3. Create a webhook endpoint, 4. Set environment variables, 5. Create plans in the admin UI, 6. Go live, Setting Up Stripe Billing

### Community 93 - "MCP Server Setup"
Cohesion: 0.29
Nodes (7): Available Tools (26), Build the CLI binary, Environment Variables, MCP Server Setup, Prerequisites, Usage with Claude Code, Usage with Claude Desktop

### Community 94 - "cmdDoctor"
Cohesion: 0.40
Nodes (5): checkConfigIntegration(), cmdDoctor(), Context, MongoDB, GetEnv()

### Community 95 - "admin.go"
Cohesion: 0.53
Nodes (5): Time, TenantListItem, UserDetail, UserListItem, UserMembershipDetail

### Community 96 - "counter_test.go"
Cohesion: 0.53
Nodes (5): T, TestConcurrentIncrements(), TestResendEmailsIncrement(), TestStripeAPICallsIncrement(), TestSwapResets()

### Community 97 - "emitter_test.go"
Cohesion: 0.53
Nodes (5): T, TestEventStruct(), TestEventTypeConstants(), TestNoopEmitterEmit(), TestNoopEmitterImplementsInterface()

### Community 98 - "CheckAndMigrate"
Cohesion: 0.60
Nodes (5): CheckAndMigrate(), Context, MongoDB, runMigrations(), sendUpgradeMessage()

### Community 99 - "LastSaaS Development Rules"
Cohesion: 0.33
Nodes (5): Build Verification, Dependent Project Deployment (CRITICAL), LastSaaS Development Rules, System Logging, Validation

### Community 100 - "server_url"
Cohesion: 0.33
Nodes (6): default, description, required, title, type, server_url

### Community 101 - "Quick Start"
Cohesion: 0.33
Nodes (6): 1. Clone the repository, 2. Run the setup script, 3. Start the backend, 4. Start the frontend, 5. Initialize the system, Quick Start

### Community 102 - "Seed"
Cohesion: 0.40
Nodes (3): Context, MongoDB, Seed()

### Community 103 - "Invitation"
Cohesion: 0.50
Nodes (4): ObjectID, Time, Invitation, InvitationStatus

### Community 104 - "SSOConnection"
Cohesion: 0.50
Nodes (4): ObjectID, Time, SSOAttributeMap, SSOConnection

### Community 105 - "package.json"
Cohesion: 0.40
Nodes (4): name, private, type, version

### Community 106 - "cmdDBStats"
Cohesion: 0.83
Nodes (3): cmdDB(), cmdDBStats(), toInt64()

### Community 107 - "NewEventDefinitionsHandler"
Cohesion: 0.50
Nodes (3): MongoDB, NewEventDefinitionsHandler(), eventDefRequest

### Community 108 - "TestPasswordValidation"
Cohesion: 0.67
Nodes (3): T, TestPasswordHashing(), TestPasswordValidation()

### Community 109 - "Announcement"
Cohesion: 0.50
Nodes (3): ObjectID, Time, Announcement

### Community 110 - "CreditBundle"
Cohesion: 0.50
Nodes (3): ObjectID, Time, CreditBundle

### Community 111 - "EventDefinition"
Cohesion: 0.50
Nodes (3): ObjectID, Time, EventDefinition

### Community 112 - "TestRoleHasPermission"
Cohesion: 0.67
Nodes (3): T, TestRoleHasPermission(), TestValidRole()

### Community 113 - "Message"
Cohesion: 0.50
Nodes (3): ObjectID, Time, Message

### Community 114 - "SystemConfig"
Cohesion: 0.50
Nodes (3): ObjectID, Time, SystemConfig

### Community 115 - "UsageEvent"
Cohesion: 0.50
Nodes (3): ObjectID, Time, UsageEvent

### Community 116 - "WebAuthnCredential"
Cohesion: 0.50
Nodes (3): ObjectID, Time, WebAuthnCredential

### Community 117 - "Seed"
Cohesion: 0.50
Nodes (3): Context, MongoDB, Seed()

### Community 118 - "React + TypeScript + Vite"
Cohesion: 0.50
Nodes (3): Expanding the ESLint configuration, React Compiler, React + TypeScript + Vite

### Community 119 - "glama.json"
Cohesion: 0.50
Nodes (3): maintainers, $schema, jonradoff

### Community 120 - "platforms"
Cohesion: 0.50
Nodes (4): compatibility, platforms, darwin, linux

### Community 125 - "author"
Cohesion: 0.67
Nodes (3): author, name, url

### Community 126 - "repository"
Cohesion: 0.67
Nodes (3): repository, type, url

### Community 127 - "Deployment"
Cohesion: 0.67
Nodes (3): Deployment, Fly.io, Other Platforms

## Knowledge Gaps
- **350 isolated node(s):** `lastsaas`, `RegisterRequest`, `LoginRequest`, `RefreshRequest`, `MFARequiredResponse` (+345 more)
  These have ≤1 connection - possible missing edges or undocumented components.
- **19 thin communities (<3 nodes) omitted from report** — run `graphify query` to explore isolated nodes.

## Suggested Questions
_Questions this graph is uniquely positioned to answer:_

- **Why does `main()` connect `main` to `Event`, `AuthHandler`, `oauth_test.go`, `Service`, `MongoDB`, `Store`, `config_test.go`, `middleware/tenant_test.go`, `RateLimiter`, `Write`, `respondWithJSON`, `PMHandler`, `ResendService`, `GetUserFromContext`, `MustConnectTestDB`, `connectDB`, `LogHandler`, `middleware/auth_test.go`, `Emitter`, `BootstrapHandler`, `RequireRole`, `BundlesHandler`, `.TrackAuthenticated`, `MessageHandler`, `BodySizeLimit`, `UsageHandler`, `cmdDoctor`, `CheckAndMigrate`, `NewEventDefinitionsHandler`, `NewAnnouncementsHandler`, `Recovery`?**
  _High betweenness centrality (0.123) - this node is a cross-community bridge._
- **Why does `setupTestServer()` connect `setupTestServer` to `GetUserFromContext`, `MustConnectTestDB`, `AuthHandler`, `Client`, `LogHandler`, `middleware/auth_test.go`, `setupIsolationEnv`, `testutil.go`, `main`, `Emitter`, `BootstrapHandler`, `RequireRole`?**
  _High betweenness centrality (0.100) - this node is a cross-community bridge._
- **Why does `Client` connect `Client` to `setupTestServer`, `ResendService`, `Event`, `cmd_mcp.go`, `MongoDB`, `setupIsolationEnv`, `main`?**
  _High betweenness centrality (0.068) - this node is a cross-community bridge._
- **Are the 148 inferred relationships involving `setupTestServer()` (e.g. with `TestIntegration_AdminCancelRootInvitation()` and `TestIntegration_AdminChangeRootMemberRole()`) actually correct?**
  _`setupTestServer()` has 148 INFERRED edges - model-reasoned connections that need verification._
- **Are the 143 inferred relationships involving `respondWithError()` (e.g. with `.CancelRootInvitation()` and `.ChangeRootMemberRole()`) actually correct?**
  _`respondWithError()` has 143 INFERRED edges - model-reasoned connections that need verification._
- **Are the 141 inferred relationships involving `respondWithJSON()` (e.g. with `.CancelRootInvitation()` and `.ChangeRootMemberRole()`) actually correct?**
  _`respondWithJSON()` has 141 INFERRED edges - model-reasoned connections that need verification._
- **Are the 77 inferred relationships involving `CreateTestUser()` (e.g. with `TestIntegration_AdminChangeRootMemberRole()` and `TestIntegration_AdminGetUser()`) actually correct?**
  _`CreateTestUser()` has 77 INFERRED edges - model-reasoned connections that need verification._