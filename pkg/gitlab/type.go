package gitlab

import "time"

type Project struct {
	ID                                        int                       `json:"id"`
	Description                               string                    `json:"description"`
	DescriptionHTML                           string                    `json:"description_html"`
	DefaultBranch                             string                    `json:"default_branch"`
	Visibility                                string                    `json:"visibility"`
	SSHURLToRepo                              string                    `json:"ssh_url_to_repo"`
	HTTPURLToRepo                             string                    `json:"http_url_to_repo"`
	WebURL                                    string                    `json:"web_url"`
	ReadmeURL                                 string                    `json:"readme_url"`
	TagList                                   []string                  `json:"tag_list"`
	Topics                                    []string                  `json:"topics"`
	Owner                                     Owner                     `json:"owner"`
	Name                                      string                    `json:"name"`
	NameWithNamespace                         string                    `json:"name_with_namespace"`
	Path                                      string                    `json:"path"`
	PathWithNamespace                         string                    `json:"path_with_namespace"`
	IssuesEnabled                             bool                      `json:"issues_enabled"`
	OpenIssuesCount                           int                       `json:"open_issues_count"`
	MergeRequestsEnabled                      bool                      `json:"merge_requests_enabled"`
	JobsEnabled                               bool                      `json:"jobs_enabled"`
	WikiEnabled                               bool                      `json:"wiki_enabled"`
	SnippetsEnabled                           bool                      `json:"snippets_enabled"`
	CanCreateMergeRequestIn                   bool                      `json:"can_create_merge_request_in"`
	ResolveOutdatedDiffDiscussions            bool                      `json:"resolve_outdated_diff_discussions"`
	ContainerRegistryEnabled                  bool                      `json:"container_registry_enabled"`
	ContainerRegistryAccessLevel              string                    `json:"container_registry_access_level"`
	SecurityAndComplianceAccessLevel          string                    `json:"security_and_compliance_access_level"`
	ContainerExpirationPolicy                 ContainerExpirationPolicy `json:"container_expiration_policy"`
	CreatedAt                                 time.Time                 `json:"created_at"`
	UpdatedAt                                 time.Time                 `json:"updated_at"`
	LastActivityAt                            time.Time                 `json:"last_activity_at"`
	CreatorID                                 int                       `json:"creator_id"`
	Namespace                                 Namespace                 `json:"namespace"`
	ImportURL                                 any                       `json:"import_url"`
	ImportType                                any                       `json:"import_type"`
	ImportStatus                              string                    `json:"import_status"`
	ImportError                               any                       `json:"import_error"`
	Permissions                               Permissions               `json:"permissions"`
	Archived                                  bool                      `json:"archived"`
	AvatarURL                                 string                    `json:"avatar_url"`
	LicenseURL                                string                    `json:"license_url"`
	License                                   License                   `json:"license"`
	SharedRunnersEnabled                      bool                      `json:"shared_runners_enabled"`
	GroupRunnersEnabled                       bool                      `json:"group_runners_enabled"`
	ForksCount                                int                       `json:"forks_count"`
	StarCount                                 int                       `json:"star_count"`
	RunnersToken                              string                    `json:"runners_token"`
	CiDefaultGitDepth                         int                       `json:"ci_default_git_depth"`
	CiForwardDeploymentEnabled                bool                      `json:"ci_forward_deployment_enabled"`
	CiForwardDeploymentRollbackAllowed        bool                      `json:"ci_forward_deployment_rollback_allowed"`
	CiAllowForkPipelinesToRunInParentProject  bool                      `json:"ci_allow_fork_pipelines_to_run_in_parent_project"`
	CiSeparatedCaches                         bool                      `json:"ci_separated_caches"`
	CiRestrictPipelineCancellationRole        string                    `json:"ci_restrict_pipeline_cancellation_role"`
	CiPipelineVariablesMinimumOverrideRole    string                    `json:"ci_pipeline_variables_minimum_override_role"`
	CiPushRepositoryForJobTokenAllowed        bool                      `json:"ci_push_repository_for_job_token_allowed"`
	PublicJobs                                bool                      `json:"public_jobs"`
	SharedWithGroups                          []SharedWithGroups        `json:"shared_with_groups"`
	RepositoryStorage                         string                    `json:"repository_storage"`
	OnlyAllowMergeIfPipelineSucceeds          bool                      `json:"only_allow_merge_if_pipeline_succeeds"`
	AllowMergeOnSkippedPipeline               bool                      `json:"allow_merge_on_skipped_pipeline"`
	AllowPipelineTriggerApproveDeployment     bool                      `json:"allow_pipeline_trigger_approve_deployment"`
	RestrictUserDefinedVariables              bool                      `json:"restrict_user_defined_variables"`
	OnlyAllowMergeIfAllDiscussionsAreResolved bool                      `json:"only_allow_merge_if_all_discussions_are_resolved"`
	RemoveSourceBranchAfterMerge              bool                      `json:"remove_source_branch_after_merge"`
	PrintingMergeRequestsLinkEnabled          bool                      `json:"printing_merge_requests_link_enabled"`
	RequestAccessEnabled                      bool                      `json:"request_access_enabled"`
	MergeMethod                               string                    `json:"merge_method"`
	SquashOption                              string                    `json:"squash_option"`
	AutoDevopsEnabled                         bool                      `json:"auto_devops_enabled"`
	AutoDevopsDeployStrategy                  string                    `json:"auto_devops_deploy_strategy"`
	ApprovalsBeforeMerge                      int                       `json:"approvals_before_merge"`
	Mirror                                    bool                      `json:"mirror"`
	MirrorUserID                              int                       `json:"mirror_user_id"`
	MirrorTriggerBuilds                       bool                      `json:"mirror_trigger_builds"`
	OnlyMirrorProtectedBranches               bool                      `json:"only_mirror_protected_branches"`
	MirrorOverwritesDivergedBranches          bool                      `json:"mirror_overwrites_diverged_branches"`
	ExternalAuthorizationClassificationLabel  any                       `json:"external_authorization_classification_label"`
	PackagesEnabled                           bool                      `json:"packages_enabled"`
	ServiceDeskEnabled                        bool                      `json:"service_desk_enabled"`
	ServiceDeskAddress                        any                       `json:"service_desk_address"`
	AutocloseReferencedIssues                 bool                      `json:"autoclose_referenced_issues"`
	SuggestionCommitMessage                   any                       `json:"suggestion_commit_message"`
	EnforceAuthChecksOnUploads                bool                      `json:"enforce_auth_checks_on_uploads"`
	MergeCommitTemplate                       any                       `json:"merge_commit_template"`
	SquashCommitTemplate                      any                       `json:"squash_commit_template"`
	IssueBranchTemplate                       string                    `json:"issue_branch_template"`
	MarkedForDeletionAt                       string                    `json:"marked_for_deletion_at"`
	MarkedForDeletionOn                       string                    `json:"marked_for_deletion_on"`
	ComplianceFrameworks                      []string                  `json:"compliance_frameworks"`
	WarnAboutPotentiallyUnwantedCharacters    bool                      `json:"warn_about_potentially_unwanted_characters"`
	PreReceiveSecretDetectionEnabled          bool                      `json:"pre_receive_secret_detection_enabled"`
	Statistics                                Statistics                `json:"statistics"`
	ContainerRegistryImagePrefix              string                    `json:"container_registry_image_prefix"`
	Links                                     Links                     `json:"_links"`
}
type Owner struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
}
type ContainerExpirationPolicy struct {
	Cadence         string    `json:"cadence"`
	Enabled         bool      `json:"enabled"`
	KeepN           any       `json:"keep_n"`
	OlderThan       any       `json:"older_than"`
	NameRegex       any       `json:"name_regex"`
	NameRegexDelete any       `json:"name_regex_delete"`
	NameRegexKeep   any       `json:"name_regex_keep"`
	NextRunAt       time.Time `json:"next_run_at"`
}
type Namespace struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Path      string `json:"path"`
	Kind      string `json:"kind"`
	FullPath  string `json:"full_path"`
	AvatarURL string `json:"avatar_url"`
	WebURL    string `json:"web_url"`
}
type ProjectAccess struct {
	AccessLevel       int `json:"access_level"`
	NotificationLevel int `json:"notification_level"`
}
type GroupAccess struct {
	AccessLevel       int `json:"access_level"`
	NotificationLevel int `json:"notification_level"`
}
type Permissions struct {
	ProjectAccess ProjectAccess `json:"project_access"`
	GroupAccess   GroupAccess   `json:"group_access"`
}
type License struct {
	Key       string `json:"key"`
	Name      string `json:"name"`
	Nickname  string `json:"nickname"`
	HTMLURL   string `json:"html_url"`
	SourceURL string `json:"source_url"`
}
type SharedWithGroups struct {
	GroupID          int    `json:"group_id"`
	GroupName        string `json:"group_name"`
	GroupFullPath    string `json:"group_full_path"`
	GroupAccessLevel int    `json:"group_access_level"`
}
type Statistics struct {
	CommitCount           int `json:"commit_count"`
	StorageSize           int `json:"storage_size"`
	RepositorySize        int `json:"repository_size"`
	WikiSize              int `json:"wiki_size"`
	LfsObjectsSize        int `json:"lfs_objects_size"`
	JobArtifactsSize      int `json:"job_artifacts_size"`
	PipelineArtifactsSize int `json:"pipeline_artifacts_size"`
	PackagesSize          int `json:"packages_size"`
	SnippetsSize          int `json:"snippets_size"`
	UploadsSize           int `json:"uploads_size"`
	ContainerRegistrySize int `json:"container_registry_size"`
}
type Links struct {
	Self          string `json:"self"`
	Issues        string `json:"issues"`
	MergeRequests string `json:"merge_requests"`
	RepoBranches  string `json:"repo_branches"`
	Labels        string `json:"labels"`
	Events        string `json:"events"`
	Members       string `json:"members"`
	ClusterAgents string `json:"cluster_agents"`
}
