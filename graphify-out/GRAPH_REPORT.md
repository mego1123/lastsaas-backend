# Graph Report - lastsaas  (2026-08-05)

## Corpus Check
- 239 files · ~187,751 words
- Verdict: corpus is large enough that graph structure adds value.

## Summary
- 2656 nodes · 6348 edges · 182 communities (145 shown, 37 thin omitted)
- Extraction: 80% EXTRACTED · 20% INFERRED · 0% AMBIGUOUS · INFERRED: 1256 edges (avg confidence: 0.8)
- Token cost: 0 input · 0 output

## Graph Freshness
- Built from commit: `50f1d1bb`
- Run `git rev-parse HEAD` and compare to check if the graph is stale.
- Run `graphify update .` after code changes (no API cost).

## Community Hubs (Navigation)
- CreateTestUser
- WebhookHandler
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
- PromotionsHandler
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
- MustConnectTestDB
- LoadingSpinner.tsx
- WebhooksHandler
- Write
- respondWithJSON
- NewMetricsCollector
- ui/index.ts
- compilerOptions
- PMHandler
- HealthHandler
- GetUserFromContext
- cmd_mcp.go
- Service
- models/health.go
- dependencies
- EventDefinitionsHandler
- devDependencies
- compilerOptions
- printJSON
- connectDB
- AllSchemas
- User
- BillingHandler
- LogHandler
- crypto_test.go
- ParseUserAgent
- cmdTenantsList
- Store
- main
- APIPage.tsx
- Features
- process.go
- emitter.go
- createAdminEnv
- BootstrapHandler
- admin.go
- ConfigHandler
- manifest.json
- LastSaaS
- BundlesHandler
- respondWithError
- init
- models_test.go
- cmdUsersGet
- Service
- PlansPage.tsx
- ReadResponseBody
- keywords
- clr
- LogCategory
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
- AnnouncementsHandler
- setupTestServer
- useTenant
- scripts
- ErrorBoundary
- models/branding.go
- User
- api_key
- Setting Up Stripe Billing
- MCP Server Setup
- cmdDoctor
- plans_test.go
- counter_test.go
- models/billing.go
- testutil.go
- LastSaaS Development Rules
- server_url
- Quick Start
- Seed
- SystemLog
- SSOConnection
- package.json
- InsertTestLogs
- Plan
- WebhookEventType
- Announcement
- CreditBundle
- EventDefinition
- MemberRole
- Message
- SystemConfig
- UsageEvent
- WebAuthnCredential
- Seed
- React + TypeScript + Vite
- glama.json
- platforms
- EventDefinitionModal.tsx
- APIVersion
- Recovery
- tsconfig.json
- author
- repository
- Deployment
- eslint
- configstore/validate.go
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
- emitter_test.go
- Invitation
- Tenant
- TestPasswordValidation
- TelemetryEvent
- stripe.go
- @eslint/js
- Reader
- MongoMetrics
- Collection
- Database
- Mutex
- FeatureUse
- FunnelStep
- KPIData
- PlanShare
- Button.tsx
- ValidConfigVarType
- Alert.tsx
- Badge.tsx
- Context
- Duration
- MongoDB
- ObjectID
- Request
- ResponseWriter
- Service
- Time
- MembershipInfo

## God Nodes (most connected - your core abstractions)
1. `setupTestServer()` - 152 edges
2. `respondWithJSON()` - 118 edges
3. `respondWithError()` - 115 edges
4. `CreateTestUser()` - 81 edges
5. `createAdminEnv()` - 76 edges
6. `MarkSystemInitialized()` - 64 edges
7. `AuthHandler` - 60 edges
8. `CreateTestTenant()` - 56 edges
9. `getErrorMessage()` - 56 edges
10. `MongoDB` - 44 edges

## Surprising Connections (you probably didn't know these)
- `main()` --calls--> `parseGlobalFlags()`  [INFERRED]
  backend/cmd/lastsaas/main.go → backend/cmd/lastsaas/output.go
- `init()` --calls--> `BillingStatus`  [INFERRED]
  backend/internal/validation/validate.go → backend/internal/models/billing.go
- `setupTestServer()` --calls--> `NewAuthHandler()`  [INFERRED]
  backend/internal/api/handlers/testhelpers_test.go → backend/internal/api/handlers/auth.go
- `main()` --calls--> `cmdDB()`  [INFERRED]
  backend/cmd/lastsaas/main.go → backend/cmd/lastsaas/cmd_db.go
- `cmdDBStats()` --calls--> `connectDB()`  [INFERRED]
  backend/cmd/lastsaas/cmd_db.go → backend/cmd/lastsaas/main.go

## Import Cycles
- None detected.

## Communities (182 total, 37 thin omitted)

### Community 0 - "CreateTestUser"
Cohesion: 0.17
Nodes (48): T, TestIntegration_AdminBilling_NonRootTenantForbidden(), TestIntegration_AdminGetMetrics_Returns(), TestIntegration_AdminListTransactions_Empty(), TestIntegration_Billing_NonRootTenantCanAccess(), TestIntegration_Billing_UnauthenticatedForbidden(), TestIntegration_BillingConfig_NilStripe_ReturnsEmptyKey(), TestIntegration_CancelSubscription_NoSubscription() (+40 more)

### Community 1 - "WebhookHandler"
Cohesion: 0.17
Nodes (16): extractInstanceFromEvent(), Context, Emitter, Event, Logger, MongoDB, ObjectID, Request (+8 more)

### Community 2 - "getErrorMessage"
Cohesion: 0.08
Nodes (35): adminApi, TableSkeleton(), TableSkeletonProps, AboutPage(), AnnouncementFormData, AnnouncementFormModal(), announcementSchema, AnnouncementsPage() (+27 more)

### Community 3 - "useAuth"
Cohesion: 0.07
Nodes (39): dompurify, dompurify, announcementsApi, setAuthToken(), BrandingThemeInjector(), generatePalette(), isValidHex(), ImpersonationBanner() (+31 more)

### Community 4 - "types/index.ts"
Cohesion: 0.04
Nodes (52): api, brandingAdminApi, brandingApi, refreshSubscribers, usageApi, BrandingContextValue, BrandingPage(), Tab (+44 more)

### Community 5 - "Service"
Cohesion: 0.11
Nodes (23): buildFunnelSteps(), Context, M, MongoDB, ObjectID, TelemetryEvent, Time, mergeBson() (+15 more)

### Community 6 - "AuthHandler"
Cohesion: 0.07
Nodes (37): AuthCodeTokenData, hashToken(), NewAuthHandler(), storeRefreshToken(), Context, Duration, Emitter, GitHubOAuthService (+29 more)

### Community 7 - "oauth_test.go"
Cohesion: 0.08
Nodes (40): GoogleOAuthService, GoogleUserInfo, MicrosoftOAuthService, MicrosoftUserInfo, mockTransport, NewGitHubOAuthService(), Context, Token (+32 more)

### Community 8 - "App.tsx"
Cohesion: 0.05
Nodes (40): bootstrapApi, tenantApi, AdminAboutPage, AdminAnnouncementsPage, AdminAPIPage, AdminBrandingPage, AdminConfigPage, AdminDashboardPage (+32 more)

### Community 9 - "Service"
Cohesion: 0.13
Nodes (13): Context, Service, IntegrationCheck, MongoDB, RWMutex, SystemMetric, WaitGroup, hostname() (+5 more)

### Community 10 - "MongoDB"
Cohesion: 0.08
Nodes (5): Collection, Context, Database, MongoDB, NewMongoDB()

### Community 11 - "Client"
Cohesion: 0.09
Nodes (23): GitHubEmail, GitHubOAuthService, GitHubUserInfo, Context, Token, Context, IntegrationCheck, LogSeverity (+15 more)

### Community 12 - "PromotionsHandler"
Cohesion: 0.25
Nodes (9): Context, MongoDB, ObjectID, Request, ResponseWriter, Service, Store, NewPromotionsHandler() (+1 more)

### Community 13 - "config_test.go"
Cohesion: 0.11
Nodes (42): expandEnvVars(), Load(), LoadEnvFile(), findConfigDir(), T, hasYAMLFiles(), setupTestEnv(), TestEnvVarExpansion() (+34 more)

### Community 14 - "HealthPage.tsx"
Cohesion: 0.09
Nodes (34): ChartCard(), ChartCardProps, avg(), CurrentStatusPanel(), CurrentStatusPanelProps, formatBytes(), formatMs(), formatPercent() (+26 more)

### Community 15 - "middleware/tenant_test.go"
Cohesion: 0.06
Nodes (51): MongoDB, Request, ResponseWriter, NewUsageHandler(), GetAPIKeyFromContext(), GetImpersonatedBy(), APIKey, Context (+43 more)

### Community 16 - "PMPage.tsx"
Cohesion: 0.10
Nodes (23): pmApi, binChartData(), EngagementTab(), EventSubTab, FlowSubTab(), formatCents(), formatNum(), formatPct() (+15 more)

### Community 17 - "setupIsolationEnv"
Cohesion: 0.11
Nodes (30): PasswordService, T, User, setupIsolationEnv(), TestIntegration_Admin_CanInviteUsers_NotAdmins(), TestIntegration_NonRootAdmin_CannotAccessDashboard(), TestIntegration_NonRootAdmin_CannotListAPIKeys(), TestIntegration_NonRootAdmin_CannotListLogs() (+22 more)

### Community 18 - "AdminHandler"
Cohesion: 0.17
Nodes (14): decodeJSON(), Context, Emitter, Logger, MongoDB, ObjectID, Request, ResponseWriter (+6 more)

### Community 19 - "JWTService"
Cohesion: 0.11
Nodes (24): AccessTokenClaims, JWTService, RefreshTokenClaims, Duration, T, newTestJWTService(), TestAccessTokenCantValidateAsRefresh(), TestDefaultTTLValues() (+16 more)

### Community 20 - "Validate"
Cohesion: 0.16
Nodes (34): formatFieldError(), T, User, TestValidate_APIKeyInvalidAuthority(), TestValidate_ConfigVarInvalidType(), TestValidate_CreditBundleZeroCredits(), TestValidate_ErrorFormatting(), TestValidate_InvitationInvalidStatus() (+26 more)

### Community 21 - "NewTOTPService"
Cohesion: 0.14
Nodes (24): TOTPService, NewTOTPService(), NewTOTPServiceWithEncryption(), T, TestGenerateRecoveryCodes_Count(), TestGenerateRecoveryCodes_Format(), TestGenerateRecoveryCodes_HashesMatchPlain(), TestGenerateRecoveryCodes_Uniqueness() (+16 more)

### Community 22 - "openapi.go"
Cohesion: 0.11
Nodes (32): apiReference(), authBadge(), authLabel(), DocsHTML(), DocsMarkdown(), Request, ResponseWriter, stripHTML() (+24 more)

### Community 23 - "RateLimiter"
Cohesion: 0.13
Nodes (21): Duration, HandlerFunc, Request, RWMutex, Time, NewDistributedRateLimiter(), NewRateLimiter(), T (+13 more)

### Community 24 - "MustConnectTestDB"
Cohesion: 0.06
Nodes (101): NewJWTService(), NewAuthMiddleware(), T, setupAuthMiddleware(), TestGetClientIPFlyClientIP(), TestGetClientIPFlyClientIPInvalid(), TestRateLimiterCleanupExpired(), TestRequireAuthAdminAPIKey() (+93 more)

### Community 25 - "LoadingSpinner.tsx"
Cohesion: 0.13
Nodes (19): authApi, billingApi, bundlesApi, ConfirmModal(), ConfirmModalProps, LoadingSpinner(), LoadingSpinnerProps, BuyCreditsPage() (+11 more)

### Community 26 - "WebhooksHandler"
Cohesion: 0.22
Nodes (10): generateRandomToken(), Logger, MongoDB, Request, ResponseWriter, NewWebhooksHandler(), validateWebhookRequest(), validateWebhookURL() (+2 more)

### Community 27 - "Write"
Cohesion: 0.15
Nodes (22): Code, Response, Request, ResponseWriter, BadRequest(), Conflict(), Forbidden(), Request (+14 more)

### Community 28 - "respondWithJSON"
Cohesion: 0.21
Nodes (10): defaultBrandingConfig(), Logger, MongoDB, Request, ResponseWriter, Store, NewBrandingHandler(), respondWithJSON() (+2 more)

### Community 29 - "NewMetricsCollector"
Cohesion: 0.15
Nodes (16): Handler, Mutex, ResponseWriter, NewMetricsCollector(), percentile(), T, TestMetricsCollectorMiddleware(), TestMetricsCollectorMiddlewareDefaultStatus() (+8 more)

### Community 30 - "ui/index.ts"
Cohesion: 0.14
Nodes (9): Card(), CardProps, paddingClasses, Input, InputProps, Select, SelectProps, Textarea (+1 more)

### Community 31 - "compilerOptions"
Cohesion: 0.07
Nodes (26): compilerOptions, allowImportingTsExtensions, erasableSyntaxOnly, jsx, lib, module, moduleDetection, moduleResolution (+18 more)

### Community 32 - "PMHandler"
Cohesion: 0.30
Nodes (8): MongoDB, Request, ResponseWriter, Service, Time, NewPMHandler(), parsePMTimeRange(), PMHandler

### Community 33 - "HealthHandler"
Cohesion: 0.20
Nodes (9): Request, ResponseWriter, Service, Time, NewHealthHandler(), parseTimeRange(), isValidEmail(), ErrorResponse (+1 more)

### Community 34 - "GetUserFromContext"
Cohesion: 0.18
Nodes (16): Context, Logger, MongoDB, Plan, Request, ResponseWriter, Service, Store (+8 more)

### Community 35 - "cmd_mcp.go"
Cohesion: 0.38
Nodes (22): buildQuery(), cmdMCP(), Context, newMCPClient(), prettyJSON(), registerAboutTools(), registerAnnouncementTools(), registerConfigTools() (+14 more)

### Community 36 - "Service"
Cohesion: 0.19
Nodes (6): Context, MongoDB, Time, CheckoutSession, Service, Subscription

### Community 37 - "models/health.go"
Cohesion: 0.10
Nodes (23): MongoMetrics, ObjectID, Time, CPUMetrics, DiskMetrics, GoRuntimeMetrics, HTTPMetrics, IntegrationCountMetrics (+15 more)

### Community 38 - "dependencies"
Cohesion: 0.09
Nodes (23): axios, dependencies, axios, @hookform/resolvers, lucide-react, react, react-dom, react-hook-form (+15 more)

### Community 39 - "EventDefinitionsHandler"
Cohesion: 0.27
Nodes (9): Context, Logger, MongoDB, ObjectID, Request, ResponseWriter, NewEventDefinitionsHandler(), EventDefinitionsHandler (+1 more)

### Community 40 - "devDependencies"
Cohesion: 0.09
Nodes (23): eslint-plugin-react-refresh, devDependencies, eslint-plugin-react-refresh, jsdom, tailwindcss, @tailwindcss/vite, @testing-library/react, @types/node (+15 more)

### Community 41 - "compilerOptions"
Cohesion: 0.09
Nodes (22): compilerOptions, allowImportingTsExtensions, erasableSyntaxOnly, lib, module, moduleDetection, moduleResolution, noEmit (+14 more)

### Community 42 - "printJSON"
Cohesion: 0.24
Nodes (13): cmdDB(), cmdDBStats(), toInt64(), cmdFinancial(), cmdFinancialMetrics(), cmdFinancialSummary(), cmdFinancialTransactions(), cmdHealth() (+5 more)

### Community 43 - "connectDB"
Cohesion: 0.20
Nodes (22): cmdChangePassword(), cmdConfig(), cmdConfigGet(), cmdConfigList(), cmdConfigReset(), cmdConfigSet(), cmdSendMessage(), cmdSetup() (+14 more)

### Community 44 - "AllSchemas"
Cohesion: 0.23
Nodes (20): AllSchemas(), announcementsSchema(), apiKeysSchema(), configVarsSchema(), creditBundlesSchema(), customPagesSchema(), eventDefinitionsSchema(), financialTransactionsSchema() (+12 more)

### Community 46 - "BillingHandler"
Cohesion: 0.21
Nodes (10): Context, Emitter, Logger, MongoDB, Request, ResponseWriter, Service, Store (+2 more)

### Community 47 - "LogHandler"
Cohesion: 0.25
Nodes (9): getFirst(), M, MongoDB, Request, ResponseWriter, SystemLog, NewLogHandler(), LogHandler (+1 more)

### Community 48 - "crypto_test.go"
Cohesion: 0.09
Nodes (36): DecryptSecret(), EncryptSecret(), ParseEncryptionKey(), T, TestDecryptInvalidBase64(), TestDecryptInvalidKeyLength(), TestDecryptTooShortCiphertext(), TestEncryptDecryptRoundTrip() (+28 more)

### Community 49 - "ParseUserAgent"
Cohesion: 0.26
Nodes (17): parseBrowser(), parseOS(), ParseUserAgent(), T, TestParseUserAgentAndroid(), TestParseUserAgentBrowserOnly(), TestParseUserAgentChrome(), TestParseUserAgentChromeOS() (+9 more)

### Community 50 - "cmdTenantsList"
Cohesion: 0.36
Nodes (11): cmdTenants(), cmdTenantsGet(), cmdTenantsList(), countMembersPerTenant(), Context, MongoDB, ObjectID, Tenant (+3 more)

### Community 51 - "Store"
Cohesion: 0.24
Nodes (7): Context, Duration, MongoDB, RWMutex, New(), Store, ConfigVar

### Community 52 - "main"
Cohesion: 0.08
Nodes (28): main(), Emitter, Logger, MongoDB, NewAPIKeysHandler(), NewResendService(), Context, Service (+20 more)

### Community 53 - "APIPage.tsx"
Cohesion: 0.17
Nodes (12): APIKeysSection(), APIPage(), CreateKeyModal(), formatDate(), timeAgo(), WebhookDetailModal(), WebhookFormModal(), WebhooksSection() (+4 more)

### Community 54 - "Features"
Cohesion: 0.12
Nodes (16): Admin Interface, API Keys, Authentication & Identity, Billing & Credits (Stripe), Built-in API Documentation, CLI Administration, Features, MCP Server (AI Admin Access) (+8 more)

### Community 55 - "process.go"
Cohesion: 0.35
Nodes (14): capitalizeStr(), cmdRestart(), cmdStart(), cmdStop(), ensurePIDDir(), findProjectRoot(), isDirAt(), mustFindProjectRoot() (+6 more)

### Community 56 - "emitter.go"
Cohesion: 0.31
Nodes (6): Time, NewNoopEmitter(), Emitter, Event, EventType, NoopEmitter

### Community 57 - "createAdminEnv"
Cohesion: 0.17
Nodes (28): T, TestIntegration_APIKeys_NonRootTenantForbidden(), TestIntegration_CreateAPIKey_InvalidAuthority(), TestIntegration_CreateAPIKey_MissingName(), TestIntegration_CreateAPIKey_Success(), TestIntegration_CreateAPIKey_UserAuthority(), TestIntegration_DeleteAPIKey_NotFound(), TestIntegration_DeleteAPIKey_Success() (+20 more)

### Community 58 - "BootstrapHandler"
Cohesion: 0.23
Nodes (8): Handler, MongoDB, Request, ResponseWriter, RWMutex, NewBootstrapHandler(), BootstrapHandler, bootstrapStatusResponse

### Community 59 - "admin.go"
Cohesion: 0.17
Nodes (13): AuthMethod, MemberRole, Time, MemberRole, Time, ChangeRoleRequest, InviteMemberRequest, MemberResponse (+5 more)

### Community 60 - "ConfigHandler"
Cohesion: 0.31
Nodes (7): Logger, MongoDB, Request, ResponseWriter, Store, NewConfigHandler(), ConfigHandler

### Community 61 - "manifest.json"
Cohesion: 0.14
Nodes (13): description, display_name, documentation, homepage, license, long_description, manifest_version, name (+5 more)

### Community 62 - "LastSaaS"
Cohesion: 0.14
Nodes (13): API Documentation, Configuration, Environment Variables, Fork It and Keep Building with AI, How It Compares, LastSaaS, License, Optional (+5 more)

### Community 63 - "BundlesHandler"
Cohesion: 0.30
Nodes (8): Logger, MongoDB, Request, ResponseWriter, NewBundlesHandler(), validateBundleRequest(), bundleRequest, BundlesHandler

### Community 64 - "respondWithError"
Cohesion: 0.19
Nodes (12): Request, ResponseWriter, ResponseWriter, respondWithError(), Emitter, Logger, MongoDB, Request (+4 more)

### Community 65 - "init"
Cohesion: 0.43
Nodes (6): ObjectID, Time, ValidAPIKeyAuthority(), init(), APIKey, APIKeyAuthority

### Community 66 - "models_test.go"
Cohesion: 0.28
Nodes (12): T, TestAllWebhookEventTypesNoDuplicates(), TestAuthMethodConstants(), TestBillingStatusConstants(), TestUserHasAuthMethod(), TestUserHasAuthMethodEmpty(), TestUserIsLockedFuture(), TestUserIsLockedNil() (+4 more)

### Community 67 - "cmdUsersGet"
Cohesion: 0.20
Nodes (15): cmdUsers(), cmdUsersGet(), cmdUsersList(), cmdUsersRevokeSessions(), cmdUsersSetActive(), Context, MongoDB, ObjectID (+7 more)

### Community 68 - "Service"
Cohesion: 0.29
Nodes (4): Context, MongoDB, New(), Service

### Community 69 - "PlansPage.tsx"
Cohesion: 0.11
Nodes (25): plansApi, telemetryApi, getSessionId(), useTelemetry(), BundleFormModal(), BundleFormModalProps, formatPrice(), PlanFormModal() (+17 more)

### Community 70 - "ReadResponseBody"
Cohesion: 0.22
Nodes (22): T, TestIntegration_AdminCancelRootInvitation(), TestIntegration_AdminChangeRootMemberRole(), TestIntegration_AdminDashboard(), TestIntegration_AdminGetTenant(), TestIntegration_AdminGetTenantNotFound(), TestIntegration_AdminGetUser(), TestIntegration_AdminGetUserNotFound() (+14 more)

### Community 71 - "keywords"
Cohesion: 0.17
Nodes (12): keywords, admin, ai-native, billing, dashboard, go, health, monitoring (+4 more)

### Community 72 - "clr"
Cohesion: 0.22
Nodes (17): buildLogFilter(), cmdLogs(), Context, M, MongoDB, SystemLog, Time, logsFollow() (+9 more)

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

### Community 84 - "AnnouncementsHandler"
Cohesion: 0.36
Nodes (5): MongoDB, Request, ResponseWriter, NewAnnouncementsHandler(), AnnouncementsHandler

### Community 85 - "setupTestServer"
Cohesion: 0.23
Nodes (22): T, TestIntegration_BootstrapStatus(), TestIntegration_BootstrapStatusAfterInit(), TestIntegration_ChangePassword(), TestIntegration_ChangePasswordWrongCurrent(), TestIntegration_GetMe(), TestIntegration_GetMeNoToken(), TestIntegration_LoginNonexistentUser() (+14 more)

### Community 86 - "useTenant"
Cohesion: 0.13
Nodes (20): messagesApi, setTenantHeader(), AdminLayout(), AdminRoute(), TenantContext, TenantContextType, TenantProvider(), useTenant() (+12 more)

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
Cohesion: 0.50
Nodes (4): checkConfigIntegration(), cmdDoctor(), Context, MongoDB

### Community 95 - "plans_test.go"
Cohesion: 0.21
Nodes (21): T, TestIntegration_AdminCannotDeletePlan(), TestIntegration_ArchivePlan_Success(), TestIntegration_ArchiveSystemPlan_Forbidden(), TestIntegration_AssignPlan_Success(), TestIntegration_CreatePlan_DuplicateName(), TestIntegration_CreatePlan_MissingName(), TestIntegration_CreatePlan_Success() (+13 more)

### Community 96 - "counter_test.go"
Cohesion: 0.53
Nodes (5): T, TestConcurrentIncrements(), TestResendEmailsIncrement(), TestStripeAPICallsIncrement(), TestSwapResets()

### Community 97 - "models/billing.go"
Cohesion: 0.36
Nodes (8): ObjectID, Time, BillingStatus, DailyMetric, FinancialTransaction, InvoiceCounter, StripeMapping, TransactionType

### Community 98 - "testutil.go"
Cohesion: 0.17
Nodes (18): APIKeyAuthority, M, TestMain(), ConnectTestDB(), CreateTestAPIKey(), CreateTestInvitation(), findAndSetConfigDir(), APIKey (+10 more)

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

### Community 103 - "SystemLog"
Cohesion: 0.47
Nodes (5): ObjectID, Time, LogCategory, LogSeverity, SystemLog

### Community 104 - "SSOConnection"
Cohesion: 0.50
Nodes (4): ObjectID, Time, SSOAttributeMap, SSOConnection

### Community 105 - "package.json"
Cohesion: 0.40
Nodes (4): name, private, type, version

### Community 106 - "InsertTestLogs"
Cohesion: 0.32
Nodes (12): T, TestIntegration_LogsDateRange(), TestIntegration_LogsEmptyResults(), TestIntegration_LogsFilterByCategory(), TestIntegration_LogsFilterBySeverity(), TestIntegration_LogsListDefault(), TestIntegration_LogsMultiSeverityFilter(), TestIntegration_LogsPagination() (+4 more)

### Community 107 - "Plan"
Cohesion: 0.39
Nodes (7): ObjectID, Time, CreditResetPolicy, EntitlementType, EntitlementValue, Plan, PricingModel

### Community 108 - "WebhookEventType"
Cohesion: 0.52
Nodes (6): ObjectID, Time, ValidWebhookEventType(), Webhook, WebhookDelivery, WebhookEventType

### Community 109 - "Announcement"
Cohesion: 0.50
Nodes (3): ObjectID, Time, Announcement

### Community 110 - "CreditBundle"
Cohesion: 0.50
Nodes (3): ObjectID, Time, CreditBundle

### Community 111 - "EventDefinition"
Cohesion: 0.50
Nodes (3): ObjectID, Time, EventDefinition

### Community 112 - "MemberRole"
Cohesion: 0.27
Nodes (9): ObjectID, Time, RoleHasPermission(), T, TestRoleHasPermission(), TestValidRole(), ValidRole(), MemberRole (+1 more)

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

### Community 121 - "EventDefinitionModal.tsx"
Cohesion: 0.28
Nodes (7): Modal(), ModalProps, EventDefinitionModal(), FormData, Props, schema, EventDefinition

### Community 125 - "author"
Cohesion: 0.67
Nodes (3): author, name, url

### Community 126 - "repository"
Cohesion: 0.67
Nodes (3): repository, type, url

### Community 127 - "Deployment"
Cohesion: 0.67
Nodes (3): Deployment, Fly.io, Other Platforms

### Community 129 - "configstore/validate.go"
Cohesion: 0.47
Nodes (5): ValidateEnumValue(), validateTemplate(), ValidateValue(), enumOption, ConfigVarType

### Community 153 - "emitter_test.go"
Cohesion: 0.53
Nodes (5): T, TestEventStruct(), TestEventTypeConstants(), TestNoopEmitterEmit(), TestNoopEmitterImplementsInterface()

### Community 154 - "Invitation"
Cohesion: 0.50
Nodes (4): ObjectID, Time, Invitation, InvitationStatus

### Community 155 - "Tenant"
Cohesion: 0.40
Nodes (4): ObjectID, Time, BillingStatus, Tenant

### Community 156 - "TestPasswordValidation"
Cohesion: 0.67
Nodes (3): T, TestPasswordHashing(), TestPasswordValidation()

### Community 157 - "TelemetryEvent"
Cohesion: 0.50
Nodes (3): ObjectID, Time, TelemetryEvent

### Community 158 - "stripe.go"
Cohesion: 0.67
Nodes (3): ObjectID, CheckoutLineItem, CheckoutRequest

### Community 169 - "Button.tsx"
Cohesion: 0.29
Nodes (6): Button(), ButtonProps, ButtonSize, ButtonVariant, sizeClasses, variantClasses

### Community 170 - "ValidConfigVarType"
Cohesion: 0.47
Nodes (5): ObjectID, Time, ValidConfigVarType(), ConfigVar, ConfigVarType

### Community 171 - "Alert.tsx"
Cohesion: 0.40
Nodes (3): AlertProps, AlertVariant, variantClasses

### Community 172 - "Badge.tsx"
Cohesion: 0.40
Nodes (4): Badge(), BadgeProps, BadgeVariant, variantClasses

## Knowledge Gaps
- **350 isolated node(s):** `RegisterRequest`, `LoginRequest`, `RefreshRequest`, `MFARequiredResponse`, `VerifyEmailRequest` (+345 more)
  These have ≤1 connection - possible missing edges or undocumented components.
- **37 thin communities (<3 nodes) omitted from report** — run `graphify query` to explore isolated nodes.

## Suggested Questions
_Questions this graph is uniquely positioned to answer:_

- **Why does `setupTestServer()` connect `setupTestServer` to `CreateTestUser`, `AuthHandler`, `Client`, `middleware/tenant_test.go`, `setupIsolationEnv`, `AdminHandler`, `MustConnectTestDB`, `WebhooksHandler`, `GetUserFromContext`, `BillingHandler`, `LogHandler`, `main`, `emitter.go`, `createAdminEnv`, `BootstrapHandler`, `respondWithError`, `ReadResponseBody`, `plans_test.go`, `testutil.go`, `InsertTestLogs`?**
  _High betweenness centrality (0.158) - this node is a cross-community bridge._
- **Why does `main()` connect `main` to `WebhookHandler`, `oauth_test.go`, `PromotionsHandler`, `config_test.go`, `AdminHandler`, `RateLimiter`, `MustConnectTestDB`, `WebhooksHandler`, `Write`, `respondWithJSON`, `HealthHandler`, `GetUserFromContext`, `EventDefinitionsHandler`, `connectDB`, `BillingHandler`, `LogHandler`, `ConfigHandler`, `BundlesHandler`, `respondWithError`, `.TrackAuthenticated`?**
  _High betweenness centrality (0.102) - this node is a cross-community bridge._
- **Why does `Client` connect `Client` to `cmd_mcp.go`, `MongoDB`, `crypto_test.go`, `setupIsolationEnv`, `main`, `setupTestServer`?**
  _High betweenness centrality (0.073) - this node is a cross-community bridge._
- **Are the 148 inferred relationships involving `setupTestServer()` (e.g. with `TestIntegration_AdminCancelRootInvitation()` and `TestIntegration_AdminChangeRootMemberRole()`) actually correct?**
  _`setupTestServer()` has 148 INFERRED edges - model-reasoned connections that need verification._
- **Are the 115 inferred relationships involving `respondWithJSON()` (e.g. with `.CancelRootInvitation()` and `.ChangeRootMemberRole()`) actually correct?**
  _`respondWithJSON()` has 115 INFERRED edges - model-reasoned connections that need verification._
- **Are the 112 inferred relationships involving `respondWithError()` (e.g. with `.CancelRootInvitation()` and `.ChangeRootMemberRole()`) actually correct?**
  _`respondWithError()` has 112 INFERRED edges - model-reasoned connections that need verification._
- **Are the 77 inferred relationships involving `CreateTestUser()` (e.g. with `TestIntegration_AdminChangeRootMemberRole()` and `TestIntegration_AdminGetUser()`) actually correct?**
  _`CreateTestUser()` has 77 INFERRED edges - model-reasoned connections that need verification._